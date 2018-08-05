package model

import "os"

type File struct {

	File *os.File

	Type FileType

	FullPath string

	Name string
}
