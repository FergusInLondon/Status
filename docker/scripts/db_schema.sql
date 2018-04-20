USE `status`;

CREATE TABLE `checks` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`domain` VARCHAR(255) NOT NULL UNIQUE,
	`last_performed` DATETIME NOT NULL,
	`is_up` BOOLEAN NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `incidents` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`check_id` INT NOT NULL,
	`description` TEXT,
	`down_detection` DATETIME NOT NULL,
	`up_detection` DATETIME,
	PRIMARY KEY (`id`)
);

ALTER TABLE `incidents` ADD CONSTRAINT `incidents_fk0` FOREIGN KEY (`check_id`) REFERENCES `checks`(`id`);

INSERT INTO `checks` (`domain`, `last_performed`, `is_up`) VALUES ('fergus.london', NOW(), '1');
INSERT INTO `checks` (`domain`, `last_performed`, `is_up`) VALUES ('musings.fergus.london', NOW(), '0');