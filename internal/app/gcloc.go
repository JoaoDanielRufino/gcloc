package app

import (
	"github.com/spf13/cobra"

	"github.com/JoaoDanielRufino/gcloc/internal/app/subcommands"
	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/internal/flags"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc"
)

type commandRunnerFunc func(cmd *cobra.Command, args []string) error

func NewGClocCmd() *cobra.Command {
	gclocCmd := &cobra.Command{
		Use:     "gcloc",
		Short:   "GCloc is a simple tool to count lines of code of many programming languages",
		Version: constants.Version,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Args: cobra.MaximumNArgs(1),
		RunE: run(),
	}

	flags.InitFlags(gclocCmd.Flags(), gclocFlags)

	languagesCmd := subcommands.NewLanguagesCmd()
	gclocCmd.AddCommand(languagesCmd)

	return gclocCmd
}

func run() commandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		params, err := getParams(cmd, args)
		if err != nil {
			return err
		}

		gc, err := gcloc.NewGCloc(params, constants.Languages)
		if err != nil {
			return err
		}

		return gc.Run()
	}
}

func getParams(cmd *cobra.Command, args []string) (gcloc.Params, error) {
	pathToScan := "./"

	if len(args) == 1 {
		pathToScan = args[0]
	}

	byFile, err := cmd.Flags().GetBool(constants.ByFileFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	excludePaths, err := cmd.Flags().GetStringSlice(constants.ExcludePathsFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	excludeExtensions, err := cmd.Flags().GetStringSlice(constants.ExcludeExtensionsFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	includeExtensions, err := cmd.Flags().GetStringSlice(constants.IncludeExtensionsFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	orderByLang, err := cmd.Flags().GetBool(constants.OrderByLangFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	orderByFile, err := cmd.Flags().GetBool(constants.OrderByFileFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	orderByCode, err := cmd.Flags().GetBool(constants.OrderByCodeFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	orderByLine, err := cmd.Flags().GetBool(constants.OrderByLineFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	orderByBlank, err := cmd.Flags().GetBool(constants.OrderByBlankFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	orderByComment, err := cmd.Flags().GetBool(constants.OrderByCommentFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	order, err := cmd.Flags().GetString(constants.OrderFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	params := gcloc.Params{
		Path:              pathToScan,
		ByFile:            byFile,
		ExcludePaths:      excludePaths,
		ExcludeExtensions: excludeExtensions,
		IncludeExtensions: includeExtensions,
		OrderByLang:       orderByLang,
		OrderByFile:       orderByFile,
		OrderByCode:       orderByCode,
		OrderByLine:       orderByLine,
		OrderByBlank:      orderByBlank,
		OrderByComment:    orderByComment,
		Order:             order,
	}

	return params, nil
}
