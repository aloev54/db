--Add constraints for library db

\c library

--Authors constr
ALTER TABLE authors 
ADD CONSTRAINT chk_author_name_length CHECK (length(Aname) >= 3),
ADD CONSTRAINT chk_birth_date CHECK (Birth IS NULL OR Birth <= current_date);

--Books constr
ALTER TABLE books 
ADD CONSTRAINT chk_title_length CHECK (length(Title) >= 3),
ADD CONSTRAINT chk_genre_values CHECK (Genre IS NULL OR Genre IN (
    'Роман', 'Фантастика', 'Детектив', 'Поэзия', 'Драма', 'Научная', 'Историческая'
));

--Readers constr
ALTER TABLE readers
ADD CONSTRAINT chk_reader_name_length CHECK (length(Rname) >= 3);

ALTER TABLE loans
ADD CONSTRAINT pk_loans PRIMARY KEY (Ano, Bno, Rno),
ADD CONSTRAINT fk_loan_author FOREIGN KEY (Ano) REFERENCES authors(Ano),
ADD CONSTRAINT fk_loan_book FOREIGN KEY (Bno) REFERENCES books(Bno),
ADD CONSTRAINT fk_loan_reader FOREIGN KEY (Rno) REFERENCES readers(Rno),
ADD CONSTRAINT chk_loan_dates CHECK (Rdate IS NULL OR Rdate >= Ldate);