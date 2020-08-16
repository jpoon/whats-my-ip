# wheres-my-server

## Background

My machine is assigned a different IP address every time it joins VPN. 
As a means of discovering what the current IP address is, it periodically sends a heartbeat to a server (that has a static IP).

## Usage

On the server:
```
$ docker-compose start
$ go run server.go
```

On the client:
```
$ go run client.go
```
