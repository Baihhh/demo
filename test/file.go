package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var file_name string = "F:\\实习\\kodo\\mongoDemo\\test\\test.txt"
var file_write string = "F:\\实习\\kodo\\mongoDemo\\test\\writer.txt"

func FileOpen() {
	file, e := os.Open(file_name)
	if e != nil {
		fmt.Println(e)
	}
	buf := make([]byte, 1024)
	for {
		len, _ := file.Read(buf)
		if len == 0 {
			break
		}
		fmt.Println(string(buf))
	}

	buf1 := make([]byte, 1024)
	offset := 0
	for {
		// 方法可以手动指定每次读取位置的偏移量。而不是默认设置。
		len1, _ := file.ReadAt(buf1, int64(offset))
		offset = offset + len1
		if len1 == 0 {
			break
		}
		fmt.Println(string(buf1))
	}
	file.Close()
}

func FileOpenFile() {
	openFile, e := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if e != nil {
		fmt.Println(e)
	}
	buf := make([]byte, 32)
	for {
		len, _ := openFile.Read(buf)
		if len == 0 {
			break
		}
		fmt.Println(string(buf))
	}
	openFile.Close()
}

func FileBufferReader() {
	openFile, e := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0777)
	if e != nil {
		fmt.Println(e)
	}
	reader := bufio.NewReader(openFile)
	for {
		line, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(string(line))
	}
	// //按照指定分隔符读取
	// for {
	// 	s, e := reader.ReadString('\n')
	// 	fmt.Println(s)
	// 	if e == io.EOF {
	// 		break
	// 	}
	// 	if e != nil {
	// 		fmt.Println(e)
	// 	}
	// }
	openFile.Close()
}

func FileWriteString() {
	openFile, e := os.OpenFile(file_write, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if e != nil {
		fmt.Println(e)
	}
	str := "overwrite to file"
	openFile.WriteString(str)
	openFile.Close()
}

func FileBufferWriter() {
	f, err := os.OpenFile(file_write, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	content := map[string]string{
		"hello":  "111",
		"world":  "222",
		"world1": "333",
		"world2": "444",
	}
	bw := bufio.NewWriter(f)
	for k, v := range content {
		bw.WriteString(k + ":" + v + "\n")
	}
	bw.Flush()
}
