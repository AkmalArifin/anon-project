CREATE TABLE `users` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE,
  `email` varchar(255) UNIQUE,
  `password` varchar(255),
  `photo_profile` varchar(255),
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `log` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `message` varchar(255),
  `is_starred` integer,
  `created_at` datetime
);

CREATE TABLE `password_resets` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `token` varchar(255),
  `expires_at` datetime,
  `created_at` datetime
);

ALTER TABLE `log` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `password_resets` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
