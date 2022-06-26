package plugin

import (
	"github.com/coreservice-io/service-util/basic"
	"github.com/coreservice-io/service-util/basic/conf"
	"github.com/coreservice-io/service-util/plugin/mail_plugin"
)

func initSmtpMail() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Smtp.Enable {

		mail_conf := mail_plugin.Config{
			Host:     toml_conf.Smtp.Host,
			Port:     toml_conf.Smtp.Port,
			Password: toml_conf.Smtp.Password,
			UserName: toml_conf.Smtp.Username,
		}

		basic.Logger.Infoln("init smtp mail plugin with config:", mail_conf)
		return mail_plugin.Init(&mail_conf)
	}
	return nil
}
