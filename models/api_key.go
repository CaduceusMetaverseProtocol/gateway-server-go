package models

import (
    "database/sql"
)

// ApiKey is uniquely identified for REST
type ApiKey struct {
    Id        int64  `json:"-" db:"id"`
    Memo      string `json:"memo" db:"memo"`
    AccessKey string `json:"access_key" db:"access_key"`
    SecretKey string `json:"secret_key,omitempty" db:"secret_key"`
    CreatedAt int64  `json:"created_at" db:"created_at"`
    UpdatedAt int64  `json:"updated_at" db:"updated_at"`
    AccessIPs string `json:"access_ips" db:"access_ips"`
}

func (key *ApiKey) Insert() (sql.Result, error) {
    return DB.Exec("insert into api_key (memo,access_key,secret_key,created_at,access_ips) values (?,?,?,?,?)",
        key.Memo, key.AccessKey, key.SecretKey, key.CreatedAt, key.AccessIPs)
}

func GetApiKeyByAk(ak string) (*ApiKey, error) {
    var key = ApiKey{}
    err := DB.Get(&key, "select * from api_key where access_key=?", ak)
    return &key, err
}
