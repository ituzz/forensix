package cretutil

import (
	"fmt"
	"log"
)

type Oututil interface {
	out(hfs []HashFileStruct) error
}

type OutInstance func() Oututil

var outAdapters = make(map[string]OutInstance)

// 注册
func OutRegister(outType string, adapter OutInstance) {
	if adapter == nil {
		panic(fmt.Errorf("Oututil: Register adapter is nil"))
	}
	if _, ok := outAdapters[outType]; ok {
		panic("Oututil: Oututil called twice for adapter " + outType)
	}
	outAdapters[outType] = adapter
}

func NewOututil(adapterName string, config interface{}) (adapter Oututil, err error) {
	instanceFunc, ok := outAdapters[adapterName]
	if !ok {
		log.Printf("Oututil: unknown adapter name %s", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
