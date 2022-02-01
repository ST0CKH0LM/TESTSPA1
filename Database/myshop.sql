-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 01, 2022 at 06:05 PM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 7.3.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `myshop`
--

-- --------------------------------------------------------

--
-- Table structure for table `db_order`
--

CREATE TABLE `db_order` (
  `o_id` int(11) NOT NULL,
  `o_user` varchar(255) NOT NULL,
  `s_day` int(11) NOT NULL,
  `o_time` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `db_order`
--

INSERT INTO `db_order` (`o_id`, `o_user`, `s_day`, `o_time`) VALUES
(1, 'test2', 3, '2022-02-01 15:54:44'),
(2, 'test2', 3, '2022-02-01 15:55:25'),
(3, 'test2', 3, '2022-02-01 15:58:33'),
(4, 'test2', 3, '2022-02-01 15:59:03'),
(5, 'test2', 9, '2022-02-01 15:59:08'),
(6, 'test2', 2, '2022-02-01 16:02:37'),
(7, 'test2', 2, '2022-02-01 16:03:13'),
(8, 'test2', 4, '2022-02-01 16:03:25'),
(9, 'test2', 3, '2022-02-01 16:03:56'),
(10, 'test2', 6, '2022-02-01 16:04:00'),
(11, 'test', 5, '2022-02-01 17:01:19');

-- --------------------------------------------------------

--
-- Table structure for table `db_spa`
--

CREATE TABLE `db_spa` (
  `s_id` int(11) NOT NULL,
  `s_day` int(11) NOT NULL,
  `s_qty` int(11) NOT NULL,
  `s_max` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `db_spa`
--

INSERT INTO `db_spa` (`s_id`, `s_day`, `s_qty`, `s_max`) VALUES
(1, 2, 2, 9),
(2, 1, 0, 3),
(3, 10, 0, 7),
(4, 8, 0, 6),
(5, 4, 1, 4),
(6, 5, 1, 3),
(7, 3, 2, 2),
(8, 6, 1, 3),
(9, 9, 1, 2),
(10, 7, 0, 5);

-- --------------------------------------------------------

--
-- Table structure for table `db_user`
--

CREATE TABLE `db_user` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `u_status` int(11) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `db_user`
--

INSERT INTO `db_user` (`id`, `username`, `password`, `u_status`) VALUES
(1, 'test', 'test', 1),
(3, 'test2', 'test2', 1),
(6, 'admin', 'admin', 2),
(26, 'test3', 'test3', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `db_order`
--
ALTER TABLE `db_order`
  ADD PRIMARY KEY (`o_id`);

--
-- Indexes for table `db_spa`
--
ALTER TABLE `db_spa`
  ADD PRIMARY KEY (`s_id`);

--
-- Indexes for table `db_user`
--
ALTER TABLE `db_user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `db_order`
--
ALTER TABLE `db_order`
  MODIFY `o_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `db_spa`
--
ALTER TABLE `db_spa`
  MODIFY `s_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `db_user`
--
ALTER TABLE `db_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
