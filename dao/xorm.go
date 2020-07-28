package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)





func Connect()*xorm.Engine {
	x,err:=xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/xorm_test")
	if err=x.Ping();err!=nil{
		log.Printf("err :%v",err)
	}

	return x
}
func SelectUser(){
	conn:=Connect()
	conn.SetTableMapper(core.SnakeMapper{})
	user:=&User{}
	conn.Where("id=?",1).Get(user)
	log.Printf("user is :%v ",user)
}
func UpdateUser(){
	conn:=Connect()
	conn.SetTableMapper(core.SnakeMapper{})
	user:=&User{Username: "user2",Nickname: "zz5",Password: "1234"}
	affectRows, err :=conn.Where("id=?",2).Update(user)
	log.Printf("affectRows is :%v err is %v ",affectRows,err)
}
func InsertUser(){
	conn:=Connect()
	conn.SetTableMapper(core.SnakeMapper{})
	user:=&User{Username: "user2",Nickname: "zzz",Password: "123456"}
	affectRows, err := conn.Insert(user)
	log.Printf("affectRows is :%v err is %v ",affectRows,err)
}
func DeleteUser(){
	conn:=Connect()
	conn.SetTableMapper(core.SnakeMapper{})
	user:=&User{Username: "user2",Nickname: "zzz"}
	affectRows, err := conn.Delete(user)
	log.Printf("affectRows is :%v err is %v ",affectRows,err)
}
func SelectAll(){
	conn:=Connect()
	conn.SetTableMapper(core.SnakeMapper{})
	users:=make([]User,0)
	conn.Where("id>?",1).Find(&users)
	log.Printf("user is :%v len:%d ",users,len(users))
}
func SessionUse(){
	conn:=Connect()
	session := conn.NewSession()
	defer session.Close()
	 session.Begin()
	user:=&User{Username: "user4",Nickname: "zzz2",Password: "123456"}
	_, err := session.Insert(user)
	if err!=nil{
		session.Rollback()
	}
     session.Commit()
}