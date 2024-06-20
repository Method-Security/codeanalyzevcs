# Capabilities

codeanalyzevcs is designed as an orchestration tool that can help security teams kick off Continuous Integration (CI) jobs within their existing Version Control Systems to perform various code analysis capabilities. It is intended to work hand in hand with [codeanalyze](https://github.com/Method-Security/codeanalyze) which is responsible for actually conducting the scans within the CI pipelines.

codeanalyzevcs does assume that CI pipelines have been preconfigured to use codeanalyze, but allows for the triggering of these pipelines as part of various security automation tasks.

Each of the below pages will provide you with a more in depth look at the codeanalyzevcs capabilities.

- [Gitlab](./gitlab.md)

## Top Level Flags

codeanalyzevcs has several top level flags that can be used on any subcommand. These include:

```bash
Flags:
  -h, --help                 help for codeanalyzevcs
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output
```

## Version Command

Run `codeanalyzevcs version` to get the exact version information for your binary

## Output Formats

For more information on the various output formats that are supported by codeanalyzevcs, see the [Output Formats](https://method-security.github.io/docs/output.html) page in our organization wide documentation.
