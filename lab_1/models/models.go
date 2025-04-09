package models

type Author struct {
	Ano int
	Aname string
	Birth string
	Country string
}

type Book struct {
	Bno int
	Title string
	Genre string
}

type Reader struct {
	Rno int
	Rname string
	Email string
}

type Loan struct {
	Ano int
	Bno int
	Rno int
	Ldate string
	Rdate string
}