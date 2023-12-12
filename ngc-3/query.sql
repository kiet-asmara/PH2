CREATE TABLE Inventories (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100),
    ItemCode VARCHAR(5) UNIQUE,
    Stock INT CHECK (Stock >= 0),
    Description VARCHAR(255),
    Status ENUM ('active', 'broken')
);

INSERT INTO Inventories (Name, ItemCode, Stock, Description, Status)
VALUES
    ('Gun', 'GUN02', 20, 'A gun', 'active'),
    ('Shield', 'SLD13', 5, 'Shield to block', 'broken'),
    ('Shirt', 'SHR22', 36, 'T-shirt', 'active');