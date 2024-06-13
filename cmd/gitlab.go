package cmd

import (
	"fmt"
	"os"

	"github.com/Method-Security/codeanalyzevcs/internal/gitlab"
	"github.com/spf13/cobra"
)

func (a *CodeAnalyzeVCS) InitGitlabCommand() {
	a.GitlabCmd = &cobra.Command{
		Use:   "gitlab",
		Short: "Run a gitlab pipeline for code analysis",
		Long:  `Run a gitlab pipeline for code analysis`,
		Run: func(cmd *cobra.Command, args []string) {
			// Validate branch flag
			branch, err := cmd.Flags().GetString("branch")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate code analyze flag
			codeAnalyzeType, err := cmd.Flags().GetString("code-analyze-type")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate config type flag
			configType, err := cmd.Flags().GetString("config-type")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate config value flag
			configValue, err := cmd.Flags().GetString("config-value")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate gitlab url flag
			gitlabURL, err := cmd.Flags().GetString("vcs-url")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate project ID flag
			projectID, err := cmd.Flags().GetString("project-id")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate timeout flag
			timeout, err := cmd.Flags().GetInt("timeout")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate job name flag
			jobName, err := cmd.Flags().GetString("job-name")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			// Validate target artifact path flag
			targetArtifactPath, err := cmd.Flags().GetString("target-artifact-path")
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			pipelineToken := os.Getenv("REPO_PIPELINE_TOKEN")
			if pipelineToken == "" {
				err = fmt.Errorf("REPO_PIPELINE_TOKEN environment variable is not set")
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			readAPIToken := os.Getenv("REPO_READ_API_TOKEN")
			if readAPIToken == "" {
				err = fmt.Errorf("REPO_READ_API_TOKEN environment variable is not set")
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
				return
			}

			opts := gitlab.ExecuteCodeAnalyzeOptions{
				ReadAPIToken:       readAPIToken,
				PipelineToken:      pipelineToken,
				Branch:             branch,
				CodeAnalyzeType:    codeAnalyzeType,
				ConfigValue:        configValue,
				ConfigType:         configType,
				GitlabURL:          gitlabURL,
				ProjectID:          projectID,
				Timeout:            timeout,
				JobName:            jobName,
				TargetArtifactPath: targetArtifactPath,
			}

			report, err := gitlab.ExecuteCodeAnalyze(cmd.Context(), opts)
			if err != nil {
				errorMessage := err.Error()
				a.OutputSignal.ErrorMessage = &errorMessage
				a.OutputSignal.Status = 1
			}
			a.OutputSignal.Content = report

		},
	}

	a.GitlabCmd.Flags().String("branch", "", "Branch to run the pipeline on")
	a.GitlabCmd.Flags().String("code-analyze-type", "", "Type of code analysis to run (semgrep)")
	a.GitlabCmd.Flags().String("config-type", "", "Config type to use for the pipeline")
	a.GitlabCmd.Flags().String("config-value", "", "Config value to use for the pipeline")
	a.GitlabCmd.Flags().String("vcs-url", "", "VCS URL")
	a.GitlabCmd.Flags().String("project-id", "", "Project ID")
	a.GitlabCmd.Flags().Int("timeout", 300, "The amount of time in seconds to wait for the pipeline to complete before timing out (default: 300)")
	a.GitlabCmd.Flags().String("job-name", "codeanalyze", "The job name in GitLab CI to reference for downloading artifacts default: codeanalyze)")
	a.GitlabCmd.Flags().String("target-artifact-path", "codeanalyze-output.json", "The codeanalyze output artifact path (default: codeanalyze-output.json)")

	a.RootCmd.AddCommand(a.GitlabCmd)
}
