package cretutil

import "testing"

func TestMD5(t *testing.T) {
	filePath := []string{"E://demo.txt", "E://demo.txt"}

	certUtil, _ := NewCretutil("MD5", nil)

	t.Log(certUtil.HashFiles(filePath))

}
