# Basic RDBMS in Go

This project demonstrates a basic Relational Database Management System (RDBMS) built in Go. It supports fundamental operations such as CREATE, READ, UPDATE, and DELETE (CRUD), along with basic JOIN functionality. The system uses in-memory storage for simplicity and is intended as an educational tool to understand the principles behind building a database management system.

## Features

- Create tables with dynamic columns
- Insert data into tables
- Select data with basic filtering and join capabilities
- Update and delete data
- In-memory data storage
- Basic data persistence using JSON files

## Getting Started

### Prerequisites

- Go (version 1.16 or later recommended) installed on your machine
- Basic knowledge of Go programming and concepts

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/basic-rdbms-go.git
cd basic-rdbms-go
```

###To run the application, use the following command from the root of the project directory:

go run main.go

###USAGE

##Creating a Table

- To create a table, define the table structure with columns and types, and call the CreateTable method:
```bash
db.CreateTable("users", []model.Column{
    {Name: "id", Type: "int"},
    {Name: "name", Type: "string"},
})
```

##Inserting Data
Insert data into the table by specifying the table name and data in key-value pairs:
```bash
db.Insert("users", model.Row{"id": 1, "name": "Alice"})
```
##Selecting Data
- Retrieve data from a table, optionally using conditions for filtering:

```bash
query := model.JoinQuery{
    Table1:   "users",
    Table2:   "orders",
    Select:   []string{"users.name", "orders.order_info"},
    JoinCol:  "user_id",
    WhereCol: "users.id",
    WhereVal: 1,
}
db.JoinTables(query)
```

##Updating Data
Update data in a table by specifying the table name, row index, and new data:

```bash
db.Update("users", 0, model.Row{"name": "Updated Alice"})
```

###Deleting Data
Delete data by specifying the table name and row index:
```bash
db.Delete("users", 0)
```
