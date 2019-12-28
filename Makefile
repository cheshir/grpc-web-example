compileproto:
	protoc -I api/proto api/proto/currenttimer.proto --go_out=plugins=grpc:api/gen/currenttime && \
	protoc -I api/proto currenttimer.proto --js_out=import_style=commonjs:static --grpc-web_out=import_style=commonjs,mode=grpcwebtext:static
