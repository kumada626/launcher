# Metadata CLI
[![Build Status][build-image]][build-url]
[![Latest Release][version-image]][version-url]
[![Go Report Card][goreport-image]][goreport-url]

> CLI for reading/writing project metadata

## Usage

```bash
$ go get github.com/screwdriver-cd/meta-cli
$ cd $GOPATH/src/github.com/screwdriver-cd/meta-cli
$ go build -a -o meta
$ ./meta set aaa bbb
$ ./meta get aaa
bbb
$ ./meta set foo[2].bar[1] baz
[null,null,{"bar":[null,"baz"]}]
```

## Testing

```bash
$ go get github.com/screwdriver-cd/meta-cli
$ go test -cover github.com/screwdriver-cd/meta-cli/...
```

## License

Code licensed under the BSD 3-Clause license. See LICENSE file for terms.

[version-image]: https://img.shields.io/github/tag/screwdriver-cd/meta-cli.svg
[version-url]: https://github.com/screwdriver-cd/meta-cli/releases
[build-image]: https://cd.screwdriver.cd/pipelines/67/badge
[build-url]: https://cd.screwdriver.cd/pipelines/67
[goreport-image]: https://goreportcard.com/badge/github.com/Screwdriver-cd/meta-cli
[goreport-url]: https://goreportcard.com/report/github.com/Screwdriver-cd/meta-cli
