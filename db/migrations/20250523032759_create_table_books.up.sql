CREATE TABLE books (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title varchar(255) NOT NULL ,
    description text NOT NULL ,
    image_path text NOT NULL ,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;