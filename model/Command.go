package model

import "regexp"

type Command struct {

	Pattern *regexp.Regexp

	Target string

	Backup string

}