// main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nullseed/logruseq"

	"github.com/sirupsen/logrus"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Setup Logrus with Seq logging
	setupLogging()
}

func setupLogging() {
	// Get the SEQ_URL environment variable
	seqURL := os.Getenv("SEQ_URL")

	// Set up Logrus with logreseq hook
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Add the Seq hook to logrus
	hook := logruseq.NewSeqHook(seqURL)

	// Add the hook to Logrus
	logrus.AddHook(hook)
}

func main() {
	r := gin.Default()

	// Sample route with logging
	r.GET("/ping", func(c *gin.Context) {
		logrus.Info("Ping route accessed")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
