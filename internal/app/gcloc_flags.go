package app

import (
	"reflect"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/internal/flags"
)

var gclocFlags = map[string]flags.Flag{
	constants.ExcludePathsFlag: {
		ShortName:    "e",
		Kind:         reflect.Slice,
		DefaultValue: []string{},
		Description:  "Exclude directories or files from being scanned",
	},
	constants.ExcludeExtensionsFlag: {
		Kind:         reflect.Slice,
		DefaultValue: []string{},
		Description:  "Exclude extensions from being scanned",
	},
	constants.ByFileFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show result by file",
	},
}
