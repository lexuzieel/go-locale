package locale

import (
	"math"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func checkIfInitializaed() {
	if fallbackLanguage == language.Und {
		panic("locale is not initialized. Did you forget to run Initialize()?")
	}
}

func getLocalizer(tag language.Tag) *i18n.Localizer {
	checkIfInitializaed()

	localizer := localizers[tag]

	// If localizer was not found, fallback
	// to the one for the default language
	if localizer == nil {
		return localizers[fallbackLanguage]
	}

	return localizer
}

// Convert an array of strings to a map of strings.
// Omit last element if the number of elements is odd.
//
// For example, when the input is ["a", "b", "c"]
// this function will return the following map: {"a": "b"}
func parseArgs(args []any) map[string]any {
	var argMap = make(map[string]any, 0)

	for i := 0; i < 2*int(math.Floor(float64(len(args))/2)); i += 2 {
		argMap[args[i].(string)] = args[i+1]
	}

	return argMap
}
