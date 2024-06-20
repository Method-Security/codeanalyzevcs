# Gitlab

The `codeanalyzevcs gitlab` command allows you to orchestrate code analysis CI jobs within the [Gitlab](https://gitlab.com/) version control system, by leveraging the [Gitlab APIs](https://docs.gitlab.com/ee/api/rest/).

## Authentication

The Gitlab API requires [authentication](https://docs.gitlab.com/ee/api/rest/#authentication) for all interactions. For `codeanalyzevcs gitlab`, we require two tokens, one with permission to kick off the CI pipeline (the token will need one of Developer, Maintainer, or Owner permissions on the repository per the [Gitlab documentation](https://docs.gitlab.com/ee/user/permissions.html#:~:text=Run%20CI/CD%20pipeline%20for%20a%20protected%20branch)) and another token with the ability to clone the repository.

These can be the same token if you so choose, but they have been split up to provide flexibility in the access and permissions that need to be provided to the `codeanalyzevcs gitlab` command.

The command will read these tokens from environment variables:

- `REPO_PIPELINE_TOKEN`
- `REPO_READ_API_TOKEN`

## Usage

```bash
codeanalyzevcs gitlab \
  --vcs-url https://gitlab.com/api/v4 \
  --project-id 12345 \
  --branch develop \
  --code-analyze-type semgrep \
  --config-type template \
  --config-value secrets
```

## Help Text

```bash
$ codeanalyze gitlab -h
Flags:
      --branch string                 Branch to run the pipeline on
      --code-analyze-type string      Type of code analysis to run (semgrep)
      --config-type string            Config type to use for the pipeline
      --config-value string           Config value to use for the pipeline
  -h, --help                          help for gitlab
      --job-name string               The job name in GitLab CI to reference for downloading artifacts default: codeanalyze) (default "codeanalyze")
      --project-id string             Project ID
      --target-artifact-path string   The codeanalyze output artifact path (default: codeanalyze-output.json) (default "codeanalyze-output.json")
      --timeout int                   The amount of time in seconds to wait for the pipeline to complete before timing out (default: 300) (default 300)
      --vcs-url string                VCS URL
```
