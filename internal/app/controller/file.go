package controller

import (
	"bufio"
	"db-go-gin/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

var uploadTempPath = path.Join("./static", "temp")

// UploadFile 上传文件
func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	index := ctx.PostForm("index")
	hash := ctx.PostForm("hash")

	// 获取uploads下所有的文件夹
	nameList, err := ioutil.ReadDir("./static")
	// 循环判断hash是否在文件里如果有就返回上传已完成
	for _, name := range nameList {
		tmpName := strings.Split(name.Name(), "-")[0]
		if tmpName == hash {
			return
		}
	}

	chunksPath := path.Join(uploadTempPath, hash, "/")
	isPathExists, err := utils.PathExists(chunksPath)
	if !isPathExists {
		err = os.MkdirAll(chunksPath, os.ModePerm)
	}
	destFile, err := os.OpenFile(path.Join(chunksPath+"/"+hash+"-"+index), syscall.O_CREAT|syscall.O_WRONLY, 0777)
	f, _ := file.Open()
	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(destFile)
	buf := make([]byte, 1024*1024) // 1M buf
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			writer.Flush()
			break
		} else if err != nil {
			return
		} else {
			writer.Write(buf[:n])
		}
	}

	defer f.Close()
	defer destFile.Close()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("第%s:%s块上传完成\n", index, destFile.Name())
}

// Chunks 合并文件
func Chunks(ctx *gin.Context) {
	size, _ := strconv.ParseInt(ctx.PostForm("size"), 10, 64)
	hash := ctx.PostForm("hash")
	name := ctx.PostForm("name")

	toSize, _ := utils.GetDirSize(path.Join(uploadTempPath, hash, "/"))
	if size != toSize {
		fmt.Fprintf(ctx.Writer, "文件上传错误")
	}
	chunksPath := path.Join(uploadTempPath, hash, "/")
	files, _ := ioutil.ReadDir(chunksPath)
	// 排序
	filesSort := make(map[string]string)
	for _, f := range files {
		nameArr := strings.Split(f.Name(), "-")
		filesSort[nameArr[1]] = f.Name()
	}
	saveFile := path.Join("./static", name)
	if exists, _ := utils.PathExists(saveFile); exists {
		os.Remove(saveFile)
	}
	fs, _ := os.OpenFile(saveFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	var wg sync.WaitGroup
	filesCount := len(files)
	if filesCount != len(filesSort) {
		fmt.Fprintf(ctx.Writer, "文件上传错误2")
	}
	wg.Add(filesCount)
	for i := 0; i < filesCount; i++ {
		// 这里一定要注意按顺序读取不然文件就会损坏
		fileName := path.Join(chunksPath, "/"+filesSort[strconv.Itoa(i)])
		data, err := ioutil.ReadFile(fileName)
		fmt.Println(err)
		fs.Write(data)

		wg.Done()
	}
	wg.Wait()
	os.RemoveAll(path.Join(chunksPath, "/"))
	defer fs.Close()
}
