package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ledisdb/ledisdb/ledis"
	"github.com/mylxsw/adanos-alert/misc"
	"github.com/mylxsw/adanos-alert/pkg/connector"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/log"
	"github.com/mylxsw/eloquent"
	"github.com/mylxsw/eloquent/query"
	"github.com/mylxsw/glacier"
	"github.com/mylxsw/glacier/starter/application"
	"github.com/mylxsw/mysql-guard/config"
	"github.com/mylxsw/mysql-guard/job"
	"github.com/mylxsw/mysql-guard/models"
	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"

	lediscfg "github.com/ledisdb/ledisdb/config"
)

func main() {
	app := application.Create("1.0.0")

	frameworkLogger := log.Module("glacier")
	frameworkLogger.LogLevel(level.Info)
	app.Logger(frameworkLogger)

	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:  "db_host",
		Usage: "数据库主机地址",
		Value: "localhost",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:  "db_user",
		Usage: "数据库连接用户名",
		Value: "root",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:  "db_password",
		Usage: "数据库连接密码",
		Value: "",
	}))
	app.AddFlags(altsrc.NewIntFlag(cli.IntFlag{
		Name:  "db_port",
		Usage: "数据库连接用户名",
		Value: 3306,
	}))
	app.AddFlags(altsrc.NewBoolFlag(cli.BoolFlag{
		Name:  "killer",
		Usage: "是否启用 Killer",
	}))
	app.AddFlags(altsrc.NewStringSliceFlag(cli.StringSliceFlag{
		Name:  "killer_match_command",
		Usage: "匹配的命令列表，支持的命令 Query, Sleep, Binlog Dump, Connect, Delayed insert, Execute, Fetch, Init DB, Kill, Prepare, Processlist, Quit, Reset stmt, Table Dump",
		Value: &cli.StringSlice{"Query"},
	}))
	app.AddFlags(altsrc.NewIntFlag(cli.IntFlag{
		Name:  "killer_busy_time",
		Usage: "Kill 处于某个状态时间超过 busy-time 秒的连接",
		Value: 30,
	}))
	app.AddFlags(altsrc.NewBoolFlag(cli.BoolFlag{
		Name:  "deadlock_logger",
		Usage: "是否启用 Deadlock Logger",
	}))
	app.AddFlags(altsrc.NewBoolFlag(cli.BoolFlag{
		Name:  "test",
		Usage: "只输出日志，不会对数据库产生真实影响",
	}))
	app.AddFlags(altsrc.NewStringSliceFlag(cli.StringSliceFlag{
		Name:  "adanos_server",
		Usage: "adanos 服务地址",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:  "adanos_token",
		Usage: "adanos server token",
	}))
	app.AddFlags(altsrc.NewStringFlag(cli.StringFlag{
		Name:  "data_dir",
		Usage: "本地数据库存储目录",
		Value: "/tmp/mysql-guard",
	}))

	app.Singleton(func(cli glacier.FlagContext) *config.Config {
		return &config.Config{
			DBHost:     cli.String("db_host"),
			DBPort:     cli.Int("db_port"),
			DBUser:     cli.String("db_user"),
			DBPassword: cli.String("db_password"),

			Killer:              cli.Bool("killer"),
			KillerMatchCommands: cli.StringSlice("killer_match_command"),
			KillerBusyTime:      cli.Int("killer_busy_time"),

			DeadlockLogger: cli.Bool("deadlock_logger"),

			AdanosServer: cli.StringSlice("adanos_server"),
			AdanosToken:  cli.String("adanos_token"),

			TestMode: cli.Bool("test"),
			DataDir:  cli.String("data_dir"),
		}
	})

	// Ledis DB
	app.Singleton(func(conf *config.Config) (*ledis.Ledis, error) {
		cfg := lediscfg.NewConfigDefault()
		cfg.DataDir = conf.DataDir
		cfg.Databases = 1

		return ledis.Open(cfg)
	})
	app.Singleton(func(ld *ledis.Ledis) (*ledis.DB, error) {
		return ld.Select(0)
	})

	app.Singleton(func(conf *config.Config) (*sql.DB, error) {
		return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/information_schema?parseTime=true", conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBPort))
	})
	app.Singleton(func(db *sql.DB) query.Database { return db })
	app.Singleton(func(db query.Database) eloquent.Database { return eloquent.DB(db) })
	app.Prototype(models.NewProcesslistModel)

	app.Provider(&job.ServiceProvider{})

	app.Main(func(conf *config.Config) {
		log.WithFields(log.Fields{
			"config": conf,
		}).Debug("config loaded")
		if len(conf.AdanosServer) > 0 {
			log.Module(job.ModuleProcessKiller).Formatter(formatter.NewJSONFormatter())
			log.Module(job.ModuleProcessKiller).Writer(AdanosLogWriter{conf: conf})

			log.Module(job.ModuleDeadlockLogger).Formatter(formatter.NewJSONFormatter())
			log.Module(job.ModuleDeadlockLogger).Writer(AdanosLogWriter{conf: conf})
		}
	})

	if err := app.Run(os.Args); err != nil {
		log.Errorf("application error: %v", err)
	}
}

type AdanosLogWriter struct {
	conf *config.Config
}

func (lw AdanosLogWriter) Write(le level.Level, module string, message string) error {
	meta := make(map[string]interface{})
	tags := make([]string, 0)

	meta["log_level"] = le.GetLevelName()
	meta["log_module"] = module
	meta["log_server"] = misc.ServerIP()
	meta["db_host"] = lw.conf.DBHost
	meta["db_port"] = lw.conf.DBPort

	tags = append(tags, module)

	go func() {
		if err := connector.Send(lw.conf.AdanosServer, lw.conf.AdanosToken, meta, tags, "mysql-guard", message); err != nil {
			log.Errorf("send message to adanos server failed: %v", err)
		}
	}()

	return nil
}

func (lw AdanosLogWriter) ReOpen() error { return nil }

func (lw AdanosLogWriter) Close() error { return nil }
