package main

import (
	"flag"
	"log"
	test "logserver/compress"
)

func main() {

	var (
		src string
		dest string
	)

	// 添加参数
	flag.StringVar(&src,"src","","要压缩的文件或者目录")
	flag.StringVar(&dest,"dest","","要锁生成的文件")
	if len(flag.Args()) == 0 {
		flag.Usage()
		log.Fatal()
	}
	flag.Parse()

	//
	fw:= test.NewTgzPacker()
	fw.Pack(src,dest)
	if err:=fw.Pack(src,dest);err != nil {
		log.Fatal(err)
	}
}
