CREATE TABLE pages (
    id int PRIMARY KEY  AUTO_INCREMENT,
    page int NOT NULL,
    text text NOT NULL,
    book_id int,
    FOREIGN KEY (book_id) REFERENCES books(id)
) ENGINE = InnoDB;