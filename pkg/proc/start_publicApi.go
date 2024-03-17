/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package proc

import (
	"github.com/codoworks/go-boilerplate/pkg/api/routers"
	"github.com/codoworks/go-boilerplate/pkg/clients/service"
)

func StartPublicApi() {
	serviceCli := service.GetClient()
	config := serviceCli.GetConfig()
	routers.InitPublicAPIRouter()
	routers.PublicAPIRouter().Start(config.Host, config.PublicApiPort)
}
