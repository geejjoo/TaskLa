CREATE TABLE STORAGE
(
    name varchar(255) PRIMARY KEY,
    available bool NOT NULL
);
CREATE TABLE PRODUCT
(
    articular serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    size int NOT NULL
);
CREATE TABLE RESERVE
(
    id serial PRIMARY KEY,
    articular int NOT NULL,
    product varchar(255) NOT NULL,
    storage varchar(255) NOT NULL,
    quantity int NOT NULL,
    reserved int,
    FOREIGN KEY (articular) REFERENCES PRODUCT(articular),
    FOREIGN KEY (storage) REFERENCES STORAGE(name)
);