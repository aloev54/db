--Create db
DROP DATABASE IF EXISTS library;
CREATE DATABASE library;

--Conn to db
\c library

--Authors table
CREATE TABLE authors (
    Ano SERIAL PRIMARY KEY,
    Aname VARCHAR(100) NOT NULL,
    Birth DATE,
    Country VARCHAR(50)
);

--Books table
CREATE TABLE books (
    Bno SERIAL PRIMARY KEY,
    Title VARCHAR(100) NOT NULL,
    Genre VARCHAR(50)
);

--Readers table
CREATE TABLE readers (
    Rno SERIAL PRIMARY KEY,
    Rname VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE
);

--Loan table
CREATE TABLE loans (
    Ano INT NOT NULL,
    Bno INT NOT NULL,
    Rno INT NOT NULL,
    Ldate DATE NOT NULL,
    Rdate DATE
);

