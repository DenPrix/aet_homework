package utils

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func InitDatabase() {
	fullSqlName := Config.Username + ":" + Config.Password + "@/" + Config.Database
	db, err := sql.Open("mysql", fullSqlName)
	if err != nil {
		log.Println(err)
	}
	Database = db
}

func CloseDatabase() {
	Database.Close()
}

func PrepareDatabase(MethodData func() string) *sql.Stmt {
	stmt, err := Database.Prepare(MethodData())
	if err != nil {
		log.Println(err)
	}
	return stmt
}

func ExecDatabaseWithId(contact Contact, stmt *sql.Stmt, id string) int64 {
	res, err := stmt.Exec(contact.Name, contact.Phone, contact.Email, id)
	if err != nil {
		log.Println(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return lastId
}

func ExecDatabase(contact Contact, stmt *sql.Stmt) int64 {
	res, err := stmt.Exec(contact.Name, contact.Phone, contact.Email)
	if err != nil {
		log.Println(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return lastId
}

func ExecDatabaseIdOnly(id string, stmt *sql.Stmt) int64 {
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return lastId
}

func GetData() string {
	return "SELECT * FROM " + Config.Database + ".contacts WHERE id = ?"
}

func InsertData() string {
	return "insert into " + Config.Database + ".contacts (name, phone, email) values (?, ?, ?)"
}

func UpdateData() string {
	return "update " + Config.Database + ".contacts set name = ?, phone = ?, email = ? where id = ?"
}

func DeleteData() string {
	return "delete from " + Config.Database + ".contacts where id = ?"
}
