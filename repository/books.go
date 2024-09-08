package repository

import (
	"code/structs"
	"database/sql"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"strings"
	"time"
)

func InsertBooks(db *sql.DB, books structs.Book) (err error) {
	count := 0

	fmt.Println(books)

	//Checking books name availability

	sqls := "SELECT count(1) from books WHERE title = $1"

	errors := db.QueryRow(sqls, books.Title).Scan(&count)

	if errors != sql.ErrNoRows && errors != nil {
		return fmt.Errorf("error checking title existence: %w", errors)
	}

	if count > 0 {
		return fmt.Errorf("title already exists")
	}

	//Checking category availability

	sqls = "SELECT count(1) from categories WHERE name = $1"

	errors = db.QueryRow(sqls, books.CategoryID).Scan(&count)

	if errors != sql.ErrNoRows && errors != nil {
		return fmt.Errorf("error checking category existence: %w", errors)
	}

	if count > 0 {
		return fmt.Errorf("category doesnt exist")
	}

	if books.TotalPage > 100 {
		books.Thickness = "Tebal"
	} else {
		books.Thickness = "Tipis"
	}

	if !IsValidURL(books.ImageURL) {
		return fmt.Errorf("URL INVALID")
	}

	if books.ReleaseYear < 1980 || books.ReleaseYear > 2024 {
		return fmt.Errorf("INVALID RELEASE YEAR (1980 - 2024), Current %d", books.ReleaseYear)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	books.ID = rng.Intn(1000000) // Generates a random ID between 0 and 999999

	sqlQuery := `
	INSERT INTO books(
		id, title, category_id, description, image_url, release_year, 
		price, total_page, thickness, created_at, created_by, 
		modified_at, modified_by
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
	)`

	// Execute the query with the book data
	_, errors = db.Exec(sqlQuery,
		books.ID, books.Title, books.CategoryID, books.Description,
		books.ImageURL, books.ReleaseYear, books.Price,
		books.TotalPage, books.Thickness, books.CreatedAt,
		books.CreatedBy, books.ModifiedAt, books.ModifiedBy)

	if errors != sql.ErrNoRows && errors != nil {
		fmt.Println("errorss :", errors)
		return fmt.Errorf("Cannot add : %s", books.Title)
	}

	return nil
}

func GetAllBooks(db *sql.DB, books *[]structs.Book) (err error) {
	sql := "SELECT title FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Errors (44) : %w", err)
		return fmt.Errorf("errors (44): Something went wrong")
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book
		err = rows.Scan(&book.Title)
		if err != nil {
			return
		}

		*books = append(*books, book)
	}

	return nil
}

func GetBooksDet(db *sql.DB, books *structs.Book) (err error) {
	sqls := "SELECT * FROM books WHERE id = $1"

	errors := db.QueryRow(sqls, books.ID).Scan(&books.ID,
		&books.Title,
		&books.CategoryID,
		&books.Description,
		&books.ImageURL,
		&books.ReleaseYear,
		&books.Price,
		&books.TotalPage,
		&books.Thickness,
		&books.CreatedAt,
		&books.CreatedBy,
		&books.ModifiedAt,
		&books.ModifiedBy)

	if errors != nil {
		return fmt.Errorf("error checking category existence: %w", errors)
	}

	return nil
}

func DeleteBook(db *sql.DB, books *structs.Book) (err error) {
	sqls := fmt.Sprintf("DELETE FROM books WHERE id = %d", books.ID)

	res, errors := db.Exec(sqls)

	if errors != nil {
		fmt.Println("error : ", errors)
		return fmt.Errorf("cannot delete %d", books.ID)
	}

	cnt, errors := res.RowsAffected()

	if cnt == 0 {
		return fmt.Errorf("error checking category existence")
	}
	if errors != nil {
		fmt.Println("errorss : ", errors)
		return fmt.Errorf("cannot delete %d", books.ID)
	}

	return nil
}

func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	// Split the host and port (if any)
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return false
		}
	}

	// Check if the host is a valid domain or IP address
	if net.ParseIP(host) != nil {
		return true
	}
	if strings.Contains(host, ".") && len(host) > 1 {
		return true
	}

	return false
}
