package replace

import (
	"file-replace/model"
	)

type DefaultDispatchReplace struct {
	Cmd             model.Command
	DefaultReplace  Replace
	replaceStrategy map[model.FileType]Replace
}

func (r *DefaultDispatchReplace) DispatchReplace(rc chan *model.File) {
	for file := range rc {
		replace := r.replaceStrategy[file.Type]
		if nil == replace {
			replace = r.DefaultReplace
		}
		go replace.Replace(r.Cmd, file)
	}
}

func (r *DefaultDispatchReplace) PutReplaceStrategy(re Replace) {
	if nil == r.replaceStrategy {
		r.replaceStrategy = map[model.FileType]Replace{}
	}
	r.replaceStrategy[re.SupportType()] = re
}
