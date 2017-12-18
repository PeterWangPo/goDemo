package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql", "write_user:write_pwd@tcp(172.16.104.207:3307)/mia_test2")
	if err != nil {
		fmt.Println("db err %s", err)
	} else {
		fmt.Println("connect ok")
	}
	f, err2 := os.Open("./c.txt")
	if err2 != nil {
		fmt.Println("open txt error %s", err2)
	}
	input := bufio.NewScanner(f)
	// e := []string{}
	// f := []string{}
	for input.Scan() {
		queryidfa := input.Text() + "|"
		rows, err4 := db.Query("SELECT id,idfa FROM cpa_platform_activation_lists where idfa =? limit 1", queryidfa)
		if err4 != nil {
			fmt.Println("Query err %s", err4)
			os.Exit(1)
		}
		fmt.Println("query ok")
		// fmt.Println(rows.Next())
		// fmt.Println("done...")
		for rows.Next() {
			var id int
			var idfa string
			err5 := rows.Scan(&id, &idfa)
			if err5 != nil {
				fmt.Println("loop err %s", err5)
			}
			fmt.Println(id, idfa)
		}
	}
}
