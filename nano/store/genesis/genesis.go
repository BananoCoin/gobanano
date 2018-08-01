package genesis

import (
	"fmt"

	"github.com/BananoCoin/gobanano/nano"
	"github.com/BananoCoin/gobanano/nano/block"
	"github.com/BananoCoin/gobanano/nano/internal/util"
	"github.com/BananoCoin/gobanano/nano/node/proto"
)

type Genesis struct {
	Block         block.OpenBlock
	Balance       nano.Balance
	WorkThreshold uint64
}

var (
	Live = Genesis{
		Block: block.OpenBlock{
			SourceHash:     util.MustDecodeHex32("2514452A978F08D1CF76BB40B6AD064183CF275D3CC5D3E0515DC96E2112AD4E"),
			Representative: util.MustDecodeHex32("2514452A978F08D1CF76BB40B6AD064183CF275D3CC5D3E0515DC96E2112AD4E"),
			Address:        util.MustDecodeHex32("2514452A978F08D1CF76BB40B6AD064183CF275D3CC5D3E0515DC96E2112AD4E"),
			Work:           0xfa055f79fa56abcf,
			Signature:      util.MustDecodeHex64("533dcab343547b93c4128e779848dea5877d3278cb5ea948bb3a9aa1ae0db293de6d9da4f69e8d1ddfa385f9b4c5e4f38dfa42c00d7b183560435d07afa18900"),
		},
		Balance:       nano.ParseBalanceInts(0xffffffffffffffff, 0xffffffffffffffff),
		WorkThreshold: uint64(0xfffffe0000000000),
	}

	Beta = Genesis{
		Block: block.OpenBlock{
			SourceHash:     util.MustDecodeHex32("a59a47cc4f593e75ae9ad653fda9358e2f7898d9acc8c60e80d0495ce20fba9f"),
			Representative: util.MustDecodeHex32("a59a47cc4f593e75ae9ad653fda9358e2f7898d9acc8c60e80d0495ce20fba9f"),
			Address:        util.MustDecodeHex32("a59a47cc4f593e75ae9ad653fda9358e2f7898d9acc8c60e80d0495ce20fba9f"),
			Work:           0x000000000f0aaeeb,
			Signature:      util.MustDecodeHex64("a726490e3325e4fa59c1c900d5b6eebb15fe13d99f49d475b93f0aacc5635929a0614cf3892764a04d1c6732a0d716ffeb254d4154c6f544d11e6630f201450b"),
		},
		Balance:       nano.ParseBalanceInts(0xffffffffffffffff, 0xffffffffffffffff),
		WorkThreshold: uint64(0xfffffe0000000000),
	}
)

func Get(network proto.Network) (Genesis, error) {
	switch network {
	case proto.NetworkLive:
		return Live, nil
	case proto.NetworkBeta:
		return Beta, nil
	}

	return Genesis{}, fmt.Errorf("unsupported network: %s", network)
}
