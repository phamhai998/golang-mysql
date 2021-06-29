package model

type Person struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
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

func UpdatePersonAPI(person Person) {
	tx := DB.MustBegin()
	tx.NamedExec("UPDATE person SET Name=? WHERE id=?", person)
}
