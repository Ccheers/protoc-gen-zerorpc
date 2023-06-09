package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.0.1"

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-zeroapi %v\n", version)
		return
	}

	var flags flag.FlagSet
	outDir := flags.String("out", "./dest", "output directory (required)")
	options := protogen.Options{
		ParamFunc: flags.Set,
	}
	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			genZeroLayout(*outDir, gen, f)
			genZeroLogic(*outDir, gen, f)
			genClient(*outDir, gen, f)
			genServer(*outDir, gen, f)
		}
		return nil
	})
}
