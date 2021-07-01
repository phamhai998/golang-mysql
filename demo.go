// package main

// import (
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// )

// // var schema = `
// // CREATE TABLE place (
// //     country text,
// //     city text NULL,
// // 	telcode integer
// // )`

// type Person struct {
// 	ID    int
// 	Name  string `db:"name"`
// 	Email string `db:"email"`
// }

// type S struct {
// 	Name string
// }

// func connectDB() *sqlx.DB {
// 	DB, err := sqlx.Connect("mysql", "root:phamhai@tcp(127.0.0.1:3307)/golang-mysql")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// defer DB.Close()
// 	return DB
// }

// func createTable(schema string, DB *sqlx.DB) {
// 	DB.MustExec(schema)
// }

// func insertDB(class Person, DB *sqlx.DB) {
// 	tx := DB.MustBegin()
// 	tx.NamedExec("INSERT INTO myclass (id, name, email) VALUES (:id, :name, :email)", class)
// 	tx.Commit()
// }

// func insertMultipleDB(data []Person, DB *sqlx.DB) {
// 	_, err := DB.NamedExec(`INSERT INTO person (id, name, email) VALUES (:id, :name, :email)`, data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func search(s string, DB *sqlx.DB) {
// 	class := []Person{}
// 	DB.Select(&class, "SELECT * FROM myclass WHERE name=?", s)
// 	demo := class[0]
// 	fmt.Print(demo)
// }

// func show(DB *sqlx.DB) {
// 	class := []Person{}
// 	DB.Select(&class, "SELECT * FROM person")
// 	// demo := class[2]
// 	fmt.Print(class[0])
// }

// func UpdatePersonAPI(DB *sqlx.DB, person Person) {
// 	// demo := Person{}
// 	err := DB.Get(&person, `UPDATE person SET name=? WHERE id=?`, person.Name, person.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func SearchPersonAPI(s S, DB *sqlx.DB) {
// 	listPerson := []Person{}
// 	DB.Select(&listPerson, "SELECT * FROM person WHERE name=?", s.Name)
// 	demo := listPerson[0]
// 	fmt.Print(demo)
// }

// type ID struct {
// 	ID int `db:"id"`
// }

// func DeletePersonAPI(DB *sqlx.DB, id ID) {
// 	tx := DB.MustBegin()
// 	tx.NamedExec("SET SQL_SAFE_UPDATES=0", id)
// 	tx.NamedExec("DELETE FROM person WHERE id=:id", id)
// }

// func main() {
// 	fmt.Println("Go MySQL run")
// 	DB := connectDB()
// 	// createTable(schema, DB)
// 	// class := MyClass{1, "pham", "phamlonghai@gmail.com"}
// 	// insertDB(class, DB)
// 	// s := "hai"
// 	// search(s, DB)
// 	// show(DB)
// 	// person := Person{3, "hailong", "hai@gmail.com"}
// 	// UpdatePersonAPI(DB, person)
// 	id := ID{4}
// 	DeletePersonAPI(DB, id)
// 	// s := S{"hai"}
// 	// SearchPersonAPI(s, DB)
// 	// data := []Person{
// 	// 	{1, "hai", "hai@gmail.com"},
// 	// 	{2, "pham", "hai@gmail.com"},
// 	// 	{3, "hai long", "hailong@gmail.com"},
// 	// }
// 	// insertMultipleDB(data, DB)
// }
