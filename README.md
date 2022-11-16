# Config - A minimalist Go configuration library

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/config)](https://pkg.go.dev/github.com/go-zoox/config)
[![Build Status](https://github.com/go-zoox/config/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/config/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/config)](https://goreportcard.com/report/github.com/go-zoox/config)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/config/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/config?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/config.svg)](https://github.com/go-zoox/config/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/config.svg?label=Release)](https://github.com/go-zoox/config/tags)

### Features
* [ ] Support Types
  * [x] YAML
  * [x] JSON
  * [x] TOML
  * [x] INI
  * [ ] DotEnv
  * [x] ENV
* [ ] Support Object Storage
  * [ ] S3
  * [ ] ALI OSS
  * [ ] Tencent COS
* [ ] Support Database
  * [ ] Redis
  * [ ] MongoDB
  * [ ] Etcd
* [ ] Support Configuration Center
  * [ ] Apollo
  * [ ] Nacos 

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/config
```

## Getting Started

```go
type Config struct {
  Version string `config:"version"`
  Server  struct {
    Ports   []int64 `config:"ports"`
    Cleanup string  `config:"cleanup"`
  } `config:"server"`
  Logger struct {
    Level string `config:"level"`
    Trace bool   `config:"trace"`
  } `config:"logger"`
  Rules []struct {
    Host    string `config:"host"`
    Backend struct {
      ServiceName string `config:"service_name"`
      ServicePort int64  `config:"service_port"`
    } `config:"backend"`
  } `config:"rules"`
}

func main() {
  var cfg Config
  if err := Load(&cfg); err != nil {
    panic(err)
  }
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
