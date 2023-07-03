package cretutil

import (
	"fmt"
	"log"
	"os"
	"time"
)

type OutTxt struct {
}

func NewOutTxt() Oututil {
	oututil := OutTxt{}
	return &oututil
}

func (o *OutTxt) Out(hfs []HashFileStruct) error {
	file, err := os.Create(fmt.Sprintf("./%s.txt", time.Now().Format("20060102")))
	if err != nil {
		log.Printf("创建TXT文件失败: %v", err)
		return err
	}
	defer file.Close()
	for _, v := range hfs {
		file.WriteString(fmt.Sprintf("%s   %s: %s\n", v.FilePath, v.HashType, v.HashStr))
	}
	return nil
}

func init() {
	OutRegister("TXT", NewOutTxt)
}
