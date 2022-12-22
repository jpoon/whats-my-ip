# wheres-my-server

Basic Dynamic DNS

My machine is assigned a different IP address every time it joins VPN. 
As a means of discovering what the current IP address is, it periodically sends a heartbeat to a server (that has a static IP).

## Usage

On the server:
```
$ cd server
$ docker-compose start
```

On the client:
```
# configure client.go to talk to above server
$ go run client.go
```
