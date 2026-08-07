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

// LoadByOrganizationIDAndTitle lets the Comp Plus+ template installer resume a
// partially completed installation without creating duplicate documents.
func (d *Document) LoadByOrganizationIDAndTitle(
    ctx context.Context,
    conn pg.Querier,
    scope Scoper,
    organizationID gid.GID,
    title string,
) error {
    q := `
WITH latest_versions AS (
    SELECT DISTINCT ON (document_id)
        document_id,
        title,
        document_type
    FROM document_versions
    ORDER BY document_id, major DESC, minor DESC
)
SELECT
    documents.id,
    documents.organization_id,
    documents.current_published_major,
    documents.current_published_minor,
    documents.write_mode,
    documents.status,
    documents.archived_at,
    documents.created_at,
    documents.updated_at,
    COALESCE(lv.title, '') AS title,
    COALESCE(lv.document_type, 'OTHER') AS document_type
FROM documents
LEFT JOIN latest_versions lv ON lv.document_id = documents.id
WHERE
    %s
    AND documents.deleted_at IS NULL
    AND documents.organization_id = @organization_id
    AND lv.title = @title
ORDER BY documents.created_at ASC
LIMIT 1;
`

    q = fmt.Sprintf(q, scope.SQLFragment())
    args := pgx.NamedArgs{
        "organization_id": organizationID,
        "title": title,
    }
    maps.Copy(args, scope.SQLArguments())

    rows, err := conn.Query(ctx, q, args)
    if err != nil {
        return fmt.Errorf("cannot query Comp Plus+ document by organization and title: %w", err)
    }

    document, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Document])
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return ErrResourceNotFound
        }
        return fmt.Errorf("cannot collect Comp Plus+ document by organization and title: %w", err)
    }

    *d = document
    return nil
}
