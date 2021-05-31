package game

import (
	"context"
	"errors"
	"fmt"

	pbGlobal "github.com/east-eden/server/proto/global"
	"github.com/east-eden/server/services/game/player"
)

func (m *MsgRegister) handleDelHero(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_DelHero)
	if !ok {
		return errors.New("handelDelHero failed: recv message body error")
	}
	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleDelHero.AccountExecute failed: %w", err)
	}

	pl.HeroManager().DelHero(msg.Id)
	return nil
}

func (m *MsgRegister) handleQueryHeroAtt(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_QueryHeroAtt)
	if !ok {
		return errors.New("handelQueryHeroAtt failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleQueryHeroAtt failed: %w", err)
	}

	h := pl.HeroManager().GetHero(msg.HeroId)
	if h == nil {
		return fmt.Errorf("handleQueryHeroAtt failed: cannot find hero<%d>", msg.HeroId)
	}

	h.GetAttManager().CalcAtt()
	pl.HeroManager().SendHeroAtt(h)
	return nil
}

func (m *MsgRegister) handleHeroLevelup(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_HeroLevelup)
	if !ok {
		return errors.New("handelHeroLevelup failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleHeroLevelup failed: %w", err)
	}

	return pl.HeroManager().HeroLevelup(msg.GetHeroId(), msg.GetStuffItems())
}

func (m *MsgRegister) handleHeroPromote(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_HeroPromote)
	if !ok {
		return errors.New("handleHeroPromote failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleHeroPromote failed: %w", err)
	}

	return pl.HeroManager().HeroPromote(msg.GetHeroId())
}

func (m *MsgRegister) handleHeroStarup(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_HeroStarup)
	if !ok {
		return errors.New("handleHeroStarup failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleHeroStarup failed: %w", err)
	}

	return pl.HeroManager().HeroStarup(msg.GetHeroId())
}

func (m *MsgRegister) handleHeroTalentChoose(ctx context.Context, p ...interface{}) error {
	acct := p[0].(*player.Account)
	msg, ok := p[1].(*pbGlobal.C2S_HeroTalentChoose)
	if !ok {
		return errors.New("handleHeroTalentChoose failed: recv message body error")
	}

	pl, err := m.am.GetPlayerByAccount(acct)
	if err != nil {
		return fmt.Errorf("handleHeroTalentChoose failed: %w", err)
	}

	return pl.HeroManager().HeroTalentChoose(msg.GetHeroId(), msg.GetTalentId())
}
