package hero

import (
	"sync"

	"github.com/east-eden/server/internal/att"
	"github.com/east-eden/server/services/game/item"
	"github.com/east-eden/server/services/game/rune"
	"github.com/east-eden/server/store"
)

// hero create pool
var heroPool = &sync.Pool{New: newPoolHeroV1}

func NewPoolHero() Hero {
	return heroPool.Get().(Hero)
}

func GetHeroPool() *sync.Pool {
	return heroPool
}

func ReleasePoolHero(x interface{}) {
	heroPool.Put(x)
}

type Hero interface {
	store.StoreObjector

	GetOptions() *Options
	GetEquipBar() *item.EquipBar
	GetAttManager() *att.AttManager
	GetRuneBox() *rune.RuneBox

	AddExp(int64) int64
	AddLevel(int32) int32
	BeforeDelete()
	CalcAtt()
}

func NewHero(opts ...Option) Hero {
	h := NewPoolHero()

	for _, o := range opts {
		o(h.GetOptions())
	}

	return h
}
