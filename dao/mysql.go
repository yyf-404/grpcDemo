package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type  Employee struct{
	 last_name string
	 id int
	 gender string
	 did int
	 email string
}
func ConnectMysql(){
	db, err := sql.Open("mysql", "root:123456.123@tcp(101.37.79.100:3306)/cache")
	if err!=nil{
		panic(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	var employee Employee
	rows, e := db.Query("selecat * from employee ")
	if e != nil {
		panic(errors.New("query incur error"))
	}
	for rows.Next(){
		e := rows.Scan(&employee.id,&employee.last_name,&employee.email, &employee.gender, &employee.did)
		if e != nil{
		panic(e)
		}
		fmt.Println(employee)
	}
	rows.Close()
	//单行查询操作
	//db.QueryRow("select * from user where id=1").Scan(
}
