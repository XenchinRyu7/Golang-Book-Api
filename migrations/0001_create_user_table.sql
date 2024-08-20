CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    -- UUID type stored as CHAR(36)
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);