\c library

\copy authors FROM './data/authors.txt' DELIMITER '|' NULL '';
\copy books FROM './data/books.txt' DELIMITER '|' NULL '';
\copy readers FROM './data/readers.txt' DELIMITER '|' NULL '';
\copy loans FROM './data/loans.txt' DELIMITER '|' NULL '';
