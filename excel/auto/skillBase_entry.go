package auto

import (
	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

var skillBaseEntries *SkillBaseEntries //SkillBase.csv全局变量

// SkillBase.csv属性表
type SkillBaseEntry struct {
	Id              int32           `json:"Id,omitempty"`              // 主键
	SkillLv         int32           `json:"SkillLv,omitempty"`         //技能等级
	Name            string          `json:"Name,omitempty"`            //名称
	Scopedesc       string          `json:"Scopedesc,omitempty"`       //技能范围
	Desc            string          `json:"Desc,omitempty"`            //描述
	EndMoveScope    decimal.Decimal `json:"EndMoveScope,omitempty"`    //随机移动
	Type            int32           `json:"Type,omitempty"`            //类型
	Icon            string          `json:"Icon,omitempty"`            //技能图标
	WaitAct         string          `json:"WaitAct,omitempty"`         //act动作
	ActAnimate      string          `json:"ActAnimate,omitempty"`      //Act触发
	ContinueAnimate string          `json:"ContinueAnimate,omitempty"` //Act持续
	PointingShow    string          `json:"PointingShow,omitempty"`    //指向表现配置
	SpeedDesc       string          `json:"SpeedDesc,omitempty"`       //描述
	AtbSpeed        decimal.Decimal `json:"AtbSpeed,omitempty"`        //Act条速度
	Rage            int32           `json:"Rage,omitempty"`            //怒气减少
	CostMP          int32           `json:"CostMP,omitempty"`          //MP消耗
	FirstCD         int32           `json:"FirstCD,omitempty"`         //初始CD
	GeneralCD       int32           `json:"GeneralCD,omitempty"`       //回转CD
	Limit           int32           `json:"Limit,omitempty"`           //次数限制
	Range           decimal.Decimal `json:"Range,omitempty"`           //施法范围
	TargetType      int32           `json:"TargetType,omitempty"`      //目标类型
	SkillLaunch     int32           `json:"SkillLaunch,omitempty"`     //发起类型
	TargetLength    decimal.Decimal `json:"TargetLength,omitempty"`    //范围长
	TargetWide      decimal.Decimal `json:"TargetWide,omitempty"`      //范围宽
	RangeType       int32           `json:"RangeType,omitempty"`       //目标范围
	Scope           int32           `json:"Scope,omitempty"`           //作用对象
	TimelineID      []int32         `json:"TimelineID,omitempty"`      //表现ID
	TriggerId       []int32         `json:"TriggerId,omitempty"`       //触发器类型
	TriggerRelation int32           `json:"TriggerRelation,omitempty"` //触发器条件关系
}

// SkillBase.csv属性表集合
type SkillBaseEntries struct {
	Rows map[int32]*SkillBaseEntry `json:"Rows,omitempty"` //
}

func init() {
	excel.AddEntryLoader("SkillBase.csv", (*SkillBaseEntries)(nil))
}

func (e *SkillBaseEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {

	skillBaseEntries = &SkillBaseEntries{
		Rows: make(map[int32]*SkillBaseEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &SkillBaseEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

		skillBaseEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil

}

func GetSkillBaseEntry(id int32) (*SkillBaseEntry, bool) {
	entry, ok := skillBaseEntries.Rows[id]
	return entry, ok
}

func GetSkillBaseSize() int32 {
	return int32(len(skillBaseEntries.Rows))
}

func GetSkillBaseRows() map[int32]*SkillBaseEntry {
	return skillBaseEntries.Rows
}
