package file

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"runtime"
)

/**
 * 获取文件字节大小
 * @param string fileName file name
 * @return int
 */
func FileSize(fileName string) int {
	if info, err := os.Stat(fileName); err == nil {
		return int(info.Size())
	}
	return 0
}

/**
 * 获取文件md5
 * @param string fileName file name
 * @return （string, error）
 */
func FileMd5(fileName string) (string, error) {
	fileObj, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer fileObj.Close()
	reader := bufio.NewReaderSize(fileObj, 1024*1024*1)
	md5 := md5.New()
	_, err = io.Copy(md5, reader)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5.Sum(nil)), nil
}

/**
 * 当前文件
 */
func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}
