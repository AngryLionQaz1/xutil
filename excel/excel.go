package excel

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"xutil/exp"
	"xutil/prt"
)

/**
  读取Excel 标题信息
*/

func ReadTitle(path string) {

	file, e := excelize.OpenFile(path)
	exp.Exp(e)

	rows := file.GetRows("Sheet1")
	for i := 0; i < 1; i++ {
		for _, colCell := range rows[i] {
			//fmt.Print(colCell, "\t")
			prt.PrintlnRight(colCell)
		}
		fmt.Println()
	}

}
