package test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func WriterIO() {
	// ytes.Buffer是一个可变字节的类型，可以让我们很容易的对字节进行操作，比如读写、追加等。bytes.Buffer实现了io.Writer和io.Reader接口，所以我们可以很容易地进行读写操作，而不用关注具体实现。
	var b bytes.Buffer
	b.Write([]byte("你好"))
	fmt.Fprint(&b, ",", "hello")
	b.WriteTo(os.Stdout) // 你好,hello
}

func WriterWriter() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 strings.Builder
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}

func WriterStringWriter() {
	w := os.Stdout

	n, err := io.WriteString(w, "hello\n")

	if err != nil {
		panic(err)
	}

	fmt.Printf("n:%d\n", n)
}

func WriterAvailable() {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	// 刷新
	w.Flush()
}

func WriterBytes() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	// ReadByte 从缓冲区读取并返回下一个字节。 如果没有可用的字节，则返回io.EOF。
	c, err := b.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	fmt.Println(b.String())
}
