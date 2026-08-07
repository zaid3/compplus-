package coredata

import (
    "context"
    "errors"
    "fmt"
    "maps"

    "github.com/jackc/pgx/v5"
    "go.gearno.de/kit/pg"
    "go.probo.inc/probo/pkg/gid"
)

// LoadByOrganizationIDAndReferenceID resolves a framework reference inside one
// organization. Upstream LoadByReferenceID is tenant-scoped; Comp Plus+ supports
// several independent organizations in one tenant, so template imports must not
// accidentally resolve another organization's framework with the same reference.
func (f *Framework) LoadByOrganizationIDAndReferenceID(
    ctx context.Context,
    conn pg.Querier,
    scope Scoper,
    organizationID gid.GID,
    referenceID string,
) error {
    q := `
SELECT
    id,
    organization_id,
    reference_id,
    name,
    description,
    light_logo_file_id,
    dark_logo_file_id,
    created_at,
    updated_at
FROM
    frameworks
WHERE
    %s
    AND organization_id = @organization_id
    AND reference_id = @reference_id
LIMIT 1;
`

    q = fmt.Sprintf(q, scope.SQLFragment())
    args := pgx.NamedArgs{
        "organization_id": organizationID,
        "reference_id": referenceID,
    }
    maps.Copy(args, scope.SQLArguments())

    rows, err := conn.Query(ctx, q, args)
    if err != nil {
        return fmt.Errorf("cannot query framework by organization and reference id: %w", err)
    }

    framework, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Framework])
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return ErrResourceNotFound
        }
        return fmt.Errorf("cannot collect framework by organization and reference id: %w", err)
    }

    *f = framework
    return nil
}
