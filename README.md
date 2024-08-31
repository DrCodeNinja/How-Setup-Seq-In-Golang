### Configure Seq in Golang Application

## Step1 - Setting Up Seq in Docker

Seq is a powerful log server that helps you collect, search, and analyze application logs. Running Seq in Docker is an efficient way to deploy it quickly without worrying about manual installations. This guide will walk you through the complete setup of Seq in Docker.

# Prerequisites

- Docker installed on your machine.
- Docker Compose (optional, but recommended for easier management).
- A basic understanding of Docker commands.

# 1-1 Create a Docker Network (Optional)

Creating a dedicated Docker network is good practice for better communication between containers.

```bash
docker network create seq_network
```

# 1-2 Pull the Seq Docker Image

Seq provides an official Docker image available on Docker Hub. Pull the latest version using the command below:

```bash
docker pull datalust/seq
```

# 1-3 Run Seq Container

Run the Seq container with the following command:

```bash
docker run -d --name seq -p 5341:80 -v "$(pwd)/seq-data:/data" -e ACCEPT_EULA=Y datalust/seq:latest
```

- [--name] seq gives the container a name.
- [--network seq-network] attaches the container to the custom network.
- [-p 5341:5341] maps the container’s port 5341 to the host’s port 5341.
- [-v seq-data:/data] mounts a volume named seq-data for persistent storage.
- [datalust/seq:latest] specifies the image to use.

# 1-4 Access Seq

Open your web browser and go to:
```bash
http://localhost:5341
```

You should see the Seq web interface.
