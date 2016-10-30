package jq

import (
	"fmt"

	jq "github.com/threatgrid/jq-go"
)

func Process(rawjson []byte, filter string) ([][]byte, error) {
	vm, err := jq.Compile(filter)
	if err != nil {
		return nil, fmt.Errorf("Can't compile %s", err)
	}
	var results [][]byte
	for ret := range vm.Run(rawjson, &err) {
		results = append(results, ret)
	}
	if err != nil {
		return nil, fmt.Errorf("Can't run %s", err)
	}
	return results, nil
}
