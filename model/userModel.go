package model

type User struct {
	Id      int
	Name    string
	Age     int
	Address int
}

const TABLENAME = "users"

func (user *User) CreateUser() error {
	OpenSQL()
	statement := "insert into users (name,age,address) values(?,?,?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return err
	}

	defer stmt.Close()
	stmt.Exec(user.Name, user.Age, user.Address)

	if err != nil {
		return err
	}
	return err
}
