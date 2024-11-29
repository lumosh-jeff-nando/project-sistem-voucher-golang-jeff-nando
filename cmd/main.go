package main

import (
	"github.com/project-sistem-voucher/config"
	routes "github.com/project-sistem-voucher/router"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}

func main() {
	routes.Server().Run()
}
