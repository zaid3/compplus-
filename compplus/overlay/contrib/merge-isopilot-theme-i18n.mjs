// Copyright (c) 2026 Probo Inc <hello@probo.com>.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Merge the ISOPilot theme-control labels into the console app's default
// ("app") i18next namespace catalogs. The theme toggle is app-wide chrome, so
// per contrib/claude/i18n.md its strings belong in apps/console/src/_locales.
// We merge rather than overwrite so upstream keys are preserved, and adding the
// keys at build time keeps them working after the overlay is applied over a
// fresh upstream checkout. The script is idempotent: re-running it is a no-op.

import { readFileSync, writeFileSync } from "node:fs";

const CATALOG_DIR = "apps/console/src/_locales";

const THEME_KEYS = {
  "en-US": {
    options: { system: "System", light: "Light", dark: "Dark" },
    control: {
      label: "Theme: {{theme}}. Click to change theme.",
      title: "Theme: {{theme}}",
    },
  },
  "fr-FR": {
    options: { system: "Système", light: "Clair", dark: "Sombre" },
    control: {
      label: "Thème : {{theme}}. Cliquez pour changer de thème.",
      title: "Thème : {{theme}}",
    },
  },
};

let changed = 0;
for (const [locale, theme] of Object.entries(THEME_KEYS)) {
  const path = `${CATALOG_DIR}/${locale}.json`;
  const catalog = JSON.parse(readFileSync(path, "utf8"));
  const upstreamTheme = catalog.theme ?? {};

  catalog.theme = {
    ...theme,
    ...upstreamTheme,
    options: { ...theme.options, ...(upstreamTheme.options ?? {}) },
    control: { ...theme.control, ...(upstreamTheme.control ?? {}) },
  };

  writeFileSync(path, `${JSON.stringify(catalog, null, 2)}\n`);
  changed += 1;
  console.log(`merged theme i18n keys into ${path}`);
}

if (changed === 0) {
  console.error("no console locale catalogs were updated");
  process.exit(1);
}
