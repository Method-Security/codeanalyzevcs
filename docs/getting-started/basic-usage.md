# Basic Usage

Depending on the command you are running, you may need to export environment variables that will be used to authenticate to the version control system. Please see the specific command documentation for the environment variables that are used.

## Binaries

Running as a binary means you don't need to do anything additional for codeanalyzevcs to leverage the environment variables you have already exported.

## Docker

Running codeanalyzevcs within a Docker container requires that you pass the required credential environment variables into the container. Here is an example of doing that with the `codeanalyzevcs gitlab` command:

```bash
docker run \
  -e REPO_PIPELINE_TOKEN=$REPO_PIPELINE_TOKEN \
  -e REPO_READ_API_TOKEN=$REPO_READ_API_TOKEN \
  codeanalyzevcs gitlab \
  --vcs-url https://gitlab.com/api/v4 \
  --project-id 12345 \
  --branch develop \
  --code-analyze-type semgrep \
  --config-type template \
  --config-value secrets
```
