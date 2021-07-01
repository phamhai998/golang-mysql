package model

type Person struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
type S struct {
	Name string `db:"name"`
}
type ID struct {
	ID int `db:"id"`
}

func GetPersonAPI() []Person {
	listPerson := []Person{}
	DB.Select(&listPerson, "SELECT * FROM person")

	// fmt.Print(listPerson)
	return listPerson
}

func CreatePersonAPI(person Person) {
	tx := DB.MustBegin()
	tx.NamedExec("INSERT INTO person (id, name, email) VALUES (:id, :name, :email)", person)
	tx.Commit()
}

func SearchPersonAPI(s S) []Person {
	listPerson := []Person{}
	DB.Select(&listPerson, "SELECT * FROM person WHERE name=?", s.Name)
	return listPerson
}

func UpdatePersonAPI(person Person) {
	tx := DB.MustBegin()
	tx.NamedExec("SET SQL_SAFE_UPDATES=0", person)
	tx.NamedExec("UPDATE person SET name=:name, email=:email WHERE id=:id", person)
	tx.Commit()
}

func DeletePersonAPI(id ID) {
	tx := DB.MustBegin()
	tx.NamedExec("SET SQL_SAFE_UPDATES=0", id)
	tx.NamedExec("DELETE FROM person WHERE id=:id", id)
}
