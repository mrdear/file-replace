package model

// 文本类型

type FileType int8

const (
	TEXT_FILE FileType = 1
)

func SelectType(name string) FileType {
	return TEXT_FILE
}


