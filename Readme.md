# Esbuild Internal

This repository exposes and renames the `internal/` packages of [esbuild](https://github.com/evanw/esbuild) so you can use them in your own project.

The tagged releases will be kept in sync with the versions that esbuild publishes.

**Note** These packages were intentionally internal. The APIs can break without notice and the packages may not be built to suit your use cases. Please review [this issue](https://github.com/evanw/esbuild/issues/201) in the esbuild project before using this package.

## Development

### Upgrading to a new version

1. Change the `VERSION` in the [Makefile](./Makefile)
2. Run `make`
