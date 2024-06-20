# codeanalyzevcs Documentation

Hello and welcome to the codeanalyzevcs documentation. While we always want to provide the most comprehensive documentation possible, we thought you may find the below sections a helpful place to get started.

- The [Getting Started](./getting-started/basic-usage.md) section provides onboarding material
- The [Development](./development/setup.md) header is the best place to get started on developing on top of and with codeanalyzevcs
- See the [Docs](./docs/index.md) section for a comprehensive rundown of codeanalyzevcs capabilities

# About codeanalyzevcs

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
