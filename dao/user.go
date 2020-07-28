package dao
type User struct {
	Id int32  /*`xorm:"INT(11)"`*/
	Username string /*`xorm:"VARCHAR(15)"`*/
	Password string /*`xorm:"VARCHAR(32)"`*/
	Nickname string /*`xorm:"VARCHAR(15)"`*/
}