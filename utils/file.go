package utils

import "os"

//文件相关函数

/**
判断文件是否存在,存在返回false,不存在返回true
*/
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

/**
新建问价夹
*/
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

/**
打开文件
*/
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
