package localization

import (
	"testing"

	"github.com/kercre123/wire-pod/chipper/pkg/vars"
)

func TestGetTextVietnamese(t *testing.T) {
	origLang := vars.APIConfig.STT.Language
	defer func() {
		vars.APIConfig.STT.Language = origLang
	}()

	vars.APIConfig.STT.Language = "vi-VN"
	got := GetText(STR_WEATHER_FORECAST)
	if got != "dự báo" {
		t.Fatalf("expected Vietnamese translation 'dự báo', got %q", got)
	}
}
