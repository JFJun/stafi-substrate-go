package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
)

/*
币种基本配置类型
https://github.com/paritytech/substrate/blob/master/ss58-registry.json
*/
type BasicTypes struct {
	Specification string            `json:"specification"`
	Schema        map[string]string `json:"schema"`
	Registry      []ChainRegistry   `json:"registry"`
}

type ChainRegistry struct {
	Prefix          int64    `json:"prefix"`
	Network         string   `json:"network"`
	DisplayName     string   `json:"displayName"`
	Symbols         []string `json:"symbols"`
	Decimals        []int    `json:"decimals"`
	StandardAccount string   `json:"standardAccount"`
	Website         string   `json:"website"`
}

func InitBasicTypes(filePath string) (*BasicTypes, error) {
	if filePath == "" {
		filePath = "ss58-registry.json"
	}
	cc, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read json file error: %v", err)
	}
	var bt BasicTypes
	err = json.Unmarshal(cc, &bt)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal Basic type error: %v", err)
	}
	return &bt, nil
}

func (bt *BasicTypes) GetChainDecimal(chainName string) (int, error) {
	if len(bt.Registry) == 0 {
		return -1, fmt.Errorf("do not set base type registry")
	}
	for _, reg := range bt.Registry {
		if strings.ToLower(reg.Network) == strings.ToLower(chainName) {
			if len(reg.Decimals) != 0 {
				return reg.Decimals[0], nil
			}
		}
	}
	return -1, errors.New("do not find this chain decimal")
}

var defaultPrefix = []byte{0x2a}

/*
错误时返回默认的prefix
*/
func (bt *BasicTypes) GetChainPrefix(chainName string) ([]byte, error) {
	if len(bt.Registry) == 0 {
		return defaultPrefix, fmt.Errorf("do not set base type registry")
	}
	for _, reg := range bt.Registry {
		if strings.ToLower(reg.Network) == strings.ToLower(chainName) {
			return big.NewInt(reg.Prefix).Bytes(), nil
		}
	}
	return defaultPrefix, errors.New("do not find this chain decimal")
}
