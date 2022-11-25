package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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
