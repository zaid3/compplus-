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

import { Combobox, ComboboxItem, InfiniteScrollTrigger } from "@probo/ui";
import { Suspense, useEffect, useState } from "react";
import {
  type Control,
  Controller,
  type FieldPath,
  type FieldValues,
} from "react-hook-form";
import { useTranslation } from "react-i18next";
import { graphql, useLazyLoadQuery } from "react-relay";

import type { AuditProgramSelectFieldSelectedQuery } from "#/__generated__/core/AuditProgramSelectFieldSelectedQuery.graphql";
import { usePaginatedAuditPrograms } from "#/hooks/graph/usePaginatedAuditPrograms";

type SelectedAuditProgram = {
  id: string;
  name: string;
  framework?: { id: string } | null;
};

type Props<T extends FieldValues> = {
  organizationId: string;
  frameworkId?: string | null;
  control: Control<T>;
  name: FieldPath<T>;
  disabled?: boolean;
  selectedAuditProgram?: SelectedAuditProgram | null;
};

const selectedAuditProgramQuery = graphql`
  query AuditProgramSelectFieldSelectedQuery($id: ID!) {
    node(id: $id) {
      ... on AuditProgram {
        id
        name
        framework {
          id
        }
      }
    }
  }
`;

export function AuditProgramSelectField<T extends FieldValues>({
  organizationId,
  frameworkId,
  control,
  name,
  disabled,
  selectedAuditProgram,
}: Props<T>) {
  const { t } = useTranslation("organizations/audit-programs");
  const { data, loadNext, hasNext, isLoadingNext }
    = usePaginatedAuditPrograms(organizationId);
  const [search, setSearch] = useState("");

  useEffect(() => {
    setSearch("");
  }, [frameworkId]);

  const allPrograms
    = data?.auditPrograms?.edges
      ?.map(edge => edge.node)
      .filter((node): node is NonNullable<typeof node> => node !== null) ?? [];

  const programs = allPrograms.filter(
    program => !frameworkId || program.framework?.id === frameworkId,
  );

  const filteredPrograms = programs.filter(program =>
    program.name.toLowerCase().includes(search.toLowerCase()),
  );

  return (
    <Controller
      control={control}
      name={name}
      render={({ field }) => {
        const selectedFromList = field.value
          ? allPrograms.find(program => program.id === field.value)
          : null;
        const selectedFromProp
          = selectedAuditProgram && selectedAuditProgram.id === field.value
            ? selectedAuditProgram
            : null;
        const selected = selectedFromList ?? selectedFromProp ?? null;
        const needsFetch = Boolean(field.value) && !selected;

        return (
          <Suspense
            fallback={(
              <AuditProgramCombobox
                fieldName={String(field.name)}
                value={field.value ?? ""}
                selected={null}
                search={search}
                onSearch={setSearch}
                onChange={field.onChange}
                onBlur={field.onBlur}
                programs={filteredPrograms}
                frameworkId={frameworkId}
                disabled={disabled}
                hasNext={hasNext}
                isLoadingNext={isLoadingNext}
                loadNext={loadNext}
                noneLabel={t("auditProgramSelect.none")}
              />
            )}
          >
            {needsFetch
              ? (
                  <AuditProgramComboboxWithFetchedSelected
                    fieldName={String(field.name)}
                    value={field.value}
                    search={search}
                    onSearch={setSearch}
                    onChange={field.onChange}
                    onBlur={field.onBlur}
                    programs={filteredPrograms}
                    frameworkId={frameworkId}
                    disabled={disabled}
                    hasNext={hasNext}
                    isLoadingNext={isLoadingNext}
                    loadNext={loadNext}
                    noneLabel={t("auditProgramSelect.none")}
                  />
                )
              : (
                  <AuditProgramCombobox
                    fieldName={String(field.name)}
                    value={field.value ?? ""}
                    selected={selected}
                    search={search}
                    onSearch={setSearch}
                    onChange={field.onChange}
                    onBlur={field.onBlur}
                    programs={filteredPrograms}
                    frameworkId={frameworkId}
                    disabled={disabled}
                    hasNext={hasNext}
                    isLoadingNext={isLoadingNext}
                    loadNext={loadNext}
                    noneLabel={t("auditProgramSelect.none")}
                  />
                )}
          </Suspense>
        );
      }}
    />
  );
}

function AuditProgramComboboxWithFetchedSelected({
  value,
  frameworkId,
  onChange,
  onSearch,
  ...props
}: {
  fieldName: string;
  value: string;
  search: string;
  onSearch: (query: string) => void;
  onChange: (value: string) => void;
  onBlur: () => void;
  programs: SelectedAuditProgram[];
  frameworkId?: string | null;
  disabled?: boolean;
  hasNext: boolean;
  isLoadingNext: boolean;
  loadNext: (count: number) => void;
  noneLabel: string;
}) {
  const data = useLazyLoadQuery<AuditProgramSelectFieldSelectedQuery>(
    selectedAuditProgramQuery,
    { id: value },
    { fetchPolicy: "store-or-network" },
  );

  const fetched
    = data.node?.id
      ? {
          id: data.node.id,
          name: data.node.name ?? "",
          framework: data.node.framework,
        }
      : null;

  useEffect(() => {
    if (!fetched || !frameworkId || !fetched.framework?.id) {
      return;
    }
    if (fetched.framework.id !== frameworkId) {
      onChange("");
      onSearch("");
    }
  }, [fetched, frameworkId, onChange, onSearch]);

  return (
    <AuditProgramCombobox
      {...props}
      value={value}
      selected={fetched}
      frameworkId={frameworkId}
      onChange={onChange}
      onSearch={onSearch}
    />
  );
}

function AuditProgramCombobox({
  fieldName,
  value,
  selected,
  search,
  onSearch,
  onChange,
  onBlur,
  programs,
  frameworkId,
  disabled,
  hasNext,
  isLoadingNext,
  loadNext,
  noneLabel,
}: {
  fieldName: string;
  value: string;
  selected: SelectedAuditProgram | null;
  search: string;
  onSearch: (query: string) => void;
  onChange: (value: string) => void;
  onBlur: () => void;
  programs: SelectedAuditProgram[];
  frameworkId?: string | null;
  disabled?: boolean;
  hasNext: boolean;
  isLoadingNext: boolean;
  loadNext: (count: number) => void;
  noneLabel: string;
}) {
  useEffect(() => {
    if (!value || !frameworkId || !selected?.framework?.id) {
      return;
    }
    if (selected.framework.id !== frameworkId) {
      onChange("");
      onSearch("");
    }
  }, [frameworkId, onChange, onSearch, selected, value]);

  return (
    <Combobox
      id={fieldName}
      name={fieldName}
      onBlur={onBlur}
      placeholder={noneLabel}
      value={search || selected?.name || ""}
      onSearch={onSearch}
      disabled={disabled}
    >
      <ComboboxItem
        onClick={() => {
          onChange("");
          onSearch("");
        }}
      >
        {noneLabel}
      </ComboboxItem>
      {programs.map(program => (
        <ComboboxItem
          key={program.id}
          onClick={() => {
            onChange(program.id);
            onSearch(program.name);
          }}
        >
          {program.name}
        </ComboboxItem>
      ))}
      {hasNext && (
        <InfiniteScrollTrigger
          loading={isLoadingNext}
          onView={() => loadNext(50)}
        />
      )}
    </Combobox>
  );
}
