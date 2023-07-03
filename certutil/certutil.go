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
	HashFiles(filePath []string) []HashFileStruct

	//对单个文件进行哈希
	HashFile(filePath string) ([]HashFileStruct, error)
}

type CertInstance func() Certutil

var certAdapters = make(map[string]CertInstance)

// 注册哈希类型
func CertRegister(hashType string, adapter CertInstance) {
	if adapter == nil {
		panic(fmt.Errorf("Certutil: Register adapter is nil"))
	}
	if _, ok := certAdapters[hashType]; ok {
		panic("Certutil: Register called twice for adapter " + hashType)
	}
	certAdapters[hashType] = adapter
}

func NewCretutil(adapterName string, config interface{}) (adapter Certutil, err error) {
	instanceFunc, ok := certAdapters[adapterName]
	if !ok {
		log.Printf("Certutil: unknown adapter name %s", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
