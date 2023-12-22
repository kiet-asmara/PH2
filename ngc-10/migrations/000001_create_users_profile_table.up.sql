CREATE TABLE users (
	id serial PRIMARY KEY,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL
);

CREATE TABLE profiles (
	id serial PRIMARY KEY,
	first_name VARCHAR ( 255 ) UNIQUE NOT NULL,
	last_name VARCHAR ( 255 ) NOT NULL,
	address VARCHAR ( 255 ) NOT NULL,
    phone_number VARCHAR ( 255 ) NOT NULL,
    user_id INT NOT NULL,
      FOREIGN KEY (user_id)
      REFERENCES users (id)
);