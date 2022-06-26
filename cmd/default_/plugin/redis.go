package plugin

import (
	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/basic/conf"
	"github.com/coreservice-io/service-util/plugin/redis_plugin"
)

func initRedis() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Redis.Enable {
		mail_conf := redis_plugin.Config{
			Address:   toml_conf.Redis.Host,
			UserName:  toml_conf.Redis.Username,
			Password:  toml_conf.Redis.Password,
			Port:      toml_conf.Redis.Port,
			KeyPrefix: toml_conf.Redis.Prefix,
			UseTLS:    toml_conf.Redis.Use_tls,
		}

		basic.Logger.Infoln("init smtp mail plugin with config:", mail_conf)
		return redis_plugin.Init(&mail_conf)
	}

	return nil
}
