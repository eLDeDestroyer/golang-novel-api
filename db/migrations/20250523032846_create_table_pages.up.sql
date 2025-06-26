CREATE TABLE pages (
    id int PRIMARY KEY  AUTO_INCREMENT,
    page varchar(255) NOT NULL,
    text int NOT NULL,
    book_id int,
    FOREIGN KEY (book_id) REFERENCES books(id)
) ENGINE = InnoDB;