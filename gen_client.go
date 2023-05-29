package main

import (
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func genClient(outDir string, gen *protogen.Plugin, file *protogen.File) {
	clientDir := outDir + "/client"
	err := os.MkdirAll(clientDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	rootPackage := protogen.GoImportPath(goPackage(outDir))
	for _, service := range file.Services {
		genClientFile(clientDir, rootPackage, gen, service, file.GoImportPath)
	}
	return
}

func genClientFile(clientDir string, rootPackage protogen.GoImportPath, gen *protogen.Plugin, svc *protogen.Service, svcPath protogen.GoImportPath) {
	filename := clientDir + "/" + strings.ToLower(svc.GoName) + ".go"
	clientPackage := rootPackage + "/client"
	g := gen.NewGeneratedFile(filename, clientPackage)
	g.P("package client")
	g.P()
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("type ", svc.GoName, "ClientImpl struct {")
	g.P("cli ", zrpcPkg.Ident("Client"))
	g.P("}")

	// 构造函数
	g.P("func New", svc.GoName, "Client(cli ", zrpcPkg.Ident("Client"), ") ", svcPath.Ident(svc.GoName+"Client"), " {")
	g.P("return &", svc.GoName, "ClientImpl{")
	g.P("cli: cli,")
	g.P("}")
	g.P("}")

	// 方法
	for _, method := range svc.Methods {
		g.P("func (x *", svc.GoName, "ClientImpl) ", method.GoName, "(ctx ", contextPkg.Ident("Context"), ", in *", method.Input.GoIdent, ", opts ...", grpcPkg.Ident("CallOption"), ") (*", method.Output.GoIdent, ", error) {")
		g.P("client := ", svcPath.Ident("New"+svc.GoName+"Client(x.cli.Conn())"))
		g.P("return client.", method.GoName, "(ctx, in, opts...)")
		g.P("}")
	}
	return
}
