package plugin

import "github.com/coreservice-io/service-util/plugin/reference_plugin"

//example 3 cache instance
func initReference() error {
	//default instance
	err := reference_plugin.Init()
	if err != nil {
		return err
	}

	return nil
}
