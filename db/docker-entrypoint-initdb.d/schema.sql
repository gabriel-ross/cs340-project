DROP TABLE IF EXISTS `Pokemon`;
DROP TABLE IF EXISTS `Types`;
DROP TABLE IF EXISTS `Moves`;
DROP TABLE IF EXISTS `Generations`;
DROP TABLE IF EXISTS `Pokemon_Moves`;

CREATE TABLE `Pokemon` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `primary_type` int(11) NOT NULL,
    `secondary_type` int(11),
    `generation` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`primary_type`) REFERENCES `Types` (`id`),
    FOREIGN KEY (`secondary_type`) REFERENCES `Types` (`id`)
);

-- TODO: write cascad logic