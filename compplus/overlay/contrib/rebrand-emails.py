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

from pathlib import Path


def replace_exact(path: Path, replacements: dict[str, str]) -> None:
    text = path.read_text()
    for old, new in replacements.items():
        if old not in text and new not in text:
            raise SystemExit(f"expected email branding text missing in {path}: {old!r}")
        text = text.replace(old, new)
    path.write_text(text)


# Plain-text transactional templates are generated separately from the React
# HTML templates. They contain no license headers, so every customer-facing
# occurrence can be safely rebranded without touching upstream attribution.
templates = Path("packages/emails/templates")
for path in templates.glob("*.txt"):
    text = path.read_text()
    text = text.replace("Probo", "ISO Pilot")
    text = text.replace("{{.SenderCompanyHeadquarterAddress}}\n", "")
    text = text.replace("Powered By ISO Pilot", "ISO Pilot")
    path.write_text(text)

# A few React templates have product wording in their message bodies or title.
# Patch only those exact visible literals so copyright/license attribution and
# internal package identifiers remain unchanged.
replace_exact(
    Path("packages/emails/src/PasswordReset.tsx"),
    {
        "You have requested a password reset for your Probo account.":
            "You have requested a password reset for your ISO Pilot account.",
    },
)
for email_path in (
    Path("packages/emails/src/DocumentSigning.tsx"),
    Path("packages/emails/src/DocumentApproval.tsx"),
):
    replace_exact(
        email_path,
        {
            "This process is managed securely by Probo, acting as the compliance partner":
                "This process is managed securely by ISO Pilot, acting as the compliance partner",
        },
    )
replace_exact(
    Path("packages/emails/src/MagicLink.tsx"),
    {
        '<EmailLayout subject="Probo Magic Link">':
            '<EmailLayout subject="ISO Pilot Magic Link">',
    },
)

# Rebrand the common presenter defaults used by every HTML/text email. Keep the
# upstream Go package/module identifiers intact; only literal customer-facing
# values are changed.
emails_go = Path("packages/emails/emails.go")
replace_exact(
    emails_go,
    {
        'SenderCompanyName:               "Probo"': 'SenderCompanyName:               "ISO Pilot"',
        'SenderCompanyWebsiteURL:         "https://www.probo.com"': 'SenderCompanyWebsiteURL:         "https://app.isopilot.co.uk"',
        'SenderCompanyHeadquarterAddress: "Probo Inc, 490 Post St, STE 640, San Francisco, CA, 94102, US"': 'SenderCompanyHeadquarterAddress: ""',
        'subjectInvitation                             = "Invitation to join %s on Probo"': 'subjectInvitation                             = "Invitation to join %s on ISO Pilot"',
    },
)

print("ISO Pilot email branding applied")
