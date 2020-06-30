package job

import (
	"fmt"

	"github.com/ledisdb/ledisdb/ledis"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/eloquent"
	"github.com/mylxsw/mysql-guard/parser"
)

const ModuleDeadlockLogger = "deadlock-logger"

func deadlockLogger(eloquentDB eloquent.Database, localDB *ledis.DB) error {
	logger := log.Module(ModuleDeadlockLogger)

	results, err := eloquentDB.Query(eloquent.Raw("SHOW ENGINE INNODB STATUS"), func(row eloquent.Scanner) (interface{}, error) {
		var typ, name, status string
		if err := row.Scan(&typ, &name, &status); err != nil {
			return nil, err
		}

		return status, nil
	})
	if err != nil {
		logger.Errorf("show engine innodb status failed: %v", err)
		return nil
	}

	deadlocks := parser.ParseInnoDBDeadlocks(results.Index(0).(string))
	if len(deadlocks) == 0 {
		return nil
	}

	for k, locks := range deadlocks {
		statusKey := fmt.Sprintf("deadlock-%s", k)
		r, _ := localDB.Get([]byte(statusKey))
		if r != nil {
			continue
		}

		log.WithFields(log.Fields{
			"deadlocks": locks,
		}).Info("found deadlocks")

		_ = localDB.Set([]byte(statusKey), []byte("true"))
	}

	return nil
}
