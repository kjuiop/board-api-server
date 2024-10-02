CREATE DATABASE IF NOT EXISTS board;

CREATE TABLE IF NOT EXISTS member (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `nickname` VARCHAR(255) NOT NULL,
    `adminYn` CHAR(1) NOT NULL DEFAULT 'N',
    `deleteYn` CHAR(1) NOT NULL DEFAULT 'N',
    `status` VARCHAR(20) NOT NULL,
    `createdAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
