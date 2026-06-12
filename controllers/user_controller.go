package controllers

import (
	"context"
	"net/http"
	"rapatln_backend/config"
	"rapatln_backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Register
func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if user exists
	var existingUser models.User
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Email:    input.Email,
		Password: string(hashedPassword),
		Name:     input.Name,
		Position: "Member",
		Bio:      "Welcome to RapatIn!",
		Role:     "user",
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

// Login
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /user/:id - Get specific user
func GetUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PUT /user/:id
func UpdateUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":             input.Name,
			"position":         input.Position,
			"bio":              input.Bio,
			"profileImagePath": input.ProfileImagePath,
			"email":            input.Email,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedUser models.User
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

// Check if user is admin helper
func isAdmin(adminIDStr string) bool {
	if adminIDStr == "" {
		return false
	}
	id, err := primitive.ObjectIDFromHex(adminIDStr)
	if err != nil {
		return false
	}
	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return false
	}
	return user.Role == "admin"
}

// GET /api/admin/stats
func GetAdminStats(c *gin.Context) {
	adminID := c.GetHeader("X-Admin-ID")
	if !isAdmin(adminID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Admin access required."})
		return
	}

	usersCol := config.GetCollection("users")
	meetingsCol := config.GetCollection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	totalUsers, err := usersCol.CountDocuments(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalMeetings, err := meetingsCol.CountDocuments(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{
		"totalUsers":    totalUsers,
		"totalMeetings": totalMeetings,
	}})
}

// GET /api/admin/users
func AdminGetAllUsers(c *gin.Context) {
	adminID := c.GetHeader("X-Admin-ID")
	if !isAdmin(adminID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Admin access required."})
		return
	}

	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	users := []models.User{}
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// PUT /api/admin/users/:id/role
func AdminUpdateUserRole(c *gin.Context) {
	adminID := c.GetHeader("X-Admin-ID")
	if !isAdmin(adminID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Admin access required."})
		return
	}

	targetID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var input struct {
		Role string `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Role != "user" && input.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role value"})
		return
	}

	collection := config.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(ctx, bson.M{"_id": targetID}, bson.M{"$set": bson.M{"role": input.Role}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User role updated successfully"})
}

// GET /api/admin/meetings
func AdminGetAllMeetings(c *gin.Context) {
	adminID := c.GetHeader("X-Admin-ID")
	if !isAdmin(adminID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Admin access required."})
		return
	}

	collection := config.GetCollection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	meetings := []models.Meeting{}
	for cursor.Next(ctx) {
		var meeting models.Meeting
		cursor.Decode(&meeting)
		meetings = append(meetings, meeting)
	}

	c.JSON(http.StatusOK, gin.H{"data": meetings})
}

// DELETE /api/admin/meetings/:id
func AdminDeleteMeeting(c *gin.Context) {
	adminID := c.GetHeader("X-Admin-ID")
	if !isAdmin(adminID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Admin access required."})
		return
	}

	meetingID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID format"})
		return
	}

	collection := config.GetCollection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": meetingID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meeting deleted successfully"})
}
