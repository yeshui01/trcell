// --- this file generated by tools, don't edit it!!!! ---
package ormdef

type CsGlobal struct {
	ID      	int32	`gorm:"primary_key" json:"id"`
	OpenTime	int64	`gorm:"open_time" json:"open_time"`
}
func (t *CsGlobal) TableName() string {
	return "cs_global"
}