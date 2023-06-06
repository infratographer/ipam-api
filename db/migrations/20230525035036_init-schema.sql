-- +goose Up
-- create "ip_block_types" table
CREATE TABLE "ip_block_types" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" text NOT NULL, "owner_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "ipblocktype_created_at" to table: "ip_block_types"
CREATE INDEX "ipblocktype_created_at" ON "ip_block_types" ("created_at");
-- create index "ipblocktype_owner_id" to table: "ip_block_types"
CREATE INDEX "ipblocktype_owner_id" ON "ip_block_types" ("owner_id");
-- create index "ipblocktype_updated_at" to table: "ip_block_types"
CREATE INDEX "ipblocktype_updated_at" ON "ip_block_types" ("updated_at");
-- create "ip_blocks" table
CREATE TABLE "ip_blocks" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "prefix" text NOT NULL, "location_id" character varying NOT NULL, "parent_block_id" character varying NOT NULL, "allow_auto_subnet" boolean NOT NULL DEFAULT true, "allow_auto_allocate" boolean NOT NULL DEFAULT true, "block_type_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "ip_blocks_ip_block_types_ip_block_type" FOREIGN KEY ("block_type_id") REFERENCES "ip_block_types" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "ipblock_block_type_id" to table: "ip_blocks"
CREATE INDEX "ipblock_block_type_id" ON "ip_blocks" ("block_type_id");
-- create index "ipblock_created_at" to table: "ip_blocks"
CREATE INDEX "ipblock_created_at" ON "ip_blocks" ("created_at");
-- create index "ipblock_updated_at" to table: "ip_blocks"
CREATE INDEX "ipblock_updated_at" ON "ip_blocks" ("updated_at");
-- create "ip_addresses" table
CREATE TABLE "ip_addresses" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "ip" text NOT NULL, "node_id" character varying NOT NULL, "node_owner_id" character varying NOT NULL, "reserved" boolean NOT NULL DEFAULT true, "block_id" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "ip_addresses_ip_blocks_ip_block" FOREIGN KEY ("block_id") REFERENCES "ip_blocks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "ipaddress_created_at" to table: "ip_addresses"
CREATE INDEX "ipaddress_created_at" ON "ip_addresses" ("created_at");
-- create index "ipaddress_node_owner_id_block_id_node_id" to table: "ip_addresses"
CREATE INDEX "ipaddress_node_owner_id_block_id_node_id" ON "ip_addresses" ("node_owner_id", "block_id", "node_id");
-- create index "ipaddress_updated_at" to table: "ip_addresses"
CREATE INDEX "ipaddress_updated_at" ON "ip_addresses" ("updated_at");

-- +goose Down
-- reverse: create index "ipaddress_updated_at" to table: "ip_addresses"
DROP INDEX "ipaddress_updated_at";
-- reverse: create index "ipaddress_node_owner_id_block_id_node_id" to table: "ip_addresses"
DROP INDEX "ipaddress_node_owner_id_block_id_node_id";
-- reverse: create index "ipaddress_created_at" to table: "ip_addresses"
DROP INDEX "ipaddress_created_at";
-- reverse: create "ip_addresses" table
DROP TABLE "ip_addresses";
-- reverse: create index "ipblock_updated_at" to table: "ip_blocks"
DROP INDEX "ipblock_updated_at";
-- reverse: create index "ipblock_created_at" to table: "ip_blocks"
DROP INDEX "ipblock_created_at";
-- reverse: create index "ipblock_block_type_id" to table: "ip_blocks"
DROP INDEX "ipblock_block_type_id";
-- reverse: create "ip_blocks" table
DROP TABLE "ip_blocks";
-- reverse: create index "ipblocktype_updated_at" to table: "ip_block_types"
DROP INDEX "ipblocktype_updated_at";
-- reverse: create index "ipblocktype_owner_id" to table: "ip_block_types"
DROP INDEX "ipblocktype_owner_id";
-- reverse: create index "ipblocktype_created_at" to table: "ip_block_types"
DROP INDEX "ipblocktype_created_at";
-- reverse: create "ip_block_types" table
DROP TABLE "ip_block_types";
