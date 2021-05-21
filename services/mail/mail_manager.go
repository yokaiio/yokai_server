package mail

import (
	"context"
	"errors"
	"runtime/debug"
	"sync"
	"time"

	"bitbucket.org/funplus/server/define"
	"bitbucket.org/funplus/server/services/mail/mailbox"
	"bitbucket.org/funplus/server/store"
	"bitbucket.org/funplus/server/utils"
	"bitbucket.org/funplus/server/utils/cache"
	"github.com/hellodudu/task"
	log "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var (
	mailBoxCacheExpire = 10 * time.Minute // 邮箱cache缓存10分钟
	ErrInvalidOwner    = errors.New("invalid owner")
)

type MailManager struct {
	m              *Mail
	cacheMailBoxes *cache.Cache
	mailBoxPool    sync.Pool
	wg             utils.WaitGroupWrapper
}

func NewMailManager(ctx *cli.Context, m *Mail) *MailManager {
	manager := &MailManager{
		m:              m,
		cacheMailBoxes: cache.New(mailBoxCacheExpire, mailBoxCacheExpire),
	}

	// 邮箱池
	manager.mailBoxPool.New = mailbox.NewMailBox

	// 邮箱缓存删除时处理
	manager.cacheMailBoxes.OnEvicted(func(k, v interface{}) {
		v.(*mailbox.MailBox).Stop()
		manager.mailBoxPool.Put(v)
	})

	// 初始化db
	store.GetStore().AddStoreInfo(define.StoreType_Mail, "mail", "_id")
	if err := store.GetStore().MigrateDbTable("mail", "owner_id", "mail_list._id"); err != nil {
		log.Fatal().Err(err).Msg("migrate collection mail failed")
	}

	log.Info().Msg("MailManager init ok ...")
	return manager
}

func (m *MailManager) Run(ctx *cli.Context) error {
	<-ctx.Done()
	log.Info().Msg("MailManager context done...")
	return nil
}

func (m *MailManager) Exit(ctx *cli.Context) {
	m.wg.Wait()
	log.Info().Msg("MailManager exit...")
}

// 获取邮箱数据
func (m *MailManager) getMailBox(ownerId int64) (*mailbox.MailBox, error) {
	if ownerId == -1 {
		return nil, ErrInvalidOwner
	}

	mb, ok := m.cacheMailBoxes.Get(ownerId)

	// 缓存没有，从db加载
	if !ok {
		mb = m.mailBoxPool.Get()
		mailbox := mb.(*mailbox.MailBox)
		mailbox.Init(m.m.ID)
		err := mailbox.Load(ownerId)
		if !utils.ErrCheck(err, "mailbox Load failed when MailManager.getMailBox", ownerId) {
			m.mailBoxPool.Put(mb)
			return nil, err
		}

		m.cacheMailBoxes.Set(ownerId, mb, mailBoxCacheExpire)
	}

	m.wg.Wrap(func() {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				log.Error().Msgf("catch exception:%v, panic recovered with stack:%s", err, stack)
			}

			// 立即删除缓存
			m.cacheMailBoxes.Delete(mb.(*mailbox.MailBox).Id)
		}()

		ctx := utils.WithSignaledCancel(context.Background())

		mb.(*mailbox.MailBox).InitTask()
		mb.(*mailbox.MailBox).ResetTaskTimeout()
		err := mb.(*mailbox.MailBox).TaskRun(ctx)
		utils.ErrPrint(err, "mailbox run failed", mb.(*mailbox.MailBox).Id)
	})

	return mb.(*mailbox.MailBox), nil
}

func (m *MailManager) AddTask(ctx context.Context, ownerId int64, fn task.TaskHandler) error {
	mb, err := m.getMailBox(ownerId)
	if err != nil {
		return err
	}

	return mb.AddTask(ctx, fn, mb)
}

// 创建新邮件
func (m *MailManager) CreateMail(ctx context.Context, ownerId int64, mail *define.Mail) error {
	fn := func(c context.Context, p ...interface{}) error {
		mailBox := p[0].(*mailbox.MailBox)

		if err := mailBox.CheckAvaliable(c); err != nil {
			return err
		}

		return mailBox.AddMail(c, mail)
	}

	return m.AddTask(ctx, ownerId, fn)
}

// 删除邮件
func (m *MailManager) DelMail(ctx context.Context, ownerId int64, mailId int64) error {
	fn := func(c context.Context, p ...interface{}) error {
		mailBox := p[0].(*mailbox.MailBox)

		if err := mailBox.CheckAvaliable(c); err != nil {
			return err
		}

		return mailBox.DelMail(c, mailId)
	}

	return m.AddTask(ctx, ownerId, fn)
}

// 查询玩家邮件
func (m *MailManager) QueryPlayerMails(ctx context.Context, ownerId int64) ([]*define.Mail, error) {
	retMails := make([]*define.Mail, 0)

	fn := func(c context.Context, p ...interface{}) error {
		mailBox := p[0].(*mailbox.MailBox)

		if err := mailBox.CheckAvaliable(c); err != nil {
			return err
		}

		retMails = mailBox.GetMails(c)
		return nil
	}

	err := m.AddTask(ctx, ownerId, fn)
	return retMails, err
}

// 读取邮件
func (m *MailManager) ReadMail(ctx context.Context, ownerId int64, mailId int64) error {
	fn := func(c context.Context, p ...interface{}) error {
		mailBox := p[0].(*mailbox.MailBox)

		return mailBox.ReadMail(c, mailId)
	}

	return m.AddTask(ctx, ownerId, fn)
}

// 获取附件
func (m *MailManager) GainAttachments(ctx context.Context, ownerId int64, mailId int64) error {
	fn := func(c context.Context, p ...interface{}) error {
		mailBox := p[0].(*mailbox.MailBox)

		return mailBox.GainAttachments(c, mailId)
	}

	return m.AddTask(ctx, ownerId, fn)
}
