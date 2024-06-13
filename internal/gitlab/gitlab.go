package gitlab

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/xanzy/go-gitlab"
)

type Report struct {
	Branch          string      `json:"branch" yaml:"branch"`
	GitlabURL       string      `json:"gitlab_url" yaml:"gitlab_url"`
	ProjectID       string      `json:"project_id" yaml:"project_id"`
	ProjectName     string      `json:"project_name" yaml:"project_name"`
	CodeAnalyzeType string      `json:"code_analyze_type" yaml:"code_analyze_type"`
	ConfigType      string      `json:"config_type" yaml:"config_type"`
	ConfigValue     string      `json:"config_value" yaml:"config_value"`
	Results         interface{} `json:"results" yaml:"results"`
}

type ExecuteCodeAnalyzeOptions struct {
	PipelineToken      string
	ReadAPIToken       string
	Branch             string
	CodeAnalyzeType    string
	ConfigValue        string
	ConfigType         string
	GitlabURL          string
	ProjectID          string
	Timeout            int
	JobName            string
	TargetArtifactPath string
}

func ExecuteCodeAnalyze(ctx context.Context, opts ExecuteCodeAnalyzeOptions) (Report, error) {
	client, err := gitlab.NewClient(opts.ReadAPIToken, gitlab.WithBaseURL(opts.GitlabURL))
	if err != nil {
		return Report{}, fmt.Errorf("failed to create GitLab client: %v", err)
	}

	// Get project information using the project ID
	project, _, err := client.Projects.GetProject(opts.ProjectID, nil)
	if err != nil {
		return Report{}, fmt.Errorf("failed to get project: %v", err)
	}

	// Trigger pipeline
	triggerOptions := &gitlab.RunPipelineTriggerOptions{
		Ref:   gitlab.Ptr(opts.Branch),
		Token: gitlab.Ptr(opts.PipelineToken),
		Variables: map[string]string{
			"ONLY_ANALYZE":     "true",
			"CODEANALYZE_TYPE": opts.CodeAnalyzeType,
			"CONFIG_TYPE":      opts.ConfigType,
			"CONFIG_VALUE":     opts.ConfigValue,
		},
	}

	pipeline, _, err := client.PipelineTriggers.RunPipelineTrigger(opts.ProjectID, triggerOptions)
	if err != nil {
		return Report{}, fmt.Errorf("failed to trigger pipeline: %v", err)
	}

	fmt.Printf("Pipeline triggered successfully: %v\n", pipeline.WebURL)

	// Wait for the pipeline to complete
	pipelineID := pipeline.ID
	timeoutDuration := time.Duration(opts.Timeout) * time.Second
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	start := time.Now()

	for {
		select {
		case <-ctx.Done():
			return Report{}, fmt.Errorf("context canceled")
		case <-ticker.C:
			pipeline, _, err := client.Pipelines.GetPipeline(opts.ProjectID, pipelineID)
			if err != nil {
				return Report{}, fmt.Errorf("failed to get pipeline status: %v", err)
			}

			if pipeline.Status == "success" || pipeline.Status == "failed" || pipeline.Status == "canceled" {
				fmt.Printf("Pipeline completed with status: %s\n", pipeline.Status)
				goto DownloadArtifacts
			}

			if time.Since(start) > timeoutDuration {
				return Report{}, fmt.Errorf("pipeline timed out")
			}
		}
	}

DownloadArtifacts:
	// Get jobs associated with the pipeline
	jobs, _, err := client.Jobs.ListPipelineJobs(opts.ProjectID, pipelineID, nil)
	if err != nil {
		return Report{}, fmt.Errorf("failed to list pipeline jobs: %v", err)
	}

	// Download the target artifact only for the job that matches jobName, which there will only be one
	for _, job := range jobs {
		if job.Status == "success" && job.ArtifactsFile.Filename != "" && job.Name == opts.JobName {
			fmt.Printf("Downloading artifacts for job: %s\n", job.Name)

			var reader *bytes.Reader
			reader, _, err = client.Jobs.DownloadSingleArtifactsFile(opts.ProjectID, job.ID, opts.TargetArtifactPath)
			if err != nil {
				return Report{}, fmt.Errorf("failed to download artifacts: %v", err)
			}

			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(reader)
			if err != nil {
				return Report{}, fmt.Errorf("failed to read from reader: %v", err)
			}

			var artifactContent interface{}
			err = json.Unmarshal(buf.Bytes(), &artifactContent)
			if err != nil {
				return Report{}, fmt.Errorf("failed to unmarshal artifact content: %v", err)
			}

			report := Report{
				Branch:          opts.Branch,
				GitlabURL:       opts.GitlabURL,
				ProjectID:       opts.ProjectID,
				ProjectName:     project.Name,
				CodeAnalyzeType: opts.CodeAnalyzeType,
				ConfigType:      opts.ConfigType,
				ConfigValue:     opts.ConfigValue,
				Results:         artifactContent,
			}

			return report, nil

		}
	}

	return Report{}, nil
}
