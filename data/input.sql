CREATE TABLE `USER`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `ip_address` INT  NOT NULL DEFAULT 0 COMMENT 'ip_address',
    `nickname`    VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user note',
    `description` VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'user description',
    `creator_email` VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'creator email',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT 'delete time',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='user table';