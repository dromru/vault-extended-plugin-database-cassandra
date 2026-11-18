package main

import (
	"github.com/hashicorp/vault/sdk/database/dbplugin/v5"

	cassandra "git.drom.ru/security/vault-extended-plugin-database-cassandra"
)

func main() {
	Run()
}

// Run starts the RPC server for the plugin.
func Run() {
	dbplugin.ServeMultiplex(cassandra.New)
}
