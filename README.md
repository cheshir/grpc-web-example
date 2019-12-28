# GRPC-web example

First of all build server that serves GRPC service and file server (for frontend part).

Generate GRPC server and client:

`make compileproto`

Compile frontend:

`cd static`

`make build`

To run altogether (from root project directory):

`docker-compose up`

Then check result in [browser console](http://:8080).
