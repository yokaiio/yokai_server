package auto

import (
	"github.com/east-eden/server/excel"
	"github.com/east-eden/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

var skillTriggerEntries *SkillTriggerEntries //SkillTrigger.csv全局变量

// SkillTrigger.csv属性表
type SkillTriggerEntry struct {
	Id                int32             `json:"Id,omitempty"`                // 主键
	CheckUinit        int32             `json:"CheckUinit,omitempty"`        //条件检查单位
	CheckHeroId       []int32           `json:"CheckHeroId,omitempty"`       //条件检查英雄ID
	CheckSkillId      []int32           `json:"CheckSkillId,omitempty"`      //条件检查技能ID
	CheckBuffId       []int32           `json:"CheckBuffId,omitempty"`       //条件检查技能ID
	Checktime         int32             `json:"Checktime,omitempty"`         //触发时机类型
	CheckCondition    int32             `json:"CheckCondition,omitempty"`    //触发条件类型
	CheckIntParameter int32             `json:"CheckIntParameter,omitempty"` //整数型条件参数
	CheckParameter    []decimal.Decimal `json:"CheckParameter,omitempty"`    //条件参数数组
}

// SkillTrigger.csv属性表集合
type SkillTriggerEntries struct {
	Rows map[int32]*SkillTriggerEntry `json:"Rows,omitempty"` //
}

func init() {
	excel.AddEntryLoader("SkillTrigger.csv", (*SkillTriggerEntries)(nil))
}

func (e *SkillTriggerEntries) Load(excelFileRaw *excel.ExcelFileRaw) error {

	skillTriggerEntries = &SkillTriggerEntries{
		Rows: make(map[int32]*SkillTriggerEntry, 100),
	}

	for _, v := range excelFileRaw.CellData {
		entry := &SkillTriggerEntry{}
		err := mapstructure.Decode(v, entry)
		if !utils.ErrCheck(err, "decode excel data to struct failed", v) {
			return err
		}

		skillTriggerEntries.Rows[entry.Id] = entry
	}

	log.Info().Str("excel_file", excelFileRaw.Filename).Msg("excel load success")
	return nil

}

func GetSkillTriggerEntry(id int32) (*SkillTriggerEntry, bool) {
	entry, ok := skillTriggerEntries.Rows[id]
	return entry, ok
}

func GetSkillTriggerSize() int32 {
	return int32(len(skillTriggerEntries.Rows))
}

func GetSkillTriggerRows() map[int32]*SkillTriggerEntry {
	return skillTriggerEntries.Rows
}
