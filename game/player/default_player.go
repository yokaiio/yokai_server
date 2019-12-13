package player

import (
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/logger"
	"github.com/yokaiio/yokai_server/game/blade"
	"github.com/yokaiio/yokai_server/game/db"
	"github.com/yokaiio/yokai_server/game/hero"
	"github.com/yokaiio/yokai_server/game/item"
	"github.com/yokaiio/yokai_server/game/token"
	"github.com/yokaiio/yokai_server/internal/define"
	"github.com/yokaiio/yokai_server/internal/global"
	"github.com/yokaiio/yokai_server/internal/utils"
)

type DefaultPlayer struct {
	ds *db.Datastore
	wg utils.WaitGroupWrapper

	itemManager  *item.ItemManager
	heroManager  *hero.HeroManager
	tokenManager *token.TokenManager
	bladeManager *blade.BladeManager

	ID       int64  `gorm:"type:bigint(20);primary_key;column:id;default:-1;not null"`
	ClientID int64  `gorm:"type:bigint(20);column:client_id;default:-1;not null"`
	Name     string `gorm:"type:varchar(32);column:name;not null"`
	Exp      int64  `gorm:"type:bigint(20);column:exp;default:0;not null"`
	Level    int32  `gorm:"type:int(10);column:level;default:1;not null"`
}

func newDefaultPlayer(id int64, name string, ds *db.Datastore) Player {
	p := &DefaultPlayer{
		ds:       ds,
		ID:       id,
		ClientID: -1,
		Name:     name,
		Exp:      0,
		Level:    1,
	}

	p.itemManager = item.NewItemManager(p, ds)
	p.heroManager = hero.NewHeroManager(p, ds)
	p.tokenManager = token.NewTokenManager(p, ds)
	p.bladeManager = blade.NewBladeManager(p, ds)

	return p
}

func defaultMigrate(ds *db.Datastore) {
	ds.ORM().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(DefaultPlayer{})
	item.Migrate(ds)
	hero.Migrate(ds)
	token.Migrate(ds)
	blade.Migrate(ds)
}

func (p *DefaultPlayer) TableName() string {
	return "player"
}

func (p *DefaultPlayer) GetType() int32 {
	return define.Plugin_Player
}

func (p *DefaultPlayer) GetID() int64 {
	return p.ID
}

func (p *DefaultPlayer) GetLevel() int32 {
	return p.Level
}

func (p *DefaultPlayer) GetClientID() int64 {
	return p.ClientID
}

func (p *DefaultPlayer) GetName() string {
	return p.Name
}

func (p *DefaultPlayer) GetExp() int64 {
	return p.Exp
}

func (p *DefaultPlayer) SetClientID(id int64) {
	p.ClientID = id
}

func (p *DefaultPlayer) SetName(name string) {
	p.Name = name
}

func (p *DefaultPlayer) SetExp(exp int64) {
	p.Exp = exp
}

func (p *DefaultPlayer) SetLevel(level int32) {
	p.Level = level
}

func (p *DefaultPlayer) HeroManager() *hero.HeroManager {
	return p.heroManager
}

func (p *DefaultPlayer) ItemManager() *item.ItemManager {
	return p.itemManager
}

func (p *DefaultPlayer) TokenManager() *token.TokenManager {
	return p.tokenManager
}

func (p *DefaultPlayer) BladeManager() *blade.BladeManager {
	return p.bladeManager
}

func (p *DefaultPlayer) LoadFromDB() {
	p.wg.Wrap(p.heroManager.LoadFromDB)
	p.wg.Wrap(p.itemManager.LoadFromDB)
	p.wg.Wrap(p.tokenManager.LoadFromDB)
	p.wg.Wrap(p.bladeManager.LoadFromDB)
	p.wg.Wait()
}

func (p *DefaultPlayer) AfterLoad() {
	items := p.itemManager.GetItemList()
	for _, v := range items {
		if v.GetEquipObj() == -1 {
			continue
		}

		if err := p.HeroManager().PutonEquip(v.GetEquipObj(), v.GetID(), v.Entry().EquipPos); err != nil {
			logger.Warn("Hero puton equip error when loading from db:", err)
		}
	}
}

func (p *DefaultPlayer) Save() {
	p.ds.ORM().Save(p)
}

func (p *DefaultPlayer) ChangeExp(add int64) {
	if p.Level >= define.Player_MaxLevel {
		return
	}

	// overflow
	if (p.Exp + add) < 0 {
		return
	}

	p.Exp += add
	for {
		levelupEntry := global.GetPlayerLevelupEntry(p.Level + 1)
		if levelupEntry == nil {
			break
		}

		if p.Exp < levelupEntry.Exp {
			break
		}

		p.Exp -= levelupEntry.Exp
		p.Level++
	}

	p.heroManager.HeroSetLevel(p.Level)
	p.ds.ORM().Model(p).Updates(DefaultPlayer{
		Exp:   p.Exp,
		Level: p.Level,
	})
}

func (p *DefaultPlayer) ChangeLevel(add int32) {
	if p.Level >= define.Player_MaxLevel {
		return
	}

	nextLevel := p.Level + add
	if nextLevel > define.Player_MaxLevel {
		nextLevel = define.Player_MaxLevel
	}

	levelupEntry := global.GetPlayerLevelupEntry(nextLevel)
	if levelupEntry == nil {
		return
	}

	p.Level = nextLevel
	p.ds.ORM().Model(p).Updates(DefaultPlayer{
		Level: p.Level,
	})

	p.heroManager.HeroSetLevel(p.Level)
}
