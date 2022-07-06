package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 打包成zip文件
func Zip(src_dir string, zip_file_name string) {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zip_file_name)

	// 创建：zip文件
	zipfile, _ := os.Create(zip_file_name)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == src_dir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, src_dir+string(filepath.Separator))

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}

type FileItem struct {
	Name string
	Body []byte
}

func ZipAppendFile(zipPath string, filename string, content []byte) *bytes.Buffer {
	file := FileItem{Name: filename, Body: content}
	return ZipAppendFiles(zipPath, &file)
}

func ZipAppendFiles(zipPath string, files ...*FileItem) *bytes.Buffer {
	zipReader, _ := zip.OpenReader(zipPath)
	defer zipReader.Close()
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	defer w.Close()
	for _, zipItem := range zipReader.File {
		zipItemReader, _ := zipItem.Open()
		header, _ := zip.FileInfoHeader(zipItem.FileInfo())
		header.Name = zipItem.Name
		targetItem, _ := w.CreateHeader(header)
		io.Copy(targetItem, zipItemReader)
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(file.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	return buf
}
