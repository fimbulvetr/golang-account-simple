package entity

type Account struct {
	ID        string `xorm:"varchar(50)"`
	Username  string `xorm:"varchar(50)"`
	Type      string `xorm:"varchar(50)"`
}