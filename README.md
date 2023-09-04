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

# Verify Nakama is Running

Visit `localhost:7351` in a web browser to access the Nakama console. For local development, use `admin:password` as your login
credentials.

The Account tab on the left will give you access to a valid account ID.

The API Explorer tab on the left will allow you to make requests to Cardinal.

# Nakama Console

You can verify the Nakama server is running by visiting `localhost:7351` in a web browser. For local development, use 
`admin:password` as our login credentials.

## API Explorer

The API Explorer (on the sidebar) allows you to make requests to both Nakama and your Cardinal server.

### Creating a User ID

Before using any endpoints, you need to populate the User ID field (between the endpoint dropdown and the submit button). 
The user ID `00000000-0000-0000-0000-000000000000` is a special admin user ID that will always be defined. Alternatively, 
a new user can be created by selecting the `Authenticate Device` endpoint from the dropdown. Populate the request boy
with a payload like:

```json
{
  "account": {
    "id": "123456789123456789"
  },
  "create": true,
  "username": "some-username"
}
```

And hit `Submit` no. User ID is required for this endpoint.

To get the User ID of your newly created account, click the `Accounts` item in the sidebar. Copy the relevant User ID 
and paste it into the User ID field on the API Explorer to hit other endpoints.

### Claiming a Persona Tag

A persona tag is essentially a cardinal based user. To create a persona tag in your cardinal game, select the `nakama/claim-persona` 
endpoint from the dropdown. Make sure to paste in a valid User ID into the User ID field. Set the request body to:
```json
{
  "persona_tag": "some-persona-tag",
}
```

and hit Submit. You should see a response like:

```json
{
  "persona_tag": "some-persona-tag",
  "status": "pending",
  "tick": 2567,
  "tx_hash": "0x6bc26694dee4c4163335e4fe01d73eab2da071f38b991ae8424fa52de330c228"
}
```

This means cardinal received the request, and the request is pending. Change your endpoint to `nakama/show-persona` and hit
Submit (no request body needed) to verify the claim-persona operation was successful. The response body should be similar
to the `nakama/claim-persona` response, except "status" should now say "accepted".

This mean both Nakama and Cardinal are aware of your Nakama user and the related Persona Tag.

### Custom Transactions

Once your persona tag has been set up, you send transactions to your custom cardinal transaction endpoints. If
you set up a transaction like this:

```go
package main

import (
	"pkg.world.dev/world-engine/cardinal/ecs"
	"pkg.world.dev/world-engine/cardinal/ecs/inmem"
)

type MoveMsg struct {
	Dx, Dy int
}

type MoveReply struct {
	FinalX, FinalY int
}

var MoveTx = ecs.NewTransactionType[MoveMsg, MoveReply]("move")

func main() {
	world := inmem.NewECSWorld()
	world.RegisterTransaction(MoveTx)
}
```

The dropdown will contain an entry with `tx-move`. The request body for that transaction could be:
```json
{
	"Dx": 100,
	"Dy": 200,
}
```

Hit submit, and the transaction will be sent to your cardinal implementation. What your game does with the transaction
depends on what Systems you've defined.

## Storage

The `Storage` item in the sidebar allows you to view Nakama user data. Storage objects with a key name of `persona_tag`
describe which persona tag has been associated with user ID.

# Cardinal Editor

The Cardinal Editor is a web-based companion app that makes game development of Cardinal easier. It allows you to inspect the state of Cardinal in real-time without any additional code.

To work with the Cardinal Editor, you must first start the Cardinal server in dev mode:

```bash
mage dev
```

Then, open the [Cardinal Editor](https://editor.world.dev) in a web browser.
