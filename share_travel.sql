-- MySQL dump 10.13  Distrib 8.0.25, for macos11.3 (x86_64)
--
-- Host: localhost    Database: share_travel
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `event`
--

DROP TABLE IF EXISTS `event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `event` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `auth_key` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `date` datetime NOT NULL,
  `pool` int NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key` (`auth_key`)
) ENGINE=InnoDB AUTO_INCREMENT=90 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event`
--

LOCK TABLES `event` WRITE;
/*!40000 ALTER TABLE `event` DISABLE KEYS */;
INSERT INTO `event` VALUES (71,'0fQasfadsfas','テスト','2021-08-17 00:00:00',0,'2021-08-26 23:11:25',NULL),(72,'mWamTb7T0dP23eHB','テスト','2021-08-31 00:00:00',0,'2021-08-26 23:42:13',NULL),(73,'0GuSqS9rMQdaRfwx','テスト','2021-08-31 00:00:00',0,'2021-08-28 12:21:59',NULL),(74,'wwiG5JYbCBfUAK4v','テスト','2021-08-31 00:00:00',0,'2021-08-28 12:50:18',NULL),(75,'4XXv4b7lAy9A5u6u','テスト','2021-08-31 00:00:00',0,'2021-08-28 12:50:58',NULL),(76,'Drsfoxq4rlBIeagK','テスト','2021-08-31 00:00:00',0,'2021-08-28 12:54:16',NULL),(77,'76C2N1WigfEV0AUw','テスト','2021-08-31 00:00:00',0,'2021-08-28 12:56:17',NULL),(78,'EA1x20dIWI9nuGFk','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:28:11',NULL),(79,'jI1Yx3DgWYFE5JIJ','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:31:36',NULL),(80,'yTFCLS9JzYfmKeNa','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:33:39',NULL),(81,'mQ7ghhm6xxTdAlGp','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:34:58',NULL),(82,'BDstwFMcbGhjbQlk','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:35:33',NULL),(83,'WFW0xVw9oAdAgqls','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:38:39',NULL),(84,'ir3aJBuxgWSTtnSd','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:39:54',NULL),(85,'XHc4jMSBBw9FD1qA','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:41:05',NULL),(86,'4JD0ekgJ5fwyAO7G','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:44:12',NULL),(87,'nghNtVo5KDx1LVOY','テスト','2021-08-31 00:00:00',0,'2021-08-28 15:45:41',NULL),(88,'1j6U4sfjneQt2Mk1','テスト','2021-08-31 00:00:00',0,'2021-08-28 16:32:42',NULL),(89,'Ak5HGmvqMPKXMHD3','テスト','2021-08-30 00:00:00',447,'2021-08-29 17:05:54','2021-08-29 16:53:40');
/*!40000 ALTER TABLE `event` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `expense`
--

DROP TABLE IF EXISTS `expense`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `expense` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `event_id` int NOT NULL,
  `temporarily_member` int NOT NULL,
  `total` int NOT NULL,
  `remarks` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `name` varchar(64) NOT NULL DEFAULT '',
  `pool` int DEFAULT '0',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `expense`
--

LOCK TABLES `expense` WRITE;
/*!40000 ALTER TABLE `expense` DISABLE KEYS */;
INSERT INTO `expense` VALUES (58,89,34,435425,'テスト','テスト',25,'2021-08-30 01:53:14',NULL),(59,89,35,435425,'特になし','テスト2',25,'2021-08-30 01:53:39',NULL),(60,89,35,7800,'','交通費',0,'2021-08-30 01:54:30','2021-08-29 18:46:42');
/*!40000 ALTER TABLE `expense` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member`
--

DROP TABLE IF EXISTS `member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `member` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `event_id` int NOT NULL,
  `name` varchar(32) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member`
--

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;
INSERT INTO `member` VALUES (26,87,'テスト','2021-08-28 16:32:13',NULL),(27,88,'テスト','2021-08-28 16:35:50',NULL),(28,88,'テスト','2021-08-28 16:36:42',NULL),(29,88,'テスト','2021-08-28 16:36:56',NULL),(30,88,'テスト','2021-08-28 16:37:15',NULL),(31,88,'テスト','2021-08-28 16:38:04',NULL),(32,88,'テスト2','2021-08-28 16:38:13',NULL),(33,88,'テスト','2021-08-29 12:25:38',NULL),(34,89,'テスト','2021-08-29 17:06:59',NULL),(35,89,'テスト','2021-08-29 17:12:09',NULL);
/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `member_expense`
--

DROP TABLE IF EXISTS `member_expense`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `member_expense` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `member_id` int NOT NULL,
  `expense_id` int NOT NULL,
  `price` int NOT NULL,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `member_expense`
--

LOCK TABLES `member_expense` WRITE;
/*!40000 ALTER TABLE `member_expense` DISABLE KEYS */;
INSERT INTO `member_expense` VALUES (122,34,58,217700,'2021-08-30 01:53:14',NULL),(123,35,58,217700,'2021-08-30 01:53:14',NULL),(124,34,59,50000,'2021-08-30 01:53:39',NULL),(125,35,59,385400,'2021-08-30 01:53:39',NULL),(126,34,60,2800,'2021-08-30 01:54:30','2021-08-29 18:46:42'),(127,35,60,5000,'2021-08-30 01:54:30','2021-08-29 18:46:42');
/*!40000 ALTER TABLE `member_expense` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-08-30  3:58:07
