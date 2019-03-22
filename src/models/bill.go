package models

type Bill struct {
	Id          int       `orm:(id)`
	UserInfoId  int `orm:(user_info_id)`
	Goods       string    `orm:(goods)`
	OrderNum    string    `orm:(order_num)`
	CreateTime  string    `orm:(create_time)`
	Description string    `orm:(description)`
}
