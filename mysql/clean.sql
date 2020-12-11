CREATE DATABASE  IF NOT EXISTS `clean` DEFAULT CHARACTER SET utf8mb4;

USE `clean`;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `ID` char(36) NOT NULL,
    `Name` VARCHAR(255) NOT NULL DEFAULT '',
    `Email` VARCHAR(255) NOT NULL,
    `Password` VARCHAR(255) NOT NULL,
    `CreatedDate` datetime NOT NULL,
    `DeletedDate` datetime DEFAULT NULL,
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `preupload`;

CREATE TABLE `preupload` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Title` text NOT NULL,
  `ShotDate` datetime,
  `CreatedDate` datetime NOT NULL,
  `DeletedDate` datetime NOT NULL,
  `Path` text NOT NULL,
  PRIMARY KEY (`ID`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `upload`;

CREATE TABLE `upload` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `Title` text NOT NULL,
  `ShotDate` datetime,
  `CreatedDate` datetime NOT NULL,
  `DeletedDate` datetime NOT NULL,
  `Path` text NOT NULL,
  PRIMARY KEY (`ID`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

