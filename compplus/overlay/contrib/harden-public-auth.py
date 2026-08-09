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


def guarded_replace(path: str, old: str, new: str) -> None:
    file_path = Path(path)
    text = file_path.read_text()
    if new in text:
        return
    if old not in text:
        raise SystemExit(f"Expected auth source block not found in {path}")
    file_path.write_text(text.replace(old, new, 1))


# A magic link proves control of a mailbox, but it must not silently create an
# account. New identities go through explicit signup + verification so signup
# throttles and the operator kill switch cannot be bypassed.
guarded_replace(
    "pkg/iam/auth_service.go",
    '''\t\terr := identity.LoadByEmail(ctx, tx, payload.Data.Email)\n\t\t\tif err != nil {\n\t\t\t\tif errors.Is(err, coredata.ErrResourceNotFound) {\n\t\t\t\t\tidentity = &coredata.Identity{\n\t\t\t\t\t\tID:                   gid.New(gid.NilTenant, coredata.IdentityEntityType),\n\t\t\t\t\t\tEmailAddress:         payload.Data.Email,\n\t\t\t\t\t\tEmailAddressVerified: true,\n\t\t\t\t\t\tCreatedAt:            now,\n\t\t\t\t\t\tUpdatedAt:            now,\n\t\t\t\t\t}\n\n\t\t\t\t\tif err := identity.Insert(ctx, tx); err != nil {\n\t\t\t\t\t\treturn fmt.Errorf("cannot create identity: %w", err)\n\t\t\t\t\t}\n\t\t\t\t} else {\n\t\t\t\t\treturn fmt.Errorf("cannot load identity by email: %w", err)\n\t\t\t\t}\n\t\t\t} else if !identity.EmailAddressVerified {''',
    '''\t\terr := identity.LoadByEmail(ctx, tx, payload.Data.Email)\n\t\t\tif err != nil {\n\t\t\t\tif errors.Is(err, coredata.ErrResourceNotFound) {\n\t\t\t\t\treturn NewInvalidCredentialsError("invalid magic link")\n\t\t\t\t}\n\n\t\t\t\treturn fmt.Errorf("cannot load identity by email: %w", err)\n\t\t\t} else if !identity.EmailAddressVerified {''',
)

# Upstream creates a root session during signup before email verification. Close
# that new session and do not set a browser cookie: the user must verify the
# mailbox and then authenticate normally.
guarded_replace(
    "pkg/server/api/connect/v1/session_resolvers.go",
    '''\tw := gqlutils.HTTPResponseWriterFromContext(ctx)\n\tr.sessionCookie.Set(w, session)\n\n\treturn &types.SignUpPayload{\n\t\tIdentity: types.NewIdentity(identity),\n\t}, nil\n}\n\n// SignOut is the resolver for the signOut field.''',
    '''\tif err := r.iam.SessionService.CloseSession(ctx, session.ID); err != nil {\n\t\tr.logger.ErrorCtx(ctx, "cannot close unverified signup session", log.Error(err))\n\t\treturn nil, gqlutils.Internal(ctx)\n\t}\n\n\treturn &types.SignUpPayload{\n\t\tIdentity: types.NewIdentity(identity),\n\t}, nil\n}\n\n// SignOut is the resolver for the signOut field.''',
)
