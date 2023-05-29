package main

import (
	"flag"
	"fmt"


	"github.com/Ccheers/protoc-gen-zerorpc/example/internal/config"
	"github.com/Ccheers/protoc-gen-zerorpc/example/internal/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(*configFile)

	logx.AddGlobalFields(logx.Field("service", c.Name))

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

        if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
        }
	    
		panic("pb.RegisterXXXX")
    })
    defer s.Stop()
	// 添加 rpc 中间件
	// s.AddUnaryInterceptors(rpcserver.ContextTransformer, rpcserver.LoggerInterceptor)

	ctx := svc.NewServiceContext(c)

	fmt.Printf("Starting server at %s...\n", c.ListenOn)
	server.Start()
}
%!(EXTRA string=github.com/Ccheers/protoc-gen-zerorpc/example)