package generate

import (
	"fmt"
	"lab_1/models"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func GenerateAllData() error {
	if err := os.MkdirAll("./data", 0755); err != nil {
		return fmt.Errorf("ошибка создания папки data: %w", err)
	}

	// Генерация данных
	authors := generateAuthors(1200)
	if err := writeToFile("./data/authors.txt", authorsToLines(authors)); err != nil {
		return fmt.Errorf("ошибка записи авторов: %w", err)
	}

	books := generateBooks(1500)
	if err := writeToFile("./data/books.txt", booksToLines(books)); err != nil {
		return fmt.Errorf("ошибка записи книг: %w", err)
	}

	readers := generateReaders(1300)
	if err := writeToFile("./data/readers.txt", readersToLines(readers)); err != nil {
		return fmt.Errorf("ошибка записи читателей: %w", err)
	}

	loans := generateLoans(2000, 1200, 1500, 1300)
	if err := writeToFile("./data/loans.txt", loansToLines(loans)); err != nil {
		return fmt.Errorf("ошибка записи выдач: %w", err)
	}

	return nil
}

func generateAuthors(count int) []models.Author {
	countries := []string{"Россия", "США", "Великобритания", "Франция", "Германия", "Япония"}
	authors := make([]models.Author, count)
	for i := 0; i < count; i++ {
		authors[i] = models.Author{
			Aname:   gofakeit.Name(),
			Birth:   gofakeit.DateRange(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()).Format("2006-01-02"),
			Country: countries[gofakeit.Number(0, len(countries)-1)],
		}
	}
	return authors
}

func generateBooks(count int) []models.Book {
	genres := []string{"Роман", "Фантастика", "Детектив", "Поэзия", "Драма", "Научная", "Историческая"}
	books := make([]models.Book, count)
	for i := 0; i < count; i++ {
		books[i] = models.Book{
			Title: gofakeit.BookTitle(),
			Genre: genres[gofakeit.Number(0, len(genres)-1)],
		}
	}
	return books
}

func generateReaders(count int) []models.Reader {
	readers := make([]models.Reader, count)
	for i := 0; i < count; i++ {
		readers[i] = models.Reader{
			Rname: gofakeit.Name(),
			Email: gofakeit.Email(),
		}
	}
	return readers
}

func generateLoans(count, maxAuthors, maxBooks, maxReaders int) []models.Loan {
	loans := make([]models.Loan, count)
	for i := 0; i < count; i++ {
		loanDate := gofakeit.DateRange(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now())
		var returnDate string
		if gofakeit.Bool() { // 50% вероятность что книга возвращена
			returnDate = loanDate.Add(time.Duration(gofakeit.Number(1, 30)) * 24 * time.Hour).Format("2006-01-02")
		}

		loans[i] = models.Loan{
			Ano:   gofakeit.Number(1, maxAuthors),
			Bno:   gofakeit.Number(1, maxBooks),
			Rno:   gofakeit.Number(1, maxReaders),
			Ldate: loanDate.Format("2006-01-02"),
			Rdate: returnDate,
		}
	}
	return loans
}

func writeToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func authorsToLines(authors []models.Author) []string {
	lines := make([]string, len(authors))
	for i, a := range authors {
		lines[i] = fmt.Sprintf("%s|%s|%s", a.Aname, a.Birth, a.Country)
	}
	return lines
}

func booksToLines(books []models.Book) []string {
	lines := make([]string, len(books))
	for i, b := range books {
		lines[i] = fmt.Sprintf("%s|%s", b.Title, b.Genre)
	}
	return lines
}

func readersToLines(readers []models.Reader) []string {
	lines := make([]string, len(readers))
	for i, r := range readers {
		lines[i] = fmt.Sprintf("%s|%s", r.Rname, r.Email)
	}
	return lines
}

func loansToLines(loans []models.Loan) []string {
	lines := make([]string, len(loans))
	for i, l := range loans {
		lines[i] = fmt.Sprintf("%d|%d|%d|%s|%s", l.Ano, l.Bno, l.Rno, l.Ldate, l.Rdate)
	}
	return lines
}