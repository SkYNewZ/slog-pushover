package slogpushover

import (
	"fmt"
	"log/slog"

	slogcommon "github.com/samber/slog-common"
)

// Converter is a function that converts a slog.Record to a string.
type Converter func(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) string

// DefaultConverter is the default Converter used by Handler.
func DefaultConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) string {
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record) // aggregate all attributes
	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)       // developer formatters

	// handler formatter
	message := fmt.Sprintf("%s\n------------\n\n", record.Message)
	message += attrToPushoverMessage("", attrs)
	return message
}

func attrToPushoverMessage(base string, attrs []slog.Attr) string {
	message := ""

	for i := range attrs {
		attr := attrs[i]
		k := base + attr.Key
		v := attr.Value
		kind := attr.Value.Kind()

		if kind == slog.KindGroup {
			message += attrToPushoverMessage(k+".", v.Group())
		} else {
			message += fmt.Sprintf("%s: %s\n", k, slogcommon.ValueToString(v))
		}
	}

	return message
}
