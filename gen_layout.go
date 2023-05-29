package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
)

// 生成骨架
//api
//├── bookstore.go                   // main入口定义
//├── etc
//│   └── bookstore-api.yaml         // 配置文件
//└── internal
//├── config
//│   └── config.go              // 定义配置
//├── handler
//│   ├── addhandler.go          // 实现addHandler
//│   ├── checkhandler.go        // 实现checkHandler
//│   └── routes.go              // 定义路由处理
//├── logic
//│   ├── addlogic.go            // 实现AddLogic
//│   └── checklogic.go          // 实现CheckLogic
//├── svc
//│   └── servicecontext.go      // 定义ServiceContext

func genZeroLayout(outDir string, gen *protogen.Plugin, file *protogen.File) {
	err := os.MkdirAll(outDir+"/internal/logic", os.ModePerm)
	if err != nil {
		panic(err)
	}
	buildConfigFile(outDir)
	buildConfig(outDir)
	buildSvc(outDir)
	buildMain(outDir, gen, file)
}

func buildSvc(outDir string) {
	err := os.MkdirAll(outDir+"/internal/svc", os.ModePerm)
	if err != nil {
		panic(err)
	}
	filename := outDir + "/internal/svc/servicecontext.go"
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}

	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `package svc

import (
	"%s/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
`
	_, err = fw.Write([]byte(fmt.Sprintf(content, goPackage(outDir))))
	if err != nil {
		panic(err)
	}
}

func buildMain(outDir string, gen *protogen.Plugin, file *protogen.File) {
	filename := outDir + "/main.go"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	_ = os.WriteFile(filename, []byte(""), os.ModePerm)

	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(outDir))
	rootPackage := goPackage(outDir)
	configPkg := protogen.GoImportPath(rootPackage + "/internal/config")
	svcPkg := protogen.GoImportPath(rootPackage + "/internal/svc")
	serverPkg := protogen.GoImportPath(rootPackage + "/internal/server")
	g.P("// Code generated by protoc-gen-zerorpc.")
	g.P("package main")
	g.P()
	g.P(`
import (
	"flag"
	"fmt"

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
)`)

	g.P("var configFile = flag.String(\"f\", \"etc/config.yaml\", \"the config file\")")
	g.P("")
	g.P("func main() {")
	g.P("flag.Parse()")
	g.P("")
	g.P("var c ", configPkg.Ident("Config"))
	g.P("conf.MustLoad(*configFile, &c)")
	g.P("g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(*configFile)")
	g.P("")
	g.P("logx.AddGlobalFields(logx.Field(\"service\", c.Name))")
	g.P("")
	g.P("svcCtx := ", svcPkg.Ident("NewServiceContext(c)"))

	g.P("s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {")
	for _, service := range file.Services {
		g.P(file.GoImportPath.Ident(fmt.Sprintf("Register%sServer", service.GoName)), "(grpcServer, ", serverPkg.Ident(fmt.Sprintf("New%sServer", service.GoName)), "(svcCtx))")
	}
	g.P(`
        if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
        }
    })
    defer s.Stop()
	// 添加 rpc 中间件
	// s.AddUnaryInterceptors(rpcserver.ContextTransformer, rpcserver.LoggerInterceptor)


	fmt.Printf("Starting server at %s...\n", c.ListenOn)
	s.Start()
}`)
}

func buildConfig(outDir string) {
	filename := outDir + "/internal/config/config.go"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	err = os.MkdirAll(outDir+"/internal/config", os.ModePerm)
	if err != nil {
		panic(err)
	}
	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
}
`
	_, err = fw.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}

func buildConfigFile(outDir string) {
	filename := outDir + "/etc/config.yaml"
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return
	}
	err = os.MkdirAll(outDir+"/etc", os.ModePerm)
	if err != nil {
		panic(err)
	}
	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	const content = `# 服务名
Name: helloworld-rpc
# 监听端口
ListenOn: 0.0.0.0:8080
# 服务注册
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: helloworld-rpc
#监控
Prometheus:
  Host: 0.0.0.0
  Port: 9091
  Path: /metrics
#链路追踪
Telemetry:
  Name: helloworld-rpc
  Endpoint: http://simple-prod-collector.observability.svc:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger
# 日志
Log:
  ServiceName: helloworld-rpc
  Level: error
`
	_, err = fw.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
