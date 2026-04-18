package validators

import "testing"

func TestRemoveSpecialCharacters(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"12345-678", "12345678"},
		{"12345678", "12345678"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := RemoveSpecialCharacters(tt.input); got != tt.want {
			t.Errorf("RemoveSpecialCharacters(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
