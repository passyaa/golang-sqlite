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
This command will build a Go application into a Docker image in 1 or 2 minutes depending on the network and start the application container.
```bash
docker-compose up --build
```

### 3. Verify the Setup
Once the containers are up and running, you can verify the API by accessing the following endpoint in Postman and use Authorization with Basic Auth Username : spadmin & Password : admin

1. Get all users: http://localhost:8080/api/v1/users
2. Get all groups: http://localhost:8080/api/v1/groups

### 4. List API Endpoints
You can Download Colletion Postman [Postman Collection](https://github.com/passyaa/golang-sqlite/blob/main/GolangApp-API.postman_collection.json),

Here are some of the available API endpoints:

| Method | Endpoint                                  | Description                                                                           |
|--------|-------------------------------------------|---------------------------------------------------------------------------------------|
| GET    | /api/v1/users/:id                         | Fetch a single user by ID                                                             |
| GET    | /api/v1/users                             | Fetch all users                                                                       |
| POST   | /api/v1/users                             | Create a new user                                                                     |
| PUT    | /api/v1/users/:id                         | Update a user by ID                                                                   |
| DELETE | /api/v1/users/:id                         | Delete a user by ID                                                                   |
| PUT    | /api/v1/users/:id/enable                  | Enable a user                                                                         |
| PUT    | /api/v1/users/:id/disable                 | Disable a user                                                                        |
| PUT    | /api/v1/users/:id/reset_password          | Reset user password                                                                   |
| GET    | /api/v1/groups/:id                        | Fetch a single group by ID                                                            |
| GET    | /api/v1/groups                            | Fetch all groups                                                                      |
| POST   | /api/v1/groups                            | Create a new group                                                                    |
| POST   | /api/v1/users/:id/groups/:group_id        | Assign a group to a user                                                              |
| DELETE | /api/v1/users/:id/groups/:group_id        | Remove assigned group from a user                                                     |
| DELETE | /api/v1/groups/:group_id                  | Remove a group                                                                        |

### 5. List URL With User Interface (HTML and CSS Bootstrap)
| Method | Endpoint                                  | Description                                                                           |
|--------|-------------------------------------------|---------------------------------------------------------------------------------------|
| GET    | /                                         | Home - Display homepage with login option                                             |
| GET    | /login                                    | Login page                                                                            |
| POST   | /login                                    | Handle login form submission                                                          |
| GET    | /profile/:id                              | Profile page - Redirected after successful login, displays user profile by ID         |
| GET    | /logout                                   | Logout - Log the user out and redirect to home page                                   |

### 6. Stopping the Containers
To stop the running containers, press Ctrl+C in the terminal where Docker Compose is running. You can also use the following command to stop and remove the containers:

```bash
docker-compose down
```

### 7. Additional Commands
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

### 8. Customizing the Database Initialization
If you want to customize the initial database schema or seed data, you can modify the code in config.go where the SQLite database is initialized and migrated automatically.

### 9. Troubleshooting
- If you encounter any issues with the container not starting, ensure Docker and Docker Compose are installed correctly and check for any error messages in the terminal.
- Make sure port 8080 (for the API) is not in use by other applications.


## Author

This project was created by [Kamal](https://www.linkedin.com/in/maulana-kamal-pasya/).

I developed this project as part of my personal learning journey in Golang. If you find it useful, feel free to use it or contribute to it. You can reach me at [maulana.kp@netpoleons.com] or connect with me on [LinkedIn](https://www.linkedin.com/in/maulana-kamal-pasya/).

Thank you for checking out my work! cheers 🍻 🍻