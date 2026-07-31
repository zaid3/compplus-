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

import { safeOpenUrl } from "@probo/helpers";
import { usePageTitle } from "@probo/hooks";
import { Badge, Button, IconBell2, IconCheckmark1, IconFolder2, IconMedal, IconPageTextLine, IconPencil, IconPeopleAdd, IconSettingsGear2, IconShield, IconStore, PageHeader, TabLink, Tabs } from "@probo/ui";
import { useTranslation } from "react-i18next";
import { type PreloadedQuery, usePreloadedQuery } from "react-relay";
import { Outlet } from "react-router";
import { graphql } from "relay-runtime";

import type { CompliancePageLayoutQuery } from "#/__generated__/core/CompliancePageLayoutQuery.graphql";
import { useOrganizationId } from "#/hooks/useOrganizationId";

export const compliancePageLayoutQuery = graphql`
  query CompliancePageLayoutQuery($organizationId: ID!) {
    organization: node(id: $organizationId) {
      __typename
      ... on Organization {
        canListDocuments: permission(action: "core:document:list")
        canListAudits: permission(action: "core:audit:list")
        canListThirdParties: permission(action: "core:thirdParty:list")
        compliancePage: compliancePortal {
          # eslint-disable-next-line relay/unused-fields
          id
          active
          publicUrl
          canUpdatePortal: permission(action: "compliance-portal:portal:update")
          canListReferences: permission(action: "compliance-portal:portal-reference:list")
          canListCommitmentGroups: permission(action: "compliance-portal:commitment-group:list")
          canListFiles: permission(action: "compliance-portal:portal-file:list")
          canListAccess: permission(action: "compliance-portal:portal-access:list")
          canListMailingUpdates: permission(action: "compliance-portal:mailing-list-update:list")
        }
      }
    }
  }
`;

export function CompliancePageLayout(props: { queryRef: PreloadedQuery<CompliancePageLayoutQuery> }) {
  const { queryRef } = props;

  const organizationId = useOrganizationId();
  const { t } = useTranslation("organizations/compliance-page");

  usePageTitle(t("layout.title"));

  const { organization } = usePreloadedQuery<CompliancePageLayoutQuery>(compliancePageLayoutQuery, queryRef);
  if (organization.__typename !== "Organization") {
    throw new Error("invalid type for node");
  }

  const compliancePage = organization.compliancePage;
  const compliancePageUrl = compliancePage?.publicUrl || null;
  const prefix = `/organizations/${organizationId}/compliance-page`;

  return (
    <div className="space-y-6">
      <PageHeader
        title={t("layout.title")}
        description={t("layout.description")}
      >
        <Badge variant={compliancePage?.active ? "success" : "danger"}>
          {compliancePage?.active
            ? t("layout.status.active")
            : t("layout.status.inactive")}
        </Badge>
        {compliancePage?.active && compliancePageUrl && (
          <Button
            variant="secondary"
            onClick={() => safeOpenUrl(compliancePageUrl)}
          >
            {t("layout.actions.open")}
          </Button>
        )}
      </PageHeader>

      <Tabs>
        {(compliancePage?.canUpdatePortal || compliancePage?.canListReferences) && (
          <TabLink to={prefix} end>
            <IconSettingsGear2 className="size-4" />
            {t("layout.tabs.overview")}
          </TabLink>
        )}
        {compliancePage?.canUpdatePortal && (
          <TabLink to={`${prefix}/brand`}>
            <IconPencil className="size-4" />
            {t("layout.tabs.brand")}
          </TabLink>
        )}
        {compliancePage?.canListReferences && (
          <TabLink to={`${prefix}/references`}>
            <IconCheckmark1 className="size-4" />
            {t("layout.tabs.references")}
          </TabLink>
        )}
        {compliancePage?.canListCommitmentGroups && (
          <TabLink to={`${prefix}/commitments`}>
            <IconShield className="size-4" />
            {t("layout.tabs.commitments")}
          </TabLink>
        )}
        {organization.canListAudits && (
          <TabLink to={`${prefix}/audits`}>
            <IconMedal className="size-4" />
            {t("layout.tabs.audits")}
          </TabLink>
        )}
        {organization.canListDocuments && (
          <TabLink to={`${prefix}/documents`}>
            <IconPageTextLine className="size-4" />
            {t("layout.tabs.documents")}
          </TabLink>
        )}
        {compliancePage?.canListFiles && (
          <TabLink to={`${prefix}/files`}>
            <IconFolder2 className="size-4" />
            {t("layout.tabs.files")}
          </TabLink>
        )}
        {organization.canListThirdParties && (
          <TabLink to={`${prefix}/third-parties`}>
            <IconStore className="size-4" />
            {t("layout.tabs.subprocessors")}
          </TabLink>
        )}
        {compliancePage?.canListAccess && (
          <TabLink to={`${prefix}/access`}>
            <IconPeopleAdd className="size-4" />
            {t("layout.tabs.access")}
          </TabLink>
        )}
        {compliancePage?.canListMailingUpdates && (
          <TabLink to={`${prefix}/mailing-list`}>
            <IconBell2 className="size-4" />
            {t("layout.tabs.mailingList")}
          </TabLink>
        )}
      </Tabs>

      <Outlet />
    </div>
  );
}
