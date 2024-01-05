

-- DDL for User table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amount DECIMAL(10, 2) DEFAULT 0.00
);

-- DDL for Product table
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    stock INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);


CREATE TABLE stores (
    Store_id      serial primary key,
	Store_name    VARCHAR(100),
	Store_address VARCHAR(100),
	Longitude     VARCHAR(100),
	Latitude      VARCHAR(100),
	Rating        DEC(10,2)
);

-- DDL for Transaction table
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    product_id INT REFERENCES products(product_id),
    quantity INT NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    store_id int REFERENCES stores(store_id)
);

INSERT INTO users (username, password, deposit_amount) VALUES -- pass 123456
    ('jimmy', '$2a$10$AzOh0n4exrsrULZWz95liuc1yr5bzFBbTta9k2y6GlQRbedlcEcUG', 100),
    ('jane', '$2a$10$AzOh0n4exrsrULZWz95liuc1yr5bzFBbTta9k2y6GlQRbedlcEcUG', 100);

INSERT INTO products (product_name, stock, price) VALUES
    ('phone', 20000, 15),
    ('shirt', 20, 5),
    ('tv', 50, 15),
    ('pants', 10, 5);

INSERT INTO transactions (user_id, product_id, quantity, total_amount) VALUES
    (1, 1, 2, 30.00),
    (2, 3, 1, 15.00);

INSERT INTO stores (store_name,Store_address, Longitude, Latitude, Rating) VALUES
    ("shop1",	"D street 5",	"106.816666",	"-6.2",	4.00);