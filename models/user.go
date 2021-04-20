package models

type User struct {
	tableName struct{} `pg:"user_table"`
	Id        string   `json:"id"`
	Password  string   `json:"password"`
	Role      string   `json:"role"`
}
