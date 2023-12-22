CREATE TABLE members (
	id INT AUTO_INCREMENT PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    full_name varchar(255) NOT NULL,
    age int NOT NULL CHECK (age >= 17),
    occupation varchar(100) NOT NULL,
    role varchar(10) NOT NULL
);