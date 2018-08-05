package main

import (
	"file-replace/walk"
	"file-replace/model"
	"time"
	"file-replace/replace"
	"regexp"
)

func main() {

	dirwalk := &walk.DirWalk{
		DirPath: "/Users/niuli/Blog/nl101531.github.io/source/_posts",
		FilterType:".md",
	}
	textReplace := &replace.TextFileReplace{}

	pattern, _ := regexp.Compile("oobu4m7ko.bkt.clouddn.com")

	dispatch := replace.DefaultDispatchReplace{
		Cmd: model.Command{
			Pattern: pattern,
			Target:"imgblog.mrdear.cn",
			//Backup:"/Users/niuli/.Trash/",
		},
		DefaultReplace: textReplace,
	}
	dispatch.PutReplaceStrategy(textReplace)

	rc := make(chan *model.File, 20)
	// 扫描文件
	go dirwalk.WalkToChan(rc)
	// 处理文件
	go dispatch.DispatchReplace(rc)

	time.Sleep(10 * time.Second)
}
