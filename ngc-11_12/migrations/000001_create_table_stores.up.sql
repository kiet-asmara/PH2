CREATE TABLE stores (
    Store_id      serial primary key,
	Store_name    VARCHAR(100),
	Store_address VARCHAR(100),
	Longitude     VARCHAR(100),
	Latitude      VARCHAR(100),
	Rating        DEC(10,2)
);