CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username varchar(255) NOT NULL UNIQUE ,
    name varchar(255) NOT NULL ,
    password varchar(255) NOT NULL ,
    is_premium BOOLEAN DEFAULT FALSE
) ENGINE = InnoDB;

