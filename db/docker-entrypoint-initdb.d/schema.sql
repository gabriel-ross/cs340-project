.
-- Reset db

DROP TABLE IF EXISTS `Pokemon`;
DROP TABLE IF EXISTS `Types`;
DROP TABLE IF EXISTS `Moves`;
DROP TABLE IF EXISTS `Generations`;
DROP TABLE IF EXISTS `Pokemon_Moves`;

-- Define schema

CREATE TABLE `Generations` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `Types` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);


CREATE TABLE `Pokemon` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `primary_type` int(11) NOT NULL,
    `secondary_type` int(11),
    `generation` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`primary_type`) REFERENCES `Types` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`secondary_type`) REFERENCES `Types` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
); 

CREATE TABLE `Moves` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `type` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`type`) REFERENCES `Types` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `Pokemon_Moves` (
    `pokemon_id` int(11) NOT NULL,
    `move_id` int(11) NOT NULL,
    PRIMARY KEY (`pokemon_id`, `move_id`),
    FOREIGN KEY (`pokemon_id`) REFERENCES `Pokemon` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`move_id`) REFERENCES `Moves` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Insert preliminary data

INSERT INTO `Types` (`name`) VALUES
('Normal'),
('Fire'),
('Water'),
('Grass'),
('Electric'),
('Ice'),
('Fighting'),
('Poison'),
('Ground'),
('Flying'),
('Psychic'),
('Bug'),
('Rock'),
('Ghost'),
('Dark'),
('Dragon'),
('Steel'),
('Fairy');

INSERT INTO `Generations` (`name`) VALUES
('I'),
('II'),
('III'),
('IV'),
('V'),
('VI'),
('VII'),
('VIII');