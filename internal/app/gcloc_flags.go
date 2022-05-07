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
	constants.OrderByLangFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by language",
	},
	constants.OrderByCodeFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by lines of code",
	},
	constants.OrderByLineFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by lines count",
	},
	constants.OrderByBlankFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by blank lines",
	},
	constants.OrderByCommentFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by comments",
	},
	constants.OrderFlag: {
		ShortName:    "o",
		Kind:         reflect.String,
		DefaultValue: "DESC",
		Description:  "Sorting order <ASC,DESC>",
	},
}
