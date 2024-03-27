package test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const (
	Database       = "fileserver"
	CollectionName = "fs"
)

func connectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		bucket, err := gridfs.NewBucket(
			client.Database(Database),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 也可以使用UploadFromStreamWithID   但是需要自己指定id等
		uploadStream, err := bucket.OpenUploadStream(handler.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer uploadStream.Close()

		// 将文件对象的数据从 file 复制到了 GridFS 中，从而实现了文件的上传操作
		if _, err := io.Copy(uploadStream, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 获取上传后的文件 ID
		fileID := uploadStream.FileID
		fmt.Println("File ID:", fileID)

		fmt.Fprintf(w, "File %s uploaded successfully.\nFileID is %s", string(handler.Filename), fileID)
	} else {
		w.Write([]byte(string("bed Method.")))
	}
}

func uploadFile2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		bucket, err := gridfs.NewBucket(
			client.Database(Database),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		objectID, err := bucket.UploadFromStream(handler.Filename, io.Reader(file))
		if err != nil {
			panic(err)
		}

		fmt.Println("File ID:", objectID)
		fmt.Fprintf(w, "File %s uploaded successfully.\nFileID is %s", string(handler.Filename), objectID)
		// ObjectID 转换为十六进制字符串
		// w.Write(objectID[:])
	} else {
		w.Write([]byte(string("bed Method.")))
	}

}

// func downloadFile(w http.ResponseWriter, r *http.Request) {
// 	fileID := r.URL.Path[len("/download/"):]
// 	if fileID == "" {
// 		http.Error(w, "File ID not provided", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Printf(fileID)

// 	bucket, err := gridfs.NewBucket(
// 		client.Database(Database),
// 	)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	objectID, err := primitive.ObjectIDFromHex(fileID)
// 	if err != nil {
// 		http.Error(w, "Invalid file ID", http.StatusBadRequest)
// 		return
// 	}
// 	fileBuffer := bytes.NewBuffer(nil)
// 	downloadStream, err := bucket.DownloadToStream(objectID, fileBuffer)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Fprint(w, downloadStream)
// 	w.Header().Set("Content-Disposition", "attachment; filename="+fileID)
// 	w.Header().Set("Content-Type", "application/octet-stream")

// }

func downloadFile2(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Path[len("/download2/"):]
	if fileID == "" {
		http.Error(w, "File ID not provided", http.StatusBadRequest)
		return
	}

	bucket, err := gridfs.NewBucket(
		client.Database(Database),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.14.0/mongo/gridfs#DownloadStream
	downloadStream, err := bucket.OpenDownloadStream(objectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer downloadStream.Close()

	file := downloadStream.GetFile()

	w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := io.Copy(w, downloadStream); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

type Trainer struct {
	Name string
	Age  int
	City string
}

func TestMongo() {
	client = connectDB()
	defer client.Disconnect(context.Background())
	fmt.Println(client)

	// test
	// collection := client.Database("test").Collection("trainer")
	// fmt.Println(collection)
	// ash := Trainer{"Ash", 10, "Pallet Town"}
	// insertResult, err := collection.InsertOne(context.TODO(), ash)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/upload2", uploadFile2)
	http.HandleFunc("/test", test)
	// http.HandleFunc("/download/", downloadFile)
	http.HandleFunc("/download2/", downloadFile2)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
