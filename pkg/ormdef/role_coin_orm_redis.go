// --- this file generated by tools, don't edit it!!!! ---
package ormdef

type RdsRoleCoin struct {
	RoleID 	int64	`redis:"RoleID"`
	Coin   	string	`redis:"Coin"`
	UpdTime	int64	`redis:"UpdTime"`
	DbStatus	int8	`redis:"DbStatus"`
}
