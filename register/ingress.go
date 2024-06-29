package register

import (
	guidehttp "github.com/advanced-go/guidance/http"
	guidemod "github.com/advanced-go/guidance/module"
	observhttp "github.com/advanced-go/observation/http"
	observmod "github.com/advanced-go/observation/module"
	opshttp "github.com/advanced-go/operations/http"
	opsmod "github.com/advanced-go/operations/module"
	"github.com/advanced-go/stdlib/host"
)

func IngressExchange() error {
	err := host.RegisterExchange(guidemod.Authority, host.NewAccessLogIntermediary(guidemod.RouteName, guidehttp.Exchange))
	if err != nil {
		return err
	}
	err = host.RegisterExchange(observmod.Authority, host.NewAccessLogIntermediary(observmod.RouteName, observhttp.Exchange))
	if err != nil {
		return err
	}
	err = host.RegisterExchange(opsmod.Authority, host.NewAccessLogIntermediary(opsmod.RouteName, opshttp.Exchange))
	if err != nil {
		return err
	}
	//err = host.RegisterExchange(actvmod.Authority, host.NewAccessLogIntermediary(actvmod.RouteName, actvhttp.Exchange))
	//if err != nil {
	//	return err
	//}
	return nil
}
