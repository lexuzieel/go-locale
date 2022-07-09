package locale

import "golang.org/x/text/language"

type FluentChain struct {
	messageId string       // Message id as defined in the file
	language  language.Tag // Translation language tag
	args      []any        // Template arguments as key-value pairs
}

func (l *FluentChain) String() string {
	return GetMessage(l.messageId, l.language, l.args)
}

// Set the message id as specified in the file
func (l *FluentChain) Message(id string) *FluentChain {
	l.messageId = id

	if l.language == language.Und {
		l.language = fallbackLanguage
	}

	return l
}

// Constructs an instance of FluentChain and calls Message() on it
func Message(id string) *FluentChain {
	return (&FluentChain{}).Message(id)
}

// Set the language for the given fluent call chain
func (l *FluentChain) In(tag language.Tag) *FluentChain {
	l.language = tag
	return l
}

// Constructs an instance of FluentChain and calls In() on it
func In(tag language.Tag) *FluentChain {
	return (&FluentChain{}).In(tag)
}

// Change the template arguments for the given fluent call chain
func (l *FluentChain) With(args ...any) *FluentChain {
	l.args = args
	return l
}

// Constructs an instance of FluentChain and calls With() on it
func With(args ...any) *FluentChain {
	return (&FluentChain{}).With(args...)
}
