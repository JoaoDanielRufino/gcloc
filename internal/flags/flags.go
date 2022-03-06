package flags

import (
	"reflect"

	"github.com/spf13/pflag"
)

type Flag struct {
	ShortName    string
	Kind         reflect.Kind
	DefaultValue interface{}
	Description  string
}

func InitFlags(flagSet *pflag.FlagSet, flags map[string]Flag) {
	for flagName, flag := range flags {
		switch flag.Kind {
		case reflect.Slice:
			flagSet.StringSliceP(flagName, flag.ShortName, flag.DefaultValue.([]string), flag.Description)
		case reflect.Bool:
			flagSet.BoolP(flagName, flag.ShortName, flag.DefaultValue.(bool), flag.Description)
		}
	}
}
