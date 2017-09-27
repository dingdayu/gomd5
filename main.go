package main

import (
	"fmt"
	"crypto/md5"
	"os"
	"io"
	"time"
	"math"
	"flag"
)

const filechunk = 8192 // we settle for 8KB

func main() {

	filePath := flag.String("file", "", "需要计算MD5的文件路径")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("请输入文件路径！")
		return
	}

	fmt.Println("整文件MD5:")
	md51(*filePath)
	fmt.Println("分块文件MD5:")
	md52(*filePath)
}

func md51(filePath string)  {
	start_time := time.Now() // get current time
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	md5h := md5.New()
	io.Copy(md5h, file)
	fmt.Printf("File MD5: %x \n", md5h.Sum([]byte(""))) //md5

	// 文件大小
	fileInfo, err := os.Stat(filePath)
	fileSize := fileInfo.Size() //获取size
	fmt.Println("File Size: " , fileSize / 1024 / 1024, "M")

	elapsed := time.Since(start_time)
	fmt.Println("Elapsed: ", elapsed)
}

func md52(filePath string)  {
	start_time := time.Now() // get current time
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 文件大小
	fileInfo, err := os.Stat(filePath)
	fileSize := fileInfo.Size() //获取size
	fmt.Println("File Size: " , fileSize / 1024 / 1024, "M")

	blocks := uint64(math.Ceil(float64(fileSize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(fileSize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	fmt.Printf("File MD5: %x\n", hash.Sum(nil))

	elapsed := time.Since(start_time)
	fmt.Println("Elapsed: ", elapsed)
}