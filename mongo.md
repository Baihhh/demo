[MongoDB超详细教程（保姆级）_mongodb教程-CSDN博客](https://blog.csdn.net/efew212efe/article/details/124524863)

## 基础

（1）数据量大

（2）写入操作频繁（读写都很频繁）

（3）**价值较低的数据，对事务性要求不高**

对于这样的数据，我们更适合使用 MongoDB来实现数据的存储


应用不需要事务及复杂join支持

新应用，需求会变，数据模型无法确定，想快速迭代开发

应用需要2000-3000以上的读写QPS（更高也可以）

应用需要TB甚至PB级别数据存储

应用要求存储的数据不丢失

应用需要99.999%高可用

应用需要大量的地理位置查询、文本查


### sql vs mongo

![在这里插入图片描述](https://img-blog.csdnimg.cn/cbb1e84fe31140869e830da43d2bd352.png)



### BSON

BSON（Binary Serialized Document Format）是一种类json的一种**二进制形式的存储格式**，简称 Binary JSON；BSON和JSON一样，支持内嵌的文档对象和数组对象，但是BSON有JSON没有的一些数据类型，如Date和Bin Data类型。

BSON采用了类似于C语言结构体的名称、对表示方法，支持内嵌的文档对象和数组对象，具有**轻量性、可遍历性、高效性**的三个特点，可以**有效描述非结构化数据和结构化数据**。这种格式的优点是灵活性高，但它的**缺点是空间利用率不是很理想**。

BSON中，除了基本JSON类型： string，integer，boolean，double，null，array和object，mongo还使用了特殊的数据类型。这些类型包括 date， object id， binary data， regular expression和code。每一个驱动都以特定语言的方式实现了这些类型，查看你的驱动的文档来获取详细信息

BSON数据类型参考列表：

![在这里插入图片描述](https://img-blog.csdnimg.cn/287a55514fe441a69eca966656081761.png)

提示：
shell默认使用64位浮点型数值。{“x”:3.14或{“x”:3}。对于整型值，可以使用NumberInt（4字节符号整数）或 NumberLong（8字节符号整数），{“x”:NumberInt(“3” ){“x”:NumberLong(“3”)}

### 特点

**高性能**

MongoDB提供高性能的数据持久性。特别是对嵌入式数据模型的支持减少了数据库系统上I/O活动。

索引支持更快的查询，并且可以包含来自嵌入式文档和数组的键。

- 文本索引解决搜索的需求、TTL索引解决历史数据自动过期的需求、地理位置索引可用于构建各种O2O应用

- mmapv1、 wiredtiger、 mongorocks（ rocks）、 In-memory等多引擎支持满足各种场景需求

- **Gridfs解决文件存储的需求**

**高可用性**

MongoDB的复制工具称为副本集（ replica set），它可提供自动故障转移和数据冗余

**高扩展性**
MongoDB提供了水平可扩展性作为其核心功能的一部分。

分片将数据分布在一组集群的机器上。（海量数据存储，服务能力水平扩展）

从3.4开始，MoηgoDB支持基于片键创建数据区域。在一个平衡的集群中， MongoDB将一个区域所覆盖的读写只定向到该区域内的那些片。

**丰富的查询支持**

MongoDB支持丰富的査询语言，支持读和写操作（CRUD），比如数据聚合、文本搜索和地理空间查询等



## gridFS

[GridFS — Go Driver (mongodb.com)](https://www.mongodb.com/docs/drivers/go/current/fundamentals/gridfs/#std-label-golang-gridfs)

GridFS 不是将文件存储在单个文档中，而是将文件分开 分成部分或块 [[1\]](https://www.mongodb.com/docs/v4.4/core/gridfs/#footnote-chunk-disambiguation)，并将每个块存储为 一个单独的文档。默认情况下，GridFS 使用默认的块大小为 255 kB; 也就是说，GridFS 将文件划分为 255 kB 的块，但例外 最后一个块。最后一个块仅根据需要大小。 同样，不大于块大小的文件只有一个 最后一个块，仅使用所需的空间加上一些额外的空间 元数据。

GridFS 使用两个集合来存储文件。**一个集合存储 文件块，其他存储文件元数据**。该部分[GridFS 集合](https://www.mongodb.com/docs/v4.4/core/gridfs/#std-label-gridfs-collections)详细描述了每个集合。

当您在 GridFS 中查询文件时，驱动程序将重新组合块 根据需要。您可以对通过 GridFS 存储的文件执行范围查询。 您还可以从文件的任意部分访问信息，例如 至于“跳到”视频或音频文件的中间。

**GridFS 不仅可用于存储超过 16 MB 的文件，还可用于存储 用于存储您想要访问的任何文件，而无需加载 将整个文件放入内存中。**

### GridFS 的工作原理[![img](https://www.mongodb.com/docs/drivers/go/current/assets/link.svg)](https://www.mongodb.com/docs/drivers/go/current/fundamentals/gridfs/#how-gridfs-works)

GridFS 将文件组织在一个**存储桶**中，存储桶是一组 MongoDB 集合 包含文件块和描述它们的信息。这 存储桶包含以下集合：

- 存储二进制文件块的集合。`chunks`
- 存储文件元数据的集合。`files`

当您创建新的 GridFS 存储桶时，驱动程序会创建上述存储桶 收集。默认存储桶名称为集合名称前缀



### 上传文件

您可以通过以下方式之一将文件上传到 GridFS 存储桶中：

- 使用从输入流读取的方法。`UploadFromStream()`   （reads from an input stream）
- 使用写入输出流的方法。`OpenUploadStream()`  （writes to an output stream）

[options package - go.mongodb.org/mongo-driver/mongo/options - Go Packages](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.14.0/mongo/options#UploadOptions)

#### UploadFromStream

To upload a file with an input stream

**参数**

- file name
- file as a parameter`io.Reader`
- An optional parameter to modify the behavior of `opts``UploadFromStream()`

```go
file, err := os.Open("path/to/file.txt")
uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "first"}})

objectID, err := bucket.UploadFromStream("file.txt", io.Reader(file),
   uploadOpts)
if err != nil {
   panic(err)
}

fmt.Printf("New file uploaded with ID %s", objectID)
```



#### OpenUploadStream

To upload a file with an output stream

**参数**

- file name
- An optional `opts` parameter to modify the behavior of `OpenUploadStream()`

```go
file, err := os.Open("path/to/file.txt")
if err != nil {
	panic(err)
}

// Defines options that specify configuration information for files
// uploaded to the bucket
uploadOpts := options.GridFSUpload().SetChunkSizeBytes(200000)

// Writes a file to an output stream
uploadStream, err := bucket.OpenUploadStream("file.txt", uploadOpts)
if err != nil {
	panic(err)
}
fileContent, err := io.ReadAll(file)
if err != nil {
	panic(err)
}
var bytes int
if bytes, err = uploadStream.Write(fileContent); err != nil {
	panic(err)
}
fmt.Printf("New file uploaded with %d bytes written", bytes)
```

### 下载文件

- 使用该方法将文件下载到输出流。`DownloadToStream()`  download a file to an output stream
- 使用该方法打开输入流。`OpenDownloadStream()`    open an input stream

#### DownloadToStream

```go
id, err := primitive.ObjectIDFromHex("62f7bd54a6e4452da13b3e88")
fileBuffer := bytes.NewBuffer(nil)
if _, err := bucket.DownloadToStream(id, fileBuffer); err != nil {
   panic(err)
}
```

#### OpenDownloadStream

```go
id, err := primitive.ObjectIDFromHex("62f7bd54a6e4452da13b3e88")
downloadStream, err := bucket.OpenDownloadStream(id)
if err != nil {
   panic(err)
}

fileBytes := make([]byte, 1024)
if _, err := downloadStream.Read(fileBytes); err != nil {
   panic(err)
}
```



### 获取文章信息

**文章信息**

- The file ID
- The file length
- The maximum chunk size
- The upload date and time
- The file name
- A document in which you can store any other information`metadata`

该方法需要查询筛选器作为参数。匹配所有 集合中的文档，将空查询筛选器传递给 。`Find()``files``Find()`

```go
// 根据条件过滤
filter := bson.D{{"length", bson.D{{"$gt", 1500}}}}
cursor, err := bucket.Find(filter)
if err != nil {
   panic(err)
}

type gridfsFile struct {
   Name   string `bson:"filename"`
   Length int64  `bson:"length"`
}
var foundFiles []gridfsFile
if err = cursor.All(context.TODO(), &foundFiles); err != nil {
   panic(err)
}

for _, file := range foundFiles {
   fmt.Printf("filename: %s, length: %d\n", file.Name, file.Length)
}
```

 