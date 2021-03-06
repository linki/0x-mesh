// demo/add_order is a short program that adds an order to 0x Mesh via RPC
package main

import (
	"math/big"
	"time"

	"github.com/0xProject/0x-mesh/constants"
	"github.com/0xProject/0x-mesh/ws"
	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/plaid/go-envvar/envvar"
	log "github.com/sirupsen/logrus"
)

type clientEnvVars struct {
	// RPCAddress is the address of the 0x Mesh node to communicate with.
	RPCAddress string `envvar:"RPC_ADDRESS"`
}

var testOrder = &zeroex.SignedOrder{
	MakerAddress:          common.HexToAddress("0x5409ed021d9299bf6814279a6a1411a7e866a631"),
	TakerAddress:          constants.NullAddress,
	SenderAddress:         constants.NullAddress,
	FeeRecipientAddress:   common.HexToAddress("0xa258b39954cef5cb142fd567a46cddb31a670124"),
	MakerAssetData:        common.Hex2Bytes("f47261b0000000000000000000000000871dd7c2b4b25e1aa18728e9d5f2af4c4e431f5c"),
	TakerAssetData:        common.Hex2Bytes("f47261b00000000000000000000000000b1ba0af832d7c05fd64161e0db78e85978e8082"),
	Salt:                  big.NewInt(1548619145450),
	MakerFee:              big.NewInt(0),
	TakerFee:              big.NewInt(0),
	MakerAssetAmount:      big.NewInt(1000),
	TakerAssetAmount:      big.NewInt(2000),
	ExpirationTimeSeconds: big.NewInt(time.Now().Add(48 * time.Hour).Unix()),
	ExchangeAddress:       constants.NetworkIDToContractAddresses[constants.TestNetworkID].Exchange,
}

func main() {
	env := clientEnvVars{}
	if err := envvar.Parse(&env); err != nil {
		panic(err)
	}

	client, err := ws.NewClient(env.RPCAddress)
	if err != nil {
		log.WithError(err).Fatal("could not create client")
	}

	hash, err := client.AddOrder(testOrder)
	if err != nil {
		log.WithError(err).Fatal("error from AddOrder")
	} else {
		log.Printf("added valid order: %s", hash.Hex())
	}
}
