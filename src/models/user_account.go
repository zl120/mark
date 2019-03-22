package models

type UserAccount struct {
	Id          int    `orm:(id)`
	UserName    string `orm:(username)`
	Pwd         string `orm:(pwd)`
	UserInfoId  int    `orm:(user_info_id)`
	Status      int    `orm:(status)`
	Description string `orm:(string)`
	Role        int    `orm:(role)`
}
