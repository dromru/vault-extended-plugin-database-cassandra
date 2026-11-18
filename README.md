# Vault database secrets engine - Extended Cassandra plugin

This plugin is a modification of the [official Vault plugin](https://github.com/hashicorp/vault/tree/main/plugins/database/cassandra).

## Changes

Gocql package updated to v2

## Usage

Build the plugin binary:

```
go build -o vault-extended-plugin-database-cassandra ./cmd/vault-extended-plugin-database-cassandra
```

Register it in Vault:

```
vault plugin install -address=... -sha256=... vault-extended-plugin-database-cassandra

vault secrets enable -path=database -plugin-name=cassandra-extended-database-plugin database

vault write database/config/cassandra plugin_name=vault-extended-plugin-database-cassandra \
  hosts="<host>" username=<username> password=<password>
```
