# Starter Game Shard Template

This is a sample game shard built using Cardinal and [Nakama](https://heroiclabs.com/nakama/) as the account abstraction and
transaction relayer.

# Prerequisites

## Mage Check

A mage target exists that will check for some common pre-requisites. Run the check with:

```bash
mage check
```

## Docker Compose

Docker and docker compose are required for running Nakama and both can be installed with Docker Desktop.

[Installation instructions for Docker Desktop](https://docs.docker.com/compose/install/#scenario-one-install-docker-desktop)

## Mage

[Mage](https://magefile.org/) is a cross-platform Make-like build tool.

```bash
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```

# Running the Server

To start Nakama and Cardinal:

```bash
mage start
```

To start ONLY Cardinal in dev mode (compatible with the Retool dashboard):

```bash
mage dev
```

To restart ONLY Cardinal:

```bash
mage restart
```

To stop Nakama and Cardinal:

```bash
mage stop
```

Alternatively, killing the `mage start` process will also stop Nakama and Cardinal

Note, for now, if any Cardinal endpoints have been added or removed Nakama must be relaunched (via `mage stop` and `mage start`).
We will add a future to hot reload this in the future.

# Verify the Server is Running

Visit `localhost:7351` in a web browser to access Nakama. For local development, use `admin:password` as your login
credentials.

The Account tab on the left will give you access to a valid account ID.

The API Explorer tab on the left will allow you to make requests to Cardinal.

# Cardinal Editor

The Cardinal Editor is a web-based companion app that makes game development of Cardinal easier. It allows you to inspect the state of Cardinal in real-time without any additional code.

To work with the Cardinal Editor, you must first start the Cardinal server in dev mode:

```bash
mage dev
```

Then, open the [Cardinal Editor](https://editor.world.dev) in a web browser.