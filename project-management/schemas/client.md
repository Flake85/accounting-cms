# Create Client

`POST '/client'`

Payload:
```json
{ 
    "name": "client 1",
    "address": "client address 1",
    "email": "client1@email.com"
}
```

ok response:
```json
{
    "data": {
        "id": "4033fd43-3082-11ed-b75b-c62779d44d46",
        "createdAt": "2022-09-09T15:59:09.9384356-05:00",
        "updatedAt": "2022-09-09T15:59:09.9384356-05:00",
        "deletedAt": null,
        "name": "client 1",
        "address": "client address 1",
        "email": "client1@email.com"
    },
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "client validation error"
    }
}
```

# Update Client

`PUT '/client/:id'`

Payload:
```json
{ 
    "name": "client 1",
    "address": "client address 1",
    "email": "client1@email.com"
}
```

ok response:
```json
{
    "data": {
        "id": "4033fd43-3082-11ed-b75b-c62779d44d46",
        "createdAt": "2022-09-09T20:59:09.938436Z",
        "updatedAt": "2022-09-09T16:23:26.2038114-05:00",
        "deletedAt": null,
        "name": "client 1",
        "address": "client address 1",
        "email": "client1@email.com"
    },
    "error": {}
}
```

error response: 
```json
{
    "error": {
        "message": "client validation error"
    }
}
```

# Get Client

`GET '/client/:id'`

ok response:
```json
{
    "data": {
        "id": "4033fd43-3082-11ed-b75b-c62779d44d46",
        "createdAt": "2022-09-09T20:59:09.938436Z",
        "updatedAt": "2022-09-09T21:23:26.203811Z",
        "deletedAt": null,
        "name": "client 1",
        "address": "client address 1",
        "email": "client1@email.com"
    },
    "error": {}
}
```

error response: 
```json
{
    "error": {
        "message": "invalid uuid"
    }
}
```

# Get Clients

`GET '/client'`

ok response:
```json
{
    "data": [
        {
            "id": "851ba2a8-2c8f-11ed-97f3-76a9b8cce05e",
            "createdAt": "2022-09-04T20:24:04.349551Z",
            "updatedAt": "2022-09-09T03:44:14.714409Z",
            "deletedAt": null,
            "name": "client 3",
            "address": "client 3 address st",
            "email": "client3@email.com"
        },
        {
            "id": "4033fd43-3082-11ed-b75b-c62779d44d46",
            "createdAt": "2022-09-09T20:59:09.938436Z",
            "updatedAt": "2022-09-09T21:23:26.203811Z",
            "deletedAt": null,
            "name": "client 1",
            "address": "client address 1",
            "email": "client1@email.com"
        }
    ],
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "error occurred retrieving clients"
    }
}
```

# Delete Client

`DELETE '/client/:id'`

ok response: 
```json
{
    "data": {
        "id": "4033fd43-3082-11ed-b75b-c62779d44d46",
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "0001-01-01T00:00:00Z",
        "deletedAt": null,
        "name": "",
        "address": "",
        "email": ""
    },
    "error": {}
}
```

error response: 
```json
{
    "error": {
        "message": "invalid uuid"
    }
}
```

# Undelete a Deleted Client

`PUT 'client_deleted/:id'`

ok response: 
```json
{
    "data": {
        "id": "4033fd43-3082-11ed-b75b-c62779d44d46",
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "2022-09-09T16:50:34.8277274-05:00",
        "deletedAt": null,
        "name": "",
        "address": "",
        "email": ""
    },
    "error": {}
}
```

error response: 
```json
{
    "error": {
        "message": "invalid uuid"
    }
}
```

# Get Deleted Client

    GET '/client_deleted/:id'

ok response:
```json
{
    "data": {
        "id": "68dc3071-2fe6-11ed-a1a2-8e371d999111",
        "createdAt": "2022-09-09T02:23:36.659836Z",
        "updatedAt": "2022-09-09T02:23:36.659836Z",
        "deletedAt": "2022-09-09T02:23:56.575597Z",
        "name": "test",
        "address": "test",
        "email": "test@email.com"
    },
    "error": {}
}
```

error response: 
```json
{
    "error": {
        "message": "invalid uuid"
    }
}
```

# Get Deleted Clients

`GET /client_deleted`

ok response: 
```json
{
    "data": [
        {
            "id": "9eadc9c7-2c8f-11ed-97f3-76a9b8cce05e",
            "createdAt": "2022-09-04T20:24:47.248862Z",
            "updatedAt": "2022-09-04T20:24:47.248862Z",
            "deletedAt": "2022-09-09T02:20:06.371026Z",
            "name": "client 4",
            "address": "client 4 address rd",
            "email": "client4@email.com"
        },
        {
            "id": "68dc3071-2fe6-11ed-a1a2-8e371d999111",
            "createdAt": "2022-09-09T02:23:36.659836Z",
            "updatedAt": "2022-09-09T02:23:36.659836Z",
            "deletedAt": "2022-09-09T02:23:56.575597Z",
            "name": "test",
            "address": "test",
            "email": "test@email.com"
        },
    ],
    "error": {}
}
```
