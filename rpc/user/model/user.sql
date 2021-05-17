CREATE TABLE `user`
(
    `openId` varchar(255) NOT NULL,
    `userName` varchar(255) NOT NULL,
    `purview` int NOT NULL,
    PRIMARY KEY(`openId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;