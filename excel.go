package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	file := "./a.xlsx"
	xls, err := xlsx.OpenFile(file)
	if err != nil {
		fmt.Printf("err %s", err)
		os.Exit(1)
	}
	txt, err1 := os.OpenFile("./c.txt", os.O_APPEND, 0666)
	if err1 != nil {
		fmt.Printf("err1 %s", err1)
		os.Exit(1)
	}
	e := []string{}
	for _, sheet := range xls.Sheets {
		for _, row := range sheet.Rows {
			for k, cell := range row.Cells {
				if k == 0 {
					tmp, err3 := cell.String()
					if err3 != nil {
						e = append(e, tmp)
					} else {
						txt.WriteString(tmp + "\r\n")
					}
				}
			}
		}
	}
	if len(e) > 0 {
		txt.WriteString("err ===>\r\n")
		for _, v := range e {
			txt.WriteString(v + "\r\n")
		}
	}
	txt.Close()
}
