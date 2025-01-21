## Risk Service Review
- Refer to [API Doc](API-Doc.md) to check the full usage of the APIs

- Tests
  - Get Risks
  - Add a risk
  - Adding a risk without title, state, description shouldn't be allowed
  - Adding a risk with invalid state shouldn't be allowed
  - Get risk by id that doesn't exist should respond with 404
  
- Service has following dependencies
  - [gin](https://github.com/gin-gonic/gin) - To create the http web server
  - [uuid](https://github.com/google/uuid) - To generate new uuid(s)

- Database interface has InMemoryDatabase implementation for now , which is extensible for future

- Validation package from gin is used for request validation

- Database models and controller layer models while may seem same at the moment they are kept different
  as we do not want to expose database models directly to the client

- To keep the API outputs consistent and extensible even for a single risk
  the format is kept as a list

- Future enhancements
  - Add pagination to GET `v1/risks` API
  - Add rate limiting