package main

import (
	"KubeXCloud/global"
	"KubeXCloud/initiallize"
)

func main() {
	r := initiallize.Routers()
	initiallize.Viper()
	//initiallize.K8S()
	panic(r.Run(global.CONF.System.Addr))
}
