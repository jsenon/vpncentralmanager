# VPN Central Manager

[![CircleCI](https://circleci.com/gh/jsenon/vpncentralmanager.svg?style=svg)](https://circleci.com/gh/jsenon/vpncentralmanager)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/vpncentralmanager)](https://goreportcard.com/report/github.com/jsenon/vpncentralmanager)

## Func

- Register New Client form web portal -> Store in DB -> Send Configuration Client to VPN Server -> Send Configuration to web portal with vpn server info
- Register New VPN Server for vpn server -> Store in DB
- Receive Ack from VPN Server -> Store in DB (Available to client config)
- Retrieve all Configuration for client or vpn server

To be used with vpn slave controller https://github.com/visheyra/vpnslavecontroller

### Prerequisite

Have DynamoDB Server.

### Launch Development Env

You can launch docker-compose stack for your local dev:

- Jaeger
- DynamoDB
- Vpncentralmanager

```sh
docker-compose up -d
```

### Standalone DynamoDB Server

You can use local server for development purpose:

```sh
docker run   -v $PWD/data/:/data -p 8000:8000  --entrypoint '/bin/bash'  amazon/dynamodb-local -c 'java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb -dbPath /data'
```

### Run

Docker image available on github `jsenon/vpncentralmanager:latest`

### Usage

Register New VPN Server

```sh
./grpc_cli call localhost:50051 GetConfig "ippublic : '192.168.0.5', keypublic : 'fdjlfdsjkf989dfjsklj', hostname : 'Server1'"
```

Update VPN Server State

```sh
./grpc_cli call localhost:50051 GetAck "status : 'Ready', serverid : 'VcnfxJcuibnfbmta'"
```

Register New Client

```sh
./grpc_cli call localhost:50051 GetClientDemand "keypublic : 'fgdfgfgfdg867gfdgdf', hostname : 'Myclient1'"
```

Get All Client

```sh
./grpc_cli call localhost:50051 GetAllConfig "id: 'fsdfsd', type: 'client'"
```

Get All VPN Server

```sh
./grpc_cli call localhost:50051 GetAllConfig "id: 'fsdfsd', type: 'vpnserver'"
```

### DynamoDB

Server Definition

```json
{
    "Server": {"S": "54klkl45klk"},
    "ServerName": {"S": "Myserver"},
    "AddressVpn": {"S": "127.0.0.1"},
    "AddressPub": {"S": "127.0.0.1"},
    "PublicKey": {"S": "dsfsdf098dsjlkj"},
    "Status": {"S": "OK"}
}
```

Client Definition

```json
{
    "Client": {"S": "Myclient"},
    "AddressVpn": {"S": "127.0.0.1"},
    "ClientName": {"S": "dfsdf98dfsdf"},
    "PublicKey": {"S": "dsfsdf098dsjlkj"},
    "Status": {"S": "OK"}
}
```

### TIPS

### TODO

- [x] Register new VPN Server
- [x] Ack func, change status
- [x] Send all client configuration to new VPN Server
- [ ] Ip management of server to be review in case of only one empty ip between two in db
- [x] Do we send generated server vpn name in order to be identified for further communication?

- [x] Regisster new client request
- [x] Send client configuration (Allowed IP/32 corresponding to client vpnipaddr + PublicKey) to vpn server
- [x] Send complete client configuration to web portal
- [ ] Delete a client from web portal

- [x] Rest for well-known
- [x] Rest for healthz
- [x] Opencensus
- [x] Opencensus exporter
- [x] Opentracing
- [x] Prometheus

- [ ] Refactor func too complex