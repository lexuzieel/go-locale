package locale

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// Singleton instance of i18n.Bundle
var bundle *i18n.Bundle

// A list of localizers for different languages
var localizers = make(map[string]*i18n.Localizer, 0)

// Fallback language to be used when no suitable
// localization string is not found
var fallbackLanguage string

// Load localization files from directory specified by the <directoryPath>
// and prepare them for usage.
//
// NOTE: This function allocates resources only once.
// On consequent calls it uses loaded resources.
//
//  //go:embed locale/*.yml
//  var localeFS embed.FS
//  err := locale.Initialize("en", localeFS, "locale")
func Initialize(
	defaultLanguage string,
	filesystem fs.ReadDirFS,
	directoryPath string,
) error {
	// If a bundle has already been initialized just return it
	if bundle != nil {
		return nil
	}

	tag, err := language.Parse(defaultLanguage)
	if err != nil {
		return err
	}

	fallbackLanguage = defaultLanguage
	bundle = i18n.NewBundle(tag)

	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if directoryPath == "" {
		directoryPath = "."
	}

	// Read all the files from the directory
	entries, err := filesystem.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		// Load each file into the bundle
		_, err = bundle.LoadMessageFileFS(
			filesystem,
			filepath.Join(directoryPath, entry.Name()),
		)
		if err != nil {
			return err
		}

		// Get file name and use it as the key for the localizer
		lang := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		localizers[lang] = i18n.NewLocalizer(bundle, lang)
	}

	return nil
}

func GetLanguages() []string {
	checkIfInitializaed()

	keys := []string{}

	for k := range localizers {
		keys = append(keys, k)
	}

	return keys
}

func GetDefaultLanguage() string {
	checkIfInitializaed()

	return fallbackLanguage
}

func GetMessage(id string, lang string, args []any) string {
	if id == "" {
		return "<no message>"
	}

	return getLocalizer(lang).MustLocalize(&i18n.LocalizeConfig{
		TemplateData: parseArgs(args),
		DefaultMessage: &i18n.Message{
			ID:    id,
			Other: fmt.Sprintf("<%s>", id),
		},
	})
}
