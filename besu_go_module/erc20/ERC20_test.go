package erc20

import (
	"fmt"
	"go_module/transaction"
	"go_module/transaction/info"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var acnt1Pk = info.BesuKey["account1"]["privateKey"][2:]
var acnt1Addr = info.BesuKey["account1"]["address"]
var acnt2Addr = info.BesuKey["account2"]["address"]
var auth = transaction.CreateAuth(acnt1Pk)
var url = info.BesuNetwork["url"]
var client, _ = ethclient.Dial(url)
var instance *Erc20 = nil

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	tokenAddr := "0xDF3149B11d7457eA472a56463bB9c0928fC25946"
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, _ = NewErc20(tokenAddress, client)

}

func TestERC20Deploy(t *testing.T) {

	addr, tx, _, err := DeployErc20(auth, client, "TEST_TOKEN", "TST", 2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Token Address : %s\n", addr.Hex())
	fmt.Printf("Transaction Hash : %s\n", tx.Hash())
}

func TestERC20ToltalSupply(t *testing.T) {

	ts, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Total Supply : %d\n", ts)

}

func TestERC20Mint(t *testing.T) {

	var amt int64 = 100
	toAddress := common.HexToAddress(acnt1Addr)
	tx, err := instance.Mint(auth, toAddress, big.NewInt(amt))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Transaction Hash : %s\n", tx.Hash())

}

func TestERC20BalanceOf(t *testing.T) {

	addr := common.HexToAddress(acnt2Addr)
	balance, err := instance.BalanceOf(&bind.CallOpts{}, addr)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s's Balance : %s\n", acnt1Addr, balance)
}

func TestERC20Transfer(t *testing.T) {

	var amt int64 = 5
	toAddress := common.HexToAddress(acnt2Addr)
	tx, err := instance.Transfer(auth, toAddress, big.NewInt(amt))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Transaction Hash : %s\n", tx.Hash())

}
