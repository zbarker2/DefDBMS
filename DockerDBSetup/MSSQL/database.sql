CREATE DATABASE students;
GO
USE students;
GO
CREATE TABLE students(
    StudentID INT NOT NULL IDENTITY,
    FirstName varchar(255) NOT NULL,
    LastName varchar(255) NOT NULL,
    PRIMARY KEY (StudentID)
);
GO
INSERT INTO students(FirstName, LastName)
Values('Luke', 'Skywalker'), ('Han', 'Solo')