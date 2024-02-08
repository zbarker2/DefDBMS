CREATE DATABASE students;
\c students
CREATE TABLE students(
    StudentID INT GENERATED ALWAYS AS IDENTITY,
    FirstName varchar(255) NOT NULL,
    LastName varchar(255) NOT NULL,
    PRIMARY KEY (StudentID)
);

INSERT INTO students(FirstName, LastName)
Values('Luke', 'Skywalker'), ('Han', 'Solo')