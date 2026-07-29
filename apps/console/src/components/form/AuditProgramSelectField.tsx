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

import { Option, Select } from "@probo/ui";
import { type Control, Controller, type FieldPath, type FieldValues } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useLazyLoadQuery } from "react-relay";

import type { AuditProgramGraphOptionsQuery } from "#/__generated__/core/AuditProgramGraphOptionsQuery.graphql";
import { auditProgramOptionsQuery } from "#/hooks/graph/AuditProgramGraph";

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
  const data = useLazyLoadQuery<AuditProgramGraphOptionsQuery>(
    auditProgramOptionsQuery,
    { organizationId },
    { fetchPolicy: "store-or-network" },
  );

  const programs
    = data?.node?.auditPrograms?.edges
      ?.map(edge => edge.node)
      .filter((node): node is NonNullable<typeof node> => node !== null)
      .filter(
        program =>
          !frameworkId || program.framework?.id === frameworkId,
      ) ?? [];

  return (
    <Controller
      control={control}
      name={name}
      render={({ field }) => (
        <Select
          id={String(name)}
          variant="editor"
          disabled={disabled}
          placeholder={t("auditProgramSelect.none")}
          onValueChange={(value) => {
            field.onChange(value === "__none__" ? "" : value);
          }}
          className="w-full"
          value={field.value ? field.value : "__none__"}
        >
          <Option value="__none__">{t("auditProgramSelect.none")}</Option>
          {programs.map(program => (
            <Option key={program.id} value={program.id}>
              {program.name}
            </Option>
          ))}
        </Select>
      )}
    />
  );
}
