-- MySQL dump 10.16  Distrib 10.1.38-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: libraryapi
-- ------------------------------------------------------
-- Server version	10.1.38-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `book_transactions`
--

DROP TABLE IF EXISTS `book_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `book_transactions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `nis` int(10) unsigned NOT NULL,
  `book_id` int(10) unsigned DEFAULT NULL,
  `status` longtext NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book_transactions`
--

LOCK TABLES `book_transactions` WRITE;
/*!40000 ALTER TABLE `book_transactions` DISABLE KEYS */;
INSERT INTO `book_transactions` VALUES (1,'Muhammad Yusup',11333,1,'Dipinjam','2019-10-11 06:06:04','2019-10-11 06:06:04',NULL),(4,'Edy',123123,4,'Dikembalikan','2019-10-12 14:54:49','2019-10-12 15:15:47',NULL);
/*!40000 ALTER TABLE `book_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `isbn` int(10) unsigned DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `author` varchar(100) DEFAULT NULL,
  `category_id` int(10) unsigned DEFAULT NULL,
  `bookshelf_id` int(10) unsigned DEFAULT NULL,
  `publisher` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,19912312,'entah apa yang merasuki','Gagak',2,2,'burung gagak','2019-10-11 00:16:21','2019-10-12 15:15:25',NULL),(2,0,'burung gagk','gagagkagak',8,4,'','2019-10-12 01:34:40','2019-10-12 10:25:15','2019-10-12 14:15:15'),(3,0,'bisnis','entah siapa',8,8,'','2019-10-12 01:36:08','2019-10-12 10:16:32','2019-10-12 14:15:17'),(4,998123,'kitab sutasoma','mpu tantular',3,6,'erlangga','2019-10-12 14:51:25','2019-10-12 14:52:13',NULL),(5,1283,'salah','slah',2,2,'sss','2019-10-12 15:15:40','2019-10-12 15:15:40','2019-10-12 15:15:42');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bookshelves`
--

DROP TABLE IF EXISTS `bookshelves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bookshelves` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bookshelves`
--

LOCK TABLES `bookshelves` WRITE;
/*!40000 ALTER TABLE `bookshelves` DISABLE KEYS */;
INSERT INTO `bookshelves` VALUES (2,'A1','2019-10-10 05:58:43','2019-10-10 05:58:43',NULL),(3,'A2','2019-10-10 23:51:44','2019-10-10 23:51:44',NULL),(4,'A3','2019-10-10 23:54:02','2019-10-11 00:00:35',NULL),(5,'X2','2019-10-11 00:01:07','2019-10-11 00:01:07','2019-10-11 00:01:14'),(6,'A4','2019-10-11 13:40:09','2019-10-12 00:36:40',NULL),(7,'B1','2019-10-11 13:42:17','2019-10-11 13:42:17',NULL),(8,'B2','2019-10-12 00:37:39','2019-10-12 00:37:39',NULL),(9,'BX','2019-10-12 14:09:27','2019-10-12 14:09:27','2019-10-12 14:09:30');
/*!40000 ALTER TABLE `bookshelves` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(125) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (2,'Sastrawati','2019-10-10 07:49:43','2019-10-11 13:25:42',NULL),(3,'Hukum','2019-10-10 23:19:56','2019-10-10 23:26:37',NULL),(4,'','2019-10-10 23:28:02','2019-10-10 23:28:02','2019-10-10 23:49:48'),(5,'','2019-10-10 23:37:16','2019-10-10 23:37:16','2019-10-10 23:49:54'),(6,'','2019-10-10 23:40:33','2019-10-10 23:40:33','2019-10-10 23:49:59'),(7,'','2019-10-10 23:44:42','2019-10-10 23:44:42','2019-10-10 23:50:03'),(8,'Ekonomi','2019-10-10 23:50:39','2019-10-12 15:08:52',NULL),(9,'Militer','2019-10-11 03:44:14','2019-10-11 03:44:14',NULL),(10,'Pengetahuan Umum','2019-10-11 03:47:03','2019-10-11 03:47:03',NULL),(11,'Sejarah','2019-10-11 03:49:25','2019-10-11 03:49:25',NULL),(12,'referensi','2019-10-11 03:50:14','2019-10-11 03:50:14','2019-10-12 03:04:27'),(13,'Coli','2019-10-12 14:09:13','2019-10-12 14:09:13','2019-10-12 14:09:15'),(14,'absurd','2019-10-12 15:09:24','2019-10-12 15:09:24','2019-10-12 15:09:26');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `guestbooks`
--

DROP TABLE IF EXISTS `guestbooks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `guestbooks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(125) NOT NULL,
  `nis` int(10) unsigned NOT NULL,
  `purpose` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `guestbooks`
--

LOCK TABLES `guestbooks` WRITE;
/*!40000 ALTER TABLE `guestbooks` DISABLE KEYS */;
INSERT INTO `guestbooks` VALUES (1,'Someone',11300,'nyantai aja dulu','2019-10-11 06:09:02','2019-10-11 06:09:02',NULL),(2,'Edy',12893,'Santuy','2019-10-12 14:08:45','2019-10-12 14:08:45',NULL);
/*!40000 ALTER TABLE `guestbooks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `role` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin@dangdut.com','qwerqwer','qwerqwer','qwerqwer','operator','2019-10-10 04:42:36','2019-10-10 04:42:36',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-10-12 22:43:39
