CREATE DATABASE students;

USE students;

CREATE TABLE students(
    StudentID INT NOT NULL AUTO_INCREMENT,
    FirstName varchar(255) NOT NULL,
    LastName varchar(255) NOT NULL,
    PRIMARY KEY (StudentID)
);

INSERT INTO students(FirstName, LastName)
Values('Luke', 'Skywalker'), ('Han', 'Solo')