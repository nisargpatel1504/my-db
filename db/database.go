package db

import (
	"fmt"
	"my-db/model"
	"my-db/storage"
)

type Database struct {
    Tables map[string]model.Table
}

func NewDatabase() *Database {
    return &Database{
        Tables: make(map[string]model.Table),
    }
}

func (db *Database) CreateTable(name string, columns []model.Column) {
    db.Tables[name] = model.Table{Name: name, Columns: columns, Rows: make([]model.Row, 0)}
 
}

func (db *Database) Insert(tableName string, rowData model.Row) {
    table, exists := db.Tables[tableName]
    if !exists {
        fmt.Println("Table does not exist.")
        return
    }
    table.Rows = append(table.Rows, rowData)
    db.Tables[tableName] = table  // Update table
    storage.WriteTable(table)     // Persist data
}

func (db *Database) Select(tableName string) {
    table, exists := db.Tables[tableName]
    if !exists {
        fmt.Println("Table does not exist.")
        return
    }
    for _, row := range table.Rows {
        fmt.Println(row)
    }
}

func (db *Database) Update(tableName string, rowIndex int, newData model.Row) {
    table, exists := db.Tables[tableName]
    if !exists {
        fmt.Println("Table does not exist.")
        return
    }
    if rowIndex >= len(table.Rows) {
        fmt.Println("Row index out of range.")
        return
    }
    for key, value := range newData {
        table.Rows[rowIndex][key] = value
    }
    storage.WriteTable(table) // Persist data
}

func (db *Database) Delete(tableName string, rowIndex int) {
    table, exists := db.Tables[tableName]
    if !exists {
        fmt.Println("Table does not exist.")
        return
    }
    if rowIndex >= len(table.Rows) {
        fmt.Println("Row index out of range.")
        return
    }
    table.Rows = append(table.Rows[:rowIndex], table.Rows[rowIndex+1:]...)
    db.Tables[tableName] = table  // Update table
    storage.WriteTable(table)     // Persist data
}
func (db *Database) JoinTables(query model.JoinQuery) {
    t1, exists1 := db.Tables[query.Table1]
    
    if !exists1 {
        fmt.Println("Table", query.Table1, "does not exist.")
        return
    }

    t2, exists2 := db.Tables[query.Table2]
    if !exists2 {
        fmt.Println("Table", query.Table2, "does not exist.")
        return
    }

    fmt.Println("Join Results:")
    for _,char := range query.Select{
        fmt.Printf("%-10s | ", char)
    }
    // Header
    fmt.Println()
    fmt.Println("----------------------------------------------------")

    // Iterate through rows of both tables and find matching join column values
    for _, row1 := range t1.Rows {
        // fmt.Println("row1    :",row1,"joincol        :",query.JoinCol);
        for _, row2 := range t2.Rows {
            // fmt.Print("row2        :",row2,"joincol         :",row2[query.WhereCol])
            if row1[query.JoinCol] == row2[query.WhereCol] {
                
                if query.WhereCol != "" && row1[query.JoinCol] != query.WhereVal {
                
                    continue  // Skip rows that do not meet the WHERE condition
                }
                // Format and print each matching row
            
            printSelectedRow(row1, row2, t1.Columns, t2.Columns, query.Select)
            }
        }
    }
}
func printSelectedRow(row1, row2 model.Row, cols1, cols2 []model.Column, selectCols []string) {
    for _, sel := range selectCols {
        val := ""
        found := false
        // Check first table
        for _, col := range cols1 {
            if sel == col.Name {
                val = fmt.Sprintf("%v", row1[col.Name])
                found = true
                break
            }
        }
        // Check second table
        if !found {
            for _, col := range cols2 {
                if sel == col.Name {
                    val = fmt.Sprintf("%v", row2[col.Name])
                    break
                }
            }
        }
        fmt.Printf("%-12s ", val)
    }
}