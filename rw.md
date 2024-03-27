在 Go 中，输入输出操作是通过能读能写的**字节流**数据模型来实现的

> java中有字节流也有字符流

[io package - io - Go Packages](https://pkg.go.dev/io)   io相关



### Reader 接口

1. **io.Reader**: 最常见的 Reader **接口**，用于从**数据源读取数据**。

   ```go
   type Reader interface {
   	Read(p []byte) (n int, err error)
   }
   
   func LimitReader(r Reader, n int64) Reader
   func MultiReader(readers ...Reader) Reader
   func TeeReader(r Reader, w Writer) Reader
   // TeeReader 返回一个 Reader，该 Reader 将它从 r 读取的内容写入 w。 通过它执行的所有从 r 读取的内容都与 对应写入 W。没有内部缓冲 - 写入必须在读取完成之前完成。 写入时遇到的任何错误都报告为读取错误。
   ```

   

2. **io.ByteReader**: 也是一个接口，在 io.Reader 基础上增加了读取单个**字节**的方法。

   ```go
   type ByteReader interface {
       // ReadByte 从数据源读取一个字节
       ReadByte() (byte, error)
   }
   ```

   

3. **io.RuneReader**: 在 io.Reader 基础上增加了读取单个 **Unicode 字符**的方法。如果没有字符可用，将设置 ERR

   ```go
   type RuneReader interface {
   	ReadRune() (r rune, size int, err error)
   }
   ```

   

4. **bufio.Reader**: 对 io.Reader 进行**缓冲读取**，提高读取效率。

   ```go
   func (b *Reader) Buffered() int  // 返回可以read的字节个数
   func (b *Reader) Discard(n int) (discarded int, err error)   // 跳过 超过则报错
   func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
   func (b *Reader) ReadBytes(delim byte) ([]byte, error)
   func (b *Reader) ReadRune() (r rune, size int, err error)
   func (b *Reader) ReadString(delim byte) (string, error)
   ```

   

5. **strings.Reader**: 从**字符串**中读取数据的 Reader。

   ```go
   func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
   func (r *Reader) ReadByte() (byte, error)
   ```

   

6. **bytes.Reader**： 通常用于对字节片段进行读取，例如从内存中的字节片段中读取数据。



### Writer 接口

1. **io.Writer**: 最常见的 Writer 接口，用于向数据目标写入数据。

   ```go
   type Writer interface {
   	Write(p []byte) (n int, err error)
   }
   
   func MultiWriter(writers ...Writer) Writer  //多写入
   ```

   

2. **io.ByteWriter**: 在 io.Writer 基础上增加了写入单个字节的方法。

   ```go
   type ByteWriter interface {
   	WriteByte(c byte) error
   }
   ```

   

3. **io.StringWriter**: 在 io.Writer 基础上增加了写入字符串的方法。

   ```go
   type StringWriter interface {
   	WriteString(s string) (n int, err error)
   }
   ```

   

4. **bufio.Writer**: 对 io.Writer 进行缓冲写入，提高写入效率。

   ```go
   func (b *Writer) Available() int     // 返回可用的字节数
   func (b *Writer) AvailableBuffer() []byte   //AvailableBuffer 返回一个容量为 b.Available（） 的空缓冲区。 此缓冲区旨在追加到 和 传递给紧随其后的 Writer.Write 调用。 缓冲区仅在对 b 执行下一次写入操作之前有效。
   func (b *Writer) Buffered() int   // 返回已经写了的字节数
   
   //也有和 io.Writer相同的 如WriteByte、 WriteRune等
   ```

   

5. **strings.Builder**：对字符串的写入操作

   ```go
   func (b *Builder) WriteString(s string) (int, error)
   func (b *Builder) WriteRune(r rune) (int, error)
   func (b *Builder) WriteByte(c byte) error
   ```

   

6. **bytes.Buffer**: 在内存中创建的字节缓冲区，实现了 io.ReaderFrom 和 io.WriterTo 接口，可以方便地进行读写。

   类似



### 适用场景


各种 Reader 和 Writer 在 Go 中有着不同的使用场景，它们通常用于不同的输入/输出源以及数据处理需求。以下是对 `io.Reader`、`bufio.Reader` 和 `strings.Reader` 的常见使用场景的简要描述：

####  `io`

- **通用读写**：io.Reader /  io.Writer 是一个通用的接口，可以从任何实现了该接口的数据源中读取数据和写入数据。适用于从文件、网络连接、内存缓冲区等各种数据源。

- **通用输入源读取**: `io.Reader` 是一个通用的接口，可以从任何实现了该接口的数据源中读取数据。它可以从文件、网络连接、内存缓冲区等各种输入源读取数据。
- **数据流处理**: 当你需要处理数据流，并且不关心数据源的具体类型时，可以使用 `io.Reader或者io.Writer`。例如，在解析 HTTP 请求体或者处理加密数据时就会用到它。

#### `bufio`

- **缓冲读写**:  带有一个内部缓冲区，可以提高对底层数据源的读取效率。它适用于大部分需要**大量数据**的场景，比如文件读写、网络数据流等。
- **逐行读取文本**: `bufio.Reader` 提供了 `ReadString` 和 `ReadBytes` 等方法，可以方便地逐行读取文本数据。这在处理文本文件或者协议中的文本数据时非常有用。

#### `strings`

- **从字符串读取数据**: `strings.Reader` 将字符串包装为 `io.Reader` 接口，可以方便地从字符串中读取数据。它常用于将字符串作为输入源，并进行处理或解析。
- **测试和模拟**: 在测试场景中，有时我们需要模拟文件或网络连接的读取行为。使用 `strings.Reader` 可以轻松地创建一个模拟的输入源，以便进行单元测试。





**io**

Go标准库的io包也是基于Unix这种输入和输出的理念，大部分的接口都是扩展了io.Writer和io.Reader，大部分的类型也都选择地实现了io.Writer和io.Reader这两个接口，然后把数据的输入和输出，抽象为流的读写。所以只要实现了这两个接口，都可以使用流的读写功能。

io.Writer和io.Reader两个接口的高度抽象，让我们不用再面向具体的业务，我们只关注，是读还是写。只要我们定义的方法函数可以接收这两个接口作为参数，那么我们就可以进行流的读写，而不用关心如何读、写到哪里去，这也是面向接口编程的好处。



### **string vs bytes**

- string可以直接比较，而 []byte不可以，所以 []byte不可以当map的key值。
- 因为无法修改string中的某个字符，需要粒度小到操作一个字符时，用 []byte。
- string值不可为nil，所以如果你想要通过返回nil表达额外的含义，就用 []byte。
- []byte切片这么灵活，想要用切片的特性就用 []byte。
- 需要大量字符串处理的时候用 []byte，性能好很多。







1. **文件读取与写入**：
   - 读取文件内容并写入其他地方：使用 `os.File` 打开文件作为 Reader，然后使用其他的 Writer（如 `os.Stdout`、`os.Stdout` 或 `bytes.Buffer`）将内容写入到其他位置。
   - 从一个文件复制到另一个文件：使用 `os.File` 分别打开两个文件，一个作为 Reader，另一个作为 Writer，然后通过 `io.Copy` 函数将内容从一个文件复制到另一个文件。
2. **网络通信**：
   - 从网络连接读取数据：使用 `net.Conn` 作为 Reader，然后将数据写入到适当的 Writer（如 `os.Stdout`、`os.Stdout` 或 `bytes.Buffer`）。
   - 向网络连接写入数据：使用 `net.Conn` 作为 Writer，然后从适当的 Reader（如 `os.Stdin`、`os.File` 或 `bytes.Buffer`）读取数据并写入到网络连接中。
3. **HTTP 请求和响应**：
   - 从 HTTP 请求体中读取数据：使用 `http.Request.Body` 作为 Reader，然后将数据写入到其他地方（如文件、网络连接或内存缓冲区）。
   - 将数据写入 HTTP 响应体中：使用 `http.ResponseWriter` 作为 Writer，然后从其他地方（如文件、网络连接或内存缓冲区）读取数据并写入到响应体中。
4. **内存缓冲区操作**：
   - 使用 `bytes.Buffer` 或 `bytes.Reader` 进行内存数据的读写操作。这些接口通常用于在内存中临时存储和操作数据。
5. **字符串操作**：
   - 使用 `strings.Reader` 从字符串中读取数据，并将数据写入到其他地方（如文件、网络连接或内存缓冲区）。
   - 使用 `io.WriteString` 将字符串写入到 Writer（如文件、网络连接或内存缓冲区）中。

