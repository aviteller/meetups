-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 15, 2019 at 06:57 PM
-- Server version: 10.3.16-MariaDB
-- PHP Version: 7.3.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `meetups`
--

-- --------------------------------------------------------

--
-- Table structure for table `meetups`
--

CREATE TABLE `meetups` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `sub_title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `contact_email` varchar(255) DEFAULT NULL,
  `is_liked` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `meetups`
--

INSERT INTO `meetups` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `sub_title`, `description`, `image_url`, `address`, `contact_email`, `is_liked`) VALUES
(3, NULL, NULL, NULL, 'Coding bootcamp', 'learn to code quickly', 'talking about coding', 'https://iceclog.com/wp-content/uploads/2016/09/596px-Internet1.jpg', '123 fake street', 'fake@1.com', 1),
(4, NULL, NULL, NULL, 'Swimming', 'Swimming', 'lots of swimming', 'https://www.bhf.org.uk/-/media/images/heart-matters/winter-2016/activity/swimmingpool-620x400-noexp.jpg?la=en', '123 fake street', 'fake@3.com', 1),
(12, '2019-08-15 10:19:40', '2019-08-15 10:19:40', NULL, 'screenshot', 'screenshot', 'screenshot', 'https://user-media-prod-cdn.itsre-sumo.mozilla.net/uploads/gallery/images/2019-04-22-20-55-35-f6d128.png', 'screenshot', 'screenshot@screenshot.com', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `meetups`
--
ALTER TABLE `meetups`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_meetups_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `meetups`
--
ALTER TABLE `meetups`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
