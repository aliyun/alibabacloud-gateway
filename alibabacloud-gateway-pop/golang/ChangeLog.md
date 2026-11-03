### 2026-11-03 Version 0.1.5

* Sign POP V4 without endpoint-util, openapi-util, tea-utils, tea-xml, or the small Darabonba utility packages.
* Keep returned errors as `*tea.SDKError`, so an HTTP 4xx response stays non-retryable for the existing retry check.
* Parse POP error XML with the Darabonba builtin. SM3 still uses `github.com/tjfoc/gmsm`.
