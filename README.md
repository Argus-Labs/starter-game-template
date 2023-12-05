# Starter Game Shard Template

This is a sample game shard built using Cardinal and [Nakama](https://heroiclabs.com/nakama/) as the account abstraction and
transaction relayer.

# Prerequisites

## World CLI

The [World CLI](https://github.com/Argus-Labs/world-cli) is a tool for creating, managing, and deploying World Engine projects. 

Install the latest world-cli release with:

```bash
curl https://install.world.dev/cli! | bash
```

## Docker Compose

Docker and docker compose are required for running Nakama and both can be installed with Docker Desktop.

[Installation instructions for Docker Desktop](https://docs.docker.com/compose/install/#scenario-one-install-docker-desktop)

# Cloning the Starter Game Template

To use the starter-game-template as a template for your own project, navigate to the directory where you want your project to live and run:

```bash
world cardinal create
```
You will be prompted for a game name. A copy of the starter-game-template will be created in the current directory.

# Running the Server

Navigate to thew newly created project and run:

```bash
world cardinal start
```

This command will use the `world.toml` config specified in your root project directory to run the following containers:
- Cardinal
- Redis
- Nakama
- Nakam's DB

To stop the containers, run:

```bash
world cardinal stop
```

# Interacting with Cardinal

## Via the Cardinal Editor

The Cardinal Editor is a web-based companion app that makes game development of Cardinal easier. It allows you to inspect the state of Cardinal in real-time without any additional code.

Then, open the [Cardinal Editor](https://editor.world.dev) in a web browser.

To start, there will be no data stored in Cardinal. As you interact with Cardinal (e.g. by creating a Persona Tag via Nakama), your new Cardinal state will show up in the Cardinal Editor.

## Via Nakama

With the containers running visit `localhost:7351` in a web browser to access the Nakama console. For local development, use `admin:password` as your login credentials.

The Account tab on the left will give you access to a valid account ID.

The API Explorer tab on the left will allow you to make requests to Cardinal.

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

And hit `Submit` no User ID is required for this endpoint.

To get the User ID of your newly created account, click the `Accounts` item in the sidebar. Copy the relevant User ID and paste it into the User ID field on the API Explorer to hit other endpoints.

### Claiming a Persona Tag

A persona tag is essentially a cardinal based user. To create a persona tag in your cardinal game, select the `nakama/claim-persona` 
endpoint from the dropdown. Make sure to paste in a valid User ID into the User ID field. Set the request body to:
```json
{
  "personaTag": "some-persona-tag"
}
```

and hit Submit. You should see a response like:

```json
{
  "personaTag": "some-persona-tag",
  "status": "pending",
  "tick": 2567,
  "txHash": "0x6bc26694dee4c4163335e4fe01d73eab2da071f38b991ae8424fa52de330c228"
}
```

This means cardinal received the request, and the request is pending. Change your endpoint to `nakama/show-persona` and hit
Submit (no request body needed) to verify the claim-persona operation was successful. The response body should be similar
to the `nakama/claim-persona` response, except "status" should now say "accepted".

This mean both Nakama and Cardinal are aware of your Nakama user and the related Persona Tag.

### Custom Messages

Once your persona tag has been set up, you can send messages to your custom cardinal message endpoints. If
you set up a message like this:

```go
package main

import (
	"pkg.world.dev/world-engine/cardinal"
)

type MoveMsg struct {
	Dx, Dy int
}

type MoveReply struct {
	FinalX, FinalY int
}

var MoveTx = cardinal.NewMessageType[MoveMsg, MoveReply]("move")

func main() {
	world, err := cardinal.NewWorld(cardinal.WithDisableSignatureVerification())
	if err != nil {
		panic(err)
	}
	// ...
	world.RegisterMessages(MoveTx)
	// ...
}
```

The dropdown will contain an entry with `tx/game/move`. The request body for that message could be:
```json
{
	"Dx": 100,
	"Dy": 200,
}
```

Hit submit, and the message will be sent to your cardinal implementation. What your game does with the message depends on the Systems you've defined.

# Your World

The code in <your-project-name>/cardinal/... powers your cardinal project. The entry point of your game is main.go, and that files contains sample code for setting up Components, Messages, Queries, and Systems. 

[Check out the official World Engine documentation](https://world.dev/Cardinal/API-Reference/Components) for more details on how to build your World!