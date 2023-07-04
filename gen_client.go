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
		g.P("func (x *", svc.GoName, "ClientImpl) ", clientSignature(g, method), " {")
		g.P("client := ", svcPath.Ident("New"+svc.GoName+"Client(x.cli.Conn())"))
		var args []string
		args = append(args, "ctx")
		if !method.Desc.IsStreamingClient() {
			args = append(args, "in")
		}
		args = append(args, "opts...")
		g.P("return client.", method.GoName, "(", strings.Join(args, ", "), ")")
		g.P("}")
	}
	return
}

func clientSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	return method.GoName + "(" + strings.Join(clientArgs(g, method), ", ") + ") " + clientReturn(g, method)
}

func clientReturn(g *protogen.GeneratedFile, method *protogen.Method) string {
	var ret []string
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		ret = append(ret, "*"+g.QualifiedGoIdent(method.Output.GoIdent))
	} else {
		ret = append(ret, g.QualifiedGoIdent(method.Output.GoIdent.GoImportPath.Ident(method.Parent.GoName+"_"+method.GoName+"Client")))
	}
	ret = append(ret, "error")
	return "(" + strings.Join(ret, ", ") + ")"
}

func clientArgs(g *protogen.GeneratedFile, method *protogen.Method) []string {
	var args []string
	args = append(args, "ctx "+g.QualifiedGoIdent(contextPkg.Ident("Context")))
	if !method.Desc.IsStreamingClient() {
		args = append(args, "in *"+g.QualifiedGoIdent(method.Input.GoIdent))
	}
	args = append(args, "opts ..."+g.QualifiedGoIdent(grpcPkg.Ident("CallOption")))
	return args
}
