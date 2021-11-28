package dbCore

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type Test struct {
    Id int
    Name string
    Created_at string
}

func GetName() {
    db, err := sql.Open("mysql", "dbgo:Test_dbgo_12345@tcp(127.0.0.1:3306)/test_db")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM test")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    for rows.Next() {
        var test Test
        err := rows.Scan(&test.Id, &test.Name, &test.Created_at)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(test.Id, test.Name, test.Created_at)
    }
    
    err = rows.Err()
    if err != nil{
        panic(err.Error())
    }
}


func InsertName(args []string) {
    db, err := sql.Open("mysql", "dbgo:Test_dbgo_12345@tcp(127.0.0.1:3306)/test_db")
    if err != nil{
        panic(err.Error())
    }
    defer db.Close()

    for _, name := range args {
        stmtInsert, err := db.Prepare("INSERT INTO test(name) VALUES(?)")
        if err != nil{
            panic(err.Error())
        }
        defer stmtInsert.Close()

        result, err := stmtInsert.Exec(name)
        if err != nil{
            panic(err.Error())
        }

        lastInsertID, err := result.LastInsertId()
        if err != nil{
            panic(err.Error())
        }
        fmt.Println(lastInsertID)
    }
}

func DeleteName(args []string) {
    db, err := sql.Open("mysql", "dbgo:Test_dbgo_12345@tcp(127.0.0.1:3306)/test_db")
    if err != nil{
        panic(err.Error())
    }
    defer db.Close()
    
    for _, name := range args {
        stmtDelete, err := db.Prepare("DELETE FROM test WHERE name=?")
        if err != nil {
	    panic(err.Error())
        }
        defer stmtDelete.Close()

        result, err := stmtDelete.Exec(name)
        if err != nil {
	    panic(err.Error())
        }

        rowsAffect, err := result.RowsAffected()
        if err != nil {
	    panic(err.Error())
        }
        fmt.Println(rowsAffect)
    }
}
