buf
========
This GitHub action reads the gRPC schema file in the repository and uses the
[buf] to lint and statically analyze the codes.

For more details, please check the [`action.yml`] file.

[buf]: https://buf.build/
[`action.yml`]: ./action.yml

Usages
--------
On change of `VERSION` file:
```yaml
on: push

jobs:
  buf:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: portone-io/actions/buf@main
```

References
--------
```yaml
- uses: actions/checkout@v4
- uses: portone-io/actions/buf@main
  with:
    # Directory where the `*.proto` files reside.
    # (default: "./protobuf")
    input-directory: foo/bar/protobuf

    # The version of the buf cli to use.
    # (ex: "latest", "1.28.0", default: "1.28.1")
    buf-version: 1.28.1
```
