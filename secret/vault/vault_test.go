package vault

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func getSecretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
func TestSet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "123", filepath: getSecretpath(), key: "mohit", plainText: "brother"},
	}
	for _, test := range testSuit {
		v := GetVault(test.encodingKey, test.filepath)
		err := v.Set(test.key, test.plainText)
		if err != nil {
			t.Error("error in Set")
		}
	}
}

func TestGet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "123", filepath: getSecretpath(), key: "mohit", plainText: "brother"},
		{encodingKey: "123", filepath: getSecretpath() + "ds", key: "mohit", plainText: ""},
		{encodingKey: "123", filepath: getSecretpath(), key: "google", plainText: ""},
	}
	for _, test := range testSuit {
		v := GetVault(test.encodingKey, test.filepath)
		plainText, _ := v.Get(test.key)
		if plainText != test.plainText {
			t.Error("error in Get")
		}
	}
}
