
# Contributing

We welcome contributions of all kinds including code, issues or documentation.

##  Contributing Code

To contribute code, please follow this steps:

1. Communicate with us on the issue you want to work on
2. Make your changes
3. Test your changes
4. Update the documentation if needed and examples too
5. Ensure to run `make reviewable test` without issues
6. Open a pull request

Ensure to use a good commit hygiene and follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.

##  Development environment

Get git submodules:

```console
make submodules
```

Run code-generation pipeline:

```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build binary:

```console
make build
```
