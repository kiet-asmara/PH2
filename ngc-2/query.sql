CREATE TABLE Heroes (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100),
    Universe VARCHAR(100),
    Skill VARCHAR(100),
    ImageURL VARCHAR(150)
);

CREATE TABLE Villain (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100),
    Universe VARCHAR(100),
    ImageURL VARCHAR(150)
);

INSERT INTO Heroes (Name, Universe, Skill, ImageURL)
VALUES
    ('Batman', 'DC', 'Money, Batmobile', 'batman.jpg'),
    ('Superman', 'DC', 'Lazers, Flying', 'superman.jpg'),
    ('Spiderman', 'Marvel', 'Spider Sense, Webs', 'spiderman.jpg');

INSERT INTO Villain (Name, Universe, ImageURL)
VALUES
    ('Bane', 'DC', 'bane.jpg'),
    ('Joker', 'DC', 'joker.jpg'),
    ('Venom', 'Marvel', 'venom.jpg');
