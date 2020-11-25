CREATE DATABASE IF NOT EXISTS medicareUsers; 

USE medicareUsers

CREATE TABLE IF NOT EXISTS users( 
    UserId INT NOT NULL AUTO_INCREMENT, 
    Email VARCHAR(50) NOT NULL, 
    Pass VARCHAR(250) NOT NULL,
    PRIMARY KEY (UserId) 
    );

INSERT INTO users (Email, Pass)
    VALUES 
    ("user1@email.com", "$2a$04$bfHERBBuv09alXt9Yu5Ym.wGTmAhCEjbXTmePNxGAt09PU9Oj2cVm"),
    ("user2@email.com", "$2a$04$bfHERBBuv09alXt9Yu5Ym.wGTmAhCEjbXTmePNxGAt09PU9Oj2cVm"),
    ("user3@email.com", "$2a$04$bfHERBBuv09alXt9Yu5Ym.wGTmAhCEjbXTmePNxGAt09PU9Oj2cVm"),
    ("user4@email.com", "$2a$04$bfHERBBuv09alXt9Yu5Ym.wGTmAhCEjbXTmePNxGAt09PU9Oj2cVm"),
    ("user5@email.com", "$2a$04$bfHERBBuv09alXt9Yu5Ym.wGTmAhCEjbXTmePNxGAt09PU9Oj2cVm");
