package model

type Column struct {
    Name string
    Type string  // Types like 'int', 'string', etc.
}

type Row map[string]interface{}  // Each row is a map of column names to values

type JoinQuery struct {
    Table1    string
    Table2    string
    Select    []string// Column on which to select
    JoinCol   string  // Column on which to join
    WhereCol  string  // Column to apply the where condition
    WhereVal  interface{}  // Value for the where condition
}
type Table struct {
    Name    string
    Columns []Column
    Rows    []Row
}
