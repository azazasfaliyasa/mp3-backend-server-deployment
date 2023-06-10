BEGIN;

CREATE DATABASE `mini_project_3`;
USE `mini_project_3`;

CREATE TABLE `roles` (
  `id` int unsigned PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(20) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL
);

CREATE TABLE `permissions` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(20) UNIQUE NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL
);

CREATE TABLE `role_permission` (
  `role_id` int unsigned,
  `permission_id` bigint unsigned
);

CREATE TABLE `users` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255) NOT NULL,
  `username` varchar(100) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL
);

CREATE TABLE `customers` (
  `id` bigint unsigned PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `avatar` varchar(255),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL
);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
ALTER TABLE `role_permission` ADD FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`);
ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);


LOCK TABLES `roles` WRITE;
INSERT INTO `roles` (`id`, `name`) VALUES (1,'super-admin'), (2,'admin');
UNLOCK TABLES;

COMMIT;