package models

import (
	"database/sql"
	"fmt"
	"time"
)

func Generate(a, p string, n int) (ans []Question) {
	fmt.Println(n)
	switch p {
	case "小学":
		Limit = 4
	case "初中":
		Limit = 6
	case "高中":
		Limit = 9
	}
	TimeStr = time.Now().Format("2006-01-02-15-04-05")
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	for len(ans) < n {
		PhaseFlag = 4
		question, answer, flag := "", "", false
		AA, BB, CC, DD := "", "", "", ""
		Root.Val, Root.Left, Root.Right = "", nil, nil
		var c [4]string
		OptCount := RandNumber(5) + 1
		InitTree(&Root, OptCount)
		question, answer, flag = TransferTree(&Root)
		if PhaseFlag != Limit || !flag || answer == "NaN" || question == "" {
			continue
		}
		if answer == "-0.00" {
			answer = "0.00"
		}
		RandChoice(answer, &c)
		AA, BB, CC, DD = c[0], c[1], c[2], c[3]
		stmt, _ := DB.Prepare("INSERT INTO Questions (PaperName, Account, Question, Answer, A, B, C, D, Choice) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if _, err := stmt.Exec(TimeStr, a, question, answer, AA, BB, CC, DD, ""); err == nil {
			ans = append(ans, Question{"", "$$" + question + "$$", "", AA, BB, CC, DD, ""})
		}
	}
	DB.Close()
	return
}

func Check(a string, q, c []string) []bool {
	n := len(q)
	var ans []bool
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	defer DB.Close()
	for i := 0; i < n; i++ {
		m := len(q[i])
		q[i] = q[i][2 : m-2]
		Sql := fmt.Sprintf("SELECT Answer FROM Questions WHERE Account = '%s' AND Question = '%s'", a, q[i])
		rows, _ := DB.Query(Sql)
		defer rows.Close()
		var s string
		for rows.Next() {
			rows.Scan(&s)
		}
		stmt, _ := DB.Prepare("UPDATE Questions SET Choice = ? WHERE Account = ? AND Question = ?")
		stmt.Exec(c[i], a, q[i])
		flag := c[i] == s
		ans = append(ans, flag)
		fmt.Println(c[i], "   ", s)
	}
	return ans

}

func CheckList(a string) (ans []string) {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	defer DB.Close()
	Sql := fmt.Sprintf("SELECT DISTINCT PaperName FROM Questions WHERE Account = '%s' ORDER BY PaperName DESC", a)
	rows, _ := DB.Query(Sql)
	defer rows.Close()
	var str string
	for rows.Next() {
		rows.Scan(&str)
		ans = append(ans, str)
	}
	return
}

func CheckPaper(a, t string) (ans []Question) {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	defer DB.Close()
	Sql := fmt.Sprintf("SELECT Question, Answer, A, B, C, D, Choice FROM Questions WHERE Account = '%s' AND PaperName= '%s'", a, t)
	rows, _ := DB.Query(Sql)
	defer rows.Close()
	i := 0
	for rows.Next() {
		var b, c, d, e, f, g, h string
		rows.Scan(&b, &c, &d, &e, &f, &g, &h)
		ans = append(ans, Question{"", b, c, d, e, f, g, h})
		i++
	}
	return
}

func DeletePaper(a, t string) bool {
	DB, _ = sql.Open("sqlite3", "sqlite/Question.db")
	defer DB.Close()
	stmt, _ := DB.Prepare("DELETE FROM Questions WHERE Account = ? AND PaperName = ?")
	stmt.Exec(a, t)
	return true

}
