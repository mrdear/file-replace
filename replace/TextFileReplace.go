package replace

import (
	"file-replace/model"
	"os"
	"fmt"
	"bufio"
	"io"
)

// 文本文件替换
type TextFileReplace struct {
}

func (t *TextFileReplace) SupportType() model.FileType {
	return model.TEXT_FILE
}

// 替换思路,创建一个新文件,按行读取旧文件,每次匹配一行替换,写入到新文件,最后使用新文件替换旧文件.
func (t *TextFileReplace) Replace(cmd model.Command, file *model.File) {
	newFileName := file.FullPath + ".bak"
	newFile, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	if nil != err {
		panic(fmt.Sprintf("create new file err, %s \n", err.Error()))
	}
	readWriter := bufio.NewReadWriter(bufio.NewReader(file.File), bufio.NewWriter(newFile))

	for true {
		line, err := readWriter.ReadString('\n')
		// 替换
		line = cmd.Pattern.ReplaceAllString(line, cmd.Target)
		// 写入新文件
		readWriter.WriteString(line)

		if err == io.EOF {
			break
		}
	}
	readWriter.Writer.Flush()
	newFile.Close()
	file.File.Close()

	// 替换旧文件
	if len(cmd.Backup) > 0 {
		err = os.Rename(file.FullPath, cmd.Backup+file.Name)
		if nil != err {
			fmt.Printf("move old file fail, err is %s \n", err.Error())
		}
	}
	err = os.Rename(newFileName, file.FullPath)
	if nil != err {
		fmt.Printf("move new file fail, err is %s \n", err.Error())
	}

	fmt.Printf("finish replace file, %s \n", file.FullPath)
}
