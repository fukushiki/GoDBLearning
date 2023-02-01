package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserID   string
	UserName string
	CreateAt time.Time
}

var DBConnection *sql.DB

func main() {
	os.Remove("./sqlite_test.db")

	// DBConnectionにDBをOpenしてアドレスとして保存
	DBConnection, err := sql.Open("sqlite3", "./sqlite_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer DBConnection.Close() // あとで閉じるからここがdeferになってる

	//DB 接続確認
	err = DBConnection.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected")
	}

	sqlStmt := `
		CREATE TABLE users {
			user_id varcher(32) NOT NULL,
			user_name varchar(100) NOT NULL,
			create_at timestamp with time zone,
			CONSTRAINT pk_users PRIMARY KEY (user_id)
		}
	`

	// sqlStmt := `
	// 	create table user (id integer not null primary key, name text);
	// 	delete from foo;
	// 	`
	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	// create_User()
	// db_Read()
	// db_Update()
	// db_Delete()

}

// func create_User() bool {
// 	tx, err := DBConnection.Begin()

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//stmt : ステートメント
// 	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	for i := 0; i < 100; i++ {
// 		_, err = stmt.Exec(i, fmt.Sprintf("こんにちは世界%03d", i))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return true
// }
