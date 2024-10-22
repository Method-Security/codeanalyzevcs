<div align="center">
<h1>codeanalyzevcs</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>

codeanalyzevcs is designed as an orchestration tool that can help security teams kick off Continuous Integration (CI) jobs within their existing Version Control Systems to perform various code analysis capabilities. It is intended to work hand in hand with [codeanalyze](https://github.com/Method-Security/codeanalyze) which is responsible for actually conducting the scans within the CI pipelines.

By leveraging codeanalyzevcs, security operators can integrate a variety of code analysis workflows into their automation needs, ensuring that they have visibility into the security of their networks all the way from source code to running binaries. Designed with data-modeling and data-integration needs in mind, codeanalyzevcs can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

We are constantly growing the types of version control systems and CI platforms that codeanalyzevcs can integrate with. For the most up to date listing, please see the documentation [here](./docs/index.md)

To learn more about codeanalyzevcs, please see the [Documentation site](https://method-security.github.io/codeanalyzevcs/) for the most detailed information.

## Quick Start

### Get codeanalyzevcs

For the full list of available installation options, please see the [Installation](./getting-started/installation.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/codeanalyzevcs`
- `docker run ghcr.io/method-security/codeanalyzevcs`
- Download the latest binary from the [Github Releases](https://github.com/Method-Security/codeanalyzevcs/releases/latest) page
- [Installation documentation](./getting-started/installation.md)

### Authentication

Depending on the command you are running, you may need to export environment variables that will be used to authenticate to the version control system. Please see the specific command documentation for the environment variables that are used.

#### Examples

```bash
codeanalyzevcs gitlab \
  --vcs-url https://gitlab.com/api/v4 \
  --project-id 12345 \
  --branch develop \
  --code-analyze-type semgrep \
  --config-type template \
  --config-value secrets
```

## Contributing

Interested in contributing to codeanalyzevcs? Please see our organization wide [Contribution](https://method-security.github.io/community/contribute/discussions.html) page.

## Want More?

If you're looking for an easy way to tie codeanalyzevcs into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, visit us [here](https://method.security)

## Community

codeanalyzevcs is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](https://github.com/Method-Security) or our organization wide documentation [here](https://method-security.github.io).

Have an idea for a Tool to contribute? Open a Discussion [here](https://github.com/Method-Security/Method-Security.github.io/discussions).

[verify]: https://github.com/Method-Security/codeanalyzevcs/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/codeanalyzevcs/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/codeanalyzevcs
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/codeanalyzevcs
[release]: https://github.com/Method-Security/codeanalyzevcs/releases
[releases]: https://github.com/Method-Security/codeanalyzevcs/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/codeanalyzevcs.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/codeanalyzevcs/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/codeanalyzevcs?logo=docker&label=docker%20pulls%20%2F%20codeanalyzevcs
[docker-pull]: https://hub.docker.com/r/methodsecurity/codeanalyzevcs
[license]: https://github.com/Method-Security/codeanalyzevcs/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
