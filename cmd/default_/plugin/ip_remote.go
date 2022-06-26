package plugin

import (
	"errors"

	"github.com/coreservice-io/ip_geo/ipstack"
	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/basic/conf"
	"github.com/coreservice-io/service-util/plugin/ip_remote_plugin"
	"github.com/coreservice-io/service-util/plugin/reference_plugin"
)

func initIpRemote() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.IpRemote.Enable {
		ip_geo_redis_config := ipstack.RedisConfig{
			Addr:     toml_conf.IpRemote.Redis.Host,
			UserName: toml_conf.IpRemote.Redis.Username,
			Password: toml_conf.IpRemote.Redis.Password,
			Port:     toml_conf.IpRemote.Redis.Port,
			Prefix:   toml_conf.IpRemote.Redis.Prefix,
			UseTLS:   toml_conf.IpRemote.Redis.Use_tls,
		}

		if toml_conf.IpRemote.Ipstack_key == "" {
			return errors.New("ip stack key empty")
		}

		reference_plugin.Init_("ip_remote")
		ip_remote_ref := reference_plugin.GetInstance_("ip_remote")

		basic.Logger.Infoln("init ecs uploader plugin with ipstack_key:", toml_conf.IpRemote.Ipstack_key,
			"ip_geo_redis_config:", ip_geo_redis_config)

		return ip_remote_plugin.Init(toml_conf.IpRemote.Ipstack_key, ip_remote_ref, ip_geo_redis_config, basic.Logger)
	}
	return nil
}
