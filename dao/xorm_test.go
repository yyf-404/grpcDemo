package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)





func TestConnect(t *testing.T) {
	Connect()
}
func TestSelectUser(t *testing.T) {
	SelectUser()
}
func TestInsertUser(t *testing.T) {
	InsertUser()
}
func TestUpdateUser(t *testing.T) {
	UpdateUser()
}
func TestDeleteUser(t *testing.T) {
	DeleteUser()
}
func TestSelectAll(t *testing.T) {
	SelectAll()
}
func TestSessionUse(t *testing.T) {
	SessionUse()
}