# slog: Pushover handler

[![tag](https://img.shields.io/github/tag/skynewz/slog-pushover.svg)](https://github.com/skynewz/slog-pushover/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://pkg.go.dev/badge/github.com/skynewz/slog-pushover)](https://pkg.go.dev/github.com/skynewz/slog-pushover)
[![Go report](https://goreportcard.com/badge/github.com/skynewz/slog-pushover)](https://goreportcard.com/report/github.com/skynewz/slog-pushover)
[![License](https://img.shields.io/github/license/skynewz/slog-pushover)](./LICENSE)

A [Pushover](https://pushover.net/) Handler for [slog](https://pkg.go.dev/log/slog) Go library. Inspired
from [samber](https://github.com/samber)'
s [slog repositories](https://github.com/samber?tab=repositories&q=slog&type=source&language=go&sort=name)

## 🚀 Install

```sh
go get github.com/SkYNewZ/slog-pushover
```

**Compatibility**: go >= 1.21

## 💡 Usage

GoDoc: [https://pkg.go.dev/github.com/skynewz/slog-pushover](https://pkg.go.dev/github.com/skynewz/slog-pushover)

### Handler options

```go
type Options struct {
	Level slog.Leveler // minimum level of messages to log (default: slog.LevelDebug)

	Token     string // Pushover application token
	Recipient string // Pushover user/group key

	Message   *pushover.Message // optional: customize Pushover message. 'Message' will be replaced by the log message
	Converter Converter         // optional: customize Pushover message builder

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}
```

### Example

```go
package main

func main() {
	handler := slogpushover.NewHandler(&slogpushover.Options{
		Level:       slog.LevelDebug,
		Token:       os.Getenv("PUSHOVER_TOKEN"),
		Recipient:   os.Getenv("PUSHOVER_RECIPIENT"),
		Message:     nil, // You can customize the message details
		Converter:   nil,
		AddSource:   true,
		ReplaceAttr: nil,
	})

	logger := slog.New(handler)
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=skynewz/slog-pushover)

## 📝 License

Copyright © 2023 [Quentin Lemaire](https://github.com/SkYNewZ).

This project is [MIT](./LICENSE) licensed.