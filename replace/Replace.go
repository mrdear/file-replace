package replace

import "file-replace/model"

type DispatchReplace interface {
	// 替换文件中指定文本处理
	DispatchReplace(rc chan *model.File)
}

type Replace interface {

	SupportType() model.FileType

	// 替换文本
	Replace(cmd model.Command, file *model.File)
}
