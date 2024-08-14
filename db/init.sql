-- DDL: Membuat tabel users dan groups
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    is_enabled BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE groups (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE user_groups (
    user_id INT,
    group_id INT,
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

-- DML: Mengisi data awal
INSERT INTO users (username, email, password, first_name, last_name) VALUES 
('john_doe', 'john@example.com', 'hashed_password', 'John', 'Doe'),
('jane_doe', 'jane@example.com', 'hashed_password', 'Jane', 'Doe');

INSERT INTO groups (name, description) VALUES 
('Admin', 'Administrative Group'),
('User', 'Regular User Group');

INSERT INTO user_groups (user_id, group_id) VALUES 
(1, 1), -- John Doe assigned to Admin
(2, 2); -- Jane Doe assigned to User
