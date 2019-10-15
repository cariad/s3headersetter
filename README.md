# s3headersetter

[![Build Status](https://travis-ci.org/cariad/s3headersetter.svg?branch=master)](https://travis-ci.org/cariad/s3headersetter) [![Go Report Card](https://goreportcard.com/badge/github.com/cariad/s3headersetter)](https://goreportcard.com/report/github.com/cariad/s3headersetter) [![MIT](https://img.shields.io/npm/l/express.svg)](https://github.com/cariad/s3headersetter/blob/master/LICENSE)

## Introduction

**s3headersetter** is a CLI application for setting the `Cache-Control` and `Content-Type` HTTP headers on S3 objects.

## Rules

The headers and values to apply must be specified in a YAML file.

For example, the following YAML will:

- Set `Cache-Control` to `max-age=3600, public` on `.html` objects.
- Set `Cache-Control` to `max-age=604800, public` on `.css` objects.
- Set `Cache-Control` to `max-age=31536000, public` on all other objects.
- Set `Content-Type` to `font/woff2` on `.woff2` objects.

```yaml
rules:
  - header:        Cache-Control
    when:
      - extension: .html
        then:      max-age=3600, public
      - extension: .css
        then:      max-age=604800, public
    else:          max-age=31536000, public

  - header:        Content-Type
    when:
      - extension: .woff2
        then:      font/woff2
```

## Downloading & installing

Release builds for Linux, macOS and Windows can be downloaded from the [releases](https://github.com/cariad/s3headersetter/releases) page.

There is no installation process; just put the executable wherever you want.

## Usage

```shell
s3headersetter -config <CONFIG FILENAME> -bucket <BUCKET> (-key-prefix <PREFIX>)
```

For example:

```shell
s3headersetter -config config-sample.yml -bucket my-bucket
```

```shell
s3headersetter -config config-sample.yml -bucket my-bucket -key-prefix www
```

## Acknowledgements

This project uses the following packages:

- [github.com/cariad/gos3headersetter](https://github.com/cariad/gos3headersetter) to interact with the headers of S3 objects.

## Licence, credit & sponsorship

This project is published under the MIT Licence.

You don't owe me anything in return, but as an indie freelance coder there are two things I'd appreciate:

- **Credit.** If your app or documentation has a credits page, please consider mentioning the projects you use.
- **Cash.** If you want *and are able to* support future development, please consider [becoming a patron](https://www.patreon.com/cariad) or [buying me a coffee](https://ko-fi.com/cariad). Thank you!
