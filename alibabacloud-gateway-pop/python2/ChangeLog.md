### 2026-11-10 Version 0.0.5

* Fix POP V4 query signing by percent-encoding parameter names and values before sorting by the encoded names.
* Keep the equals sign for empty query values and canonicalize signed header names case-insensitively.
* Use a single entry per header name. Tea Python 2 can collapse keys that differ only in case when sending the request.

### 2026-11-05 Version 0.0.4

* Python 2 is not part of this release. The package version stays 0.0.4.
