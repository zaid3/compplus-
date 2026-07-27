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

import { formatError } from "@probo/helpers";
import {
  Breadcrumb,
  Button,
  Dialog,
  DialogContent,
  DialogFooter,
  Field,
  Input,
  Label,
  Option,
  Select,
  Textarea,
  useDialogRef,
  useToast,
} from "@probo/ui";
import { type ReactNode } from "react";
import { Controller } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { graphql, useMutation } from "react-relay";
import { z } from "zod";

import type { CreateBusinessFunctionDialogMutation } from "#/__generated__/core/CreateBusinessFunctionDialogMutation.graphql";
import { AssetsMultiSelectField } from "#/components/form/AssetsMultiSelectField";
import { PeopleSelectField } from "#/components/form/PeopleSelectField";
import { ThirdPartiesMultiSelectField } from "#/components/form/ThirdPartiesMultiSelectField";
import { useFormWithSchema } from "#/hooks/useFormWithSchema";

import {
  businessFunctionClassificationOptions,
  durationMinutesHelperText,
} from "../_lib/businessFunctionHelpers";

const createBusinessFunctionMutation = graphql`
  mutation CreateBusinessFunctionDialogMutation(
    $input: CreateBusinessFunctionInput!
    $connections: [ID!]!
  ) {
    createBusinessFunction(input: $input) {
      businessFunctionEdge @prependEdge(connections: $connections) {
        node {
          id
          referenceId
          name
          classification
          mtdMinutes
          rtoMinutes
          rpoMinutes
          owner {
            id
            fullName
          }
          createdAt
          canUpdate: permission(action: "core:business-function:update")
          canDelete: permission(action: "core:business-function:delete")
        }
      }
    }
  }
`;

const schema = z.object({
  referenceId: z.string().min(1),
  name: z.string().min(1),
  classification: z.enum(["CRITICAL", "IMPORTANT", "SECONDARY", "STANDARD"]),
  mtdMinutes: z.coerce.number().int().min(0),
  rtoMinutes: z.coerce.number().int().min(0),
  rpoMinutes: z.coerce.number().int().min(0),
  impactTolerance: z.string().optional(),
  notes: z.string().optional(),
  ownerId: z.string().nullable().optional(),
  assetIds: z.array(z.string()).optional(),
  thirdPartyIds: z.array(z.string()).optional(),
});

type FormData = z.infer<typeof schema>;

interface CreateBusinessFunctionDialogProps {
  children: ReactNode;
  organizationId: string;
  connectionIds?: string[];
}

export function CreateBusinessFunctionDialog({
  children,
  organizationId,
  connectionIds,
}: CreateBusinessFunctionDialogProps) {
  const { t } = useTranslation();
  const { toast } = useToast();
  const dialogRef = useDialogRef();
  const [createBusinessFunction, isCreating] = useMutation<CreateBusinessFunctionDialogMutation>(
    createBusinessFunctionMutation,
  );

  const classificationOptions = businessFunctionClassificationOptions(
    t,
    "createBusinessFunctionDialog",
  );
  const durationHelper = durationMinutesHelperText(t, "createBusinessFunctionDialog");

  const { register, handleSubmit, formState, reset, control } = useFormWithSchema(schema, {
    defaultValues: {
      referenceId: "",
      name: "",
      classification: "STANDARD" as const,
      mtdMinutes: 0,
      rtoMinutes: 0,
      rpoMinutes: 0,
      impactTolerance: "",
      notes: "",
      ownerId: null,
      assetIds: [],
      thirdPartyIds: [],
    },
  });

  const onSubmit = (formData: FormData) => {
    createBusinessFunction({
      variables: {
        input: {
          organizationId,
          referenceId: formData.referenceId,
          name: formData.name,
          classification: formData.classification,
          mtdMinutes: formData.mtdMinutes,
          rtoMinutes: formData.rtoMinutes,
          rpoMinutes: formData.rpoMinutes,
          impactTolerance: formData.impactTolerance || undefined,
          notes: formData.notes || undefined,
          ownerId: formData.ownerId || undefined,
          assetIds: formData.assetIds,
          thirdPartyIds: formData.thirdPartyIds,
        },
        connections: connectionIds ?? [],
      },
      onCompleted() {
        toast({
          title: t("createBusinessFunctionDialog.messages.success"),
          description: t("createBusinessFunctionDialog.messages.created"),
          variant: "success",
        });
        reset();
        dialogRef.current?.close();
      },
      onError(error) {
        toast({
          title: t("createBusinessFunctionDialog.messages.error"),
          description: formatError(
            t("createBusinessFunctionDialog.errors.create"),
            error,
          ),
          variant: "error",
        });
      },
    });
  };

  return (
    <Dialog
      ref={dialogRef}
      trigger={children}
      title={(
        <Breadcrumb
          items={[
            t("createBusinessFunctionDialog.breadcrumb.businessFunctions"),
            t("createBusinessFunctionDialog.breadcrumb.create"),
          ]}
        />
      )}
      className="max-w-2xl"
    >
      <form onSubmit={e => void handleSubmit(onSubmit)(e)}>
        <DialogContent padded className="space-y-4">
          <div className="grid grid-cols-2 gap-4">
            <Field
              label={t("createBusinessFunctionDialog.fields.referenceId")}
              error={formState.errors.referenceId?.message}
              required
            >
              <Input
                {...register("referenceId")}
                placeholder={t("createBusinessFunctionDialog.fields.referenceIdPlaceholder")}
              />
            </Field>

            <Field
              label={t("createBusinessFunctionDialog.fields.name")}
              error={formState.errors.name?.message}
              required
            >
              <Input
                {...register("name")}
                placeholder={t("createBusinessFunctionDialog.fields.namePlaceholder")}
              />
            </Field>
          </div>

          <Controller
            control={control}
            name="classification"
            render={({ field }) => (
              <Field label={t("createBusinessFunctionDialog.fields.classification")} required>
                <Select
                  variant="editor"
                  placeholder={t("createBusinessFunctionDialog.fields.classificationPlaceholder")}
                  onValueChange={field.onChange}
                  value={field.value}
                  className="w-full"
                >
                  {classificationOptions.map(option => (
                    <Option key={option.value} value={option.value}>
                      {option.label}
                    </Option>
                  ))}
                </Select>
              </Field>
            )}
          />

          <div className="grid grid-cols-3 gap-4">
            <Field
              label={t("createBusinessFunctionDialog.fields.mtd")}
              error={formState.errors.mtdMinutes?.message}
              required
            >
              <Input
                {...register("mtdMinutes")}
                type="number"
                min={0}
                title={durationHelper}
              />
            </Field>

            <Field
              label={t("createBusinessFunctionDialog.fields.rto")}
              error={formState.errors.rtoMinutes?.message}
              required
            >
              <Input
                {...register("rtoMinutes")}
                type="number"
                min={0}
                title={durationHelper}
              />
            </Field>

            <Field
              label={t("createBusinessFunctionDialog.fields.rpo")}
              error={formState.errors.rpoMinutes?.message}
              required
            >
              <Input
                {...register("rpoMinutes")}
                type="number"
                min={0}
                title={durationHelper}
              />
            </Field>
          </div>

          <div className="space-y-2">
            <Label htmlFor="impactTolerance">
              {t("createBusinessFunctionDialog.fields.impactTolerance")}
            </Label>
            <Textarea
              id="impactTolerance"
              {...register("impactTolerance")}
              placeholder={t("createBusinessFunctionDialog.fields.impactTolerancePlaceholder")}
              rows={2}
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor="notes">{t("createBusinessFunctionDialog.fields.notes")}</Label>
            <Textarea
              id="notes"
              {...register("notes")}
              placeholder={t("createBusinessFunctionDialog.fields.notesPlaceholder")}
              rows={2}
            />
          </div>

          <PeopleSelectField
            organizationId={organizationId}
            control={control}
            name="ownerId"
            label={t("createBusinessFunctionDialog.fields.owner")}
            error={formState.errors.ownerId?.message}
            optional
          />

          <AssetsMultiSelectField
            organizationId={organizationId}
            control={control}
            name="assetIds"
            label={t("createBusinessFunctionDialog.fields.assets")}
          />

          <ThirdPartiesMultiSelectField
            organizationId={organizationId}
            control={control}
            name="thirdPartyIds"
            label={t("createBusinessFunctionDialog.fields.thirdParties")}
            level={null}
          />
        </DialogContent>

        <DialogFooter>
          <Button type="submit" disabled={formState.isSubmitting || isCreating}>
            {formState.isSubmitting || isCreating
              ? t("createBusinessFunctionDialog.actions.creating")
              : t("createBusinessFunctionDialog.actions.create")}
          </Button>
        </DialogFooter>
      </form>
    </Dialog>
  );
}
