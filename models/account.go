package models

import "database/sql"

type Account struct {
	Id      int64  `json:"-" db:"id"`
	Address string `json:"address" db:"address"`
	Name    string `json:"name" db:"name"`  //userName
}

func (a *Account) Insert() (sql.Result, error) {
	return DB.Exec("insert into account (address,`name`) values (?,?)", a.Address, a.Name)
}

func GetAccountByName(name string) (*Account, error) {
	var a = Account{}
	err := DB.Get(&a, "select * from account where `name`=?", name)
	return &a, err
}

func GetAccountById(id int64) (*Account, error) {
	var a = Account{}
	err := DB.Get(&a, "select * from account where id=?", id)
	return &a, err
}
func GetAccountByAddress(address string) (*Account, error) {
	var a = Account{}
	err := DB.Get(&a, "select * from account where address=?", address)
	return &a, err
}
