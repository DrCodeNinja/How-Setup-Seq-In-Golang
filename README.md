# Configure Seq in Golang Application

## Table of Contents

- [Step 1 - Setting Up Seq in Docker](#step-1---setting-up-seq-in-docker)
- [Prerequisites](#prerequisites)
  - [1-1 Create a Docker Network (Optional)](#1-1-create-a-docker-network-optional)
  - [1-2 Pull the Seq Docker Image](#1-2-pull-the-seq-docker-image)
  - [1-3 Run Seq Container](#1-3-run-seq-container)
  - [1-4 Access Seq](#1-4-access-seq)
- [Step 2 - Setting Up the Project](#step-2---setting-up-the-project)
  - [2-1 Initialize the Project](#2-1-initialize-the-project)
  - [2-2 Install Dependencies](#2-2-install-dependencies)
  - [2-3 Create the .env File](#2-3-create-the-env-file)
  - [2-4 Configuring the Project (main.go)](#2-4-configuring-the-project-maingo)
  - [2-5 Running the Application](#2-5-running-the-application)
  - [2-6 Test the Route](#2-6-test-the-route)
- [Conclusion](#conclusion)


## Step 1 - Setting Up Seq in Docker

Seq is a powerful log server that helps you collect, search, and analyze application logs. Running Seq in Docker is an efficient way to deploy it quickly without worrying about manual installations. This guide will walk you through the complete setup of Seq in Docker.

### Prerequisites

- Docker installed on your machine.
- Docker Compose (optional, but recommended for easier management).
- A basic understanding of Docker commands.

### 1-1 Create a Docker Network (Optional)

Creating a dedicated Docker network is good practice for better communication between containers.

```bash
docker network create seq_network
```

### 1-2 Pull the Seq Docker Image

Seq provides an official Docker image available on Docker Hub. Pull the latest version using the command below:

```bash
docker pull datalust/seq
```

### 1-3 Run Seq Container

Run the Seq container with the following command:

If you created seperate network use this:
```bash
docker run -d --name seq --network seq_network -p 5341:80 -v "$(pwd)/seq-data:/data" -e ACCEPT_EULA=Y datalust/seq:latest
```
else use this:
```bash
docker run -d --name seq -p 5341:80 -v "$(pwd)/seq-data:/data" -e ACCEPT_EULA=Y datalust/seq:latest
```

- [--name] seq gives the container a name.
- [--network seq-network] attaches the container to the custom network.
- [-p 5341:5341] maps the container’s port 5341 to the host’s port 5341.
- [-v seq-data:/data] mounts a volume named seq-data for persistent storage.
- [datalust/seq:latest] specifies the image to use.

### 1-4 Access Seq

Open your web browser and go to:
```bash
http://localhost:5341
```

You should see the Seq web interface.

## Step 2 - Setting Up the Project

### 2-1 Initialize the Project:

```bash
mkdir gin-golang-project
cd gin-golang-project
```
#### Initialize a new Go module

```bash
go mod init gin-golang-project
```

### 2-2 Install Dependencies:

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/joho/godotenv
go get -u github.com/sirupsen/logrus
go get -u github.com/nullseed/logruseq
```

#### Here is a simple structure for your project

```bash
gin-golang-project/
│
├── .env
├── go.mod
├── go.sum
└── main.go
```

### 2-3 Creating the [.env] File:

```bash
# .env
SEQ_URL=http://localhost:5341
PORT=8080
```

### 2-4 Configuring the Project (main.go)

```bash
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
```

### 2-5 Running the Application:

```bash
go run main.go / go run .
```

### 2-5 Test the route:

```bash
http://localhost:8080/ping
```

Copy and paste this url in the browser or make postman request and got to seq dashboard you will see the logs.

## Thanks for refer this, is this helpfull don't foget to give a star for this.






