package main

import (
	"errors"
	"testing"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantLen int
		wantErr bool
	}{
		{
			name:    "Valid UTF-8",
			input:   []byte("Hello, world!"),
			wantLen: 13,
			wantErr: false,
		},
		{
			name:    "Empty string",
			input:   []byte(""),
			wantLen: 0,
			wantErr: false,
		},
		{
			name:    "Invalid UTF-8",
			input:   []byte{0xFF, 0xFE}, // Invalid UTF-8 sequence
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "Valid UTF-8 with multibyte runes",
			input:   []byte("こんにちは世界"), // Japanese "Konnichiwa Sekai"
			wantLen: 7,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLen, gotErr := GetUTFLength(tt.input)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetUTFLength() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotLen != tt.wantLen {
				t.Errorf("GetUTFLength() = %v, want %v", gotLen, tt.wantLen)
			}
		})
	}
}
func main() {}
