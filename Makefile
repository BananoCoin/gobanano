export GO15VENDOREXPERIMENT=1

all: banano-node banano-vanity banano-wallet

banano-node: prep
	go build -o build/bin/banano-node github.com/BananoCoin/gobanano/cmd/nano-node

banano-vanity: prep
	go build -o build/bin/banano-vanity github.com/BananoCoin/gobanano/cmd/nano-vanity

banano-wallet: prep
	go build -o build/bin/banano-wallet github.com/BananoCoin/gobanano/cmd/nano-wallet

test:
	GOCACHE=off go test -v $(shell go list ./... | grep -v vendor)

prep:
	mkdir -p build/bin

clean:
	rm -rf build

loc:
	find . -name "*.go" -not -path "./vendor/*" -not -path "./nano/internal/uint128/*" -not -path "./nano/crypto/ed25519/*" | xargs wc -l
