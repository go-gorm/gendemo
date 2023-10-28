CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL DEFAULT '',
                        `extra` json DEFAULT NULL,
                        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                        `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
                        `phone` varchar(255) NOT NULL DEFAULT '',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;