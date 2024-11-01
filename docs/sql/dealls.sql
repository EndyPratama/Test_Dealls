--
-- Database: `test_dealls`
--
CREATE DATABASE IF NOT EXISTS `test_dealls` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `test_dealls`;

-- --------------------------------------------------------

--
-- Table structure for table `history_watch`
--

DROP TABLE IF EXISTS `history_watch`;
CREATE TABLE `history_watch` (
  `id` int(11) NOT NULL,
  `profile1_id` int(11) NOT NULL,
  `profile2_id` int(11) NOT NULL,
  `label` varchar(15) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `history_watch`
--

INSERT INTO `history_watch` (`id`, `profile1_id`, `profile2_id`, `label`, `created_at`, `updated_at`) VALUES
(1, 8, 19, '', '2024-11-01 06:46:41', NULL),
(2, 8, 25, '', '2024-11-01 06:46:46', NULL),
(3, 8, 16, '', '2024-11-01 06:46:48', NULL),
(4, 8, 32, '', '2024-11-01 06:46:49', NULL),
(5, 8, 32, '', '2024-11-01 06:47:08', NULL),
(6, 8, 26, '', '2024-11-01 06:47:09', NULL),
(7, 8, 23, '', '2024-11-01 06:47:11', NULL),
(8, 8, 24, '', '2024-11-01 06:47:14', NULL),
(9, 8, 17, '', '2024-11-01 06:47:16', NULL),
(10, 8, 21, '', '2024-11-01 06:47:19', NULL),
(11, 8, 25, '', '2024-11-01 06:47:22', NULL),
(12, 8, 31, '', '2024-11-01 06:48:31', NULL),
(13, 8, 30, '', '2024-11-01 07:04:35', NULL),
(14, 8, 12, '', '2024-11-01 07:05:13', NULL),
(15, 8, 31, '', '2024-11-01 07:44:03', NULL),
(16, 0, 0, 'Skip', '0000-00-00 00:00:00', NULL),
(17, 8, 30, 'Like', '2024-11-01 07:58:44', '2024-11-01 08:02:13');

-- --------------------------------------------------------

--
-- Table structure for table `likes`
--

DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `id` int(11) NOT NULL,
  `liker_id` int(11) NOT NULL,
  `liked_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `likes`
--

INSERT INTO `likes` (`id`, `liker_id`, `liked_id`, `created_at`, `updated_at`) VALUES
(1, 8, 9, '2024-11-01 02:40:19', NULL),
(3, 8, 10, '2024-11-01 02:45:17', NULL),
(12, 8, 13, '2024-11-01 05:30:09', NULL),
(13, 8, 30, '2024-11-01 07:59:01', NULL),
(14, 8, 30, '2024-11-01 08:02:13', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `matches`
--

DROP TABLE IF EXISTS `matches`;
CREATE TABLE `matches` (
  `id` int(11) NOT NULL,
  `profile1_id` int(11) NOT NULL,
  `profile2_id` int(11) NOT NULL,
  `matched_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `photo`
--

DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo` (
  `id` int(11) NOT NULL,
  `profile_id` int(11) NOT NULL,
  `photo_url` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `photo`
--

INSERT INTO `photo` (`id`, `profile_id`, `photo_url`, `created_at`, `updated_at`) VALUES
(1, 8, 'https://img.freepik.com/premium-photo/sad-anime-character-with-black-hoodie_900958-11969.jpg', '2024-10-31 21:34:57', '2024-10-31 21:35:50');

-- --------------------------------------------------------

--
-- Table structure for table `profile`
--

DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `gender` varchar(15) NOT NULL,
  `bio` text NOT NULL,
  `birthdate` date NOT NULL,
  `location` text NOT NULL,
  `subscription_id` int(11) NOT NULL DEFAULT 1,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `profile`
--

INSERT INTO `profile` (`id`, `user_id`, `name`, `gender`, `bio`, `birthdate`, `location`, `subscription_id`, `created_at`, `updated_at`) VALUES
(1, 1, 'male 1', 'male', 'happy', '2020-11-01', '-7.7526139,112.1787896', 1, '2024-10-31 20:32:22', NULL),
(6, 2, 'male 2', 'male', 'happy holiday', '2020-11-01', '-7.7526139,112.1787896', 1, '2024-10-31 20:48:29', '2024-10-31 20:50:54'),
(8, 4, 'male 3', 'male', 'happy holiday', '2020-11-01', '-7.7526139,112.1787896', 2, '2024-10-31 21:08:29', '2024-11-01 08:03:05'),
(9, 5, 'female 1', 'female', 'she', '2020-11-01', '-303231212,13214224', 1, '2024-11-01 08:08:59', NULL),
(10, 6, 'female 2', 'female', 'she', '2001-10-03', '-7.2381028414,112.5235208230', 1, '2024-11-01 08:08:59', NULL),
(11, 7, 'female 3', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(12, 8, 'female 4', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(13, 9, 'female 5', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(14, 10, 'female 6', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(15, 11, 'female 7', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(16, 12, 'female 8', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(17, 13, 'female 9', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(18, 14, 'female 10', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(19, 15, 'female 11', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(20, 16, 'female 12', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(21, 17, 'female 13', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(22, 18, 'female 14', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(23, 19, 'female 15', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(24, 20, 'female 16', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(25, 21, 'female 17', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(26, 22, 'female 18', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(27, 23, 'female 19', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(28, 24, 'female 20', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(29, 25, 'female 21', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(30, 26, 'female 22', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(31, 27, 'female 23', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(32, 28, 'female 24', 'female', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(33, 29, 'male 4', 'male', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(34, 30, 'male 5', 'male', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(35, 31, 'male 6', 'male', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(36, 32, 'male 7', 'male', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(37, 33, 'male 8', 'male', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL),
(38, 34, 'male 9', 'male', 'oke', '2004-11-03', '', 1, '2024-11-01 11:01:46', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `subscription`
--

DROP TABLE IF EXISTS `subscription`;
CREATE TABLE `subscription` (
  `id` int(11) NOT NULL,
  `name` varchar(25) NOT NULL,
  `value` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `subscription`
--

INSERT INTO `subscription` (`id`, `name`, `value`, `created_at`, `updated_at`) VALUES
(1, 'free', 10, NULL, NULL),
(2, 'premium', 999, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'male-1@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 00:43:15', NULL),
(3, 'male-2@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-10-31 18:26:15', NULL),
(4, 'male-3@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-10-31 21:04:23', NULL),
(5, 'female-1@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 08:07:51', NULL),
(6, 'female-2@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 08:07:51', NULL),
(7, 'female-3@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(8, 'female-4@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(9, 'female-5@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(10, 'female-6@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(11, 'female-7@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(12, 'female-8@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(13, 'female-9@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(14, 'female-10@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(15, 'female-11@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(16, 'female-12@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(17, 'female-13@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(18, 'female-14@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(19, 'female-15@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(20, 'female-16@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(21, 'female-17@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(22, 'female-18@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(23, 'female-19@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(24, 'female-20@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(25, 'female-21@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(26, 'female-22@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(27, 'female-23@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(28, 'female-24@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(29, 'male-4@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(30, 'male-5@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(31, 'male-6@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(32, 'male-7@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(33, 'male-8@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL),
(34, 'male-9@gmail.com', 'a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76', '2024-11-01 10:56:56', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `history_watch`
--
ALTER TABLE `history_watch`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `likes`
--
ALTER TABLE `likes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `matches`
--
ALTER TABLE `matches`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `photo`
--
ALTER TABLE `photo`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `profile`
--
ALTER TABLE `profile`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Indexes for table `subscription`
--
ALTER TABLE `subscription`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `history_watch`
--
ALTER TABLE `history_watch`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `likes`
--
ALTER TABLE `likes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `matches`
--
ALTER TABLE `matches`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `photo`
--
ALTER TABLE `photo`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `profile`
--
ALTER TABLE `profile`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=39;

--
-- AUTO_INCREMENT for table `subscription`
--
ALTER TABLE `subscription`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=35;
COMMIT;
