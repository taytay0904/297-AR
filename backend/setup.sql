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
  `modelPath` varchar(255) NOT NULL,
  `detailPath` varchar(255) NOT NULL,
  `soundPath` varchar(255) NOT NULL,
  `animationPath` varchar(255) NOT NULL,
  PRIMARY KEY (id)
); 

-- Create Table [detail]
DROP TABLE IF EXISTS `model`.`detail`;
CREATE TABLE `model`.`detail` (
  `id` int NOT NULL AUTO_INCREMENT, 
  `height` varchar(200) NOT NULL UNIQUE,
  `weight` varchar(100) NOT NULL UNIQUE,
  `built` varchar(100) NOT NULL UNIQUE,
  `designer` varchar(100) NOT NULL UNIQUE,
  `timestamp` date NOT NULL DEFAULT CURRENT_TIMESTAMP, 
   PRIMARY KEY (id)
); 

-- Create Table [modelDetailAssociate]
-- DROP TABLE IF EXISTS `model`.`modelDetailAssociate`;
-- CREATE TABLE `model`.`modelDetailAssociate` (
--   `modelID` varchar(200) NOT NULL, 
--   `detailID` varchar(200) NOT NULL UNIQUE,
--   `timestamp` date NOT NULL DEFAULT CURRENT_TIMESTAMP
-- ); 

insert into model values ("1", "Robert_Cai", "/assets/.mod/pa_drone", "fake_detail_path1", "fake_sound_path1", "fake_am_path1");
insert into model values ("2", "Kun_Su", "/assets/.mod/pa_warrior", "fake_detail_path2", "fake_sound_path2", "fake_am_path2");
insert into detail values ("1", "150m", "500kg", "iron", "Robert_Cai", DEFAULT);
insert into detail values ("2", "110m", "5100kg", "meterial", "Kun_Su", DEFAULT);

select * from model;
select * from detail;