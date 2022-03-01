package upload

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	// 文件 key
	uploadFileKey = "file"
)

//func main() {
//	http.HandleFunc("/upload", uploadHandler)
//	if err := http.ListenAndServe(":82", nil); err != nil {
//		log.Fatalf("error to start http server:%s", err.Error())
//	}
//}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 接受文件
	file, header, err := r.FormFile(uploadFileKey)
	if err != nil {
		// ignore the error handler
	}
	Upload(file, header.Filename)
}

func Upload(file multipart.File, filename string) (filePath string, fullFilePath string) {
	log.Printf("selected file name is: %s", filename)
	uploadsPath := "/Users/liyuan/LY/projects/uploads/dcswitch/"
	folderName := time.Now().Format("20060102150405")
	folderPath := filepath.Join(uploadsPath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(folderPath, 0777) //0777也可以os.ModePerm
		//os.Chmod(folderPath, 0777)
	}
	// 将文件拷贝到指定路径下，或者其他文件操作
	dst, err := os.Create(filepath.Join(folderPath, filename))
	if err != nil {
		// ignore
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		// ignore
	}
	return filepath.Join(folderName, filename), filepath.Join(folderPath, filename)
}
