CREATE TABLE `Crimes` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `heroID` INT,
  `VillainID` INT,
  `Description` TEXT,
  `CrimeTime` DATETIME
);

ALTER TABLE `Crimes` ADD FOREIGN KEY (`heroID`) REFERENCES `Heroes` (`ID`);

ALTER TABLE `Crimes` ADD FOREIGN KEY (`VillainID`) REFERENCES `Villain` (`ID`);

INSERT INTO Crimes (heroID, VillainID, Description, CrimeTime)
VALUES
    (1,1,'Batman vs Bane', '2023-11-11 11:12:01'),
    (1,2,'Batman vs Joker', '2023-12-11 11:12:01'),
    (2,3,'Superman vs Venom', '2023-10-11 11:12:01'),
    (3,2,'Spiderman vs Joker', '2023-09-11 11:12:01');