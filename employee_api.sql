CREATE DATABASE employee_api;
USE employee_api;

CREATE TABLE employees
(
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    fname VARCHAR(255),
    lname VARCHAR(255),
	department VARCHAR(255)
);


INSERT INTO employees(fname, lname, department)
VALUES('John', 'Doe', 'IT'),
('Mary', 'Rose', 'IT'),
('Mark', 'Johnson', 'Marketing'),
('Jenny', 'Doppler', 'HR');







