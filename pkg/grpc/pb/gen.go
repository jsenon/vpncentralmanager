package pb

//go:generate protoc -I . -I /usr/local/include --go_out=plugins=grpc:. advertise.proto
//go:generate protoc -I . -I /usr/local/include --go_out=plugins=grpc:. clientconfig.proto
//go:generate protoc -I . -I /usr/local/include --go_out=plugins=grpc:. newclientdemand.proto
//go:generate protoc -I . -I /usr/local/include --go_out=plugins=grpc:. ackconfig.proto
//go:generate protoc -I . -I /usr/local/include --go_out=plugins=grpc:. retrieveconfig.proto
