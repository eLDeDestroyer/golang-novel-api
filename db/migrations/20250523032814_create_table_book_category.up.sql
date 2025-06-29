CREATE TABLE book_category (
    id int PRIMARY KEY AUTO_INCREMENT,
    book_id int not null,
    category_id int not null,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
) ENGINE = InnoDB;