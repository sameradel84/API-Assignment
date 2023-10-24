CREATE DATABASE petsDB;
USE petsBD;
CREATE TABLE pets (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    age INT,
    breed VARCHAR(255)
);

INSERT INTO pets (name, age, breed) VALUES
    ('catty', 3, 'cat'),
    ('doggy', 5, 'dog');