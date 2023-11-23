openapi-to-ghp
========

This GitHub action reads the OpenAPI schema file in the repository and uses the
[OpenAPI Generator] to build it into an NPM package, then publishes it to the
package registry.

By default, it's set to build using the `typescript-axios` generator and upload
to the [GitHub Package Registry]. However, you can modify the settings to use a
different generator and upload to a different registry.

For more details, please check the [`action.yml`] file.

[OpenAPI Generator]: https://github.com/OpenAPITools/openapi-generator
[GitHub Package Registry]: https://github.com/features/packages
[`action.yml`]: ./action.yml

Usage
--------

**Push on change of `VERSION` file:**

```yaml
on:
  push:
    branches:
      - main
  workflow_dispatch:

jobs:
  openapi-to-ghp:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write
    steps:
      - uses: actions/checkout@v4
      - uses: portone-io/actions/openapi-to-ghp@build-ts-openapi-client
        with:
          input: rest/merchant_server.yaml
          github-token: ${{ secrets.GITHUB_TOKEN }}
```

**Push on every push:**

```yaml
on: [push, workflow_dispatch]

jobs:
  openapi-to-ghp:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write
    steps:
      - uses: actions/checkout@v4
      - uses: portone-io/actions/openapi-to-ghp@build-ts-openapi-client
        with:
          input: rest/merchant_server.yaml
          github-token: ${{ secrets.GITHUB_TOKEN }}
```

**Detailed:**

```yaml
- uses: actions/checkout@v4
- uses: portone-io/actions/openapi-to-ghp@main
  with:
    # Path of OpenAPI schema, YAML or JSON.
    # (ex: "src/login_server.yaml", "auth_service.swagger.json")
    input: foo/bar/openapi.yaml

    # Token to authorize GitHub Packages. Use "secrets.GITHUB_TOKEN".
    github-token: ${{ secrets.GITHUB_TOKEN }}

    # Package versioning policy, must be one of "file", "tag", "canary", and
    # "manual".
    # (default: "file")
    #
    # Version policy:
    # - On "file", content of current directory's "VERSION" file will be used as
    #   a package version.
    # - On "tag", current workflow's git tag minus the leading v is used as a
    #   package version. If current workflow is not git tag push or tag name
    #   does not start with "v", this action will do nothing but fail.
    # - On "canary", "0.0.0-canary.<COMMIT HASH>" will be used as a package
    #   version, and "canary" distribution tag will be appended to built NPM
    #   package.
    # - On "manual", "version" parameter of this action will be used as a
    #   package version.
    # - Otherwise, this action will do nothing but fail.
    #
    # Behavior on duplicated version:
    # - If "file" versioning policy is used, this action will not fail but
    #   silently skip build process even if specified version already does exist
    #   in the package registry.
    # - Otherwise, this action will explicitly fail if specified version already
    #   does exist in the package registry.
    versioning: manual

    # If "versioning" parameter is "manual", this field will be used as version of built NPM package.
    # Otherwise, ignored.
    version: 1.2.3

    # Name of built NPM package.
    # (ex: "mypackage", default: Name of repository)
    name: mypackage

    # Scope of built NPM package.
    # (ex: "myorg", default: GitHub username of repository owner)
    scope: portone-io

    # URL of NPM registry to upload.
    # (default: "https://npm.pkg.github.com")
    registry-url: https://github.com/portone-io/actions

    # The Docker tag of the openapitools/openapi-generator-cli image to use.
    # See https://hub.docker.com/r/openapitools/openapi-generator-cli/tags for available tags.
    # (ex: "latest", "v7.1.0", default: "v6.2.1")
    openapi-generator-version: v6.2.1

    # The generator you wish to use. See https://openapi-generator.tech/docs/generators/ for available generators.
    # (ex: "javascript", "typescript-fetch", default: "typescript-axios")
    generator-type: typescript-axios

    # Additional arguments to pass through to the generate command.
    # (default: "--additional-properties paramNaming=original,supportsES6=true")
    generator-cli-args: --additional-properties paramNaming=original,supportsES6=true

    # The additional file list to include in the NPM package.
    # Write a YAML object as a string,
    # with file paths in the working directory as keys and the paths to be included in the package as values.
    # (ex: "{ src-path1: dst-path1, src-path2: dst-path2 }", default: "{}")
    additional-files: |
      src-path1: dst-path1
      src-path2: dst-path2

    # JSON object to override the default package.json generated during the
    # package build process.
    # (ex: '{"scripts":{"start":"node index.js"}}', default: "{}")
    #
    # This feature is useful for customizing various aspects of the npm package,
    # such as adding additional metadata, scripts, dependencies, or any other
    # fields supported by package.json.
    #
    # The JSON object should be formatted as a string. For example, to add
    # custom scripts and dependencies, format it as: '{"scripts":{"start":"node
    # index.js"}}'. This string will be merged with the default package.json,
    # allowing for extensive customization of the package.
    #
    # If left empty (default behavior), the action will proceed with the default
    # package.json configuration without any custom overrides.
    override-packagejson: |
      {
        "exports": {
          ".": "./src/index.ts",
          "./schema.graphql": "./schema.graphql"
        }
      }
```
