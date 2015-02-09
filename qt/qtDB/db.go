package qtDB

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"qt/qtTime"
	"strings"
	"time"
)

func open() *sql.DB {
	const dbName = "qtDB"
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	logErr(err)
	return db
}

func Create(list []string) {
	db := open()
	defer db.Close()

	tx, err := db.Begin()
	logErr(err)

	var code string

	for i, l := 0, len(list); i < l; i++ {
		if strings.Contains(list[i], "r_") {
			code = strings.Split(list[i], "_")[1]
		} else {
			code = list[i]
		}
		// Table 存在检测
		sqlStmt := "create table if not exists " + code + " (id integer primary key autoincrement not null, date text not null, timeUnix integer not null, content text not null);"
		_, err := db.Exec(sqlStmt)
		logErr(err)
	}
	tx.Commit()
}

func Insert(data []string) {
	db := open()
	defer db.Close()

	tx, err := db.Begin()
	logErr(err)

	// 变量初始化
	var code string
	var market rune
	var time time.Time
	layout := "2015-01-01 09:00:00"

	for i, l := 0, len(data); i < l; i++ {
		// 数据格式化
		qtArray := strings.Split(data[i], "~")
		split := []rune(qtArray[0])

		// 区分实时数据
		if split[2] == 'r' {
			market = split[4]
			code = strings.Split(strings.Split(qtArray[0], "=")[0], "_")[2]
		} else {
			market = split[2]
			code = strings.Split(strings.Split(qtArray[0], "=")[0], "_")[1]
		}

		stmt, err := tx.Prepare("insert into " + code + "(id, date, timeUnix, content) values(?,?,?,?)")
		logErr(err)

		defer stmt.Close()

		if market == 's' {
			time = qtTime.ParseHS(qtArray[30])
		} else if market == 'h' {
			time = qtTime.ParseHK(qtArray[30])
		} else {
			time = qtTime.ParseUS(qtArray[30])
		}
		date := qtTime.GetDate(time)
		_, err = stmt.Exec(nil, date, time.Format(layout), data[i])
		logErr(err)
	}
	tx.Commit()
}

// func Query(db *sql.DB) {
// 	rows, err := db.Query("select id, name from foo")
// 	logErr(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		rows.Scan(&id, &name)
// 		fmt.Println(id, name)
// 	}
// 	rows.Close()
// }

// func main() {

// 	stmt, err = db.Prepare("select name from foo where id = ?")
// 	logErr(err)
// 	defer stmt.Close()
// 	var name string
// 	err = stmt.QueryRow("3").Scan(&name)
// 	logErr(err)
// 	fmt.Println(name)

// 	_, err = db.Exec("delete from foo")
// 	logErr(err)

// 	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
// 	logErr(err)

// 	rows, err = db.Query("select id, name from foo")
// 	logErr(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		rows.Scan(&id, &name)
// 		fmt.Println(id, name)
// 	}
// }

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
