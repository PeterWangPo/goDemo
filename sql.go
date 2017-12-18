package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog")
	// if err != nil {
	// 	fmt.Println("db err", err)
	// }
	// defer db.Close()
	// // fmt.Println(db)
	// rows, err2 := db.Query("select * from article where id = ? or id = ?", 8, 9)
	// if err2 != nil {
	// 	fmt.Println("sql err2", err2)
	// }
	// fmt.Println(rows)
	// for rows.Next() {
	// 	var id, created, viewnum int64
	// 	var title, uri, keywords, summary, content, author string
	// 	var status int8
	// 	rows.Scan(&id, &title, &uri, &keywords, &summary, &content, &author, &created, &viewnum, &status)
	// 	fmt.Println(id, title, uri, keywords, summary, content, author, created, viewnum, status)
	// }
	query()
}

func query() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog")

	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")

	checkErr(err)

	//字典类型

	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址

	columns, _ := rows.Columns()
	// fmt.Println(columns)
	scanArgs := make([]interface{}, len(columns))

	values := make([]interface{}, len(columns))

	for i := range values {

		scanArgs[i] = &values[i]

	}
	fmt.Println(scanArgs)

	for rows.Next() {

		//将行数据保存到record字典
		record := make(map[string]string)
		err = rows.Scan(scanArgs...)

		for i, col := range values {

			if col != nil {

				record[columns[i]] = string(col.([]byte))

			}

		}
		fmt.Println(record)
	}

}

func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}
