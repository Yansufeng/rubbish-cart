CREATE TABLE `rubbish`
(
    `id` int NOT NULL,
    `cartId` int NOT NULL,
    `rubbishType` int NOT NULL,
    `rubbishAmount` int NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;