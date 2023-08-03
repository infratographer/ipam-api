# Infratographer IP Address Management API

Infratographer IP Address Management implements a GraphQL API that provides a way to manage IP Addresses

## IPAM Structure

```mermaid
erDiagram
    NODE ||--o| IPBlockType : "has"
    IPBlockType {
        string Name
        id OwnerID
    }

    IPBlockType ||--o{ IPBlock : "has"
    IPBlock {
        string Prefix
        id BlockTypeID
        id LocationID
        id ParentBlockID
        boolean AllowAutoSubnet
        boolean AllowAutoAllocate
    }

    IPBlock ||--o{ IPAddress : "has"
    IPAddress {
        string IP
        id BlockID
        id NodeID
        id NodeOwnerID
        boolean Reserved
    }
```


## Development and Contributing

- [Development Guide](docs/development.md)
- [Contributing](https://infratographer.com/community/contributing/)

### Getting Started

To begin start by opening the devcontainer as outlined in the [Development Guide](docs/development.md)

**To initialize the database:**
1. `go build`
1. `./ipam-api migrate up`

**To run the api**
```sh
make go-run
```

**To subscribe to the NATS stream:**
```sh
# if not created already
nats --server=nats:4222 --creds=/nsc/nkeys/creds/LOCAL/IPAM/USER.creds stream add $NAME --subjects='com.example.>'

nats --server=nats:4222 --creds=/nsc/nkeys/creds/LOCAL/IPAM/USER.creds sub --stream=$NAME
```

**Interacting with the GraphQL queries:**
> Go to localhost:7909/playground

## Example GraphQL Queries

### Create IP Block Mutation

Input:
```graphql
mutation{
  createIPBlockType(
    input: {
        name:"super-sweet-ip-block-type",
        ownerID:"tenants-df234a22-f849-11ed-b67e-0242ac120002"
    }
  )
  {
    ip_block_type{
      name,
      id
    }
  }
}
```

Output:
```json
{
  "data": {
    "createIPBlockType": {
      "ip_block_type": {
        "name": "super-sweet-ip-block-type",
        "id": "ipamibt-9xaBQDAFmLOdceu9zO6Rj"
      }
    }
  }
}
```

### Get IP Block by ID

Input:
```graphql
query{
  ip_block_type(id:"ipamibt-9xaBQDAFmLOdceu9zO6Rj"){
    name,
    id
  }
}
```

Output:
```json
{
  "data": {
    "ip_block_type": {
      "name": "super-sweet-ip-block-type",
      "id": "ipamibt-9xaBQDAFmLOdceu9zO6Rj"
    }
  }
}
```


## Code of Conduct

[Contributor Code of Conduct](https://infratographer.com/community/code-of-conduct/). By participating in this project you agree to abide by its terms.

## Contact

To contact the maintainers, please open a [GithHub Issue](https://github.com/infratographer/ipam-api/issues/new)

## License

[Apache 2.0](LICENSE)
