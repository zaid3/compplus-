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

import { usePageTitle } from "@probo/hooks";
import { Button, Field, useToast } from "@probo/ui";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import { useMutation, usePreloadedQuery, useQueryLoader } from "react-relay";
import { Link, useNavigate } from "react-router";
import { graphql } from "relay-runtime";
import { z } from "zod";

import type { SignUpPageMutation } from "#/__generated__/iam/SignUpPageMutation.graphql";
import type { SignUpPageQuery } from "#/__generated__/iam/SignUpPageQuery.graphql";
import { useFormWithSchema } from "#/hooks/useFormWithSchema";

const signUpPageQuery = graphql`
  query SignUpPageQuery {
    signUpEnabled
  }
`;

const signUpMutation = graphql`
  mutation SignUpPageMutation($input: SignUpInput!) {
    signUp(input: $input) {
      identity {
        id
      }
    }
  }
`;

const schema = z.object({
  email: z.string().email(),
  password: z.string().min(12),
  fullName: z.string().min(2),
});

const publicSignUpError =
  "We couldn't create this account. This email may already be registered. Try signing in or reset your password. Otherwise, try again shortly.";

type FormData = z.infer<typeof schema>;

function SignUpPageContent(props: { queryRef: NonNullable<ReturnType<typeof useQueryLoader<SignUpPageQuery>>[0]> }) {
  const { t } = useTranslation();
  const { toast } = useToast();
  const navigate = useNavigate();

  usePageTitle(t("signUpPage.pageTitle"));

  const data = usePreloadedQuery<SignUpPageQuery>(signUpPageQuery, props.queryRef);

  const { register, handleSubmit, formState } = useFormWithSchema(schema, {
    defaultValues: { email: "", password: "", fullName: "" },
  });

  const [signUp, isSigningUp] = useMutation<SignUpPageMutation>(signUpMutation);

  const onSubmit = (form: FormData) => {
    if (isSigningUp) return;

    signUp({
      variables: {
        input: {
          email: form.email,
          password: form.password,
          fullName: form.fullName,
        },
      },
      onCompleted: (_, errors) => {
        if (errors) {
          // Keep account-existence details private while still giving the user
          // an actionable path when they may already have an account.
          toast({
            title: "Unable to create account",
            description: publicSignUpError,
            variant: "error",
          });
          return;
        }

        toast({
          title: t("common.success"),
          description: t("signUpPage.messages.created"),
          variant: "success",
        });

        const search = new URLSearchParams({ email: form.email }).toString();
        void navigate(`/auth/resend-verification-email?${search}`, { replace: true });
      },
      onError: () => {
        toast({
          title: "Unable to create account",
          description: publicSignUpError,
          variant: "error",
        });
      },
    });
  };

  if (!data.signUpEnabled) {
    return (
      <div className="space-y-6 w-full max-w-md mx-auto pt-8 text-center">
        <div className="space-y-2">
          <h1 className="text-3xl font-bold">{t("signUpPage.unavailable.title")}</h1>
          <p className="text-txt-tertiary">
            {t("signUpPage.unavailable.description")}
          </p>
        </div>
        <Button variant="secondary" className="w-xs h-10 mx-auto" to="/auth/login">
          {t("signUpPage.actions.backToLogin")}
        </Button>
      </div>
    );
  }

  return (
    <div className="space-y-6 w-full max-w-md mx-auto pt-8">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold">{t("signUpPage.title")}</h1>
        <p className="text-txt-tertiary">{t("signUpPage.description")}</p>
      </div>

      <form onSubmit={e => void handleSubmit(onSubmit)(e)} className="space-y-4">
        <Field
          label={t("signUpPage.fields.fullName")}
          type="text"
          autoComplete="name"
          placeholder={t("signUpPage.fields.fullNamePlaceholder")}
          {...register("fullName")}
          required
          error={formState.errors.fullName?.message}
        />

        <Field
          label={t("signUpPage.fields.email")}
          type="email"
          autoComplete="email"
          autoCapitalize="none"
          spellCheck={false}
          placeholder={t("signUpPage.fields.emailPlaceholder")}
          {...register("email")}
          required
          error={formState.errors.email?.message}
        />

        <Field
          label={t("signUpPage.fields.password")}
          type="password"
          autoComplete="new-password"
          placeholder="At least 12 characters"
          {...register("password")}
          required
          error={formState.errors.password?.message}
        />

        <Button type="submit" className="w-xs h-10 mx-auto mt-6" disabled={formState.isLoading || isSigningUp}>
          {isSigningUp
            ? t("signUpPage.actions.creating")
            : t("signUpPage.actions.signUpWithEmail")}
        </Button>
      </form>

      <div className="space-y-2 text-center text-sm text-txt-tertiary">
        <p>
          {t("signUpPage.alreadyHaveAccount")}
          {" "}
          <Link to="/auth/login" className="underline text-txt-primary hover:text-txt-secondary">
            {t("signUpPage.actions.logIn")}
          </Link>
        </p>
        <p>
          Forgot your password?{" "}
          <Link to="/auth/forgot-password" className="underline text-txt-primary hover:text-txt-secondary">
            Reset it here
          </Link>
        </p>
      </div>
    </div>
  );
}

export default function SignUpPage() {
  const [queryRef, loadQuery] = useQueryLoader<SignUpPageQuery>(signUpPageQuery);

  useEffect(() => {
    loadQuery({});
  }, [loadQuery]);

  if (!queryRef) return null;

  return <SignUpPageContent queryRef={queryRef} />;
}
