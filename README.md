## gobanano
gobanano is a low effort banano fork of the *excellent* [gonano](https://github.com/alexbakker/gonano); a nano currency node implementation, written in Go by @alexbakker.


The upstream project is **WIP**. Please see the [upstream project](https://github.com/alexbakker/gonano) for latest status and progression.

Thanks is also due for the wonderful [protocol documentation](https://github.com/alexbakker/gonano/blob/master/doc/protocol.md).

Hopefully Alex has the cycles to continue this project, it seems promising. 


### Building / Running 

Install go (>1.8) and configure a $GOPATH e.g. ~/go/. Please note the `go get` is case-sensitive.

```
$ go get github.com/BananoCoin/gobanano

$ cd go/src/github.com/BananoCoin/gobanano/

$ make
mkdir -p build/bin
go build -o build/bin/banano-node github.com/BananoCoin/gobanano/cmd/nano-node
go build -o build/bin/banano-vanity github.com/BananoCoin/gobanano/cmd/nano-vanity
go build -o build/bin/banano-wallet github.com/BananoCoin/gobanano/cmd/nano-wallet

$ build/bin/banano-node
2018/08/01 21:49:34.455560 opening badger database at /home/am/.config/gonano/node
2018/08/01 21:49:34.462027 initializing ledger
2018/08/01 21:49:34.463065 initializing node
2018/08/01 21:49:34.464098 starting node (network: live)
send (167.99.194.145:7071): keep_alive (152 bytes)
add peer: 167.99.194.145:7071
requesting frontiers from 167.99.194.145:7071
recv packet (167.99.194.145:7071): keep_alive (152 bytes)
send (167.99.194.145:7071): keep_alive (152 bytes)
send (37.17.230.147:7071): keep_alive (152 bytes)
add peer: 37.17.230.147:7071
send (164.132.197.158:7071): keep_alive (152 bytes)
add peer: 164.132.197.158:7071
send (138.68.11.251:7071): keep_alive (152 bytes)
[...]
```

The [badger](https://blog.dgraph.io/post/badger/) db is located in ~/.config/gonano/db/.
