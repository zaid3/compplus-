// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

import {
  Breadcrumb,
  Button,
  Dialog,
  DialogContent,
  DialogFooter,
  type DialogRef,
  Input,
  Textarea,
  useDialogRef,
} from "@probo/ui";
import { useState } from "react";
import { graphql } from "relay-runtime";
import { z } from "zod";

import type { CorePackSetupDialogInstallMutation } from "#/__generated__/core/CorePackSetupDialogInstallMutation.graphql";
import type { CorePackSetupDialogPreviewMutation } from "#/__generated__/core/CorePackSetupDialogPreviewMutation.graphql";
import { useFormWithSchema } from "#/hooks/useFormWithSchema";
import { useMutationWithToasts } from "#/hooks/useMutationWithToasts";

const previewTemplatePackMutation = graphql`
  mutation CorePackSetupDialogPreviewMutation(
    $input: PreviewTemplatePackInput!
  ) {
    previewTemplatePack(input: $input) {
      preview {
        packId
        version
        frameworkName
        controlsCount
        measuresCount
        documentsCount
        tasksCount
        evidenceRequestsCount
        confirmationFieldsCount
      }
    }
  }
`;

const installTemplatePackMutation = graphql`
  mutation CorePackSetupDialogInstallMutation(
    $input: InstallTemplatePackInput!
  ) {
    installTemplatePack(input: $input) {
      packId
      version
      measuresCreated
      documentsCreated
      tasksCreated
      evidenceRequestsCreated
      framework {
        id
        name
      }
    }
  }
`;

const schema = z.object({
  legalName: z.string().trim().min(1).max(255),
  address: z.string().trim().min(1).max(1000),
  services: z.string().trim().min(1).max(2000),
  locations: z.string().trim().min(1).max(1000),
  employeeCount: z.coerce.number().int().min(1).max(1_000_000),
  executiveOwner: z.string().trim().min(1).max(255),
  systemManager: z.string().trim().min(1).max(255),
  certificationTarget: z.string().optional(),
  reviewMonth: z.string().min(1),
  usesSuppliers: z.boolean(),
  usesAI: z.boolean(),
  processesPersonalData: z.boolean(),
  environmentalImpacts: z.boolean(),
});

type FormData = z.infer<typeof schema>;

type Preview = {
  packId: string;
  version: string;
  frameworkName: string;
  controlsCount: number;
  measuresCount: number;
  documentsCount: number;
  tasksCount: number;
  evidenceRequestsCount: number;
  confirmationFieldsCount: number;
};

type Props = {
  organizationId: string;
  organizationName: string;
  ref?: DialogRef;
};

const months = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

const selectedStandards = [
  "ISO/IEC 27001",
  "ISO 9001",
  "ISO 14001",
  "ISO/IEC 42001",
  "UK GDPR",
];

export function CorePackSetupDialog(props: Props) {
  const fallbackRef = useDialogRef();
  const dialogRef = props.ref ?? fallbackRef;
  const [preview, setPreview] = useState<Preview | null>(null);
  const {
    register,
    handleSubmit,
    getValues,
    formState: { errors },
  } = useFormWithSchema(schema, {
    defaultValues: {
      legalName: props.organizationName,
      address: "",
      services: "",
      locations: "",
      employeeCount: 1,
      executiveOwner: "",
      systemManager: "",
      certificationTarget: "",
      reviewMonth: "January",
      usesSuppliers: true,
      usesAI: false,
      processesPersonalData: true,
      environmentalImpacts: true,
    },
  });

  const [preparePreview, isPreparingPreview] = useMutationWithToasts<CorePackSetupDialogPreviewMutation>(
    previewTemplatePackMutation,
    {
      successMessage: "Your Core Pack preview is ready.",
      errorMessage: "The Core Pack preview could not be prepared.",
    },
  );
  const [installPack, isInstalling] = useMutationWithToasts<CorePackSetupDialogInstallMutation>(
    installTemplatePackMutation,
    {
      successMessage: "The CompPlus Core Pack has been installed.",
      errorMessage: "The CompPlus Core Pack could not be installed.",
    },
  );

  const buildInput = (data: FormData) => ({
    organizationId: props.organizationId,
    packId: "core",
    answers: {
      legalName: data.legalName,
      tradingName: null,
      address: data.address,
      services: data.services,
      locations: data.locations,
      employeeCount: data.employeeCount,
      executiveOwner: data.executiveOwner,
      systemManager: data.systemManager,
      securityOwner: data.systemManager,
      privacyOwner: data.systemManager,
      qualityOwner: data.systemManager,
      environmentalOwner: data.systemManager,
      aiOwner: data.systemManager,
      certificationTarget: data.certificationTarget || null,
      reviewMonth: data.reviewMonth,
      usesSuppliers: data.usesSuppliers,
      usesAI: data.usesAI,
      processesPersonalData: data.processesPersonalData,
      environmentalImpacts: data.environmentalImpacts,
      selectedStandards,
    },
  });

  const onPreview = async (data: FormData) => {
    await preparePreview({
      variables: {
        input: buildInput(data),
      },
      onCompleted(response) {
        setPreview(response.previewTemplatePack.preview);
      },
    });
  };

  const onInstall = async () => {
    const data = schema.parse(getValues());
    await installPack({
      variables: {
        input: buildInput(data),
      },
      onSuccess() {
        dialogRef.current?.close();
        window.location.reload();
      },
    });
  };

  return (
    <Dialog
      ref={dialogRef}
      title={(
        <Breadcrumb
          items={[
            "Frameworks",
            preview ? "Review Core Pack" : "CompPlus Fast Start",
          ]}
        />
      )}
    >
      {preview
        ? (
            <div>
              <DialogContent padded className="space-y-5">
                <div>
                  <h2 className="text-lg font-semibold text-txt-primary">
                    Ready to create your compliance workspace
                  </h2>
                  <p className="mt-1 text-sm text-txt-secondary">
                    CompPlus will prepare the shared management-system work once and reuse it across your selected standards.
                  </p>
                </div>

                <div className="grid grid-cols-2 gap-3 sm:grid-cols-3">
                  <PreviewCard label="Controls" value={preview.controlsCount} />
                  <PreviewCard label="Documents" value={preview.documentsCount} />
                  <PreviewCard label="Measures" value={preview.measuresCount} />
                  <PreviewCard label="Tasks" value={preview.tasksCount} />
                  <PreviewCard label="Evidence requests" value={preview.evidenceRequestsCount} />
                  <PreviewCard label="Items to confirm" value={preview.confirmationFieldsCount} />
                </div>

                <div className="rounded-lg border border-border-low bg-subtle p-4 text-sm">
                  <p className="font-medium text-txt-primary">What happens next</p>
                  <p className="mt-1 text-txt-secondary">
                    The documents will already contain your company details. You only review the clearly marked confirmation fields, attach evidence and approve the documents when ready.
                  </p>
                </div>

                <div className="text-xs text-txt-tertiary">
                  Pack version: {preview.version}
                </div>
              </DialogContent>
              <DialogFooter>
                <Button
                  type="button"
                  variant="secondary"
                  onClick={() => setPreview(null)}
                  disabled={isInstalling}
                >
                  Back to details
                </Button>
                <Button
                  type="button"
                  onClick={() => void onInstall()}
                  disabled={isInstalling}
                >
                  {isInstalling ? "Installing…" : "Install Core Pack"}
                </Button>
              </DialogFooter>
            </div>
          )
        : (
            <form onSubmit={event => void handleSubmit(onPreview)(event)}>
              <DialogContent padded className="max-h-[70vh] space-y-6 overflow-y-auto">
                <div>
                  <h2 className="text-lg font-semibold text-txt-primary">
                    Tell us about the organisation once
                  </h2>
                  <p className="mt-1 text-sm text-txt-secondary">
                    CompPlus reuses these answers across every policy, register, task and evidence request.
                  </p>
                </div>

                <section className="space-y-4">
                  <h3 className="text-sm font-semibold text-txt-primary">Organisation</h3>
                  <Field label="Legal company name" error={errors.legalName?.message}>
                    <Input {...register("legalName")} required />
                  </Field>
                  <Field label="Main business address" error={errors.address?.message}>
                    <Textarea {...register("address")} autogrow required />
                  </Field>
                  <Field label="Products and services" error={errors.services?.message}>
                    <Textarea
                      {...register("services")}
                      autogrow
                      required
                      placeholder="For example: education and student-support services"
                    />
                  </Field>
                  <Field label="Locations covered" error={errors.locations?.message}>
                    <Textarea
                      {...register("locations")}
                      autogrow
                      required
                      placeholder="For example: London office and approved remote-working locations"
                    />
                  </Field>
                  <Field label="Number of workers" error={errors.employeeCount?.message}>
                    <Input
                      {...register("employeeCount")}
                      type="number"
                      min={1}
                      required
                    />
                  </Field>
                </section>

                <section className="space-y-4">
                  <h3 className="text-sm font-semibold text-txt-primary">Responsibility</h3>
                  <Field
                    label="Executive owner"
                    help="The senior person who approves the management system."
                    error={errors.executiveOwner?.message}
                  >
                    <Input {...register("executiveOwner")} required />
                  </Field>
                  <Field
                    label="Compliance or management-system manager"
                    help="CompPlus will use this person as the initial owner for security, privacy, quality, environment and AI work. These roles can be changed later."
                    error={errors.systemManager?.message}
                  >
                    <Input {...register("systemManager")} required />
                  </Field>
                  <Field
                    label="Target certification or assessment date"
                    help="Optional. It can be changed later."
                    error={errors.certificationTarget?.message}
                  >
                    <Input {...register("certificationTarget")} type="date" />
                  </Field>
                  <Field label="Preferred annual review month" error={errors.reviewMonth?.message}>
                    <select
                      {...register("reviewMonth")}
                      className="h-10 w-full rounded-md border border-border-low bg-level-0 px-3 text-sm text-txt-primary"
                    >
                      {months.map(month => (
                        <option key={month} value={month}>{month}</option>
                      ))}
                    </select>
                  </Field>
                </section>

                <section className="space-y-3">
                  <h3 className="text-sm font-semibold text-txt-primary">Quick checks</h3>
                  <BooleanField
                    label="Suppliers process data or support important services"
                    registration={register("usesSuppliers")}
                  />
                  <BooleanField
                    label="The organisation develops or uses AI"
                    registration={register("usesAI")}
                  />
                  <BooleanField
                    label="The organisation processes personal data"
                    registration={register("processesPersonalData")}
                  />
                  <BooleanField
                    label="The organisation has environmental impacts such as energy, travel, waste or purchased goods"
                    registration={register("environmentalImpacts")}
                  />
                </section>

                <div className="rounded-lg border border-border-low bg-subtle p-4 text-sm text-txt-secondary">
                  Selected baseline: ISO/IEC 27001, ISO 9001, ISO 14001, ISO/IEC 42001 and UK GDPR. The separate detailed packs will be added after the shared Core Pack.
                </div>
              </DialogContent>
              <DialogFooter>
                <Button type="submit" disabled={isPreparingPreview}>
                  {isPreparingPreview ? "Preparing…" : "Review what will be created"}
                </Button>
              </DialogFooter>
            </form>
          )}
    </Dialog>
  );
}

function Field(props: {
  label: string;
  help?: string;
  error?: string;
  children: React.ReactNode;
}) {
  return (
    <label className="block space-y-1.5">
      <span className="text-sm font-medium text-txt-primary">{props.label}</span>
      {props.children}
      {props.help && <span className="block text-xs text-txt-tertiary">{props.help}</span>}
      {props.error && <span className="block text-xs text-danger">{props.error}</span>}
    </label>
  );
}

function BooleanField(props: {
  label: string;
  registration: React.InputHTMLAttributes<HTMLInputElement>;
}) {
  return (
    <label className="flex items-start gap-3 rounded-lg border border-border-low p-3">
      <input
        {...props.registration}
        type="checkbox"
        className="mt-0.5 size-4 rounded border-border-strong"
      />
      <span className="text-sm text-txt-primary">{props.label}</span>
    </label>
  );
}

function PreviewCard(props: { label: string; value: number }) {
  return (
    <div className="rounded-lg border border-border-low p-3">
      <div className="text-2xl font-semibold text-txt-primary">{props.value}</div>
      <div className="text-xs text-txt-secondary">{props.label}</div>
    </div>
  );
}
