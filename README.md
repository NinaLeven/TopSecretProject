# TopSecretProject

According to the requirements: 
* Service can create, update, get, list projects
* Participants may be added via create or update methods
* Owner-participants department validation is implemented
* Owner manager status validation is implemented
* Service is stateless: uses postgresql for storage
* Migration script is provided
* API is provided in Openapi v3 format
* Application may be run natively or in docker

In the case that the requrements are such, that ope participant must participate in no more than one project: one should modify storage scheme from 
`participant_ids text[]` to a new table binding participant and project, with unique index constaint on (project_id, participant_id). Also an additional check in the business-logic layer must be implemented along with corresponding methods in the storage layer. 

### Packages:
* [api](api/projectmanager.yaml)
* [server](internal/api)
* [logic](internal/projectmanager)
* [storage](internal/integrations/storage)
* [employees](internal/integrations/employee)

## Quickstart 

Unfortunately I had problems actually running the application with provided docker compose.
Need more time to work out the problem with networks.

`docker-compose up`

## Longstart

This is the way i was able to virify anything works

```sh
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:12
cd cmd/projectmanager 
go build
./projectmanager -c ../../configs/local.yaml
```

## Examples

```sh
curl -XPOST localhost:8080/api/projects -d '{"uid":"1234", "name": "Name!", "owner_id": "eba96253-5ff4-48e4-86d5-7197bcc7c349"}'
curl -XPATCH localhost:8080/api/projects/1234 -d '{"participant_ids": ["1088511c-18e0-4bb6-861b-8112de23be97"]}'
curl localhost:8080/api/projects
```
