package item

import "github.com/yokaiio/yokai_server/define"

type Option func(*Options)

// item options
type Options struct {
	Id      int64 `bson:"_id" redis:"_id"`
	OwnerId int64 `bson:"owner_id" redis:"owner_id"`
	TypeId  int32 `bson:"type_id" redis:"type_id"`
	Num     int32 `bson:"num" redis:"num"`

	EquipObj          int64                     `bson:"equip_obj" redis:"equip_obj"`
	Entry             *define.ItemEntry         `bson:"-" redis:"-"`
	EquipEnchantEntry *define.EquipEnchantEntry `bson:"-" redis:"-"`
}

func DefaultOptions() *Options {
	return &Options{
		Id:                -1,
		OwnerId:           -1,
		TypeId:            -1,
		Num:               0,
		EquipObj:          -1,
		Entry:             nil,
		EquipEnchantEntry: nil,
	}
}

func Id(id int64) Option {
	return func(o *Options) {
		o.Id = id
	}
}

func OwnerId(id int64) Option {
	return func(o *Options) {
		o.OwnerId = id
	}
}

func TypeId(id int32) Option {
	return func(o *Options) {
		o.TypeId = id
	}
}

func Num(n int32) Option {
	return func(o *Options) {
		o.Num = n
	}
}

func EquipObj(obj int64) Option {
	return func(o *Options) {
		o.EquipObj = obj
	}
}

func Entry(entry *define.ItemEntry) Option {
	return func(o *Options) {
		o.Entry = entry
	}
}

func EquipEnchantEntry(entry *define.EquipEnchantEntry) Option {
	return func(o *Options) {
		o.EquipEnchantEntry = entry
	}
}
