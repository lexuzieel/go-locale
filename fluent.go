package locale

type FluentChain struct {
	messageId string   // Message id as defined in the file
	language  string   // Translation file name
	args      []string // Template arguments as key-value pairs
}

func (l *FluentChain) String() string {
	return GetMessage(l.messageId, l.language, l.args)
}

// Set the message id as specified in the file
func (l *FluentChain) Message(id string) *FluentChain {
	l.messageId = id

	if l.language == "" {
		l.language = fallbackLanguage
	}

	return l
}

// Constructs an instance of FluentChain and calls Message() on it
func Message(id string) *FluentChain {
	return (&FluentChain{}).Message(id)
}

// Set the language for the given eloquent locale call chain
func (l *FluentChain) In(lang string) *FluentChain {
	l.language = lang
	return l
}

// Constructs an instance of FluentChain and calls In() on it
func In(language string) *FluentChain {
	return (&FluentChain{}).In(language)
}

// Change the template arguments for the given eloquent locale call chain
func (l *FluentChain) With(args ...string) *FluentChain {
	l.args = args
	return l
}

// Constructs an instance of FluentChain and calls With() on it
func With(args ...string) *FluentChain {
	return (&FluentChain{}).With(args...)
}
