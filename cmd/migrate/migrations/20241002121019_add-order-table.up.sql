-- Finally, create order_items table
CREATE TABLE IF NOT EXISTS order_items (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `orderId` INT UNSIGNED NOT NULL,
  `productId` INT UNSIGNED NOT NULL,
  `quantity` INT NOT NULL,
  `price` DECIMAL(10, 2) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`orderId`) REFERENCES orders(`id`),
  FOREIGN KEY (`productId`) REFERENCES products(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;