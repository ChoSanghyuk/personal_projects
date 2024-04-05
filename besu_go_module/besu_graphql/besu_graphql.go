package besu_graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
)

const url string = "http://localhost:8547"

var CallLimit int = 66

type Call struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

func BesuMultiCall(bn *big.Int, params ...Call) []string {

	n := len(params)/CallLimit + 1

	var rtn []string = make([]string, n)

	c := make(chan string)
	for t := 0; t < n; t++ {
		s := t * CallLimit
		e := s + CallLimit
		if e > len(params) {
			e = len(params)
		}
		go BesuCallTest(c, bn, params[s:e]...)
	}

	for t := 0; t < n; t++ {
		rtn[t] = <-c
	}
	return rtn
}

func BesuCallTest(c chan string, bn *big.Int, params ...Call) {
	q := makeCallMessage(bn, params...)
	rtn := graphqlSend(q)
	c <- rtn
}

func BesuCall(bn *big.Int, params ...Call) string {
	q := makeCallMessage(bn, params...)
	rtn := graphqlSend(q)
	return rtn
}

func BesuMut(params ...string) string {
	q := makeMutMessage(params...)
	rtn := graphqlSend(q)
	return rtn
}

type besuGraphQL = struct {
	Query    string                 `json:"query"`
	Variable map[string]interface{} `json:"variables"`
}

func newBesuCall(i string, b string, c string, v map[string]interface{}) besuGraphQL { // input, block, call, variable
	return besuGraphQL{
		Query:    fmt.Sprintf("query getCall( %s ) {block%s{ %s }}", i, b, c),
		Variable: v,
	}
}

func newBesuMut(i string, m string, v map[string]interface{}) besuGraphQL {
	return besuGraphQL{
		Query:    fmt.Sprintf("mutation( %s ) {%s}", i, m),
		Variable: v,
	}
}

func graphqlSend(query string) string {

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(query))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body)

}

func makeCallMessage(bn *big.Int, params ...Call) string {

	var bf string
	if bn == nil || bn.Cmp(big.NewInt(0)) == 0 {
		bf = ""
	} else {
		bf = fmt.Sprintf("(number: %v)", bn)
	}

	// RequestString form
	inf := "$input%d: CallData!"                   // input form
	cf := "call%d : call(data: $%s){data, status}" // call form

	is := make([]string, len(params))  // input slice
	cs := make([]string, len(params))  // call slice
	vm := make(map[string]interface{}) // variable map

	for i, p := range params {
		n := i + 1
		is[i] = fmt.Sprintf(inf, n)
		cs[i] = fmt.Sprintf(cf, n, fmt.Sprintf("input%d", n))
		vm[fmt.Sprintf("input%d", n)] = p
	}

	mc := newBesuCall(strings.Join(is, ", "), bf, strings.Join(cs, ", "), vm)
	r, _ := json.Marshal(mc)
	return string(r)
}

func makeMutMessage(params ...string) string {

	// RequestString form
	inf := "$input%d: Bytes!"                   // input form
	mf := "t%d : sendRawTransaction(data: $%s)" // mutation form

	is := make([]string, len(params))  // input slice
	ms := make([]string, len(params))  // mutation slice
	vm := make(map[string]interface{}) // variable map

	for i, p := range params {
		n := i + 1
		is[i] = fmt.Sprintf(inf, n)
		ms[i] = fmt.Sprintf(mf, n, fmt.Sprintf("input%d", n))
		vm[fmt.Sprintf("input%d", n)] = p
	}

	mc := newBesuMut(strings.Join(is, ", "), strings.Join(ms, ", "), vm)
	r, _ := json.Marshal(mc)
	return string(r)

}
