package test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReaderLimit() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err) // some
	}
}

func ReaderMulti() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err) // first reader second reader third reader
	}
}

func ReaderTee() {
	// TeeReader 返回一个 Reader，该 Reader 将它从 r 读取的内容写入 w。 通过它执行的所有从 r 读取的内容都与 对应写入 W。没有内部缓冲 - 写入必须在读取完成之前完成。 写入时遇到的任何错误都报告为读取错误。
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}
}

func ReaderRune() {
	reader := strings.NewReader("Hello, 世界")

	runeReader := io.RuneReader(reader)

	// 读取单个 Unicode 字符并打印
	for {
		r, size, err := runeReader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		fmt.Printf("Rune: %c, Size: %d bytes\n", r, size)
	}
}

func ReaderSection() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err) //io.Reader stream
	}
}

func ReaderReadLine() {
	reader := strings.NewReader("Line 1\nLine 2\nLine 3\n")
	bufReader := bufio.NewReader(reader)

	for {
		line, isPrefix, err := bufReader.ReadLine()
		if err != nil {
			break
		}

		fmt.Printf("Line: %s, IsPrefix: %v\n", line, isPrefix)
	}
}
