package locale

import "golang.org/x/text/language"

type FluentChain struct {
	messageId string       // Message id as defined in the file
	fallback  string       // Fallback message when no message is found
	mock      string       // Mock message to use during testing (replaces fallback)
	language  language.Tag // Translation language tag
	args      []any        // Template arguments as key-value pairs
	count     interface{}
}

func (l *FluentChain) String() string {
	id := l.messageId
	fallback := l.fallback

	if mocking && len(l.mock) > 0 {
		id = ""
		fallback = l.mock
	}

	return GetMessage(id, l.language, l.args, l.count, fallback)
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

// Change the plural count for the given fluent call chain
func (l *FluentChain) Count(count interface{}) *FluentChain {
	l.count = count
	return l
}

// Constructs an instance of FluentChain and calls Count() on it
func Count(count interface{}) *FluentChain {
	return (&FluentChain{}).Count(count)
}

// Change the plural count for the given fluent call chain
func (l *FluentChain) Mock(text string) *FluentChain {
	if mocking {
		l.mock = text
	}

	return l
}

// Constructs an instance of FluentChain and calls Mock() on it
func Mock(text string) *FluentChain {
	return (&FluentChain{}).Mock(text)
}

// Change the plural count for the given fluent call chain
func (l *FluentChain) Fallback(text string) *FluentChain {
	l.fallback = text
	return l
}

// Constructs an instance of FluentChain and calls Count() on it
func Fallback(text string) *FluentChain {
	return (&FluentChain{}).Fallback(text)
}
