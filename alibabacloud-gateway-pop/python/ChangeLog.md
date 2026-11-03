### 2026-11-03 Version 0.1.5

* Sign POP V4 without endpoint-util, openapi-util, tea-util, or tea-xml.
* Keep JSON encoding for `TeaModel`, including a model nested in a dict, by calling `to_map()`.
* Parse POP error XML with the Darabonba builtin.
