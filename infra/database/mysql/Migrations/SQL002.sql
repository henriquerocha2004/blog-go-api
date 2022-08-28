-- Active: 1659521641949@@127.0.0.1@3306@blogapi
CREATE TABLE categories (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    description VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
) ENGINE innodb;

CREATE TABLE posts (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    content BLOB NOT NULL,
    created_at DATETIME NOT NULL,
    PRIMARY KEY(id)
) ENGINE innodb;

ALTER TABLE posts ADD CONSTRAINT fk_user_post FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE category_post (
    id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
    category_id BIGINT UNSIGNED NOT NULL,
    post_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(id)
) ENGINE innodb;

ALTER TABLE category_post ADD CONSTRAINT fk_category_post FOREIGN KEY (category_id) REFERENCES categories (id); 
ALTER TABLE category_post ADD CONSTRAINT fk_post_category FOREIGN KEY (post_id) REFERENCES posts (id); 