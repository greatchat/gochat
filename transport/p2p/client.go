package p2p

import (
	goP2p "github.com/tendermint/go-p2p"
)

type client struct {
	conn goP2p.MConnection
}
