package cmd

import (
    "context"

    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/gogf/gf/v2/net/goai"
    "github.com/gogf/gf/v2/os/gcmd"

    "gf-demo-user/internal/consts"
    "gf-demo-user/internal/controller/merchant"
    "gf-demo-user/internal/controller/order"
    "gf-demo-user/internal/controller/user"
    "gf-demo-user/internal/service"
)

var (
    // Main is the main command.
    Main = gcmd.Command{
        Name:  "main",
        Usage: "main",
        Brief: "外卖订餐系统 API 服务",
        Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
            s := g.Server()
            s.Use(ghttp.MiddlewareHandlerResponse)
            s.Group("/", func(group *ghttp.RouterGroup) {
                // Group middlewares.
                group.Middleware(
                    service.Middleware().Ctx,
                    ghttp.MiddlewareCORS,
                )
                // Register route handlers.
                var (
                    userCtrl     = user.NewV1()
                    merchantCtrl = merchant.NewV1()
                    orderCtrl    = order.NewV1()
                )
                // 绑定用户控制器（包含原有认证接口 + 新增CRUD接口）
                group.Bind(
                    userCtrl,
                )
                // 绑定商户控制器
                group.Bind(
                    merchantCtrl,
                )
                // 绑定订单控制器
                group.Bind(
                    orderCtrl,
                )
                // Special handler that needs authentication.
                group.Group("/", func(group *ghttp.RouterGroup) {
                    group.Middleware(service.Middleware().Auth)
                    group.ALLMap(g.Map{
                        "/user/profile": userCtrl.Profile,
                    })
                })
            })
            // Custom enhance API document.
            enhanceOpenAPIDoc(s)
            // Just run the server.
            s.Run()
            return nil
        },
    }
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
    openapi := s.GetOpenApi()
    openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
    openapi.Config.CommonResponseDataField = `Data`

    // API description.
    openapi.Info = goai.Info{
        Title:       consts.OpenAPITitle,
        Description: consts.OpenAPIDescription,
        Contact: &goai.Contact{
            Name: "GoFrame",
            URL:  "https://goframe.org",
        },
    }
}