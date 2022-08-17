## Name:
Http Response & CORS

---
## Description:
make middleware to set http headers.
make a http response package for api responses.

---
#### Subtasks:
- [x] at the highest level, make middleware that sets the Headers for each endpoint.
    - [x] make middleware package for cors and set headers there
    - [x] content-type to application/json
- [x] make response package for building api response
    - [x] 1 response for building ok(200) responses. accepts a message to send in response
    - [x] 1 response for building error responses. accepts error code, message, & optional data to send in response