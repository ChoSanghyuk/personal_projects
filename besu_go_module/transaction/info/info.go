package info

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

var BesuKey = make(map[string]map[string]string)
var BesuNetwork = make(map[string]string)

func init() {
	// Reading JSON file
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime Caller Error")
	}
	dir := filepath.Dir(f)

	/////////////////
	k, err := os.ReadFile(dir + "/key.json")
	if err != nil {
		panic(err)
	}

	// Umarshalling JSON into struct
	json.Unmarshal(k, &BesuKey)

	n, err := os.ReadFile(dir + "/network.json")
	if err != nil {
		panic(err)
	}

	// Umarshalling JSON into struct
	json.Unmarshal(n, &BesuNetwork)

}
