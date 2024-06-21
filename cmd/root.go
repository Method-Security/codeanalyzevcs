package cmd

import (
	"errors"
	"strings"
	"time"

	"github.com/Method-Security/codeanalyzevcs/internal/config"
	"github.com/Method-Security/pkg/signal"
	"github.com/Method-Security/pkg/writer"
	"github.com/palantir/pkg/datetime"
	"github.com/palantir/witchcraft-go-logging/wlog/svclog/svc1log"
	"github.com/spf13/cobra"
)

type CodeAnalyzeVCS struct {
	Version      string
	RootFlags    config.RootFlags
	OutputConfig writer.OutputConfig
	OutputSignal signal.Signal
	RootCmd      *cobra.Command
	VersionCmd   *cobra.Command
	GitlabCmd    *cobra.Command
}

func NewCodeAnalyzeVCS(version string) *CodeAnalyzeVCS {
	codeAnalyzeVCS := CodeAnalyzeVCS{
		Version: version,
		RootFlags: config.RootFlags{
			Quiet:   false,
			Verbose: false,
		},
	}
	return &codeAnalyzeVCS
}

func (a *CodeAnalyzeVCS) InitRootCommand() {
	var outputFormat string
	var outputFile string
	a.RootCmd = &cobra.Command{
		Use:   "codeanalyzevcs",
		Short: "Code Analyze VCS CLI",
		Long:  `Code Analyze VCS CLI is used for remote triggering of code analysis in version control systems`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			format, err := validateOutputFormat(outputFormat)
			if err != nil {
				return err
			}
			var outputFilePointer *string
			if outputFile != "" {
				outputFilePointer = &outputFile
			} else {
				outputFilePointer = nil
			}
			a.OutputConfig = writer.NewOutputConfig(outputFilePointer, format)
			cmd.SetContext(svc1log.WithLogger(cmd.Context(), config.InitializeLogging(cmd, &a.RootFlags)))
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			completedAt := datetime.DateTime(time.Now())
			a.OutputSignal.CompletedAt = &completedAt
			return writer.Write(
				a.OutputSignal.Content,
				a.OutputConfig,
				a.OutputSignal.StartedAt,
				a.OutputSignal.CompletedAt,
				a.OutputSignal.Status,
				a.OutputSignal.ErrorMessage,
			)
		},
	}

	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Quiet, "quiet", "q", false, "Suppress output")
	a.RootCmd.PersistentFlags().BoolVarP(&a.RootFlags.Verbose, "verbose", "v", false, "Verbose output")
	a.RootCmd.PersistentFlags().StringVarP(&outputFile, "output-file", "f", "", "Path to output file. If blank, will output to STDOUT")
	a.RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "signal", "Output format (signal, json, yaml). Default value is signal")

	a.VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version number of gitlabctl",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(a.Version)
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}
	a.RootCmd.AddCommand(a.VersionCmd)
}

func validateOutputFormat(output string) (writer.Format, error) {
	var format writer.FormatValue
	switch strings.ToLower(output) {
	case "json":
		format = writer.JSON
	case "yaml":
		format = writer.YAML
	case "signal":
		format = writer.SIGNAL
	default:
		return writer.Format{}, errors.New("invalid output format. Valid formats are: json, yaml, signal")
	}
	return writer.NewFormat(format), nil
}
