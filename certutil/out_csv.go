package cretutil

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type OutCsv struct {
}

func NewOutCsv() Oututil {
	oututil := OutCsv{}
	return &oututil
}

func (o *OutCsv) out(hfs []HashFileStruct) error {
	file, err := os.Create(fmt.Sprintf("./%s.csv", time.Now().Format("20060102")))
	if err != nil {
		log.Printf("创建CSV文件失败: %v", err)
		return err
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	writer.Write([]string{"编号", "文件名", hfs[0].HashType})
	for k, v := range hfs {
		writer.Write([]string{strconv.Itoa(k + 1), v.FilePath, v.HashStr})
	}
	writer.Flush()
	return nil
}

func init() {
	OutRegister("CSV", NewOutCsv)
}
