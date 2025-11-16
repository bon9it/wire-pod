package vars

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestLoadIntentsVietnamese(t *testing.T) {
	origLang := APIConfig.STT.Language
	APIConfig.STT.Language = "vi-VN"
	t.Cleanup(func() {
		APIConfig.STT.Language = origLang
	})

	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("unable to determine test file path")
	}
	chipperRoot := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", ".."))
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd failed: %v", err)
	}
	if err := os.Chdir(chipperRoot); err != nil {
		t.Fatalf("chdir to chipper root failed: %v", err)
	}
	t.Cleanup(func() {
		os.Chdir(wd)
	})

	intents, err := LoadIntents()
	if err != nil {
		t.Fatalf("LoadIntents returned error: %v", err)
	}
	if len(intents) == 0 {
		t.Fatal("expected intents for vi-VN, got none")
	}
	found := false
	for _, intent := range intents {
		if intent.Name == "intent_weather_extend" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected intent_weather_extend in vi-VN intents")
	}
}
