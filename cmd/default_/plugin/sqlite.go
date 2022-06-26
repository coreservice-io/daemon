package plugin

import (
	"errors"

	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/basic/conf"
	"github.com/coreservice-io/service-util/plugin/sqlite_plugin"
	"github.com/coreservice-io/utils/path_util"
)

func initSqlite() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Sqlite.Enable {
		sqlite_abs_path, sqlite_abs_path_exist, _ := path_util.SmartPathExist(toml_conf.Sqlite.Path)
		if !sqlite_abs_path_exist {
			return errors.New(toml_conf.Sqlite.Path + " :sqlite.path not exist , please reset your sqlite.path :" + toml_conf.Sqlite.Path)
		}

		sqlite_conf := sqlite_plugin.Config{
			Sqlite_abs_path: sqlite_abs_path,
		}

		basic.Logger.Infoln("init sqlite plugin with config:", sqlite_conf)
		return sqlite_plugin.Init(&sqlite_conf, basic.Logger)
	}

	return nil
}
