package cretutil

import (
	"fmt"
	"log"
)

type HashFileStruct struct {
	HashStr  string
	FilePath string
	HashType string
}

type Certutil interface {

	//对多文件进行哈希
	HashFiles(filePath []string) ([]HashFileStruct, error)

	//对单个文件进行哈希
	HashFile(filePath string) ([]HashFileStruct, error)
}

type Instance func() Certutil

var adapters = make(map[string]Instance)

// 注册哈希类型
func Register(hashType string, adapter Instance) {
	if adapter == nil {
		panic(fmt.Errorf("Certutil: Register adapter is nil"))
	}
	if _, ok := adapters[hashType]; ok {
		panic("Certutil: Register called twice for adapter " + hashType)
	}
	adapters[hashType] = adapter
}

func NewCretutil(adapterName string, config interface{}) (adapter Certutil, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		log.Printf("Certutil: unknown adapter name %s", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
