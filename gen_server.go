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
		g.P("func (x *", svc.GoName, "ServerImpl) ", serverSignature(g, method), "{")
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			g.P("ctx := stream.Context()")
		}
		g.P("handler := ", logicPackage.Ident("New"+method.GoName+"Logic(ctx, x.svcCtx)"))

		var reqArgs []string
		if !method.Desc.IsStreamingClient() {
			reqArgs = append(reqArgs, "in")
		}
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			reqArgs = append(reqArgs, "stream")
		}
		g.P("return handler.", method.GoName, "(", strings.Join(reqArgs, ", "), ")")
		g.P("}")
	}
	return
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	return method.GoName + "(" + strings.Join(serverArgs(g, method), ", ") + ") " + serverReturn(g, method)
}

func serverReturn(g *protogen.GeneratedFile, method *protogen.Method) string {
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		return "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	}
	return "error"
}

func serverArgs(g *protogen.GeneratedFile, method *protogen.Method) []string {
	var reqArgs []string
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, "ctx "+g.QualifiedGoIdent(contextPkg.Ident("Context")))
	}
	if !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, "in *"+g.QualifiedGoIdent(method.Input.GoIdent))
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, "stream "+g.QualifiedGoIdent(method.Input.GoIdent.GoImportPath.Ident(method.Parent.GoName+"_"+method.GoName+"Server")))
	}
	return reqArgs
}
