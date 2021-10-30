package api

import (
	"fmt"

	"github.com/mernat/sso-clean-arch/api/rest"
	"github.com/mernat/sso-clean-arch/config"
)

func Run() {
	var err error
	globalEnvFileName := fmt.Sprintf("%s%s", "./envs/", "config.dev.json")
	_config, _ := config.ExtractConfiguration(globalEnvFileName)

	fmt.Printf("[%s] Serving API at %s\n", "DEV", _config.RestfulEndpoint)

	err = rest.ServeAPI(_config.RestfulEndpoint)
	if err != nil {
		panic(err)
	}
}
