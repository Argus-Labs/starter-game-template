# Starter Game Shard Template

This is a sample game shard built using Cardinal and [Nakama](https://heroiclabs.com/nakama/) as the account abstraction and
transaction relayer.

# Prerequisites

## Mage Check

A mage target exists that will check for some common preqrequisites. Run the check with:

```bash
mage check
```

## Github Access

The game code (under the `cardinal` directory) runs on top of Cardinal, which is a library in the
(currently) private [World Engine](https://github.com/Argus-Labs/world-engine) repo.

You likely have access to this repo, but the `go` binary sometimes has trouble accessing private repos on your behalf.

### GOPATH

[Configure Go to access private modules](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project#configuring-go-to-access-private-modules)

TL;DR: Add 'GOPRIVATE="github.com/argus-labs/world-engine"' to your environment variables.

### Github Credentials

In addition, configure git to use your private credentials via HTTPS or SSH:

[Providing Private Module Credentials for HTTPS](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project#providing-private-module-credentials-for-https)
OR
[Providing Private Module Credentials for SSH](https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project#providing-private-module-credentials-for-ssh)

TODO: It would be helpful if this section included the error message that a user could expect to see when their git
credentials are incorrect.

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

# Copy the Starter Game Template

You can make a full copy of this template with:

```bash
mage copy <target-directory> <module-path>
```

The <target-directory> is where you want your code to live on your local machine and the <module-path> parameter is the
repo location of your new game.

See the [go mod init](https://golang.org/ref/mod#go-mod-init) documentation for more details about the module path
parameter.
