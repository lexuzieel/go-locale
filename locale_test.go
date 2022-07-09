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

func TestMessageEnglishPluralization(t *testing.T) {
	Initialize(language.English, ymlFiles, "messages")

	assert.Equal(t, "I have 0 apples", Message("plural").Count(0).String())
	assert.Equal(t, "I have an apple", Message("plural").Count(1).String())
	assert.Equal(t, "I have 2 apples", Message("plural").Count(2).String())
	assert.Equal(t, "I have 3 apples", Message("plural").Count(3).String())
	assert.Equal(t, "I have 4 apples", Message("plural").Count(4).String())
	assert.Equal(t, "I have 5 apples", Message("plural").Count(5).String())
}

func TestMessageRussianPluralization(t *testing.T) {
	Initialize(language.Russian, ymlFiles, "messages")

	assert.Equal(t, "У меня 0 яблок", Message("plural").Count(0).String())
	assert.Equal(t, "У меня 1 яблоко", Message("plural").Count(1).String())
	assert.Equal(t, "У меня 2 яблока", Message("plural").Count(2).String())
	assert.Equal(t, "У меня 3 яблока", Message("plural").Count(3).String())
	assert.Equal(t, "У меня 4 яблока", Message("plural").Count(4).String())
	assert.Equal(t, "У меня 5 яблок", Message("plural").Count(5).String())
	assert.Equal(t, "У меня 21 яблоко", Message("plural").Count(21).String())
	assert.Equal(t, "У меня 22 яблока", Message("plural").Count(22).String())
	assert.Equal(t, "У меня 23 яблока", Message("plural").Count(23).String())
	assert.Equal(t, "У меня 24 яблока", Message("plural").Count(24).String())
	assert.Equal(t, "У меня 25 яблок", Message("plural").Count(25).String())
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
