package job

import (
	"github.com/mylxsw/container"
	"github.com/mylxsw/glacier"
	"github.com/mylxsw/glacier/cron"
	"github.com/mylxsw/mysql-guard/config"
)

type ServiceProvider struct {
	conf *config.Config `autowire:"@"`
}

func (s *ServiceProvider) Register(app container.Container) {
}

func (s *ServiceProvider) Boot(app glacier.Glacier) {
	app.Cron(func(cr cron.Manager, cc container.Container) error {
		if s.conf.Killer {
			cc.Must(cr.Add("Process Killer", "@every 3s", processKiller))
		}
		if s.conf.DeadlockLogger {
			cc.Must(cr.Add("Deadlock Logger", "@every 3s", deadlockLogger))
		}
		return nil
	})
}
