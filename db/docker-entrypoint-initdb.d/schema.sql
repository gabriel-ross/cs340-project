-- Reset db

DROP TABLE IF EXISTS `Pokemon`;
DROP TABLE IF EXISTS `Types`;
DROP TABLE IF EXISTS `Moves`;
DROP TABLE IF EXISTS `Generations`;
DROP TABLE IF EXISTS `Pokemon_Moves`;

-- Define schema
-- TODO: write FK and cascade logic

CREATE TABLE `Pokemon` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `primary_type` int(11) NOT NULL,
    `secondary_type` int(11),
    `generation` int(11) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `Types` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);

CREATE TABLE `Moves` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `type` int(11) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `Pokemon` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `primary_type` int(11) NOT NULL,
    `secondary_type` int(11),
    `generation` int(11) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `Pokemon` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `primary_type` int(11) NOT NULL,
    `secondary_type` int(11),
    `generation` int(11) NOT NULL,
    PRIMARY KEY (`id`)
);

-- Insert preliminary data
