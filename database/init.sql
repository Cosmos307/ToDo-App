CREATE DATABASE IF NOT EXISTS todo_app;

CREATE TABLE IF NOT EXISTS users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL

);

CREATE TABLE IF NOT EXISTS categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    category_id int,
    priority ENUM ('highest', 'high', 'medium', 'low', 'lowest') NOT NULL DEFAULT 'medium',
    status ENUM(
        'pending', 
        'in_progress', 
        'completed', 
        'on_hold', 
        'blocked', 
        'cancelled'
    ) NOT NULL DEFAULT 'pending',
    due_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Example data for users
INSERT INTO users (name, email, password)
VALUES
    ('Max', 'Max@email.com', '123'),
    ('Zara', 'Zara@email.com', '123');

-- Example data for categories
INSERT INTO categories (title, user_id)
VALUES 
    ('Todo-App', 1),
    ('Welt retten', 2);

-- Example data for tasks
INSERT INTO tasks (title, description, category_id, priority, status, due_date)
VALUES
    ('datenbank', 'Datenbank für die Todo-App implementieren', 1, 'highest', 'in_progress', '2024-11-30 00:00:00'),
    ('backend', 'Backend für die Todo-App implementieren', 1, 'high', 'on_hold', '2024-12-05 00:00:00'),
    ('frontend', 'Frontend für die Todo-App implementieren', 1, 'medium', 'blocked', '2024-12-10 00:00:00'),
    ('Schlafen', 'Bis Mittag schlafen und Abends früh schlafen gehen', 2, 'high', 'on_hold', '2024-12-10 00:00:00'),
    ('Essen', 'Früh essen und Abend essen', 1, 'highest', 'pending', '2024-12-10 00:00:00');
