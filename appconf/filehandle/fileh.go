package filehandle

import (
	"home/appconf/golog"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(zfile string) string {
	b, err := ioutil.ReadFile(zfile)
	if err != nil {
		golog.Info.Println("Error readFile file:", zfile, ",error:", err, "\r")
	} else {
		golog.Info.Println("Success readFile file:", zfile, ",lenth:", len(b), "\r")
	}
	str := string(b)
	return str
}

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".ini")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}
	return files, dirs, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
