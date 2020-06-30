package job

import (
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/eloquent"
	"github.com/mylxsw/eloquent/query"
	"github.com/mylxsw/mysql-guard/config"
	"github.com/mylxsw/mysql-guard/models"
)

const ModuleProcessKiller = "process-killer"

// processKiller 杀死执行时间超长的 SQL 连接
func processKiller(processlistModel *models.ProcesslistModel, eloquentDB eloquent.Database, conf *config.Config) error {
	logger := log.Module(ModuleProcessKiller)

	processes, err := processlistModel.Get(query.Builder().
		WhereIn("command", stringArrayToInterface(conf.KillerMatchCommands)...).
		Where("time", ">", conf.KillerBusyTime))
	if err != nil {
		return err
	}

	for _, p := range processes {
		if !conf.TestMode {
			err := eloquentDB.Statement("kill ?", p.Id)
			if err != nil {
				logger.WithFields(log.Fields{
					"process": p,
				}).Errorf("kill connection %d failed: %v", p.Id, err)
				continue
			}
		}

		logger.WithFields(log.Fields{
			"process": p,
		}).Infof("kill connection %d", p.Id)
	}

	return nil
}

func stringArrayToInterface(s []string) []interface{} {
	d := make([]interface{}, len(s))
	for i, v := range s {
		d[i] = v
	}

	return d
}
