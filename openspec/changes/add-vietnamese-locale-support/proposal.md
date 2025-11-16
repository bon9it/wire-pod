# Change: Vietnamese locale parity across Vector voice flows

## Why
Vietnamese is exposed as an STT option, but Vector cannot set its locale to vi-VN, the cloud voice service maps any unknown locale back to English, and KG replies often drop Vietnamese diacritics. Owners who pick Vietnamese therefore get English prompts, missing intents, or silent responses. We need a spec-level change to define what true Vietnamese support means before touching the runtime.

## What Changes
- Define Vietnamese locale requirements covering locale selection, STT/Vosk/Whisper download flows, intent/localization resources, and OpenAI/Vector TTS handling
- Capture UX expectations for setup UI, SDK settings, and logging/telemetry so owners can confirm the active locale
- Enumerate validation and testing obligations for vi-VN (unit + manual) before implementation begins

## Impact
- Affected specs: `vietnamese-locale`
- Affected code (later implementation): setup UI, SDK settings endpoints, `vector-cloud/internal/voice/process.go`, localization + STT downloaders, KG/TTS pipeline
