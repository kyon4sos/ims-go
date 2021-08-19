package service

var mode = "static"
var CommonPath = "http://172.16.82.19:8080/assets/"
var OssPath = "aliyun.com"

func init() {
	if mode == "oss" {
		CommonPath = OssPath
	}
}

func GenFilePath(filename string) string {
	return CommonPath + filename
}
