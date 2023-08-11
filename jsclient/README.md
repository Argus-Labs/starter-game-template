This is a simple js client that demonstrates transaction receipts.

after run `mage start` in the parent directory (to start cardinal and nakama), run:

`
npm install
node main.mjs
`

The js client will attempt to join the singleton match and then will print any transaciton receipts that cardinal generates.

Transactions can be sent in via the Nakama API Explorer (localhost:7351/#/apiexplorer)

