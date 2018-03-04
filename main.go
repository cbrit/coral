package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cbrit/coral/av"
)

func main() {
	// Load list of symbols
	raw, err := ioutil.ReadFile("./symbols.json")
	if err != nil {
		panic(err)
	}

	var symbolMap av.InterfaceMap
	json.Unmarshal(raw, &symbolMap)
	var interfaceSlice av.InterfaceSlice = symbolMap["symbols"].([]interface{})
	symbols := interfaceSlice.ConvertToStrings()

	fmt.Println(symbols)

	for _, s := range symbols {
		av.GetStock(s)
	}
}
