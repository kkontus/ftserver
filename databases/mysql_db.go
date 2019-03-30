package databases

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	. "ftserver/entity" // using dot is not a best practice, I only use it on entity/struct that are reused for cleaner code,
	// but in general everything else should use qualifiers
)

type DB struct {
	DB *sql.DB
}

const (
	MYSQL_HOST     = "127.0.0.1" // or localhost
	MYSQL_PORT     = "3306"
	MYSQL_USERNAME = "root"
	MYSQL_PASSWORD = "root"
	MYSQL_DBNAME   = "quiz"
)

// example connect: db, err := cdb.Connect(cdb.MYSQL_HOST, cdb.MYSQL_PORT, cdb.MYSQL_USERNAME, cdb.MYSQL_PASSWORD, cdb.MYSQL_DBNAME)
func Connect(host string, port string, username string, password string, dbname string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	db, err := sql.Open("mysql", dataSource)
	return db, err
}

func FindAll(db *sql.DB, err error) []Questions {
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT id, question, correct, answers FROM questions")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	var id int
	var question string
	var correct string
	var answers string

	questions := []Questions{}
	q := Questions{}
	for rows.Next() {
		rows.Scan(&id, &question, &correct, &answers)
		//fmt.Printf("%d : %s %s %s \n", id, question, correct, answers)
		q.Question = question
		q.Correct = correct
		q.Answers = strings.Split(answers, ",")

		questions = append(questions, q)
	}

	return questions
}
