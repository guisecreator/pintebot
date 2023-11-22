package i18_n

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type I18N struct {
	translations map[string]map[string]any
}

func NewI18N(localPath string) (*I18N, error) {
	entries, err := os.ReadDir(localPath)
	if err != nil {
		return nil, err
	}

	translations := make(map[string]map[string]any)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.Replace(entry.Name(), ".json", "", 1)

		fileContent := make(map[string]any)
		file, readErr := os.ReadFile(localPath + entry.Name())
		if err != nil {
			return nil, readErr
		}

		readErr = yaml.Unmarshal(file, &fileContent)
		translations[name] = fileContent
	}

	return &I18N{
		translations: make(map[string]map[string]any),
	}, nil
}

func (i18n *I18N) Translate(key string, args ...any) string {
	if i18n.translations[key] == nil {
		return key
	}

	return i18n.translations[key][args[0].(string)].(string)
}
