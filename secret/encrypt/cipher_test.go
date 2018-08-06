package encrypt

import "testing"

func TestCipher(t *testing.T) {
	testSuit := []struct {
		key       string
		plainText string
	}{
		{key: "demo1", plainText: "123abc"},
		{key: "demo2", plainText: "123456"},
		{key: "demo3", plainText: "asdfg"},
		{key: "demo4", plainText: "qwerty"},
	}
	for i, test := range testSuit {
		hex, err := Encrypt(test.key, test.plainText)
		if err != nil {
			t.Error("error in encrypt")
		}
		plainText, err := Decrypt(test.key, hex)
		if err != nil {
			t.Error("error in decrypt")
		}
		if test.plainText != plainText {
			t.Errorf("error in both value not match %d", i)
		}
	}
}
