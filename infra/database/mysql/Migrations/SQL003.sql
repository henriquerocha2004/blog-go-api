CREATE TABLE comments (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    content VARCHAR(255) NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    post_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(id)
) ENGINE innodb;


ALTER TABLE comments ADD CONSTRAINT fk_user_comment FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE comments ADD CONSTRAINT fk_post_comment FOREIGN KEY (post_id) REFERENCES posts (id);