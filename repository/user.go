package repository

import (
	"code/structs"
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

func InsertUser(db *sql.DB, user structs.User) (err error) {
	// Generate a random ID for the user
	count := 0

	fmt.Println(user)

	sqls := "SELECT count(1) from users WHERE username = $1"

	errors := db.QueryRow(sqls, user.Username).Scan(&count)

	if errors != sql.ErrNoRows && errors != nil {
		return fmt.Errorf("error checking username existence: %w", errors)
	}

	if count > 0 {
		return fmt.Errorf("username already exists")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	user.ID = rng.Intn(1000000) // Generates a random ID between 0 and 999999

	// Prepare the SQL query
	sqlQuery := "INSERT INTO users(id, username, password, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// Execute the query with the user data
	_, err = db.Exec(sqlQuery, user.ID, user.Username, string(user.Password), user.CreatedAt, user.Username, user.ModifiedAt, user.Username)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return fmt.Errorf("ERRORS : %w", err)
	}

	return nil
}

func GetUser(db *sql.DB, user *structs.User, encryptedPass string) (err error) {

	fmt.Println(user)

	sqls := "SELECT id, username, password from users WHERE username = $1"

	errors := db.QueryRow(sqls, user.Username).Scan(&user.ID, &user.Username, &user.Password)

	if errors != nil {
		fmt.Println("Errors (13) :", errors)
		return fmt.Errorf("errors (13): wrong password or username")
	}

	if encryptedPass != user.Password {
		return fmt.Errorf("errors (14) : wrong password or username")
	}

	user.Password = ""

	return nil
}

func KeepLogin(db *sql.DB, user string) (err error) {

	var username string

	fmt.Println(user)

	sqls := "SELECT username from users WHERE username = $1"

	errors := db.QueryRow(sqls, user).Scan(&username)

	if errors != nil {
		fmt.Println("Errors (13)", errors)
		return fmt.Errorf("errors (13): wrong password or username")
	}

	return nil
}
