#!/usr/bin/env python3
"""Fail if a shipped POP language package depends on the removed utility libraries.

main.tea and Teafile still import those libraries. Darabonba cannot express
HMAC without them, and `dara codegen` therefore emits the old calls. The
language packages in this directory are the maintained result after that
generation: private helpers replace the utility packages, and this script is
the check that a later codegen was not copied over those packages.

Python 2 is outside the V4 acceptance set and is not scanned.
"""

import pathlib
import sys

ROOT = pathlib.Path(__file__).resolve().parents[1]

MANIFESTS = [
    "golang/go.mod",
    "java/pom.xml",
    "ts/package.json",
    "php/composer.json",
    "python/setup.py",
    "csharp/core/client.csproj",
    "swift/Package.swift",
]

FORBIDDEN = (
    "tea-xml",
    "teaxml",
    "tea-util",
    "tea-utils",
    "openapi-util",
    "openapiutil",
    "endpoint-util",
    "darabonba-string",
    "darabonba-array",
    "darabonba-map",
    "darabonba-encode-util",
    "darabonba-signature-util",
    "alibabacloud_tea_xml",
    "alibabacloud_tea_util",
    "alibabacloud_openapi_util",
    "alibabacloud_endpoint_util",
    "alibabacloud_darabonba_",
    "TeaXML",
    "TeaUtil",
    "OpenApiUtil",
    "EndpointUtil",
    "DarabonbaString",
    "DarabonbaArray",
    "DarabonbaMap",
    "DarabonbaEncodeUtil",
    "DarabonbaSignatureUtil",
)


def main():
    failures = []
    for relative in MANIFESTS:
        path = ROOT / relative
        if not path.is_file():
            failures.append("%s: missing manifest" % relative)
            continue
        text = path.read_text(encoding="utf-8").lower()
        for token in FORBIDDEN:
            if token.lower() in text:
                failures.append("%s still depends on %s" % (relative, token))
    if failures:
        print("\n".join(failures))
        return 1
    print("POP language manifests do not depend on the removed utility packages")
    return 0


if __name__ == "__main__":
    sys.exit(main())
