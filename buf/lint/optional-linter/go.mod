module github.com/project-planton/project-planton/buf/lint/optional-linter

go 1.24.0

require (
	buf.build/go/bufplugin v0.9.0
	github.com/project-planton/project-planton v0.0.0
	google.golang.org/protobuf v1.36.9
)

require (
	buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go v1.36.3-20250121211742-6d880cc6cc8d.1 // indirect
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.9-20250912141014-52f32327d4b0.1 // indirect
	buf.build/gen/go/pluginrpc/pluginrpc/protocolbuffers/go v1.36.3-20241007202033-cf42259fcbfc.1 // indirect
	buf.build/go/protovalidate v1.0.0 // indirect
	buf.build/go/spdx v0.2.0 // indirect
	cel.dev/expr v0.24.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/google/cel-go v0.26.1 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/stoewer/go-strcase v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250603155806-513f23925822 // indirect
	pluginrpc.com/pluginrpc v0.5.0 // indirect
)

replace (
	github.com/project-planton/project-planton => ../../..
	github.com/bufbuild/protovalidate-go => buf.build/go/protovalidate v1.0.0
)
