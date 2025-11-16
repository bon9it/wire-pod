## ADDED Requirements

### Requirement: Vietnamese locale can be selected everywhere Vector exposes locale controls
The web setup page and SDK settings UI SHALL expose `vi-VN` as a first-class locale option, persist the choice server-side, and push it to any connected robot so its runtime locale equals `vi-VN`.

#### Scenario: Owner sets locale from SDK settings
- **GIVEN** a robot connected to Wire-Pod and the owner opens the SDK settings page
- **WHEN** the owner selects "Vietnamese" and saves
- **THEN** the backend stores `locale = vi-VN`, returns success, and the locale radio list reflects the persisted value on refresh

#### Scenario: Locale propagates to voice processor
- **GIVEN** the stored locale is `vi-VN`
- **WHEN** the Vector runtime sends the next hotword request
- **THEN** cloud voice processing reports `LanguageCode_VIETNAMESE` (not an English fallback) when creating streaming sessions

### Requirement: vi-VN STT resources are provisioned and loaded without manual steps
Selecting Vietnamese SHALL trigger the correct STT model download (Vosk or Whisper), load Vietnamese intent JSON/grammar, and reject the selection if resources cannot be downloaded.

#### Scenario: Download succeeds
- **WHEN** a user picks vi-VN in the STT dropdown
- **THEN** the UI shows download progress, the Vietnamese acoustic model is stored, and a confirmation message indicates the language is active

#### Scenario: Download fails
- **WHEN** the model download errors
- **THEN** the UI stays on the picker, surfaces an actionable error, and Wire-Pod keeps its previous STT language

### Requirement: Knowledge Graph + TTS preserve Vietnamese text and pronunciation
KG/LLM prompts SHALL instruct the model to answer in Vietnamese, responses SHALL retain Vietnamese diacritics, and Vector SHALL synthesize those answers using either a Vietnamese OpenAI voice or Vector's hardware voice without dropping text.

#### Scenario: OpenAI TTS succeeds
- **WHEN** KG produces a Vietnamese response while OpenAI TTS is available
- **THEN** the OpenAI Vietnamese voice plays through Vector and logs show "Vietnamese voice" usage

#### Scenario: OpenAI TTS unavailable
- **WHEN** OpenAI rejects a request
- **THEN** the system logs the failure and immediately falls back to Vector's built-in TTS without stripping diacritics

### Requirement: Vietnamese localization + intents cover default voice commands
Wire-Pod SHALL ship Vietnamese translations for all default prompts/strings and include vi-VN variants of the stock intent list so weather, timers, alarms, volume, and cube commands match user speech.

#### Scenario: Built-in intent resolves in Vietnamese
- **GIVEN** the vi-VN intent JSON contains "thời tiết" phrases
- **WHEN** a user asks "Vector ơi, thời tiết hôm nay sao?"
- **THEN** the weather intent fires and responds (without deferring to KG)

### Requirement: Observability + validation prove vi-VN is active
Operators SHALL see the current locale/STT language in the setup UI, and CI SHALL include unit tests for locale mapping and localization helpers before implementation is considered done.

#### Scenario: UI status block
- **WHEN** Wire-Pod is running with Vietnamese enabled
- **THEN** the setup page displays the active STT language + locale and links to logs showing Vietnamese mode

#### Scenario: Automated tests
- **WHEN** CI executes
- **THEN** it runs tests that assert `getLanguage("vi-VN") == LanguageCode_VIETNAMESE` and localization helpers return Vietnamese strings
