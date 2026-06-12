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
)

// GET /meetings
func GetMeetings(c *gin.Context) {
	meetings := []models.Meeting{}
	collection := config.GetCollection("meetings")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	userIDStr := c.Query("userId")
	if userIDStr != "" {
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err == nil {
			filter["userId"] = userID
		}
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var meeting models.Meeting
		cursor.Decode(&meeting)
		meetings = append(meetings, meeting)
	}

	c.JSON(http.StatusOK, gin.H{"data": meetings})
}

// GET /meetings/:id
func GetMeeting(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var meeting models.Meeting
	collection := config.GetCollection("meetings")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&meeting)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": meeting})
}

// POST /meetings
func CreateMeeting(c *gin.Context) {
	var input models.Meeting
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = primitive.NewObjectID()
	collection := config.GetCollection("meetings")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// PATCH /meetings/:id
func UpdateMeeting(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := config.GetCollection("meetings")

	var input models.Meeting
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateData := bson.M{}
	if input.Title != "" {
		updateData["title"] = input.Title
	}
	if input.Date != "" {
		updateData["date"] = input.Date
	}
	if input.Time != "" {
		updateData["time"] = input.Time
	}
	if input.Location != "" {
		updateData["location"] = input.Location
	}
	if input.Description != "" {
		updateData["description"] = input.Description
	}
	if input.Tags != nil {
		updateData["tags"] = input.Tags
	}

	update := bson.M{"$set": updateData}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedMeeting models.Meeting
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&updatedMeeting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated meeting: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedMeeting})
}

// DELETE /meetings/:id
func DeleteMeeting(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := config.GetCollection("meetings")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
