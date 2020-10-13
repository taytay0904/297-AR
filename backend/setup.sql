DROP database IF EXISTS model;
CREATE database model
DEFAULT CHARACTER SET utf8;
SET GLOBAL local_infile = 1;
SET default_storage_engine = InnoDB;


-- Specify Database [foodApp]
USE model;


-- Create Table [model]
DROP TABLE IF EXISTS `model`.`model`;
CREATE TABLE `model`.`model` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL UNIQUE,
  `ICON_Path` varchar(20) NOT NULL,
  `model_File_Path` varchar(255) NOT NULL,
  `detail_Info_ID` varchar(255) NOT NULL,
  `sound_Path` varchar(255) NOT NULL,
  `timestamp` date NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  PRIMARY KEY (id)
); 

-- Create Table [detail]
DROP TABLE IF EXISTS `model`.`detail`;
CREATE TABLE `model`.`detail` (
  `id` int NOT NULL AUTO_INCREMENT, 
  `height` varchar(200) NOT NULL UNIQUE,
  `weight` varchar(100) NOT NULL UNIQUE,
  `width` varchar(100) NOT NULL UNIQUE,
  `designer` varchar(100) NOT NULL UNIQUE,
  `timestamp` date NOT NULL DEFAULT CURRENT_TIMESTAMP, 
   PRIMARY KEY (id)
); 

-- Create Table [modelDetailAssociate]
DROP TABLE IF EXISTS `model`.`modelDetailAssociate`;
CREATE TABLE `model`.`modelDetailAssociate` (
  `modelID` varchar(200) NOT NULL, 
  `detailID` varchar(200) NOT NULL UNIQUE,
  `timestamp` date NOT NULL DEFAULT CURRENT_TIMESTAMP
); 

