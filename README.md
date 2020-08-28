# go-surfspot-api

A simple REST API built using No third-party packages/dependencies

## Requirements

* [x] `GET /surfspots` returns list of surfspots as JSON
* [x] `GET /surfspots/{id}` returns details of specific surfspot as JSON
* [x] `GET /admin` requires basic auth
* [x] `GET /surfspots/random` redirects (Status 302) to a random surfspot
* [x] `POST /surfspots` accepts a new surfspot to be added
* [x] `POST /surfspots` returns status 415 if content is not `application/json`
* []  `DELETE /surfspots/{id}` returns the deleted surfspot as JSON
* []  `PUT /surfspots/{id}` returns the updated surfspot as JSON

### Data Types

A surfspot object should look like this:
```json
{
  "id": "someid",
  "name": "name of the surfspot",
  "founder": "the name of the person who found the sufspotspot",
  "beach": "name of the beach the surfspot is located at",
  "difficulty": 5, 
}
```

### Persistence

There is no persistence, just temporary in-memory storage.
