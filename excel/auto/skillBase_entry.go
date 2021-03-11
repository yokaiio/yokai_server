package auto

import (
	"bitbucket.org/funplus/server/excel"
	"bitbucket.org/funplus/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

var	skillBaseEntries	*SkillBaseEntries	//SkillBase.xlsx全局变量

// SkillBase.xlsx属性表
type SkillBaseEntry struct {
	Id             	int32               	`json:"Id,omitempty"`	// 主键       
	SkillLv        	int32               	`json:"SkillLv,omitempty"`	//技能等级      
	Name           	int32               	`json:"Name,omitempty"`	//名称        
	Desc           	int32               	`json:"Desc,omitempty"`	//描述        
	Icon           	string              	`json:"Icon,omitempty"`	//图标        
	WaitAct        	string              	`json:"WaitAct,omitempty"`	//act动作     
	BeforeSkillAct 	string              	`json:"BeforeSkillAct,omitempty"`	//动作前摇      
	SkillAct       	string              	`json:"SkillAct,omitempty"`	//动作过程      
	AfterSkillAct  	string              	`json:"AfterSkillAct,omitempty"`	//动作后摇      
	Performance    	string              	`json:"Performance,omitempty"`	//动画        
	Type           	int32               	`json:"Type,omitempty"`	//类型        
	AtbSpeed       	int32               	`json:"AtbSpeed,omitempty"`	//Act条速度    
	Rage           	int32               	`json:"Rage,omitempty"`	//怒气增减      
	CostMP         	int32               	`json:"CostMP,omitempty"`	//MP消耗      
	FirstCD        	int32               	`json:"FirstCD,omitempty"`	//初始CD      
	GeneralCD      	int32               	`json:"GeneralCD,omitempty"`	//回转CD      
	Limit          	int32               	`json:"Limit,omitempty"`	//次数限制      
	Range          	int32               	`json:"Range,omitempty"`	//施法范围      
	RangeType      	int32               	`json:"RangeType,omitempty"`	//目标范围      
	TargetType     	int32               	`json:"TargetType,omitempty"`	//目标类型      
	TargetLength   	int32               	`json:"TargetLength,omitempty"`	//范围长       
	TargetWide     	int32               	`json:"TargetWide,omitempty"`	//范围宽       
	Scope          	int32               	`json:"Scope,omitempty"`	//作用对象      
	Effects        	[]int32             	`json:"Effects,omitempty"`	//效果逻辑      
}

// SkillBase.xlsx属性表集合
type SkillBaseEntries struct {
	Rows           	map[int32]*SkillBaseEntry	`json:"Rows,omitempty"`	//          
}

func  init()  {
	excel.AddEntryLoader("SkillBase.xlsx", (*SkillBaseEntries)(nil))
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

func  GetSkillBaseEntry(id int32) (*SkillBaseEntry, bool) {
	entry, ok := skillBaseEntries.Rows[id]
	return entry, ok
}

func  GetSkillBaseSize() int32 {
	return int32(len(skillBaseEntries.Rows))
}

func  GetSkillBaseRows() map[int32]*SkillBaseEntry {
	return skillBaseEntries.Rows
}
