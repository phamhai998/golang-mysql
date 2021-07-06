package model

type Person struct {
	ID         int    `db:"id" json:"id,string,omitempty"`
	NamePerson string `db:"namePerson" validate:"required,ascii,max=128"`
	Email      string `db:"email" validate:"required,email,max=128"`
	Image      string `db:"image" validate:"required,url"`
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
	tx.NamedExec("INSERT INTO person (id, namePerson, email, image) VALUES (:id, :namePerson, :email, :image)", person)
	tx.Commit()
}

func SearchPersonAPI(s S) []Person {
	listPerson := []Person{}
	DB.Select(&listPerson, "SELECT * FROM person WHERE namePerson=?", s.Name)
	return listPerson
}

func UpdatePersonAPI(person Person) {
	tx := DB.MustBegin()
	tx.NamedExec("SET SQL_SAFE_UPDATES=0", person)
	tx.NamedExec("UPDATE person SET name=:namePerson, email=:email, image=:image WHERE id=:id", person)
	tx.Commit()
}

func DeletePersonAPI(id ID) {
	tx := DB.MustBegin()
	tx.NamedExec("SET SQL_SAFE_UPDATES=0", id)
	tx.NamedExec("DELETE FROM person WHERE id=:id", id)
}
