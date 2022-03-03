package app

import (
	"github.com/spf13/cobra"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/internal/flags"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc"
)

type commandRunnerFunc func(cmd *cobra.Command, args []string) error

func NewGClocCmd() *cobra.Command {
	gclocCmd := &cobra.Command{
		Use:     "gcloc",
		Short:   "GCloc is a simple tool to count lines of code of many programming languages",
		Version: "1.0.0",
		Args:    cobra.MaximumNArgs(1),
		RunE:    run(),
	}

	flags.InitFlags(gclocCmd.Flags(), gclocFlags)

	return gclocCmd
}

func run() commandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		params, err := getParams(cmd, args)
		if err != nil {
			return err
		}

		gc, err := gcloc.NewGCloc(params)
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

	excludeDirs, err := cmd.Flags().GetStringSlice(constants.ExcludeFlag)
	if err != nil {
		return gcloc.Params{}, err
	}

	params := gcloc.Params{
		Path:        pathToScan,
		ExcludeDirs: excludeDirs,
	}

	return params, nil
}
