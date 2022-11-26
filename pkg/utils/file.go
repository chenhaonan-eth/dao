package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			// global.LOGGER.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				// global.LOGGER.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}

func ReaderText(readerfile string) (string, error) {
	//打开文件
	file, err := os.Open(readerfile)
	if err != nil {
		fmt.Println("Read file err, err =", err)
		return "", err
	}
	defer file.Close()
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return "", err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	return string(chunk), nil
}

func WirteText(savefile string, txt string) {
	f, err := os.OpenFile(savefile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	bw.WriteString(txt + "\n")
	bw.Flush()
}
