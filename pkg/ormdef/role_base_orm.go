// --- this file generated by tools, don't edit it!!!! ---
package ormdef

type RoleBase struct {
	RoleID     	int64	`gorm:"primary_key" json:"role_id"`
	UserID     	int64	`gorm:"user_id" json:"user_id"`
	RoleName   	string	`gorm:"role_name" json:"role_name"`
	CreateTime 	int64	`gorm:"create_time" json:"create_time"`
	Level      	int32	`gorm:"level" json:"level"`
	LoginTime  	int64	`gorm:"login_time" json:"login_time"`
	OfflineTime	int64	`gorm:"offline_time" json:"offline_time"`
	UpdTime    	int64	`gorm:"upd_time" json:"upd_time"`
}
func (t *RoleBase) TableName() string {
	return "role_base"
}