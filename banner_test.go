package main

import (
	"os"
	"reflect"
	"testing"
)

func createTempBannerFile(t *testing.T, content string) string {
	t.Helper()

	file, err := os.CreateTemp("", "banner-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	return file.Name()
}

func TestLoadBanner(t *testing.T) {
	// ASCII 32 = space
	// ASCII 33 = !
	content := `
        
        
        
        
        
        
        
        

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....
`

	file := createTempBannerFile(t, content)
	defer os.Remove(file)

	banner, err := loadBanner(file)
	if err != nil {
		t.Fatalf("loadBanner returned error: %v", err)
	}

	tests := []struct {
		char     rune
		expected []string
	}{
		{
			char: ' ',
			expected: []string{
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
			},
		},
		{
			char: '!',
			expected: []string{
				"._..",
				"|.|.",
				"|.|.",
				"|.|.",
				"|_|.",
				"(_).",
				"....",
				"....",
			},
		},
	}

	for _, tt := range tests {
		got, ok := banner[tt.char]

		if !ok {
			t.Errorf("character %q not found in banner", tt.char)
			continue
		}

		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("wrong banner for %q\nexpected: %v\ngot: %v",
				tt.char, tt.expected, got)
		}
	}
}

func TestLoadBanner_FileNotFound(t *testing.T) {
	_, err := loadBanner("does-not-exist.txt")

	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoadBanner_CharacterCount(t *testing.T) {
	content := `
        
        
        
        
        
        
        
        

####
#..#
#..#
####
#..#
#..#
####
....
`

	file := createTempBannerFile(t, content)
	defer os.Remove(file)

	banner, err := loadBanner(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(banner) != 2 {
		t.Errorf("expected 2 characters, got %d", len(banner))
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	file := createTempBannerFile(t, "")
	defer os.Remove(file)

	banner, err := loadBanner(file)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(banner) != 0 {
		t.Errorf("expected empty banner map, got %d entries", len(banner))
	}
}
