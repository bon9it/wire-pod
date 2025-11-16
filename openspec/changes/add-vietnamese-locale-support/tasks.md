## 1. Specification
- [x] Review existing locale/spec assets (setup UI, SDK settings, localization package)
- [x] Validate no other change covers Vietnamese locale already
- [ ] Finalize `vietnamese-locale` spec delta and run `openspec validate` once tooling is available

## 2. Implementation Prep
- [x] Update SDK settings + setup UI to expose and persist vi-VN locale + STT selections
- [x] Extend backend locale mapper to emit `LanguageCode_VIETNAMESE` and add unit tests
- [x] Ensure STT download + localization loaders handle vi-VN (models, intents, error UX)
- [x] Define TTS/KG prompt updates and OpenAI fallback handling for Vietnamese replies
- [x] Add observability (status indicators + logs) confirming Vietnamese mode is active

## 3. Verification
- [x] Add automated tests for locale mapping + localization helpers
- [ ] Manual test plan: set locale to vi-VN, trigger hotword + KG conversation, confirm Vietnamese responses end-to-end
