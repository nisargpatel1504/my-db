package storage

import (
	"encoding/json"
	"fmt"
	"my-db/model"
	"os"
)

func WriteTable(table model.Table) {
    fileName := fmt.Sprintf("%s.json", table.Name)
    data, err := json.MarshalIndent(table.Rows, "", " ")
    if err != nil {
        fmt.Println("Error while marshaling data:", err)
        return
    }

    // Open file with options: os.O_RDWR|os.O_CREATE|os.O_TRUNC
    // 0644 permissions: user can read/write, group and others can read.
    file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Write marshaled data to file
    _, err = file.Write(data)
    if err != nil {
        fmt.Println("Error writing to file:", err)
    }
}
