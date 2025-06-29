CREATE TABLE histories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id int NOT NULL ,
    book_id int NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
) ENGINE = InnoDB;