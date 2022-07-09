package locale

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

//go:embed messages/*.yml
var ymlFiles embed.FS

func TestBundleSingleton(t *testing.T) {
	assert.Nil(t, bundle)

	err := Initialize(language.English, ymlFiles, "messages")
	assert.NotNil(t, bundle)
	assert.NoError(t, err)
}

func TestLanguageListAccessor(t *testing.T) {
	Initialize(language.English, ymlFiles, "messages")

	assert.Equal(t, language.English, GetDefaultLanguage())

	assert.Contains(t, GetLanguages(), language.English)
	assert.Contains(t, GetLanguages(), language.Spanish)
	assert.NotContains(t, GetLanguages(), language.German)
}

func TestMessageAccessor(t *testing.T) {
	Initialize(language.English, ymlFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").In(language.English).With("name", "Alex").String(),
	)
	assert.Equal(t,
		"Hola, Alex!",
		Message("greeting").In(language.Spanish).With("name", "Alex").String(),
	)

	assert.Equal(t, "<unknown>", Message("unknown").String())
	assert.Equal(t, "<unknown>", Message("unknown").In(language.German).String())
}

func TestNestedMessage(t *testing.T) {
	Initialize(language.English, ymlFiles, "messages")

	assert.Equal(t,
		"This is a nested message",
		Message("nested.message.test").String(),
	)
}

//go:embed messages/*.yaml
var yamlFiles embed.FS

func TestYamlFile(t *testing.T) {
	Initialize(language.English, yamlFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
}

//go:embed messages/*.toml
var tomlFiles embed.FS

func TestTomlFile(t *testing.T) {
	Initialize(language.English, tomlFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
}

//go:embed messages/*.json
var jsonFiles embed.FS

func TestJsonFile(t *testing.T) {
	Initialize(language.English, jsonFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
}
