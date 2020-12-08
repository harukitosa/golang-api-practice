
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `posts` (
  `post_id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `author` VARCHAR(255) NOT NULL,
  `text` TEXT NOT NULL
) ENGINE = InnoDB;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `posts`;