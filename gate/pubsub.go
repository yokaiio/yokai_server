package gate

import (
	"context"

	"github.com/micro/go-micro"
	logger "github.com/sirupsen/logrus"
	pbAccount "github.com/yokaiio/yokai_server/proto/account"
	pbPubSub "github.com/yokaiio/yokai_server/proto/pubsub"
)

type PubSub struct {
	pubGateResult micro.Publisher
	g             *Gate
}

func NewPubSub(g *Gate) *PubSub {
	ps := &PubSub{
		g: g,
	}

	// create publisher
	ps.pubGateResult = micro.NewPublisher("gate.GateResult", g.mi.srv.Client())

	// register subscriber
	micro.RegisterSubscriber("game.StartGate", g.mi.srv.Server(), &subStartGate{g: g})
	micro.RegisterSubscriber("game.ExpirePlayer", g.mi.srv.Server(), &subExpirePlayer{g: g})

	return ps
}

/////////////////////////////////////
// publish handle
/////////////////////////////////////
func (ps *PubSub) PubGateResult(ctx context.Context, win bool) error {
	info := &pbAccount.LiteAccount{Id: 1, Name: "pub_client"}
	return ps.pubGateResult.Publish(ctx, &pbPubSub.PubGateResult{Info: info, Win: win})
}

/////////////////////////////////////
// subscribe handle
/////////////////////////////////////
type subStartGate struct {
	g *Gate
}

func (s *subStartGate) Process(ctx context.Context, event *pbPubSub.PubStartGate) error {
	logger.WithFields(logger.Fields{
		"event": event,
	}).Info("recv game.StartGate")
	return nil
}

type subExpirePlayer struct {
	g *Gate
}

func (s *subExpirePlayer) Process(ctx context.Context, event *pbPubSub.PubExpirePlayer) error {
	logger.WithFields(logger.Fields{
		"event": event,
	}).Info("recv game.ExpirePlayer")
	return nil
}