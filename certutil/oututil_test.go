package cretutil

import "testing"

func TestOut(t *testing.T) {

	filePath := []string{"demo.txt", "demo.txt"}
	certUtil, _ := NewCretutil("MD5", nil)
	hfs := certUtil.HashFiles(filePath)
	outUtil, _ := NewOututil("TXT", nil)
	outUtil.out(hfs)
}
