module mini_messenger/gateway

go 1.25.1

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250818200422-3122310a409c
	google.golang.org/grpc v1.75.1
	mini_messenger/auth v0.0.0-00010101000000-000000000000
	mini_messenger/user v0.0.0-00010101000000-000000000000
)

replace (
	mini_messenger/auth => ../auth
	mini_messenger/user => ../user
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.9-20250912141014-52f32327d4b0.1 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250826171959-ef028d996bc1 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
)
