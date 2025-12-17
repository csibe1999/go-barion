package barion

import (
	"encoding/json"
	"testing"
)

func TestLocaleString(t *testing.T) {
	tests := []struct {
		name   string
		locale Locale
		want   string
	}{
		{"Czech", CZ, "cs-CZ"},
		{"German", DE, "de-DE"},
		{"English", US, "en-US"},
		{"Spanish", ES, "es-ES"},
		{"French", FR, "fr-FR"},
		{"Hungarian", HU, "hu-HU"},
		{"Slovak", SK, "sk-SK"},
		{"Slovenian", SI, "sl-SI"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.locale) != tt.want {
				t.Errorf("Locale string = %v, want %v", tt.locale, tt.want)
			}
		})
	}
}

func TestLocaleJSONMarshal(t *testing.T) {
	tests := []struct {
		name   string
		locale Locale
		want   string
	}{
		{"Hungarian", HU, `"hu-HU"`},
		{"English", US, `"en-US"`},
		{"German", DE, `"de-DE"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.locale)
			if err != nil {
				t.Fatalf("json.Marshal() error = %v", err)
			}
			if string(data) != tt.want {
				t.Errorf("json.Marshal() = %v, want %v", string(data), tt.want)
			}
		})
	}
}

func TestLocaleJSONUnmarshal(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Locale
	}{
		{"Hungarian", `"hu-HU"`, HU},
		{"English", `"en-US"`, US},
		{"German", `"de-DE"`, DE},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var locale Locale
			err := json.Unmarshal([]byte(tt.input), &locale)
			if err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			if locale != tt.want {
				t.Errorf("json.Unmarshal() = %v, want %v", locale, tt.want)
			}
		})
	}
}
