package helper

import (
	"fmt"
	"my-db/model"
)

func printRow(row model.Row, columns []model.Column) {
    for _, col := range columns {
        fmt.Printf("%-12v ", row[col.Name])
    }
    fmt.Println()
}