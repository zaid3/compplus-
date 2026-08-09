import { usePageTitle } from "@probo/hooks";
import {
  Button,
  IconPlusLarge,
  PageHeader,
  TabItem,
  Tabs,
} from "@probo/ui";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import {
  type PreloadedQuery,
  usePreloadedQuery,
} from "react-relay";
import { ConnectionHandler, graphql } from "relay-runtime";

import type { DocumentsPageQuery } from "#/__generated__/core/DocumentsPageQuery.graphql";
import { useOrganizationId } from "#/hooks/useOrganizationId";

import { CreateDocumentDialog } from "./_components/CreateDocumentDialog";
import { DocumentList } from "./_components/DocumentList";

export const documentsPageQuery = graphql`
  query DocumentsPageQuery($organizationId: ID!) {
    organization: node(id: $organizationId) {
      __typename
      ... on Organization {
        canCreateDocument: permission(action: "core:document:create")
        ...DocumentListFragment @arguments(first: 50, order: { field: TITLE, direction: ASC })
      }
    }
  }
`;

export default function DocumentsPage(props: {
  queryRef: PreloadedQuery<DocumentsPageQuery>;
}) {
  const { queryRef } = props;
  const organizationId = useOrganizationId();
  const { t } = useTranslation();
  const { organization } = usePreloadedQuery<DocumentsPageQuery>(documentsPageQuery, queryRef);

  if (organization.__typename !== "Organization") {
    throw new Error("invalid type for node");
  }

  usePageTitle("Documents");

  const [tab, setTab] = useState<"ACTIVE" | "ARCHIVED">("ACTIVE");
  const [documentListConnectionId, setDocumentListConnectionId] = useState(
    ConnectionHandler.getConnectionID(
      organizationId,
      "DocumentsListQuery_documents",
      { orderBy: { direction: "ASC", field: "TITLE" } },
    ),
  );

  return (
    <div className="space-y-6">
      <PageHeader
        title={t("documentsPage.title")}
        description="Create, review and manage your organisation documents."
      >
        <div className="flex gap-2">
          {organization.canCreateDocument && tab === "ACTIVE" && (
            <CreateDocumentDialog
              connection={documentListConnectionId}
              trigger={(
                <Button icon={IconPlusLarge}>
                  {t("documentsPage.actions.new")}
                </Button>
              )}
            />
          )}
        </div>
      </PageHeader>
      <Tabs>
        <TabItem active={tab === "ACTIVE"} onClick={() => setTab("ACTIVE")}>
          {t("documentsPage.tabs.active")}
        </TabItem>
        <TabItem active={tab === "ARCHIVED"} onClick={() => setTab("ARCHIVED")}>
          {t("documentsPage.tabs.archived")}
        </TabItem>
      </Tabs>
      <DocumentList
        fKey={organization}
        onConnectionIdChange={setDocumentListConnectionId}
        tab={tab}
      />
    </div>
  );
}
