package models

import "database/sql"

type Contract struct {
    Id      int64  `json:"-" db:"id"`
    Address string `json:"address" db:"address"`
    ABI     string `json:"abi" db:"abi"`
    Name    string `json:"name" db:"name"`
}

func (c *Contract) Insert() (sql.Result, error) {
    return DB.Exec("insert into contract (address,abi,`name`) values (?,?,?)", c.Address, c.ABI, c.Name)
}

func GetContractByAddr(addr string) (*Contract, error) {
    var c = Contract{}
    err := DB.Get(&c, "select * from contract where address=?", addr)
    return &c, err
}

func GetContractByName(name string) (*Contract, error) {
    var c = Contract{}
    err := DB.Get(&c, "select * from contract where name=?", name)
    return &c, err
}

func GetContractById(id int64) (*Contract, error) {
    var c = Contract{}
    err := DB.Get(&c, "select * from contract where id=?", id)
    return &c, err
}
