package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"reflect"
)

func main() {
	dns := "root:123456@tcp(192.168.196.16:3307)/game_db"

	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	count := 0
	err = db.Get(&count, "SELECT COUNT(*) FROM account_data")
	if err != nil {
		panic(err)
	}

	var ids []int
	test(ids)
	err = db.Select(&ids, "SELECT user_id FROM account_data LIMIT 10")
	if err != nil {
		panic("ids: " + err.Error())
	}
	fmt.Println(ids)

	var list []AccountData
	err = db.Select(&list, "SELECT * FROM account_data LIMIT 10")
	if err != nil {
		panic("ids: " + err.Error())
	}
	fmt.Println(list)
}

type PlatformType int32

type EquipmentType int32

type AccountData struct {
	// @inject_tag: json:"user_id" db:"user_id"
	UserId uint32 `protobuf:"varint,1,opt,name=userId,proto3" json:"user_id" db:"user_id"` // 用户唯一 Id
	// @inject_tag: json:"platform_type" db:"platform_type"
	PlatformType PlatformType `protobuf:"varint,2,opt,name=platformType,proto3,enum=data.PlatformType" json:"platform_type" db:"platform_type"`
	// @inject_tag: json:"platform_identity" db:"platform_identity"
	PlatformIdentity string `protobuf:"bytes,3,opt,name=platformIdentity,proto3" json:"platform_identity" db:"platform_identity"`
	// @inject_tag: json:"version" db:"version"
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version" db:"version"`
	// @inject_tag: json:"tunnel" db:"tunnel"
	Tunnel string `protobuf:"bytes,5,opt,name=tunnel,proto3" json:"tunnel" db:"tunnel"`
	// @inject_tag: json:"device" db:"device"
	Device string `protobuf:"bytes,6,opt,name=device,proto3" json:"device" db:"device"`
	// @inject_tag: json:"equipment_type" db:"equipment_type"
	EquipmentType EquipmentType `protobuf:"varint,7,opt,name=equipmentType,proto3,enum=data.EquipmentType" json:"equipment_type" db:"equipment_type"`
	// @inject_tag: json:"equipment_identity" db:"equipment_identity"
	EquipmentIdentity string `protobuf:"bytes,8,opt,name=equipmentIdentity,proto3" json:"equipment_identity" db:"equipment_identity"`
	// @inject_tag: json:"country" db:"country"
	Country string `protobuf:"bytes,9,opt,name=country,proto3" json:"country" db:"country"`
	// @inject_tag: json:"time" db:"time"
	Time string `protobuf:"bytes,10,opt,name=time,proto3" json:"time" db:"time"` // 时间格式
	// @inject_tag: json:"ip" db:"ip"
	Ip string `protobuf:"bytes,11,opt,name=ip,proto3" json:"ip" db:"ip"` // ip
	// @inject_tag: json:"state" db:"state"
	State uint32 `protobuf:"bytes,12,opt,name=state,proto3" json:"state" db:"state"` // state
	// @inject_tag: json:"is_bind" db:"is_bind"
	IsBind bool `protobuf:"varint,13,opt,name=isBind,proto3" json:"is_bind" db:"is_bind"`
	// @inject_tag: json:"channel" db:"channel"
	Channel uint32 `protobuf:"varint,14,opt,name=channel,proto3" json:"channel" db:"channel"`
}

func test(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}
