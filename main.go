package ufc

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// PathExist 检测路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PathNotExist 检测路径是否不存在
func PathNotExist(path string) bool {
	return !PathExist(path)
}

// IsDir 是否为目录
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()

}

// IsFile 是否为文件（含普通、设备）
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// IsCommonFile 是否为普通文件
func IsCommonFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}

// CreateDir 创建文件夹，如果已存在则直接返回，否则按照0755权限创建
func CreateDir(path string) error {
	if PathNotExist(path) {
		return os.Mkdir(path, os.ModeDir)
	}
	return nil
}

// FileReadByte 读取指定的文件，并返回[]byte
func FileReadByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	return ioutil.ReadAll(fi)
}

// FileReadStr 读取指定的文件，并返回字符串
func FileReadStr(path string) (string, error) {
	raw, err := FileReadByte(path)
	return string(raw), err
}

// FileCopy 复制文件，会自动创建并覆盖目标文件
func FileCopy(dstName, srcName string) (written int64, err error) {
	if !IsFile(srcName) {
		return 0, errors.New("src file does not exist")
	}

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// FileCopyN 按字节复制文件，会自动创建并覆盖目标文件
func FileCopyN(dstName, srcName string, n int64) (written int64, err error) {
	if !IsFile(srcName) {
		return 0, errors.New("src file does not exist")
	}

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.CopyN(dst, src, n)
}

// IsTrue 1、t、true、True、TRUE将返回布尔值true，其他（包含发送错误）返回布尔值false
func IsTrue(v string) bool {
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}

// IsFalse 非IsTrue则是false
// 如0、f、false、False、FALSE、其他字符串或发生错误时都返回布尔值true
func IsFalse(v string) bool {
	return !IsTrue(v)
}
