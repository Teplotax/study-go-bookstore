USE simplerest;

CREATE TABLE IF NOT EXISTS books (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY
    name        VARCHAR(255),
    author      VARCHAR(255),
    publication VARCHAR(255)
);