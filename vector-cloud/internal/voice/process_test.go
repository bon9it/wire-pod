package voice

import (
	"testing"

	pb "github.com/digital-dream-labs/api/go/chipperpb"
)

func TestGetLanguageVietnamese(t *testing.T) {
	tests := []string{"vi-VN", "vi_VN"}
	for _, locale := range tests {
		code, err := getLanguage(locale)
		if err != nil {
			t.Fatalf("expected no error for %s, got %v", locale, err)
		}
		if code != pb.LanguageCode_VIETNAMESE {
			t.Fatalf("expected Vietnamese language code for %s, got %v", locale, code)
		}
	}
}

func TestGetLanguageInvalid(t *testing.T) {
	if _, err := getLanguage("invalid"); err == nil {
		t.Fatalf("expected error for invalid locale")
	}
}
