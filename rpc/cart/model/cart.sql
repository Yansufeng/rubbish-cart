CREATE TABLE `cart`
(
    `cartId` int NOT NULL,
    `cartName` varchar(255) NOT NULL,
    `state` varchar(255) NOT NULL,
    PRIMARY KEY(`cartId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;