CREATE TABLE save_book (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id int NOT NULL ,
    book_id int NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(id)
) ENGINE = InnoDB;