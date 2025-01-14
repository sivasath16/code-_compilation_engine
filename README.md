# Code Execution Engine

The Code Execution Engine is a web-based platform that allows users to write and execute code in multiple programming languages (JavaScript, Python, Java, and C++). It uses a React frontend, a Go-based backend, and Docker containers for isolated code execution.

## Features

- Write code in a web-based editor using the Monaco Editor.

- Support for JavaScript, Python, Java, and C++.

- Isolated code execution using Docker.

- Handles compilation errors, runtime errors, and successful outputs.

- Displays error messages in red for better visualization.

# Prerequisites

Ensure the following tools are installed on your system:

- Docker, Node.js, Go, RabbitMQ, MongoDB

# Run RabbitMQ and MongoDB Containers

``` 
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
```
```
docker run -d --name mongodb -p 27017:27017 mongo
```

# Execution Environment Setup
```
docker build -t code-executor-java -f docker/java/Dockerfile .
```
```
docker build -t code-executor-python -f docker/python/Dockerfile .
```
```
docker build -t code-executor-cpp -f docker/cpp/Dockerfile .
```
```
docker build -t code-executor-javascript -f docker/javascript/Dockerfile .
```

![WhatsApp Image 2025-01-13 at 22 24 53_187530ae](https://github.com/user-attachments/assets/adf3d015-423b-40c5-a1b8-026bea8874d6)
![WhatsApp Image 2025-01-13 at 22 24 53_759b58ee](https://github.com/user-attachments/assets/dfbddfa6-f8d6-47f0-9461-c955ebfb4082)

