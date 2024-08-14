# GolangApp API

This is a simple API built with Go using the Echo framework and GORM for database management. The API provides user and group management features, including user authentication. This guide will walk you through the steps to set up and run the API using Docker.

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your machine
- [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine

## Getting Started

Follow these steps to set up and run the GolangApp API in a Docker container.

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/your-username/golangapp.git
cd golangapp
```

### 2. Build and Run the Containers
This command :
```bash
docker-compose up --build
```
will build the Go application into a Docker image.
Start a MySQL database container and initialize it with the necessary tables and sample data.
Start the Go application container.

### 3. Verify the Setup
Once the containers are up and running, you can verify the API by accessing the following endpoint in your browser or using a tool like Postman:

1. Get all users: http://localhost:8080/users
2. Get all groups: http://localhost:8080/groups

### 4. List API Endpoints
Here are some of the available API endpoints:

1. GET /users/:id - Fetch a single user by ID
2. GET /users - Fetch all users
3. POST /users - Create a new user
4. PUT /users/:id - Update a user by ID
5. DELETE /users/:id - Delete a user by ID
6. PUT /users/:id/enable - Enable a user
7. PUT /users/:id/disable - Disable a user
8. PUT /users/:id/reset_password - Reset user password
9. GET /groups/:id - Fetch a single group by ID
10. GET /groups - Fetch all groups
11. POST /users/:id/groups/:group_id - Assign a group to a user
12. DELETE /groups/:group_id - Remove a group

### 5. Stopping the Containers
To stop the running containers, press Ctrl+C in the terminal where Docker Compose is running. You can also use the following command to stop and remove the containers:

```bash
docker-compose down
```

### 6. Additional Commands
To rebuild the images without using the cache:
```bash
docker-compose build --no-cache
```

To view the logs:
```bash
docker-compose logs
```

To start the containers in the background (detached mode):
```bash
docker-compose up -d
```

### 7. Customizing the Database Initialization
If you want to customize the initial database schema or seed data, edit the db/init.sql file. This file is executed when the MySQL container is first created.

### 8. Troubleshooting
- If you encounter any issues with the containers not starting, ensure Docker and Docker Compose are installed correctly and check for any error messages in the terminal.
- Make sure ports 8080 (for the API) and 3306 (for MySQL) are not in use by other applications.