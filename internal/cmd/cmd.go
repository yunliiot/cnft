package cmd

import (
	"cnft/internal/config"
	"cnft/internal/service/token"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"cnft/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			Init(ctx)
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.NftEntity,
				)
			})
			s.Run()
			return nil
		},
	}
)

func Init(ctx context.Context) {
	InitConfig(ctx)
	if config.Config.UseChain33 {
		token.InitJSONClient()
	}
}

func InitConfig(ctx context.Context) {
	data, _ := g.Cfg().Data(ctx)
	bs := gvar.New(data).Bytes()
	json.Unmarshal(bs, &config.Config)
}
