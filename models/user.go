package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Account  string
	Password string
	Phase    string
}

type MailboxConf struct {
	Title         string
	Body          string
	RecipientList []string
	Sender        string
	SPassword     string
	SMTPAddr      string
	SMTPPort      int
}

func Login(u User) bool {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	Sql := fmt.Sprintf("SELECT Account, Password FROM UserList WHERE Account = '%s' AND Password = '%s'", u.Account, u.Password)
	rows, _ := DB.Query(Sql)
	defer rows.Close()
	DB.Close()
	for rows.Next() {
		return true
	}
	return false
}

func GetCode(a string) bool {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	defer DB.Close()
	Sql := fmt.Sprintf("SELECT Account, Password FROM UserList WHERE Account = '%s'", a)
	rows, _ := DB.Query(Sql)
	defer rows.Close()
	for rows.Next() {
		return false
	}
	Sql = fmt.Sprintf("SELECT Passcode FROM UserCode WHERE Account = '%s'", a)
	rows, _ = DB.Query(Sql)
	defer rows.Close()
	for rows.Next() {
		stmt, _ := DB.Prepare("DELETE FROM UserCode WHERE Account = ?")
		stmt.Exec(a)
	}
	passcode, flag := SendEmail(a)
	if flag == false {
		return false
	}
	stmt, _ := DB.Prepare("INSERT INTO UserCode (Account, Passcode) VALUES (?, ?)")
	stmt.Exec(a, passcode)
	DB.Close()
	return true
}

func Register(a, p string) bool {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	Sql := fmt.Sprintf("SELECT Account, Passcode FROM UserCode WHERE Account = '%s'", a)
	rows, _ := DB.Query(Sql)
	defer rows.Close()
	var x, y string
	for rows.Next() {
		rows.Scan(&x, &y)
	}
	DB.Close()
	return p == y
}

func SetAccount(a, p string) {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	stmt, _ := DB.Prepare("INSERT INTO UserList (Account, Password) VALUES (?, ?)")
	stmt.Exec(a, p)
	DB.Close()
}

func ModifyPassword(a, p string) {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	stmt, _ := DB.Prepare("update UserList set Password = ? WHERE Account=?")
	stmt.Exec(p, a)
	DB.Close()
}
