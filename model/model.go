package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
}

var table = `CREATE TABLE IF NOT EXISTS USERS(
    "Id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "FirstName" TEXT,
    "LastName" TEXT,
    "Email"TEXT
);`

type database struct {
	DB *sql.DB
}

func NewDatabse(db *sql.DB) *database {
	return &database{
		DB: db,
	}
}
func ConnectDatabase() {
	db, err := sql.Open("sqlite3", "./bank.db")
	if err != nil {
		log.Fatal("Error While Connecting Database.")
	}
	stmt, err := db.Prepare(table)
	if err != nil {
		fmt.Println("Error While Creating Table.")
	}
	stmt.Exec()

}

// type Database interface {
// 	ListUser() ([]*User, error)
// 	SingleUser(Id int64) (*User, error)
// 	CreateUser(firstname string, lastname string, email string) error
// 	DeleteUser(id int64) error
// 	UpdateUser(Id int64, FirstName string, LastName string, Email string) error
// }

func (d *database) ListUser() ([]*User, error) {
	allusers := `SELECT (Id,FirstName,LastName,Email) FROM USERS;
	`
	var users []*User
	rows, err := d.DB.Query(allusers)
	if err != nil {
		panic(err)
	}
	var id int64
	var first_name string
	var last_name string
	var email string
	for rows.Next() {
		rows.Scan(&id, &first_name, &last_name, &email)
	}
	users = append(users, &User{Id: id, FirstName: first_name, LastName: last_name, Email: email})
	defer rows.Close()
	defer d.DB.Close()
	return users, nil
}

func (d *database) SingleUser(Id int64) (*User, error) {
	singleuser := `SELECT (Id,FirstName,LastName,Email) FROM USERS WHERE Id = ?;`
	var user *User
	row, err := d.DB.Query(singleuser, Id)
	if err != nil {
		panic(err)
	}
	var id int64
	var first_name string
	var last_name string
	var email string
	for row.Next() {
		row.Scan(&id, &first_name, &last_name, &email)
	}
	user = &User{Id: id, FirstName: first_name, LastName: last_name, Email: email}
	defer row.Close()
	defer d.DB.Close()
	return user, nil
}
func (d *database) CreateUser(firstname string, lastname string, email string) error {
	createnewuser := `INSERT INTO USERS (FirstName,LastName,Email)VALUES(?,?,?);`
	stmt, err := d.DB.Prepare(createnewuser)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(firstname, lastname, email)
	if err != nil {
		panic(err)
	}
	n, _ := res.RowsAffected()
	n2, _ := res.LastInsertId()
	fmt.Printf("No. Of Rows Affected : %d\n", n)
	fmt.Printf("Last Insert Id : %d\n", n2)
	defer d.DB.Close()
	return nil
}

func (d *database) DeleteUser(id int64) error {
	deleteuserbyid := `DELETE FROM USERS WHERE Id = ?;`
	stmt, err := d.DB.Prepare(deleteuserbyid)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	n, _ := res.RowsAffected()
	fmt.Println("No. Of Rows Affected : ", n)
	defer d.DB.Close()
	return nil
}

func (d *database) UpdateUser(Id int64, FirstName string, LastName string, Email string) error {
	updateuser := `UPDATE USERS
	SET FirstName = ?,
		LastName = ?,
		Email = ?
	WHERE
	   Id = ?;`
	stmt, err := d.DB.Prepare(updateuser)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(FirstName, LastName, Email, Id)
	if err != nil {
		panic(err)
	}
	n, _ := res.RowsAffected()
	fmt.Println("No. Of Rows Affected : ", n)
	defer d.DB.Close()
	return nil
}
