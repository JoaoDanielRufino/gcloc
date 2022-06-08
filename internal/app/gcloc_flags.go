package app

import (
	"reflect"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/internal/flags"
)

var gclocFlags = map[string]flags.Flag{
	constants.ByFileFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results by file",
	},
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
	constants.IncludeExtensionsFlag: {
		Kind:         reflect.Slice,
		DefaultValue: []string{},
		Description:  "Include the extensions to be scanned",
	},
	constants.OrderByLangFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by language",
	},
	constants.OrderByFileFlag: {
		Kind:         reflect.Bool,
		DefaultValue: false,
		Description:  "Show results ordered by file count",
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
		Kind:         reflect.String,
		DefaultValue: "DESC",
		Description:  "Sorting order <ASC,DESC>",
	},
	constants.OutputNameFlag: {
		Kind:         reflect.String,
		DefaultValue: "result",
		Description:  "Name of report output",
	},
	constants.OutputPathFlag: {
		ShortName:    "o",
		Kind:         reflect.String,
		DefaultValue: "",
		Description:  "Path where the report will be exported",
	},
	constants.ReportFormatsFlag: {
		Kind:         reflect.Slice,
		DefaultValue: []string{constants.PROMPT_REPORT},
		Description:  "Report formats in which the results will be exported (prompt, json)",
	},
}
