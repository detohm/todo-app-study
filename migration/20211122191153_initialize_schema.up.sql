CREATE TABLE todo (
    id INT NOT NULL AUTO_INCREMENT,
    description VARCHAR(100) NOT NULL,
    is_completed BOOLEAN NOT NULL,
    is_deleted BOOLEAN NOT NULL,

    PRIMARY KEY (id)
);
