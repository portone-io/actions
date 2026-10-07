buf-go
========
This GitHub action reads the gRPC schema file in the repository and uses the 
[buf] to build it into a Golang codes.

For more details, please check the [`action.yml`] file.

[buf]: https://buf.build/
[`action.yml`]: ./action.yml

Usages
--------
```yaml
on:
  push:
    tags:
      - 'v*'

jobs:
  buf:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: portone-io/actions/buf-go@buf-go
        with:
          deploy-key: ${{ secrets.DEPLOY_KEY_PORTONE_IO_GO }}
```

References
--------
```yaml
- uses: actions/checkout@v4
- uses: portone-io/actions/buf-go@main
  with:
    # Directory where the `*.proto` files reside.
    # (default: "./protobuf")
    input-directory: foo/bar/protobuf

    # Go module name under interface/. By default, derived from the repository.
    slug: merchant

    # Prefix of the source release tag. Defaults to v.
    # With this prefix, interface/v1.2.3 publishes interface/merchant/v1.2.3.
    tag-prefix: interface/v

    # The version of the buf cli to use.
    # (ex: "latest", "1.28.0", default: "1.28.1")
    buf-version: 1.28.1

    # Deploy key (OpenSSH private key) which have write permission to
    # portone-io/go repository
    # (ex: '-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnN ...')
    deploy-key: ${{ secrets.DEPLOY_KEY }}
```
