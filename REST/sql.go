package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func main()  {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3",":memory:")
	err = db.Ping()
	if err != nil{
		panic(err.Error())
	}

	createSchema(db)
}

//User...
type User struct{
	ID int `db:"id"`
	Name string `db: "name"`
}

func createSchema(db *sqlx.DB)  {
	schema := `CREATE TABLE user (
		id integer primary key autoincrement,
		name varchar(56));`

	_, err := db.Exec(schema)
	if err!=nil{
		panic(err.Error())
		//fmt.Println(r)
	}/*else{
		fmt.Println(err)
	}*/

	db.MustExec("INSERT INTO user (name) VALUES (?)", "Jane Doe")
	rows, err := db.Queryx("SELECT id, name FROM user")
	if err!= nil{
		panic(err.Error())
	}

	for rows.Next(){
		var u User
		//rows.Scan(&u.ID, &u.Name) ASÍ IMPRIMÍA EL U
		rows.StructScan(&u)		//ASÍ IMPRIME {0}
		fmt.Println(u)
	}
}