package slogpushover

import (
	"context"
	"log/slog"

	"github.com/gregdel/pushover"
	slogcommon "github.com/samber/slog-common"
)

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

// NewHandler creates a new slog.Handler that sends messages to Pushover.
func NewHandler(opts *Options) slog.Handler {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Level == nil {
		opts.Level = slog.LevelDebug
	}

	if opts.Token == "" {
		panic("missing Pushover token")
	}

	if opts.Recipient == "" {
		panic("missing Pushover recipient")
	}

	if opts.Message == nil {
		opts.Message = &pushover.Message{}
	}

	if opts.Converter == nil {
		opts.Converter = DefaultConverter
	}

	return &handler{
		options:   opts,
		client:    pushover.New(opts.Token),
		recipient: pushover.NewRecipient(opts.Recipient),
		message:   opts.Message,
		attrs:     make([]slog.Attr, 0),
		groups:    make([]string, 0),
	}
}

type handler struct {
	options   *Options
	client    *pushover.Pushover
	recipient *pushover.Recipient
	message   *pushover.Message
	attrs     []slog.Attr
	groups    []string
}

func (h *handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.options.Level.Level()
}

func (h *handler) Handle(_ context.Context, record slog.Record) error {
	message := h.options.Converter(h.options.AddSource, h.options.ReplaceAttr, h.attrs, h.groups, &record)
	msg := &pushover.Message{
		Message:     message,
		Title:       h.message.Title,
		Priority:    h.message.Priority,
		URL:         h.message.URL,
		URLTitle:    h.message.URLTitle,
		Timestamp:   h.message.Timestamp,
		Retry:       h.message.Retry,
		Expire:      h.message.Expire,
		CallbackURL: h.message.CallbackURL,
		DeviceName:  h.message.DeviceName,
		Sound:       h.message.Sound,
		HTML:        h.message.HTML,
		Monospace:   h.message.Monospace,
		TTL:         h.message.TTL,
	}

	_, err := h.client.SendMessage(msg, h.recipient)
	return err
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{
		options:   h.options,
		client:    h.client,
		recipient: h.recipient,
		message:   h.message,
		attrs:     slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups:    h.groups,
	}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{
		options: h.options,
		client:  h.client,
		attrs:   h.attrs,
		groups:  append(h.groups, name),
	}
}
