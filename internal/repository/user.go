package repository

import (
	"database/sql"
	"fmt"
	"forum/internal/entity"
	"log"
)

type User struct {
	db *sql.DB
}

// NewUser lower layer
// that implements interface UserRepo
func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

// CreateUser errors for repo
// goes directly to database and create new user in DB
func (u *User) CreateUser(user entity.UserModel) error {
	log.Println("HEllo CreateUser repo method")
	insert := `INSERT INTO user (username,password,email) VALUES (?,?,?);`
	_, err := u.db.Exec(insert, user.Username, user.Password, user.Email)
	if err != nil {
		return fmt.Errorf("repository: create user :%v", err)
	}
	return nil
}

func (u *User) ComparePassword(password string) (string, error) {
	query := ` SELECT password FROM user WHERE username = ?`
	var hashedPassword string
	err := u.db.QueryRow(query, password).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}

func (u *User) Login() {
	fmt.Println("Login func")
	email := "m.a_k@mail.ru"
	var id, un, e string
	err := u.db.QueryRow("SELECT id, username, email FROM user WHERE email = ?", email).Scan(&id, &un, &e)
	if err != nil {
		log.Panicf("error query row - login f - %v\n", err)
	}
	fmt.Println(id, un, e)
}

func (u *User) DeleteUser(username string) error {
	deleteQuery := `DELETE FROM user WHERE username = $1;`
	_, err := u.db.Exec(deleteQuery, username)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetUser(username string) (entity.UserModel, error) {
	selectQuery := `SELECT * from user WHERE username = $1;`
	user := entity.UserModel{}
	// var id string
	err := u.db.QueryRow(selectQuery, username).Scan(&user.UserId, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) UpdateUser(username, password string) error {
	updateQuery := `UPDATE user SET password = $1 WHERE username = $2;`
	_, err := u.db.Exec(updateQuery, password, username)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetAll() ([]entity.UserModel, error) {
	selectQuery := `SELECT * FROM user;`
	users := []entity.UserModel{}
	rows, err := u.db.Query(selectQuery)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		user := entity.UserModel{}
		if err := rows.Scan("", &user.Username, &user.Email, &user.Password, &user.UserId); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
