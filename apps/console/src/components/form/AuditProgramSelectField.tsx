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
import { useEffect, useState } from "react";
import {
  type Control,
  Controller,
  type FieldPath,
  type FieldValues,
} from "react-hook-form";
import { useTranslation } from "react-i18next";

import { usePaginatedAuditPrograms } from "#/hooks/graph/usePaginatedAuditPrograms";

type Props<T extends FieldValues> = {
  organizationId: string;
  frameworkId?: string | null;
  control: Control<T>;
  name: FieldPath<T>;
  disabled?: boolean;
};

export function AuditProgramSelectField<T extends FieldValues>({
  organizationId,
  frameworkId,
  control,
  name,
  disabled,
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
        const selected = field.value
          ? allPrograms.find(program => program.id === field.value)
          : null;

        return (
          <AuditProgramCombobox
            fieldName={String(field.name)}
            value={field.value ?? ""}
            selectedName={selected?.name ?? ""}
            search={search}
            onSearch={setSearch}
            onChange={field.onChange}
            onBlur={field.onBlur}
            programs={filteredPrograms}
            allPrograms={allPrograms}
            frameworkId={frameworkId}
            disabled={disabled}
            hasNext={hasNext}
            isLoadingNext={isLoadingNext}
            loadNext={loadNext}
            noneLabel={t("auditProgramSelect.none")}
          />
        );
      }}
    />
  );
}

function AuditProgramCombobox({
  fieldName,
  value,
  selectedName,
  search,
  onSearch,
  onChange,
  onBlur,
  programs,
  allPrograms,
  frameworkId,
  disabled,
  hasNext,
  isLoadingNext,
  loadNext,
  noneLabel,
}: {
  fieldName: string;
  value: string;
  selectedName: string;
  search: string;
  onSearch: (query: string) => void;
  onChange: (value: string) => void;
  onBlur: () => void;
  programs: { id: string; name: string; framework?: { id: string } | null }[];
  allPrograms: { id: string; name: string; framework?: { id: string } | null }[];
  frameworkId?: string | null;
  disabled?: boolean;
  hasNext: boolean;
  isLoadingNext: boolean;
  loadNext: (count: number) => void;
  noneLabel: string;
}) {
  useEffect(() => {
    if (!value || !frameworkId) {
      return;
    }
    const selected = allPrograms.find(program => program.id === value);
    if (selected && selected.framework?.id !== frameworkId) {
      onChange("");
      onSearch("");
    }
  }, [allPrograms, frameworkId, onChange, onSearch, value]);

  return (
    <Combobox
      id={fieldName}
      name={fieldName}
      onBlur={onBlur}
      placeholder={noneLabel}
      value={search || selectedName || ""}
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
