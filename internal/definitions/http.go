package definitions

import (
	"github.com/fasthttp/router"
	"github.com/sarulabs/di"
	"github.com/valyala/fasthttp"

	"companies/internal/server"
)

const (
	HttpDef = "http"
)

func GetHttpDef() di.Def {
	return di.Def{
		Name:  HttpDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			r := ctn.Get(RouterDef).(*router.Router)
			return &fasthttp.Server{
				Handler: server.Cors(r.Handler),
			}, nil
		},
	}
}
