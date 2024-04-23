package main

import (
	"my-db/db"
	"my-db/model"
)

func main() {
    database := db.NewDatabase()
	//columns for user table
    columns := []model.Column{
        {Name: "id", Type: "int"},
        {Name: "name", Type: "string"},
		{Name: "age",Type:"int"},
    }
	//Columns for employee table
	employee := []model.Column{
		{Name:"empid",Type:"int"},
		{Name:"empsalary",Type:"int"},
		{Name:"userid",Type:"int"},
	}
	
    database.CreateTable("users", columns)
    database.Insert("users", model.Row{"id": 1, "name": "Alice", "age":24})
    database.Insert("users", model.Row{"id": 2, "name": "Bob","age":30})
	database.Insert("users", model.Row{"id": 3, "name": "Patel","age":26})
    database.Update("users", 1, model.Row{"name": "Updated Bob"})
    database.Delete("users", 0)
    database.Select("users")

	//Creating new employee table
	database.CreateTable("employee",employee)
	database.Insert("employee", model.Row{"empid": 1, "empsalary": 45000, "userid":2})
    database.Insert("employee", model.Row{"empid": 2, "empsalary": 25000, "userid":1})

	query := model.JoinQuery{
        Table1:   "users",
        Table2:   "employee",
        Select: []string{"id","name","age","empsalary"},
        JoinCol:  "id",
        WhereCol: "userid",
        WhereVal: 2,
    }
    database.JoinTables(query);
	
}
