#!/usr/bin/env python3
# Copyright (c) 2026 Probo Inc <hello@probo.com>.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

"""Rebrand customer-facing locale values without ever changing JSON keys."""

from __future__ import annotations

import json
from pathlib import Path
from typing import Any

REPLACEMENTS = (
    ("Comp Plus+", "ISO Pilot"),
    ("ISOpilot", "ISO Pilot"),
    ("ISOPilot", "ISO Pilot"),
    ("Probo", "ISO Pilot"),
)


def rebrand(value: Any) -> Any:
    if isinstance(value, str):
        for old, new in REPLACEMENTS:
            value = value.replace(old, new)
        return value
    if isinstance(value, list):
        return [rebrand(item) for item in value]
    if isinstance(value, dict):
        # Deliberately preserve keys. React i18n lookups depend on stable keys
        # such as signInPage.newToProbo even when the displayed product name is
        # rebranded.
        return {key: rebrand(item) for key, item in value.items()}
    return value


def main() -> None:
    locale_files = sorted(Path("apps").glob("*/src/_locales/*.json"))
    if not locale_files:
        raise SystemExit("no application locale files found")

    for path in locale_files:
        data = json.loads(path.read_text(encoding="utf-8"))
        path.write_text(
            json.dumps(rebrand(data), ensure_ascii=False, indent=2) + "\n",
            encoding="utf-8",
        )


if __name__ == "__main__":
    main()
