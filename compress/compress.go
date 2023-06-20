package compress

//
//// 压缩文件
//func CompressedFile(file *os.File, prefix string, zw *zip.Writer) error {
//	info, err := file.Stat()
//	if err != nil || info.IsDir() {
//		return err
//	}
//	header, err := zip.FileInfoHeader(info)
//	if err != nil {
//		return err
//	}
//	header.Name = prefix + "/" + header.Name
//	writer, err := zw.CreateHeader(header)
//	if err != nil {
//		return err
//	}
//	if _, err = io.Copy(writer, file); err != nil {
//		return err
//	}
//	return nil
//}
//// 压缩目录
//func CompressedDir(file *os.File, prefix string, zw *zip.Writer) error {
//	info, _ := file.Stat()
//	prefix = prefix + "/" + info.Name()
//	dirInfo, err := file.Readdir(-1)
//	if err != nil {
//		return err
//	}
//	for _, f := range dirInfo {
//		f, err := os.Open(file.Name() + "/" + f.Name())
//		if err != nil {
//			return err
//		}
//		err = Compress(f, prefix, zw)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//func DeCompressed(src string) error {
//	s, _ := os.Open(src)
//	info, _ := s.Stat()
//	ZipReader, err := zip.NewReader(s, info.Size())
//	if err != nil {
//		return err
//	}
//	for _, f := range ZipReader.File {
//		if err := deCompressed(f); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func deCompressed(f *zip.File) error {
//	d, _ := os.Create(f.Name)
//	unzipFile, err := f.Open()
//	if err != nil {
//		return err
//	}
//	if _, err := io.Copy(d, unzipFile); err != nil {
//		return err
//	}
//	if err := unzipFile.Close(); err != nil {
//		return err
//	}
//	return d.Close()
//}
//
//// Compress
//func Compress(file *os.File, prefix string, zw *zip.Writer) error {
//	info, err := file.Stat()
//	if err != nil {
//		return err
//	}
//	// 如果是目录调用CompressedDir
//	if info.IsDir() {
//		return CompressedDir(file, prefix, zw)
//	}
//	// 如果是文件调用CompressedFile
//	return CompressedFile(file, prefix, zw)
//}




