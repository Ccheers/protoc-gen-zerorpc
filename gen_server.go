package main

import (
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func genServer(outDir string, gen *protogen.Plugin, file *protogen.File) {
	serverDir := outDir + "/internal/server"
	err := os.MkdirAll(serverDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	rootPackage := protogen.GoImportPath(goPackage(outDir))
	for _, service := range file.Services {
		genServerFile(serverDir, rootPackage, gen, service, file.GoImportPath)
	}
	return
}

func genServerFile(serverDir string, rootPackage protogen.GoImportPath, gen *protogen.Plugin, svc *protogen.Service, svcPath protogen.GoImportPath) {
	filename := serverDir + "/" + strings.ToLower(svc.GoName) + ".go"
	serverPackage := rootPackage + "/internal/server"
	svcPackage := rootPackage + "/internal/svc"
	logicPackage := rootPackage + "/internal/logic"
	g := gen.NewGeneratedFile(filename, serverPackage)
	g.P("package server")
	g.P()
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("type ", svc.GoName, "ServerImpl struct {")
	g.P("svcCtx *", svcPackage.Ident("ServiceContext"))
	g.P(svcPath.Ident("Unimplemented" + svc.GoName + "Server"))
	g.P("}")

	// 构造函数
	g.P("func New", svc.GoName, "Server(svcCtx *", svcPackage.Ident("ServiceContext"), ") ", svcPath.Ident(svc.GoName+"Server"), " {")
	g.P("return &", svc.GoName, "ServerImpl{")
	g.P("svcCtx: svcCtx,")
	g.P("}")
	g.P("}")

	// 方法
	for _, method := range svc.Methods {
		g.P("func (x *", svc.GoName, "ServerImpl) ", method.GoName, "(ctx ", contextPkg.Ident("Context"), ", in *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")
		g.P("handler := ", logicPackage.Ident("New"+method.GoName+"Logic(ctx, x.svcCtx)"))
		g.P("return handler.", method.GoName, "(in)")
		g.P("}")
	}
	return
}
