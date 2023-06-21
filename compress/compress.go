package compress

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TgzPacker struct {
}

func NewTgzPacker() *TgzPacker{
	return &TgzPacker{}
}

//判断目标文件是否存在，在压缩的时候，判断
func (tp *TgzPacker) removeTargetFile(filename string) error {
	//
	if _,err := os.Stat(filename);os.IsNotExist(err){
		return nil
	}
	return os.Remove(filename)
}

//判断目录是否存在，在解压的时候使用
func (tp *TgzPacker) dirExists(dir string) bool {
	info, err := os.Stat(dir)
	return (err == nil || os.IsExist(err)) && info.IsDir()
}
func getFiletime(filename string,currentTime time.Time) bool {
	fileinfo,_ :=os.Stat(filename)
	//oldTime := currentTime.AddDate(0, 0, -7)
	return currentTime.Before(fileinfo.ModTime().AddDate(0,0,7))
}

//压缩打包
func (tp *TgzPacker) Pack(sourceFullPath string, tarFileName string) (err error) {
	sourceInfo, err := os.Stat(sourceFullPath)
	// 校验源目录是否存在
	if err != nil {
		return err
	}
	// 删除目标tar文件
	if err = tp.removeTargetFile(tarFileName); err != nil {
		return err
	}
	// 创建写入文件句柄
	file, err := os.Create(tarFileName)
	if err != nil {
		return err
	}
	defer func() {
		// 主程序没有err，但是关闭句柄报错，则将关闭句柄的报错返回
		if err2 := file.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()
	// 创建gzip的写入句柄，对file的包装
	//gWriter := gzip.NewWriter(file)
	gWriter,_ := gzip.NewWriterLevel(file,gzip.BestSpeed)
	defer func() {
		// 主程序没有err，但是关闭句柄报错，则将关闭句柄的报错返回
		if err2 := gWriter.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()
	// 创建tar的写入句柄，对gzip的包装
	tarWriter := tar.NewWriter(gWriter)
	defer func() {
		// 主程序没有err，但是关闭句柄报错，则将关闭句柄的报错返回
		if err2 := tarWriter.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()
	// 开始压缩
	if sourceInfo.IsDir() {
		return tp.tarFolder(sourceFullPath, filepath.Base(sourceFullPath), tarWriter)
	}
	return tp.tarFile(sourceFullPath, tarWriter)
}

// 对单个文件进行打包
func (tp *TgzPacker) tarFile(sourceFullFile string,writer *tar.Writer) error {
	info, err := os.Stat(sourceFullFile)
	if err != nil {
		return err
	}
	// 创建头信息
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	// 头信息写入
	err = writer.WriteHeader(header)
	if err != nil {
		return err
	}
	// 读取源文件，将内容拷贝到tar.Writer中
	fr, err := os.Open(sourceFullFile)
	if err != nil {
		return err
	}
	defer func() {
		// 如果主程序的err为空nil，而文件句柄关闭err，则将关闭句柄的err返回
		if err2 := fr.Close(); err2 != nil && err == nil {
			err = err2
		}
	}()
	if _, err = io.Copy(writer, fr); err != nil {
		return err
	}
	return nil
}

//对目录进行打包
// sourceFullPath为待打包目录，baseName为待打包目录的根目录名称
func (tp *TgzPacker) tarFolder(sourceFullPath string, baseName string, writer *tar.Writer) error {
	// 保留最开始的原始目录，用于目录遍历过程中将文件由绝对路径改为相对路径
	currentTime := time.Now()
	baseFullPath := sourceFullPath
	return filepath.Walk(sourceFullPath,
		func(fileName string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// 创建头信息
			header, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return err
			}
			// 修改header的name，这里需要按照相对路径来
			// 说明这里是根目录，直接将目录名写入header即可
			if fileName == baseFullPath {
				header.Name = baseName
			} else {
				// 非根目录，需要对路径做处理：去掉绝对路径的前半部分，然后构造基于根目录的相对路径
				header.Name = filepath.Join(baseName, strings.TrimPrefix(fileName, baseFullPath))
			}

			if err = writer.WriteHeader(header); err != nil {
				return err
			}
			// linux文件有很多类型，这里仅处理普通文件，如业务需要处理其他类型的文件，这里添加相应的处理逻辑即可

			if !info.Mode().IsRegular() {
				return nil
			}
			// 普通文件，则创建读句柄，将内容拷贝到tarWriter中

			//判断文件modify time
			fr, err := os.Open(fileName)
			if getFiletime(fileName,currentTime) {

			}
			if err != nil {
				return err
			}
			defer fr.Close()
			if _, err := io.Copy(writer, fr); err != nil {
				return err
			}



			//fr, err := os.Open(fileName)
			//if err != nil {
			//	return err
			//}
			//defer fr.Close()
			//if _, err := io.Copy(writer, fr); err != nil {
			//	return err
			//}
			return nil
	})
}
