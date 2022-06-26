package plugin

import (
	"github.com/coreservice-io/redis_spr"
	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/basic/conf"
	"github.com/coreservice-io/service-util/plugin/spr_plugin"
)

func initSpr() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Redis.Enable {

		redis_conf := redis_spr.RedisConfig{
			Addr:     toml_conf.Redis.Host,
			UserName: toml_conf.Redis.Username,
			Password: toml_conf.Redis.Password,
			Port:     toml_conf.Redis.Port,
			Prefix:   toml_conf.Redis.Prefix,
			UseTLS:   toml_conf.Redis.Use_tls,
		}

		basic.Logger.Infoln("init redis plugin with config:", redis_conf)
		return spr_plugin.Init(&redis_conf, basic.Logger)
	}

	return nil
}
