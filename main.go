package main

import (
	"flag"
	"log"
	test "logserver/compress"
)

func main() {

	var (
		src  string
		dest string
	)
	// 添加参数
	flag.StringVar(&src, "src", "", "要压缩的文件或者目录")
	flag.StringVar(&dest, "dest", "", "要锁生成的文件")
	flag.Parse()

	fw := test.NewTgzPacker()
	if err := fw.Pack(src, dest); err != nil {
		log.Fatal(err)
	}
}
