// --- this file generated by tools, please write your code in custom block ---
package tbobj

import (
	"trcell/pkg/ormdef"
	"trcell/pkg/pborm"
	"google.golang.org/protobuf/proto"
)

type TbRoleBag struct{
	TableObjOpt
	ormMeta *ormdef.RoleBag
	// ------------------ customtag1 code begin ---------------------
	bagData ITableBlobField
	// ------------------ customtag1 code end -----------------------
}

// RoleID
func (tb *TbRoleBag) GetRoleID() int64 {
	return tb.ormMeta.RoleID
}
func (tb *TbRoleBag) SetRoleID(roleID int64) {
	tb.ormMeta.RoleID = roleID
	tb.SetDbStatus(DbStatusUpdate)
}

// Bag
func (tb *TbRoleBag) GetBag() []byte {
	return tb.ormMeta.Bag
}
func (tb *TbRoleBag) SetBag(bag []byte) {
	tb.ormMeta.Bag = bag
	tb.SetDbStatus(DbStatusUpdate)
}

// UpdTime
func (tb *TbRoleBag) GetUpdTime() int64 {
	return tb.ormMeta.UpdTime
}
func (tb *TbRoleBag) SetUpdTime(updTime int64) {
	tb.ormMeta.UpdTime = updTime
	tb.SetDbStatus(DbStatusUpdate)
}

func (tb *TbRoleBag) DumpToPbOrm() *pborm.RoleBag {
	tb.Serialize()
	pbObj := &pborm.RoleBag{
		RoleID: tb.ormMeta.RoleID,
		Bag: tb.ormMeta.Bag,
		UpdTime: tb.ormMeta.UpdTime,
	}
	pbObj.Bag = make([]byte, len(tb.ormMeta.Bag))
	copy(pbObj.Bag, tb.ormMeta.Bag)
	pbObj.DbStatus = int32(tb.dbStatus)
	return pbObj
}
func (tb *TbRoleBag) LoadFromPbOrm(obj *pborm.RoleBag) {
	tb.ormMeta.RoleID = obj.RoleID
	tb.ormMeta.Bag = obj.Bag
	tb.ormMeta.UpdTime = obj.UpdTime
	tb.ormMeta.Bag = make([]byte, len(obj.Bag))
	copy(tb.ormMeta.Bag, obj.Bag)
	tb.dbStatus = int8(obj.DbStatus)
	tb.Unserialize()
}
func NewTbRoleBag() *TbRoleBag {
	newObj := &TbRoleBag{
		ormMeta: &ormdef.RoleBag{
			RoleID: 0,
			Bag: make([]byte, 0),
			UpdTime: 0,
		},
	}
	// ------------------ customtag2 code begin ---------------------
	// ------------------ customtag2 code end -----------------------
	return newObj
}
func (tb *TbRoleBag) Serialize() {
	// ------------------ customtag3 code begin ---------------------
	if tb.bagData != nil {
		tb.SetBag(tb.bagData.ToBytes())
	}
	// ------------------ customtag3 code end -----------------------
}
func (tb *TbRoleBag) Unserialize() {
	// ------------------ customtag4 code begin ---------------------
	if tb.bagData != nil {
		if tb.GetBag() != nil {
			tb.bagData.FromBytes(tb.GetBag())
		}
	}
	// ------------------ customtag4 code end -----------------------
}
func (tb *TbRoleBag) ToBytes() ([]byte, error) {
	pbObj := tb.DumpToPbOrm()
	return proto.Marshal(pbObj)
}
func (tb *TbRoleBag) FromBytes(data []byte) error {
	pbOrm := &pborm.RoleBag{}
	err := proto.Unmarshal(data, pbOrm)
	if err != nil {
		return err
	}
	tb.LoadFromPbOrm(pbOrm)
	return nil
}
func (tb *TbRoleBag) GetOrmMeta() *ormdef.RoleBag {
	return tb.ormMeta
}
// ------------------ customtag5 code begin ---------------------
func (tb *TbRoleBag) BindModule(fieldModule ITableBlobField) {
	tb.bagData = fieldModule
}

// ------------------ customtag5 code end -----------------------
