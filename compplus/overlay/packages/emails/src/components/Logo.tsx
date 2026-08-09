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

import * as React from 'react';

export function Logo() {
  return (
    <table role="presentation" cellPadding="0" cellSpacing="0" style={{ borderCollapse: 'collapse' }}>
      <tbody>
        <tr>
          <td
            style={{
              width: '42px',
              height: '42px',
              borderRadius: '10px',
              backgroundColor: '#3b6df6',
              color: '#ffffff',
              fontSize: '24px',
              fontWeight: 700,
              lineHeight: '42px',
              textAlign: 'center',
              verticalAlign: 'middle',
            }}
          >
            ➤
          </td>
          <td
            style={{
              paddingLeft: '10px',
              color: '#111827',
              fontFamily: "-apple-system, BlinkMacSystemFont, 'Segoe UI', Arial, sans-serif",
              fontSize: '24px',
              fontWeight: 700,
              lineHeight: '30px',
              whiteSpace: 'nowrap',
              verticalAlign: 'middle',
            }}
          >
            ISO Pilot
          </td>
        </tr>
      </tbody>
    </table>
  );
}
