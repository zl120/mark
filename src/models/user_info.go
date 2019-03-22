package models

type UserInfo struct {
	Id          int    `orm:(id)`
	Name        string `orm:(name)`
	Age         int    `orm:(age)`
	Gender      string `orm:(gender)`
	IdCard      string `orm:(id_card)`
	Status      int    `orm:(status)`
	Description string `orm:(description)`
	Tel         string `orm:(tel)`
}
