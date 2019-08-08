package tx

import (
	"fmt"
	"github.com/irisnet/sdk-go/client/basic"
	"github.com/irisnet/sdk-go/client/lcd"
	"github.com/irisnet/sdk-go/client/rpc"
	"github.com/irisnet/sdk-go/client/types"
	"github.com/irisnet/sdk-go/keys"
	commontypes "github.com/irisnet/sdk-go/types"
	"github.com/irisnet/sdk-go/util"
	"math"
	"testing"
)

var (
	c TxClient
)

func TestMain(m *testing.M) {
	km, err := keys.NewKeyStoreKeyManager("./ks_1234567890.json", "1234567890")
	if err != nil {
		panic(err)
	}
	basicClient := basic.NewClient("v2.irisnet-lcd.qa.rainbow.one")
	lite := lcd.NewClient(basicClient)
	rpcClient := rpc.NewClient("tcp://35.236.146.181:30657")

	c, err = NewClient("rainbow-qa", commontypes.Testnet, km, lite, rpcClient)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestClient_SendToken(t *testing.T) {
	receiver := "faa1j3ufmgwe2cuumj7423jt4creqlcskltn6ht5w9"
	amount := fmt.Sprintf("%.0f", 0.12*math.Pow10(18))
	coins := []types.Coin{
		{
			Denom:  "iris-atto",
			Amount: amount,
		},
	}
	memo := "send from irisnet/sdk-go"
	if res, err := c.SendToken(receiver, coins, memo); err != nil {
		t.Fatal(err)
	} else {
		t.Log(util.ToJsonIgnoreErr(res))
	}
}
