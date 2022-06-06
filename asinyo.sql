-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 06, 2022 at 05:44 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 7.4.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `asinyo`
--

-- --------------------------------------------------------

--
-- Table structure for table `addresses`
--

CREATE TABLE `addresses` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `phone` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `digital_address` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `city` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `region` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `address` longtext COLLATE utf8mb4_bin NOT NULL,
  `other_information` longtext COLLATE utf8mb4_bin DEFAULT NULL,
  `default` tinyint(1) NOT NULL DEFAULT 0,
  `agent_addresses` bigint(20) DEFAULT NULL,
  `customer_addresses` bigint(20) DEFAULT NULL,
  `merchant_addresses` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password` blob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `agents`
--

CREATE TABLE `agents` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password` blob NOT NULL,
  `ghana_card` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `phone` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `digital_address` varchar(255) COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `baskets`
--

CREATE TABLE `baskets` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `agent_baskets` bigint(20) DEFAULT NULL,
  `customer_baskets` bigint(20) DEFAULT NULL,
  `merchant_baskets` bigint(20) DEFAULT NULL,
  `product_baskets` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password` blob NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `phone` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `ent_types`
--

CREATE TABLE `ent_types` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `type` varchar(255) COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `ent_types`
--

INSERT INTO `ent_types` (`id`, `type`) VALUES
(1, 'addresses'),
(2, 'admins'),
(3, 'agents'),
(4, 'baskets'),
(5, 'customers'),
(6, 'favourites'),
(14, 'merchant_stores'),
(7, 'merchants'),
(8, 'orders'),
(10, 'product_category_majors'),
(11, 'product_category_minors'),
(9, 'products'),
(12, 'retail_merchants'),
(13, 'supplier_merchants');

-- --------------------------------------------------------

--
-- Table structure for table `favourites`
--

CREATE TABLE `favourites` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `agent_favourites` bigint(20) DEFAULT NULL,
  `customer_favourites` bigint(20) DEFAULT NULL,
  `merchant_favourites` bigint(20) DEFAULT NULL,
  `product_favourites` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `merchants`
--

CREATE TABLE `merchants` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `username` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `password` blob NOT NULL,
  `type` varchar(255) COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `merchants`
--

INSERT INTO `merchants` (`id`, `created_at`, `updated_at`, `username`, `password`, `type`) VALUES
(25769803776, '2022-05-28 11:44:14', '2022-05-28 11:44:14', 'supplier@mail.com', 0x243261243136245867487232722e4953634b344562416961774f6e484f4a7537326e6a655749543853596967757a4e3949582f4e415062426e6a7071, 'supplier'),
(25769803777, '2022-05-29 11:01:24', '2022-05-29 11:01:24', 'retailer@mail.com', 0x24326124313624624a373752636b4a6f31614468737a54666e416a774f4531474c466b6e694e684c5959325672734f486673782f4c43544f7142614f, 'retailer');

-- --------------------------------------------------------

--
-- Table structure for table `merchant_stores`
--

CREATE TABLE `merchant_stores` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `about` longtext COLLATE utf8mb4_bin NOT NULL,
  `desc_title` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `description` longtext COLLATE utf8mb4_bin NOT NULL,
  `logo` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `images` longtext COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`images`)),
  `bank_account` longtext COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`bank_account`)),
  `momo_account` longtext COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`momo_account`)),
  `default_account` enum('bank','momo') COLLATE utf8mb4_bin DEFAULT NULL,
  `merchant_store` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `merchant_stores`
--

INSERT INTO `merchant_stores` (`id`, `created_at`, `updated_at`, `name`, `about`, `desc_title`, `description`, `logo`, `images`, `bank_account`, `momo_account`, `default_account`, `merchant_store`) VALUES
(55834574848, '2022-06-05 10:44:59', '2022-06-06 00:18:03', 'Seydel Farms', 'About Seydel', 'The Story', 'Description', 'http://127.0.0.1:8000/merchant/stores/25769803777/asinyo_13073c8d-c0c2-4ab2-b115-d40941aceb77.jpg', '[\"http://127.0.0.1:8000/merchant/stores/25769803777/asinyo_42a3469e-f80b-4db0-954f-ae5b72f9d4f4.jpg\",\"http://127.0.0.1:8000/merchant/stores/25769803777/asinyo_90f731fd-0858-49ec-a0e8-83cc96eae926.jpg\",\"http://127.0.0.1:8000/merchant/stores/25769803777/asinyo_0b976404-8778-4167-9de6-e9f851fb0414.png\"]', '{\"name\":\"Godsway Seyram Prikah\",\"number\":\"4654564546465\",\"bank\":\"ecobank\",\"branch\":\"Madina\"}', '{\"name\":\"Godsway Seyram Prikah\",\"number\":\"+233599827815\",\"provider\":\"MTN\"}', 'momo', 25769803777);

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `status` enum('in_progress','shipping','delivered') COLLATE utf8mb4_bin NOT NULL DEFAULT 'in_progress',
  `delivered_at` timestamp NULL DEFAULT NULL,
  `address_orders` bigint(20) DEFAULT NULL,
  `agent_orders` bigint(20) DEFAULT NULL,
  `customer_orders` bigint(20) DEFAULT NULL,
  `merchant_orders` bigint(20) DEFAULT NULL,
  `product_orders` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `price` double NOT NULL DEFAULT 0,
  `promo_price` double DEFAULT NULL,
  `quantity` int(10) UNSIGNED NOT NULL DEFAULT 1,
  `unit` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `description` longtext COLLATE utf8mb4_bin NOT NULL,
  `image` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `merchant_products` bigint(20) NOT NULL,
  `product_category_major_products` bigint(20) NOT NULL,
  `product_category_minor_products` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `created_at`, `updated_at`, `name`, `price`, `promo_price`, `quantity`, `unit`, `description`, `image`, `merchant_products`, `product_category_major_products`, `product_category_minor_products`) VALUES
(34359738368, '2022-05-28 14:26:56', '2022-05-28 14:26:56', 'Okra', 50, 0, 1, 'cup', 'Okra description', 'http://127.0.0.1:8000/products/asinyo_a7d8e6c8-324d-4a7e-813b-cc02f9610a16.jpg', 25769803776, 38654705664, 42949672960),
(34359738369, '2022-05-28 14:29:30', '2022-05-28 14:29:30', 'Tomato', 100, 0, 1, 'g', 'Tomato description', 'http://127.0.0.1:8000/products/asinyo_90e3cf2e-f83d-4885-86c5-754de9a0bd73.jpg', 25769803776, 38654705664, 42949672960),
(34359738370, '2022-05-28 14:31:38', '2022-05-28 14:31:38', 'Maize', 300, 298, 10, 'kg', 'Maize description', 'http://127.0.0.1:8000/products/asinyo_b5f50627-b2d9-477c-ae7a-99a72916019b.jpg', 25769803776, 38654705664, 42949672961),
(34359738371, '2022-05-28 14:34:58', '2022-05-28 14:34:58', 'NPK Fertilizer', 45, 0, 1, 'g', 'Fertilizer description', 'http://127.0.0.1:8000/products/asinyo_d921094f-67ea-454d-ab40-11d218e2d7f5.jpg', 25769803776, 38654705665, 42949672962),
(34359738372, '2022-05-29 08:57:05', '2022-05-29 08:57:05', 'Beans', 20, 0, 1, 'cup', 'Beans Description', 'http://127.0.0.1:8000/products/asinyo_c3c8aee2-a367-4c03-9338-f4db42b1aabb.jpg', 25769803776, 38654705664, 42949672961),
(34359738373, '2022-05-29 09:29:36', '2022-05-29 09:29:36', 'Weed Killer', 100, 0, 1, 'kg', ' Weed Killer description', 'http://127.0.0.1:8000/products/asinyo_a1314261-16ce-479d-90a8-ed8cbdac83db.jpg', 25769803776, 38654705665, 42949672963),
(34359738374, '2022-05-29 10:31:31', '2022-05-29 10:31:31', 'Garden egg', 20, 0, 2, 'kg', 'Garden egg description', 'http://127.0.0.1:8000/products/asinyo_1987d0a4-d102-4ea8-84b8-fcc0b8cfa258.jpg', 25769803776, 38654705664, 42949672960),
(34359738375, '2022-05-29 10:31:45', '2022-05-29 10:31:45', 'Garden egg', 20, 0, 2, 'kg', 'Garden egg description', 'http://127.0.0.1:8000/products/asinyo_bcc9498e-b88a-4f3c-b580-d6b39b15482a.jpg', 25769803776, 38654705664, 42949672960),
(34359738376, '2022-05-29 10:32:01', '2022-05-29 10:32:01', 'Garden egg', 20, 0, 2, 'kg', 'Garden egg description', 'http://127.0.0.1:8000/products/asinyo_00c3a8ef-f751-40fb-9b6b-07f25601f73f.jpg', 25769803776, 38654705664, 42949672960),
(34359738377, '2022-05-29 10:42:28', '2022-05-29 10:42:28', 'Garden egg', 150, 0, 3, 'ton', 'Garden egg description', 'http://127.0.0.1:8000/products/asinyo_321e3c1e-20c0-43db-bb5b-4b608e3e6e85.jpg', 25769803776, 38654705665, 42949672962),
(34359738378, '2022-05-29 10:46:36', '2022-05-29 10:46:36', 'Garden egg', 99, 0, 6, 'ton', 'Garden egg description', 'http://127.0.0.1:8000/products/asinyo_bcc83476-7c0f-443b-a801-81d2f1787f54.jpg', 25769803776, 38654705664, 42949672960),
(34359738379, '2022-05-29 10:51:01', '2022-05-29 10:51:01', 'Garden egg', 97, 0, 6, 'ton', 'Garden egg description', 'http://127.0.0.1:8000/products/asinyo_4bab818d-33a2-46b4-a983-c3257cce26d5.jpg', 25769803776, 38654705664, 42949672960),
(34359738380, '2022-05-29 12:09:29', '2022-05-29 12:09:29', 'Beans', 20, 0, 2, 'g', 'Beans Description', 'http://127.0.0.1:8000/products/asinyo_cb04f352-39a8-4af3-82d9-9b11529def20.jpg', 25769803777, 38654705664, 42949672961),
(34359738381, '2022-05-29 12:10:42', '2022-05-29 12:10:42', 'Beans', 20, 0, 2, 'g', 'Beans Description', 'http://127.0.0.1:8000/products/asinyo_a88838a4-0bd0-4ee5-9eb3-87815d47c32a.jpg', 25769803777, 38654705664, 42949672961),
(34359738382, '2022-05-29 12:15:32', '2022-05-29 12:15:32', 'Beans', 20, 0, 2, 'g', 'Beans Description', 'http://127.0.0.1:8000/products/asinyo_65ed402b-7234-4a56-9240-af6bf49bf6b9.jpg', 25769803777, 38654705664, 42949672961),
(34359738383, '2022-05-29 12:17:00', '2022-05-29 12:17:00', 'Maize', 500, 499, 100, 'kg', 'Maize description', 'http://127.0.0.1:8000/products/asinyo_edf26180-6088-4130-a9c3-93916e0158c6.jpg', 25769803777, 38654705664, 42949672961),
(34359738384, '2022-05-29 12:18:59', '2022-05-29 12:18:59', 'Weed Killer', 50, 0, 1, 'g', 'Description', 'http://127.0.0.1:8000/products/asinyo_b4850b2f-f612-4825-994e-f6c42717feb8.jpg', 25769803777, 38654705665, 42949672963),
(34359738385, '2022-05-29 14:48:20', '2022-05-29 14:48:20', 'Tomato', 200, 0, 12, 'kg', 'Description', 'http://127.0.0.1:8000/products/asinyo_106861f3-3ff7-448f-aaf4-45530222b7e0.jpg', 25769803777, 38654705664, 42949672960),
(34359738386, '2022-06-02 12:23:27', '2022-06-02 12:23:27', 'NPK', 50, 0, 1, 'kg', 'Fertilizer Description', 'http://127.0.0.1:8000/products/asinyo_c2c1f5ff-d3a4-41f7-8beb-76a326d716a8.jpg', 25769803776, 38654705665, 42949672962),
(34359738387, '2022-06-02 13:56:53', '2022-06-02 13:56:53', 'NPK', 20, 0, 10, 'kg', 'Fertilizer Description', 'http://127.0.0.1:8000/products/asinyo_60a1792b-0dd8-4487-9e2f-354400c9f221.jpg', 25769803777, 38654705665, 42949672962);

-- --------------------------------------------------------

--
-- Table structure for table `product_category_majors`
--

CREATE TABLE `product_category_majors` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `category` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `sulg` varchar(255) COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `product_category_majors`
--

INSERT INTO `product_category_majors` (`id`, `created_at`, `updated_at`, `category`, `sulg`) VALUES
(38654705664, '2022-05-28 11:41:27', '2022-05-28 11:41:27', 'Food Items', 'food-items'),
(38654705665, '2022-05-28 11:41:35', '2022-05-28 11:41:35', 'Farm Supply', 'farm-supply');

-- --------------------------------------------------------

--
-- Table structure for table `product_category_minors`
--

CREATE TABLE `product_category_minors` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `category` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `image` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `sulg` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `product_category_major_minors` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `product_category_minors`
--

INSERT INTO `product_category_minors` (`id`, `created_at`, `updated_at`, `category`, `image`, `sulg`, `product_category_major_minors`) VALUES
(42949672960, '2022-05-28 11:41:50', '2022-05-28 11:41:50', 'Vegetables', 'http://127.0.0.1:8000/products/asinyo_bcc83476-7c0f-443b-a801-81d2f1787f54.jpg', 'vegetables', 38654705664),
(42949672961, '2022-05-28 11:42:06', '2022-05-28 11:42:06', 'Cereals', 'http://127.0.0.1:8000/products/asinyo_b5f50627-b2d9-477c-ae7a-99a72916019b.jpg', 'cereals', 38654705664),
(42949672962, '2022-05-28 11:42:18', '2022-05-28 11:42:18', 'Fertilizer', 'http://127.0.0.1:8000/products/asinyo_d921094f-67ea-454d-ab40-11d218e2d7f5.jpg', 'fertilizer', 38654705665),
(42949672963, '2022-05-28 11:42:27', '2022-05-28 11:42:27', 'Pesticide', 'http://127.0.0.1:8000/products/asinyo_a1314261-16ce-479d-90a8-ed8cbdac83db.jpg', 'pesticide', 38654705665);

-- --------------------------------------------------------

--
-- Table structure for table `retail_merchants`
--

CREATE TABLE `retail_merchants` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `ghana_card` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `phone` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `digital_address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `merchant_retailer` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `retail_merchants`
--

INSERT INTO `retail_merchants` (`id`, `created_at`, `updated_at`, `ghana_card`, `last_name`, `other_name`, `phone`, `other_phone`, `address`, `digital_address`, `merchant_retailer`) VALUES
(47244640256, '2022-05-29 11:01:24', '2022-05-29 11:01:24', 'GHA-456645444-7', 'Atumfo', 'Samuel', '+233265518694', '', 'Adenta', 'GD-4568-7892', 25769803777);

-- --------------------------------------------------------

--
-- Table structure for table `supplier_merchants`
--

CREATE TABLE `supplier_merchants` (
  `id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `ghana_card` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `phone` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `other_phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `digital_address` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `merchant_supplier` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `supplier_merchants`
--

INSERT INTO `supplier_merchants` (`id`, `created_at`, `updated_at`, `ghana_card`, `last_name`, `other_name`, `phone`, `other_phone`, `address`, `digital_address`, `merchant_supplier`) VALUES
(51539607552, '2022-05-28 11:44:14', '2022-05-28 11:44:14', 'GHA-456645444-7', 'Gawuso', 'Dorcas Akua', '+233265518694', '', 'Adenta', 'DG-5455-4598', 25769803776);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `addresses`
--
ALTER TABLE `addresses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `addresses_agents_addresses` (`agent_addresses`),
  ADD KEY `addresses_customers_addresses` (`customer_addresses`),
  ADD KEY `addresses_merchants_addresses` (`merchant_addresses`);

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `agents`
--
ALTER TABLE `agents`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `ghana_card` (`ghana_card`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- Indexes for table `baskets`
--
ALTER TABLE `baskets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `baskets_agents_baskets` (`agent_baskets`),
  ADD KEY `baskets_customers_baskets` (`customer_baskets`),
  ADD KEY `baskets_merchants_baskets` (`merchant_baskets`),
  ADD KEY `baskets_products_baskets` (`product_baskets`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `ent_types`
--
ALTER TABLE `ent_types`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `type` (`type`);

--
-- Indexes for table `favourites`
--
ALTER TABLE `favourites`
  ADD PRIMARY KEY (`id`),
  ADD KEY `favourites_agents_favourites` (`agent_favourites`),
  ADD KEY `favourites_customers_favourites` (`customer_favourites`),
  ADD KEY `favourites_merchants_favourites` (`merchant_favourites`),
  ADD KEY `favourites_products_favourites` (`product_favourites`);

--
-- Indexes for table `merchants`
--
ALTER TABLE `merchants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `merchant_stores`
--
ALTER TABLE `merchant_stores`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `merchant_store` (`merchant_store`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `orders_addresses_orders` (`address_orders`),
  ADD KEY `orders_agents_orders` (`agent_orders`),
  ADD KEY `orders_customers_orders` (`customer_orders`),
  ADD KEY `orders_merchants_orders` (`merchant_orders`),
  ADD KEY `orders_products_orders` (`product_orders`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `products_merchants_products` (`merchant_products`),
  ADD KEY `products_product_category_majors_products` (`product_category_major_products`),
  ADD KEY `products_product_category_minors_products` (`product_category_minor_products`);

--
-- Indexes for table `product_category_majors`
--
ALTER TABLE `product_category_majors`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `category` (`category`);

--
-- Indexes for table `product_category_minors`
--
ALTER TABLE `product_category_minors`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `category` (`category`),
  ADD KEY `product_category_minors_product_category_majors_minors` (`product_category_major_minors`);

--
-- Indexes for table `retail_merchants`
--
ALTER TABLE `retail_merchants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `ghana_card` (`ghana_card`),
  ADD UNIQUE KEY `phone` (`phone`),
  ADD UNIQUE KEY `merchant_retailer` (`merchant_retailer`);

--
-- Indexes for table `supplier_merchants`
--
ALTER TABLE `supplier_merchants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `ghana_card` (`ghana_card`),
  ADD UNIQUE KEY `phone` (`phone`),
  ADD UNIQUE KEY `merchant_supplier` (`merchant_supplier`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `addresses`
--
ALTER TABLE `addresses`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4294967296;

--
-- AUTO_INCREMENT for table `agents`
--
ALTER TABLE `agents`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8589934592;

--
-- AUTO_INCREMENT for table `baskets`
--
ALTER TABLE `baskets`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12884901888;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17179869184;

--
-- AUTO_INCREMENT for table `ent_types`
--
ALTER TABLE `ent_types`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `favourites`
--
ALTER TABLE `favourites`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21474836480;

--
-- AUTO_INCREMENT for table `merchants`
--
ALTER TABLE `merchants`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25769803778;

--
-- AUTO_INCREMENT for table `merchant_stores`
--
ALTER TABLE `merchant_stores`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=55834574849;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30064771072;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34359738388;

--
-- AUTO_INCREMENT for table `product_category_majors`
--
ALTER TABLE `product_category_majors`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38654705666;

--
-- AUTO_INCREMENT for table `product_category_minors`
--
ALTER TABLE `product_category_minors`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=42949672964;

--
-- AUTO_INCREMENT for table `retail_merchants`
--
ALTER TABLE `retail_merchants`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47244640257;

--
-- AUTO_INCREMENT for table `supplier_merchants`
--
ALTER TABLE `supplier_merchants`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=51539607553;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `addresses`
--
ALTER TABLE `addresses`
  ADD CONSTRAINT `addresses_agents_addresses` FOREIGN KEY (`agent_addresses`) REFERENCES `agents` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `addresses_customers_addresses` FOREIGN KEY (`customer_addresses`) REFERENCES `customers` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `addresses_merchants_addresses` FOREIGN KEY (`merchant_addresses`) REFERENCES `merchants` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `baskets`
--
ALTER TABLE `baskets`
  ADD CONSTRAINT `baskets_agents_baskets` FOREIGN KEY (`agent_baskets`) REFERENCES `agents` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `baskets_customers_baskets` FOREIGN KEY (`customer_baskets`) REFERENCES `customers` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `baskets_merchants_baskets` FOREIGN KEY (`merchant_baskets`) REFERENCES `merchants` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `baskets_products_baskets` FOREIGN KEY (`product_baskets`) REFERENCES `products` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `favourites`
--
ALTER TABLE `favourites`
  ADD CONSTRAINT `favourites_agents_favourites` FOREIGN KEY (`agent_favourites`) REFERENCES `agents` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `favourites_customers_favourites` FOREIGN KEY (`customer_favourites`) REFERENCES `customers` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `favourites_merchants_favourites` FOREIGN KEY (`merchant_favourites`) REFERENCES `merchants` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `favourites_products_favourites` FOREIGN KEY (`product_favourites`) REFERENCES `products` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `merchant_stores`
--
ALTER TABLE `merchant_stores`
  ADD CONSTRAINT `merchant_stores_merchants_store` FOREIGN KEY (`merchant_store`) REFERENCES `merchants` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_addresses_orders` FOREIGN KEY (`address_orders`) REFERENCES `addresses` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `orders_agents_orders` FOREIGN KEY (`agent_orders`) REFERENCES `agents` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `orders_customers_orders` FOREIGN KEY (`customer_orders`) REFERENCES `customers` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `orders_merchants_orders` FOREIGN KEY (`merchant_orders`) REFERENCES `merchants` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `orders_products_orders` FOREIGN KEY (`product_orders`) REFERENCES `products` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_merchants_products` FOREIGN KEY (`merchant_products`) REFERENCES `merchants` (`id`) ON DELETE NO ACTION,
  ADD CONSTRAINT `products_product_category_majors_products` FOREIGN KEY (`product_category_major_products`) REFERENCES `product_category_majors` (`id`) ON DELETE NO ACTION,
  ADD CONSTRAINT `products_product_category_minors_products` FOREIGN KEY (`product_category_minor_products`) REFERENCES `product_category_minors` (`id`) ON DELETE NO ACTION;

--
-- Constraints for table `product_category_minors`
--
ALTER TABLE `product_category_minors`
  ADD CONSTRAINT `product_category_minors_product_category_majors_minors` FOREIGN KEY (`product_category_major_minors`) REFERENCES `product_category_majors` (`id`) ON DELETE NO ACTION;

--
-- Constraints for table `retail_merchants`
--
ALTER TABLE `retail_merchants`
  ADD CONSTRAINT `retail_merchants_merchants_retailer` FOREIGN KEY (`merchant_retailer`) REFERENCES `merchants` (`id`) ON DELETE NO ACTION;

--
-- Constraints for table `supplier_merchants`
--
ALTER TABLE `supplier_merchants`
  ADD CONSTRAINT `supplier_merchants_merchants_supplier` FOREIGN KEY (`merchant_supplier`) REFERENCES `merchants` (`id`) ON DELETE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
