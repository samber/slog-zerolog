
# slog: Zerolog handler

[![tag](https://img.shields.io/github/tag/samber/slog-zerolog.svg)](https://github.com/samber/slog-zerolog/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/slog-zerolog?status.svg)](https://pkg.go.dev/github.com/samber/slog-zerolog)
![Build Status](https://github.com/samber/slog-zerolog/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/slog-zerolog)](https://goreportcard.com/report/github.com/samber/slog-zerolog)
[![Coverage](https://img.shields.io/codecov/c/github/samber/slog-zerolog)](https://codecov.io/gh/samber/slog-zerolog)
[![Contributors](https://img.shields.io/github/contributors/samber/slog-zerolog)](https://github.com/samber/slog-zerolog/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/slog-zerolog)](./LICENSE)

A [Zerolog](https://github.com/rs/zerolog) Handler for [slog](https://pkg.go.dev/log/slog) Go library.

**See also:**

- [slog-multi](https://github.com/samber/slog-multi): `slog.Handler` chaining, fanout, routing, failover, load balancing...
- [slog-formatter](https://github.com/samber/slog-formatter): `slog` attribute formatting
- [slog-sampling](https://github.com/samber/slog-sampling): `slog` sampling policy
- [slog-gin](https://github.com/samber/slog-gin): Gin middleware for `slog` logger
- [slog-echo](https://github.com/samber/slog-echo): Echo middleware for `slog` logger
- [slog-fiber](https://github.com/samber/slog-fiber): Fiber middleware for `slog` logger
- [slog-chi](https://github.com/samber/slog-chi): Chi middleware for `slog` logger
- [slog-datadog](https://github.com/samber/slog-datadog): A `slog` handler for `Datadog`
- [slog-rollbar](https://github.com/samber/slog-rollbar): A `slog` handler for `Rollbar`
- [slog-sentry](https://github.com/samber/slog-sentry): A `slog` handler for `Sentry`
- [slog-syslog](https://github.com/samber/slog-syslog): A `slog` handler for `Syslog`
- [slog-logstash](https://github.com/samber/slog-logstash): A `slog` handler for `Logstash`
- [slog-fluentd](https://github.com/samber/slog-fluentd): A `slog` handler for `Fluentd`
- [slog-graylog](https://github.com/samber/slog-graylog): A `slog` handler for `Graylog`
- [slog-loki](https://github.com/samber/slog-loki): A `slog` handler for `Loki`
- [slog-slack](https://github.com/samber/slog-slack): A `slog` handler for `Slack`
- [slog-telegram](https://github.com/samber/slog-telegram): A `slog` handler for `Telegram`
- [slog-mattermost](https://github.com/samber/slog-mattermost): A `slog` handler for `Mattermost`
- [slog-microsoft-teams](https://github.com/samber/slog-microsoft-teams): A `slog` handler for `Microsoft Teams`
- [slog-webhook](https://github.com/samber/slog-webhook): A `slog` handler for `Webhook`
- [slog-kafka](https://github.com/samber/slog-kafka): A `slog` handler for `Kafka`
- [slog-parquet](https://github.com/samber/slog-parquet): A `slog` handler for `Parquet` + `Object Storage`
- [slog-zap](https://github.com/samber/slog-zap): A `slog` handler for `Zap`
- [slog-zerolog](https://github.com/samber/slog-zerolog): A `slog` handler for `Zerolog`
- [slog-logrus](https://github.com/samber/slog-logrus): A `slog` handler for `Logrus`

## 🚀 Install

```sh
go get github.com/samber/slog-zerolog/v2
```

**Compatibility**: go >= 1.21

No breaking changes will be made to exported APIs before v3.0.0.

## 💡 Usage

GoDoc: [https://pkg.go.dev/github.com/samber/slog-zerolog/v2](https://pkg.go.dev/github.com/samber/slog-zerolog/v2)

### Handler options

```go
type Option struct {
    // log level (default: debug)
    Level slog.Leveler

    // optional: zerolog logger (default: zerolog.Logger)
    Logger *zerolog.Logger

    // optional: customize json payload builder
    Converter Converter

    // optional: see slog.HandlerOptions
    AddSource   bool
    ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}
```

Other global parameters:

```go
slogzerolog.SourceKey = "source"
slogzerolog.ErrorKeys = []string{"error", "err"}
slogzerolog.LogLevels = map[slog.Level]zerolog.Level{...}
```

### Example

```go
import (
    "github.com/rs/zerolog"
    slogzerolog "github.com/samber/slog-zerolog/v2"
    "os"
    "log/slog"
)

func main() {
    zerologLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

    logger := slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &zerologLogger}.NewZerologHandler())
    logger = logger.
        With("environment", "dev").
        With("release", "v1.0.0")

    // log error
    logger.
        With("category", "sql").
        With("query.statement", "SELECT COUNT(*) FROM users;").
        With("query.duration", 1*time.Second).
        With("error", fmt.Errorf("could not count users")).
        Error("caramba!")

    // log user signup
    logger.
        With(
            slog.Group("user",
                slog.String("id", "user-123"),
                slog.Time("created_at", time.Now()),
            ),
        ).
        Info("user registration")
}
```

## 🤝 Contributing

- Ping me on twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/slog-zerolog)
- Fix [open issues](https://github.com/samber/slog-zerolog/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/slog-zerolog)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2023 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
