CREATE TABLE `log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `request_body` text,
  `response_body` text,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci