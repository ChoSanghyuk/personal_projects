package besu_graphql

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

func TestGraphqlSend(t *testing.T) {

	query := `{
		"query": "{pending {transactionCount}}"
	}
	`
	rtn := graphqlSend(query)
	fmt.Println("Response : ", rtn)
}

var call1 = Call{To: "0x2114de86c8ea1fd8144c2f1e1e94c74e498afb1b", Data: "0x0e37008a0000000000000000000000000000000000000000000000000000000000000001"}
var call2 = Call{To: "0x2114de86c8ea1fd8144c2f1e1e94c74e498afb1b", Data: "0x0e37008a0000000000000000000000000000000000000000000000000000000000000001"}

// Two Calls 소요 시간 :  5
func TestGraphqlSendWithTwoCalls(t *testing.T) {
	s := time.Now().UnixMilli()
	rtn := BesuCall(big.NewInt(0), call1, call2)
	e := time.Now().UnixMilli()
	fmt.Println(rtn)
	fmt.Println("Two Calls 소요 시간 : ", e-s)
}

func TestGraphqlSendWithThousandCalls(t *testing.T) {

	var calls []Call = make([]Call, 832)

	for i := 0; i < len(calls); i++ {
		calls[i] = call1
	}

	s := time.Now().UnixMilli()
	rtn := BesuCall(big.NewInt(24001), calls...)
	e := time.Now().UnixMilli()
	fmt.Println("rtn : ", rtn)
	fmt.Println("Graphql 단일 스레드 Thousands Calls 소요 시간 : ", e-s)
}

func TestGraphqlSendWithSingleThreadGraphQL(t *testing.T) {

	var calls []Call = make([]Call, 1000000)

	for i := 0; i < len(calls); i++ {
		calls[i] = call1
	}

	s := time.Now().UnixMilli()
	for i := 0; i < len(calls); i += CallLimit {
		se := i + CallLimit
		if se > len(calls) {
			se = len(calls)
		}
		rtn := BesuCall(big.NewInt(24001), calls[i:se]...)
		_ = rtn
	}

	e := time.Now().UnixMilli()
	// fmt.Println("rtn : ", rtn)
	fmt.Println("Graphql 단일 스레드 Thousands Calls 소요 시간 : ", e-s)
}

// Graphql 멀티스레드 Calls with 1000 소요 시간 77
// Graphql 멀티스레드 Calls with 10000 소요 시간 419
func TestGraphqlSendWithMultiCalls(t *testing.T) {
	n := 832
	var calls []Call = make([]Call, n)
	for i := 0; i < len(calls); i++ {
		calls[i] = call1
	}
	s := time.Now().UnixMilli()
	rtn := BesuMultiCall(big.NewInt(24001), calls...)
	e := time.Now().UnixMilli()
	// fmt.Println("rtn : ", rtn)
	_ = rtn
	fmt.Printf("Graphql 멀티스레드 Calls with %d 소요 시간 %d", n, e-s)
}

func TestMakeCallMessageWithZeroBN(t *testing.T) {
	rtn := makeCallMessage(big.NewInt(0), call1, call2)
	fmt.Println("==========================\n", rtn)
}

func TestMakeCallMessageWithNoBN(t *testing.T) {
	rtn := makeCallMessage(nil, call1, call2)
	fmt.Println("==========================\r", rtn)
}

func TestMakeCallMessageWithBN(t *testing.T) {
	rtn := makeCallMessage(new(big.Int).SetInt64(800), call1, call2)
	fmt.Println("==========================\r", rtn)
}

var rt1 = "0xf8a780808302ca51942114de86c8ea1fd8144c2f1e1e94c74e498afb1b80b844fe093d3b00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000014820f9da0e251b5e312ae4f9f558ddd0d19bc30a9f4c040afbd7819fa57e0cdf41573f43da032b6eadd73da42676b37bd57bbef0bba349942be3917f0fab3bd499b441a44ac"
var rt2 = "0xf8a780808302ca51942114de86c8ea1fd8144c2f1e1e94c74e498afb1b80b844fe093d3b00000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000014820f9da0a5337ea829e0fa01e261aa68666de85f208400763909eac63c2c87bb33bad691a042edfc75dd3b6ba7ebb4aea8c7eef222edb4b0fc0f5f9fc0a097955867e61f0e"

func TestMakeMutMessage(t *testing.T) {
	q := makeMutMessage(rt1, rt2)
	fmt.Println("==========================\r", q)
}

func TestGraphqlSendWithMut(t *testing.T) {
	rtn := BesuMut(rt1)
	fmt.Println(rtn)
}

func TestGraphqlSendWithManyMut(t *testing.T) {
	var ts []string = make([]string, 200)
	for i := 0; i < len(ts); i++ {
		ts[i] = rt1
	}
	rtn := BesuMut(ts...)
	fmt.Println(rtn)
}

// 소요시간 동일 결론
func TestCompareTimeWithSameRepeatAndDiffRepeat(t *testing.T) {

	n := 1000

	var calls []Call = make([]Call, n)

	for i := 0; i < len(calls); i++ {
		calls[i] = call1
	}

	s1 := time.Now().UnixMilli()
	rtn := BesuMultiCall(big.NewInt(0), calls...)
	e1 := time.Now().UnixMilli()

	// fmt.Println("rtn : ", rtn)

	var sf string = "0x0e37008a%064x"

	for i := 0; i < len(calls); i++ {
		call1.Data = fmt.Sprintf(sf, i)
		calls[i] = call1
	}

	s2 := time.Now().UnixMilli()
	rtn = BesuMultiCall(big.NewInt(0), calls...)
	e2 := time.Now().UnixMilli()

	_ = rtn
	// fmt.Println("rtn : ", rtn)
	fmt.Printf("Graphql 멀티 스레드 동일 항목 %d번 반복 소요 시간 : %d\n", n, e1-s1)
	fmt.Printf("Graphql 멀티 스레드 다른 항목 %d개 요청 소요 시간 : %d\n", n, e2-s2)
}

func TestCompareTimeWithGroupSize(t *testing.T) {

	n := 1000

	var calls []Call = make([]Call, n)

	for i := 0; i < len(calls); i++ {
		calls[i] = call1
	}
	for t := 100; t <= 800; t = t + 100 {
		CallLimit = t
		s := time.Now().UnixMilli()
		rtn := BesuMultiCall(big.NewInt(0), calls...)
		e := time.Now().UnixMilli()

		_ = rtn
		fmt.Printf("Graphql 멀티 스레드 %d개 요청을 %d의 그룹 사이즈 시의 소요 시간 : %d\n", n, t, e-s)
		time.Sleep(5000)

	}

}
