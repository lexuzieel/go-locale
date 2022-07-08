package locale

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed messages/*.yml
var ymlFiles embed.FS

func TestBundleSingleton(t *testing.T) {
	assert.Nil(t, bundle)

	err := Initialize("en", ymlFiles, "messages")
	assert.NotNil(t, bundle)
	assert.NoError(t, err)
}

func TestLanguageListAccessor(t *testing.T) {
	Initialize("en", ymlFiles, "messages")

	assert.Equal(t, "en", GetDefaultLanguage())

	assert.Contains(t, GetLanguages(), "en")
	assert.Contains(t, GetLanguages(), "es")
	assert.NotContains(t, GetLanguages(), "de")
}

func TestMessages(t *testing.T) {
	Initialize("en", ymlFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").In("en").With("name", "Alex").String(),
	)
	assert.Equal(t,
		"Hola, Alex!",
		Message("greeting").In("es").With("name", "Alex").String(),
	)

	assert.Equal(t, "<unknown>", Message("unknown").String())
	assert.Equal(t, "<unknown>", Message("unknown").In("de").String())
}

func TestNestedMessage(t *testing.T) {
	Initialize("en", ymlFiles, "messages")

	assert.Equal(t,
		"This is a nested message",
		Message("nested.message.test").String(),
	)
}

//go:embed messages/*.yaml
var yamlFiles embed.FS

func TestYamlFile(t *testing.T) {
	Initialize("en", yamlFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
}

//go:embed messages/*.toml
var tomlFiles embed.FS

func TestTomlFile(t *testing.T) {
	Initialize("en", tomlFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
}

//go:embed messages/*.json
var jsonFiles embed.FS

func TestJsonFile(t *testing.T) {
	Initialize("en", jsonFiles, "messages")

	assert.Equal(t,
		"Hello, Alex!",
		Message("greeting").With("name", "Alex").String(),
	)
}
