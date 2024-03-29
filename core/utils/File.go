package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// 文件工具
type fileUtil struct {
}

var File = new(fileUtil)

// 获取运行程序所在的路径
func (util *fileUtil) GetRunPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// 判断文件是否存在
func (util *fileUtil) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 创建目录
func (util *fileUtil) Mkdir(path string) error {
	if !util.Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

// 读取文件内所有数据
func (util *fileUtil) ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// 写文件
func (util *fileUtil) WriteFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

// 删除文件
func (util *fileUtil) DeleteFile(filename string) error {
	return os.Remove(filename)
}

// 获取目录下所有的文件列表（仅遍历，非递归）
// dir 绝对路径
// withDir 返回结果是否包含文件夹
func (util *fileUtil) ScanDir(dir string, withDir bool) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dir)
}
