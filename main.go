package main

import (
	"archive/tar"
	"flag"
	"io"
	"log"
	"os"
)

func main() {

	var src string
	var dest string
	flag.StringVar(&src,"src","","要压缩的文件或者目录")
	flag.StringVar(&dest,"dest","","要锁生成的文件")
	if len(flag.Args()) == 0 {
		flag.Usage()
		log.Fatal()
	}
	flag.Parse()
	if err := Tar([]string{"D:\\go_code\\video_server\\src\\main.go", "D:\\go_code\\video_server\\src\\t.go"}, dest); err != nil {
		log.Fatal(err)
	}
}

func Tar(src []string, dst string) error {
	// 创建tar文件
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	// 通过fw创建一个tar.Writer
	tw := tar.NewWriter(fw)
	// 如果关闭失败会造成tar包不完整
	defer func() {
		if err := tw.Close(); err != nil {
			log.Println(err)
		}
	}()

	for _, fileName := range src {
		fi, err := os.Stat(fileName)
		if err != nil {
			log.Println(err)
			continue
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		// 将tar的文件信息hdr写入到tw
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}

		// 将文件数据写入
		fs, err := os.Open(fileName)
		if err != nil {
			return err
		}
		if _, err = io.Copy(tw, fs); err != nil {
			return err
		}
		fs.Close()
	}
	return nil
}
