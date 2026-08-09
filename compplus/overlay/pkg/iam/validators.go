// Copyright (c) 2025-2026 Probo Inc <hello@probo.com>.
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

package iam

import "go.probo.inc/probo/pkg/validator"

// PasswordValidator remains compatible with existing accounts. Authentication
// must not reject a legitimate legacy password merely because the policy for
// newly-created passwords has become stronger.
func PasswordValidator() validator.ValidatorFunc {
	return passwordValidatorWithMinimum(8)
}

// StrongPasswordValidator is used whenever a user creates, resets, or changes
// a password. Twelve characters gives a materially stronger baseline without
// imposing brittle composition rules that encourage predictable passwords.
func StrongPasswordValidator() validator.ValidatorFunc {
	return passwordValidatorWithMinimum(12)
}

func passwordValidatorWithMinimum(minimum int) validator.ValidatorFunc {
	validators := []validator.ValidatorFunc{
		validator.NotEmpty(),
		validator.MaxLen(255), // Bound hashing work to mitigate password-DoS input.
		validator.MinLen(minimum),
	}

	return func(value any) *validator.ValidationError {
		for _, validate := range validators {
			if err := validate(value); err != nil {
				return err
			}
		}

		return nil
	}
}
