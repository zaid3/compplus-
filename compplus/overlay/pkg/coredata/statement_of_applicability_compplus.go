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

// LoadByOrganizationIDAndName supports idempotent Comp Plus+ SoA installation.
func (s *StatementOfApplicability) LoadByOrganizationIDAndName(
    ctx context.Context,
    conn pg.Querier,
    scope Scoper,
    organizationID gid.GID,
    name string,
) error {
    q := `
SELECT
    id,
    organization_id,
    name,
    document_id,
    created_at,
    updated_at
FROM statements_of_applicability
WHERE
    %s
    AND organization_id = @organization_id
    AND name = @name
ORDER BY created_at ASC
LIMIT 1;
`

    q = fmt.Sprintf(q, scope.SQLFragment())
    args := pgx.NamedArgs{
        "organization_id": organizationID,
        "name": name,
    }
    maps.Copy(args, scope.SQLArguments())

    rows, err := conn.Query(ctx, q, args)
    if err != nil {
        return fmt.Errorf("cannot query Comp Plus+ statement of applicability: %w", err)
    }

    statement, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[StatementOfApplicability])
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return ErrResourceNotFound
        }
        return fmt.Errorf("cannot collect Comp Plus+ statement of applicability: %w", err)
    }

    *s = statement
    return nil
}
