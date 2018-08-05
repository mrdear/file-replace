package walk

import (
	"file-replace/model"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

type DirWalk struct {
	DirPath string

	FilterType string
}

// 扫描文件,传到channel中
func (dir *DirWalk) WalkToChan(rc chan *model.File) {
	fmt.Printf("start walk dir %s \n", dir.DirPath)

	// 打开目录
	dirInfo, err := ioutil.ReadDir(dir.DirPath)
	if nil != err {
		panic(fmt.Sprintf("readDir fail,%s", err.Error()))
	}

	// 遍历目录
	for _, file := range dirInfo {
		fullFileName := dir.DirPath + "/" + file.Name()
		// 递归
		if file.IsDir() {
			fmt.Printf("walk find dir %s \n", fullFileName)
			temp := &DirWalk{
				DirPath: fullFileName,
				FilterType:dir.FilterType,
			}
			go temp.WalkToChan(rc)
			continue
		}

		if !strings.HasSuffix(file.Name(), dir.FilterType) {
			continue
		}
		fmt.Printf("walk find file %s \n", fullFileName)

		// 封装文件
		realFile, err := os.Open(fullFileName)
		if nil != err {
			panic(fmt.Sprintf("open file fail,%s", err.Error()))
		}
		tempFile := &model.File{
			File:     realFile,
			Name:     file.Name(),
			FullPath: fullFileName,
			Type:     model.SelectType(file.Name()),
		}
		rc <- tempFile
	}

	fmt.Printf("end walk dir %s", dir.DirPath)
}
