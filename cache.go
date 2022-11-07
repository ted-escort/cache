package cache

import (
	"github.com/ted-escort/utils"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

// Dir the directory to store cache files
func Dir() string {
	return "./data/cache/"
}

// FileSuffix cache file suffix. Defaults to '.bin'.
func FileSuffix() string {
	return ".bin"
}

// DirMode the permission to be set for newly created directories.
func DirMode() int {
	return 0775 // 0666
}

// Name the name of cache file
func Name(key string) string {
	return utils.Md5(key)
}

// cacheDirInit 缓存目录初始化
func cacheDirInit() (string, error) {
	// 缓存目录
	fileDir := Dir()
	// 创建目录
	createDir, _ := utils.CreateDir(fileDir)
	if !createDir {
		return "", nil
	}
	return fileDir, nil
}

// File 缓存文件
func File(key string) (string, error) {
	// 缓存目录
	fileDir, err := cacheDirInit()
	if err != nil {
		return "", err
	}
	// 缓存文件
	fileName := Name(key)
	// 后缀
	fileSuffix := FileSuffix()
	// 文件完整路径
	filename := fileDir + fileName + fileSuffix
	if !utils.FileIsExist(filename) {
		_, err := os.Create(filename)
		//_, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fs.FileMode(DirMode()))
		//defer file.Close()
		if err != nil {
			return "", err
		}
	}
	return filename, nil
}

// Get 获取缓存
func Get(key string) ([]byte, error) {
	cacheFile, err := File(key)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Stat(cacheFile)
	if err != nil {
		return nil, err
	}
	// 保质期内
	if fileInfo.ModTime().Unix() > utils.Timestamp() {
		bytes, err := utils.ReadFile(cacheFile)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	}
	return nil, utils.NewError("获取失败")
}

// Set 设置缓存
func Set(key string, value []byte) (bool, error) {
	// 获取缓存文件路径
	cacheFile, err := File(key)
	if err != nil {
		return false, err
	}
	// 写入缓存内容
	err = ioutil.WriteFile(cacheFile, value, fs.FileMode(DirMode()))
	if err != nil {
		return false, err
	}
	//duration := 31536000
	// 设置一年的“修改日期”
	err = os.Chtimes(cacheFile, time.Now(), time.Now().AddDate(1, 0, 0))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Delete 删除缓存
func Delete(key string) error {
	// 获取缓存文件路径
	cacheFile, err := File(key)
	if err != nil {
		return err
	}
	err = os.Remove(cacheFile)
	if err != nil {
		return err
	}
	return nil
}
