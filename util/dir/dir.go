package dir

import (
	"os"
	"path/filepath"
)

/**
 *	检查文件或目录是否存在
 *	@param string filename  文件或目录
 *	@return bool
 */
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

/**
 *	创建目录
 *	@param string dirPath 目录
 *	@return error
 */
func MakeDir(dirPath string) error {
	if !IsExist(dirPath) {
		return os.MkdirAll(dirPath, 0755)
	}

	return nil
}

/**
 *	获取文件目录
 *	@param string file
 *	@return string
 */
func GetFilePath(file string) string {
	if stat, err := os.Stat(file); err == nil && stat.IsDir() {
		return file
	}

	return filepath.Dir(file)
}

/**
 *	获取目录所有文件
 *	@param string dir
 *	@return slice
 */
func GetAllFilesByDir(dir string) []string {
	pattern := GetFilePath(dir) + "/*"
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil
	}
	return files
}
