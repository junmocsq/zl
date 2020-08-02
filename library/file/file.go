package file

// 文件存储读取接口
type Filer interface {
	WriteFile(bucket string, objectId string)
	ReadAbsUrl(relative string)
}
