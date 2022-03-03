package app

import (
	"reflect"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/internal/flags"
)

var gclocFlags = map[string]flags.Flag{
	constants.ExcludeFlag: {
		ShortName:    "e",
		Kind:         reflect.Slice,
		DefaultValue: []string{},
		Description:  "Exclude directories or files from being scanned",
	},
}
