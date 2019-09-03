package file

import (
	"os"
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
