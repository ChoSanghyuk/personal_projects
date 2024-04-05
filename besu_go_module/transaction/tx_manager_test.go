package transaction

import (
	"fmt"
	"go_module/transaction/info"
	store "go_module/transaction/smart_contract"
	"go_module/util"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func TestDeploy(t *testing.T) {
	pk := info.BesuKey["account1"]["privateKey"][2:]
	abi, _ := store.MapStoreMetaData.GetAbi()
	bin := store.MapStoreMetaData.Bin

	addr, txHash := Deploy(pk, abi, bin) // 0xE03Ef2490316bfF9808d936eEe70f23896F07548
	fmt.Println("생성된 컨트랙트 주소", addr)
	fmt.Println("수행된 트랜잭션 해시", txHash)
}

func TestCall(t *testing.T) {

	addr := "0xe52155361a36C7d445F2c6784B14Bf7A3C306e15" //"0x2114de86c8ea1fd8144c2f1e1e94c74e498afb1b" // "0xE03Ef2490316bfF9808d936eEe70f23896F07548"
	abi, _ := store.MapStoreMetaData.GetAbi()            // store.StoreMetaData.GetAbi()
	// method := "store"

	rtn, err := Call(nil, addr, abi, "getAge", big.NewInt(1))
	util.CheckErr(err, "Call 실패")

	// a := fmt.Sprintf("%s", rtn)
	fmt.Println(rtn)

}

const n = 1000000

// 단일 스레드 1000개 소요 시간 : 492
func TestMultiCall(t *testing.T) {
	var rtns []interface{} = make([]interface{}, n)

	addr := "0x2114de86c8ea1fd8144c2f1e1e94c74e498afb1b"
	abi, _ := store.MapStoreMetaData.GetAbi()

	address := common.HexToAddress(addr)

	input, err := abi.Pack("getAge", big.NewInt(1))
	util.CheckErr(err)

	msg := ethereum.CallMsg{
		From: common.Address{},
		To:   &address,
		Data: input,
	}

	s := time.Now().UnixMilli()
	for i := 0; i < n; i++ {
		rtn, err := callByMsgForTest(nil, msg)
		util.CheckErr(err, "Call 실패")
		rtns[i] = rtn
	}
	e := time.Now().UnixMilli()

	// fmt.Println(rtns)
	_ = rtns
	fmt.Printf("단일 스레드 %d개 소요 시간 : %d", n, e-s)

}

// Goroutine 사용 시, 1000개 소요 시간 : 296
func TestMultiWithGoGroutineCall(t *testing.T) {
	var rtns [][]byte = make([][]byte, n)

	addr := "0x2114de86c8ea1fd8144c2f1e1e94c74e498afb1b"
	abi, _ := store.MapStoreMetaData.GetAbi()

	address := common.HexToAddress(addr)

	input, err := abi.Pack("getAge", big.NewInt(1))
	util.CheckErr(err)

	msg := ethereum.CallMsg{
		From: common.Address{},
		To:   &address,
		Data: input,
	}
	c := make(chan []byte)

	s := time.Now().UnixMilli()
	for i := 0; i < n; i++ {
		go callByMsgWithGoRoutineForTest(c, nil, msg)
	}

	for i := 0; i < n; i++ {
		rtns[i] = <-c
	}

	e := time.Now().UnixMilli()

	// fmt.Println(rtns)
	_ = rtns
	fmt.Printf("Goroutine 사용 시, %d개 소요 시간 : %d", n, e-s)

}

func TestWrite(t *testing.T) {

	pk := info.BesuKey["account1"]["privateKey"][2:]
	addr := "0xe52155361a36C7d445F2c6784B14Bf7A3C306e15" //"0x2114de86c8ea1fd8144c2f1e1e94c74e498afb1b"
	abi, _ := store.MapStoreMetaData.GetAbi()

	rtn := Write(pk, addr, abi, "setAgeList", big.NewInt(1), big.NewInt(30))
	fmt.Printf("%v", rtn.Hex())
}
