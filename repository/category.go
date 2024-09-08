package repository

import (
	"code/structs"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

func GetAllCategories(db *sql.DB, categories *[]structs.Category) (err error) {

	fmt.Println(categories)

	sql := "SELECT name FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Errors (19) : %w", err)
		return fmt.Errorf("errors (19): Something went wrong")
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.Name)
		if err != nil {
			return
		}

		*categories = append(*categories, category)
	}

	return nil
}

func InsertCategories(db *sql.DB, categories structs.Category) (err error) {

	count := 0

	fmt.Println(categories)

	sqls := "SELECT count(1) from categories WHERE name = $1"

	errors := db.QueryRow(sqls, categories.Name).Scan(&count)

	if errors != sql.ErrNoRows && errors != nil {
		return fmt.Errorf("error checking name existence: %w", errors)
	}

	if count > 0 {
		return fmt.Errorf("name already exists")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	categories.ID = rng.Intn(1000000) // Generates a random ID between 0 and 999999

	// Prepare the SQL query
	sqlQuery := "INSERT INTO categories(id, name, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6)"

	// Execute the query with the category data
	_, err = db.Exec(sqlQuery, categories.ID, categories.Name, categories.CreatedAt, categories.CreatedBy, categories.ModifiedAt, categories.ModifiedBy)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return fmt.Errorf("ERRORS : %w", err)
	}

	return nil
}

func GetCategoryDet(db *sql.DB, categories *structs.Category) (err error) {
	sqls := "SELECT * FROM categories WHERE id = $1"

	errors := db.QueryRow(sqls, categories.ID).Scan(&categories.ID, &categories.Name, &categories.CreatedAt, &categories.CreatedBy, &categories.ModifiedAt, &categories.ModifiedBy)

	if errors != nil {
		return fmt.Errorf("error checking category existence: %w", errors)
	}

	return nil
}

func DeleteCategories(db *sql.DB, categories *structs.Category) (err error) {
	sqls := fmt.Sprintf("DELETE FROM categories WHERE id = %d", categories.ID)

	res, errors := db.Exec(sqls)

	if errors != nil {
		fmt.Println("error : ", errors)
		return fmt.Errorf("cannot delete %d", categories.ID)
	}

	cnt, errors := res.RowsAffected()

	if cnt == 0 {
		return fmt.Errorf("error checking category existence")
	}
	if errors != nil {
		fmt.Println("errorss : ", errors)
		return fmt.Errorf("cannot delete %d", categories.ID)
	}

	return nil

}

func GetCatBook(db *sql.DB, categories *structs.Category, names *[]string) (err error) {
	sqls := "SELECT title FROM books WHERE category_id = $1"

	rows, err := db.Query(sqls, categories.ID)

	if err != nil {
		fmt.Println("Errors (88) : %w", err)
		return fmt.Errorf("errors (88): Something went wrong")
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category
		err = rows.Scan(&category.Name)
		if err != nil {
			return
		}

		*names = append(*names, category.Name)
	}

	return nil
}
