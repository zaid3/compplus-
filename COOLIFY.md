# Deploy CompPlus on Coolify

This repository includes a source-building production stack for Coolify:

- `Dockerfile.coolify` builds the imported application source, frontends, generated APIs and Go binaries.
- `compose.coolify.yaml` runs CompPlus, PostgreSQL, SeaweedFS object storage and headless Chrome.
- `.env.coolify.example` lists the runtime variables to add in Coolify.

## Coolify resource settings

1. In Coolify, open your project and choose **New Resource**.
2. Select **Public Repository** while this GitHub repository is public. If you later make it private, connect it using the Coolify GitHub App or a deploy key.
3. Repository URL: `https://github.com/zaid3/compplus-`
4. Branch: `adding-ferpa`
5. Build Pack: **Docker Compose**
6. Base Directory: `/`
7. Docker Compose Location: `/compose.coolify.yaml`
8. Keep **Git Submodules** enabled. This imported repository still references an upstream data submodule.
9. Continue without using Raw Compose mode.

## Environment variables

Open the resource's **Environment Variables** section and add every variable from `.env.coolify.example`.

Generate five different secure values locally:

```bash
openssl rand -base64 32
```

Use a separate generated value for:

- `PROBOD_ENCRYPTION_KEY`
- `AUTH_COOKIE_SECRET`
- `AUTH_PASSWORD_PEPPER`
- `TRUST_AUTH_TOKEN_SECRET`
- `POSTGRES_PASSWORD`

Do not mark runtime API keys or passwords as build variables. They are only needed when the containers start.

## Domain

After Coolify parses the Compose stack:

1. Open the `compplus` service.
2. Assign your HTTPS domain to container port `8080`.
3. Set the same full URL in `PROBOD_BASE_URL` and `API_CORS_ALLOWED_ORIGINS`.
4. Set only the hostname, without `https://`, in `AUTH_COOKIE_DOMAIN`.

Example:

```text
PROBOD_BASE_URL=https://compliance.example.com
API_CORS_ALLOWED_ORIGINS=https://compliance.example.com
AUTH_COOKIE_DOMAIN=compliance.example.com
```

## First deployment

Deploy the resource and watch the build logs. The initial source build is large because it installs the Node and Go toolchains, builds both frontends, generates the GraphQL/MCP code and compiles the server.

When the application is healthy, visit the assigned domain and create the first administrator account. Then change:

```text
AUTH_DISABLE_SIGNUP=true
```

Redeploy to stop public account registration.

## Persistent data

The Compose stack defines persistent Docker volumes for:

- application data
- PostgreSQL
- SeaweedFS object storage

Do not delete these volumes during routine redeployments. Back up PostgreSQL and object storage before application upgrades.

## Security

Only expose the `compplus` service through Coolify's proxy. PostgreSQL, SeaweedFS and Chrome intentionally have no public port mappings.

Before holding real compliance data:

- make the GitHub repository private;
- rename it from `compplus-` to `compplus`;
- configure working SMTP;
- take scheduled database and object-storage backups;
- test restore procedures;
- review the older imported Probo version and plan an upgrade.
