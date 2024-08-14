# GolangApp API With SQLite

This is a API built with Go using the Echo framework, GORM for database management and SQLite. The API provides user and group management features, including user authentication. This guide will walk you through the steps to set up and run the API using Docker.

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your machine
- [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine

## Getting Started

Follow these steps to set up and run the GolangApp API in a Docker container.

### 1. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/passyaa/golang-sqlite.git
cd golang-sqlite
```

### 2. Build and Run the Containers
This command will build the Go application into a Docker image and start the application container.
```bash
docker-compose up --build
```

### 3. Verify the Setup
Once the containers are up and running, you can verify the API by accessing the following endpoint in Postman and use Authorization with Basic Auth Username : spadmin & Password : admin

1. Get all users: http://localhost:8080/users
2. Get all groups: http://localhost:8080/groups

### 4. List API Endpoints
You can Download Colletion Postman [Postman Collection](https://github.com/passyaa/GolangAppCURD/blob/main/golangApp_postman_collection.json),
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
11. POST /groups - Create a new groups
12. POST /users/:id/groups/:group_id - Assign a group to a user
13  . DELETE /groups/:group_id - Remove a group

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
If you want to customize the initial database schema or seed data, you can modify the code in config.go where the SQLite database is initialized and migrated automatically.

### 8. Troubleshooting
- If you encounter any issues with the container not starting, ensure Docker and Docker Compose are installed correctly and check for any error messages in the terminal.
- Make sure port 8080 (for the API) is not in use by other applications.