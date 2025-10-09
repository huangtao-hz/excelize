package main

import (
	"excelize"
	"fmt"
)

func main() {
	var err error
	defer func() {
		if err == nil {
			fmt.Println("Success")
		} else {
			fmt.Println(err)
		}
	}()
	book := excelize.NewFile()
	sheet := "Sheet1"
	book.SetSheetRow(sheet, "A2", &[]any{"你好", 13, 25})
	book.SetSheetRow(sheet, "A3", &[]any{"你好", 14, 27})
	err = book.AddTable(sheet, &excelize.Table{
		Range:     "A1:C4",
		StyleName: "TableStyleMedium2",
		Columns: []excelize.TableColumn{
			excelize.TableColumn{
				Name:           "姓名",
				TotalsRowLabel: "合计",
			},
			excelize.TableColumn{
				Name:              "年龄",
				TotalsRowFunction: "average",
			},
			excelize.TableColumn{
				Name:              "分数",
				TotalsRowFunction: "countNums",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = book.SaveAs("/Users/huangtao/abc.xlsx")

}
