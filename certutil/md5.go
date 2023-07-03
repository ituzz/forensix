package cretutil

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Md5Util struct {
}

func NewMd5Util() Certutil {

	certutil := Md5Util{}
	return &certutil
}

func (m *Md5Util) HashFile(filePath string) ([]HashFileStruct, error) {
	fileData, err := readFIle(filePath)
	if err != nil {
		return nil, err
	}
	md5Bytes := md5.Sum(fileData)
	md5Str := fmt.Sprintf("%x", md5Bytes)
	hfs := make([]HashFileStruct, 1)
	hfs[0] = newHashFIleStruct(filePath, md5Str)
	return hfs, nil
}

func (m *Md5Util) HashFiles(filePath []string) []HashFileStruct {
	hfs := make([]HashFileStruct, len(filePath))
	for k, v := range filePath {
		fileData, err := readFIle(v)
		if err == nil {
			md5Str := fmt.Sprintf("%x", md5.Sum(fileData))
			hfs[k] = newHashFIleStruct(v, md5Str)
		} else {
			hfs[k] = newHashFIleStruct(v, "ERROR")
		}
	}
	return hfs
}

func newHashFIleStruct(filePath string, HashStr string) HashFileStruct {
	return HashFileStruct{
		FilePath: filePath,
		HashStr:  HashStr,
		HashType: "MD5",
	}
}

func readFIle(filePath string) ([]byte, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("计算MD5-文件不存在或者无权限: %v", err)
		return nil, err
	}
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("计算MD5-文件读取错误: %v", err)
		return nil, err
	}
	return fileData, nil
}

func init() {
	CertRegister("MD5", NewMd5Util)
}
