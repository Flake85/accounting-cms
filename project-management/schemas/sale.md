# Create Sale

`POST '/sale'`

payload:
```json
{
    "description": "expense 1",
    "cost": 143.23
}
```

ok response:
```json
{
    "data": {
        "id": "4c63c313-312c-11ed-8176-328fd411f723",
        "createdAt": "2022-09-10T12:16:24.8305135-05:00",
        "updatedAt": "2022-09-10T12:16:24.8305135-05:00",
        "deletedAt": null,
        "description": "sale 1",
        "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
        "client": {
            "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "createdAt": "2022-09-02T18:45:30.471254Z",
            "updatedAt": "2022-09-10T12:16:24.8282679-05:00",
            "deletedAt": null,
            "name": "client 1",
            "address": "client 1 address",
            "email": "client1@email.com"
        },
        "units": 65,
        "unitCost": 9.21,
        "total": 598.65
    },
    "error": {}
}
```

error response: 
```json
{
    "error": {
        "message": "sale decode malfunction"
    }
}
```

# Update Sale

`PUT 'sale/:id'`

payload: 
```json
{
    "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
    "invoiceId": null,
    "description": "sale 1",
    "units": 65,
    "unitCost": 1.75
}
```

ok response:
```json
{
    "data": {
        "id": "4c63c313-312c-11ed-8176-328fd411f723",
        "createdAt": "2022-09-10T17:16:24.830514Z",
        "updatedAt": "2022-09-10T12:23:30.9044227-05:00",
        "deletedAt": null,
        "description": "sale 1",
        "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
        "client": {
            "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "createdAt": "2022-09-02T18:45:30.471254Z",
            "updatedAt": "2022-09-10T12:23:30.9025201-05:00",
            "deletedAt": null,
            "name": "client 1",
            "address": "client 1 address",
            "email": "client1@email.com"
        },
        "units": 65,
        "unitCost": 1.75,
        "total": 113.75
    },
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "sale decode malfunction"
    }
}
```

# Get Sale

`GET '/sale/:id'`

ok response without invoice:
```json
{
    "data": {
        "id": "4c63c313-312c-11ed-8176-328fd411f723",
        "createdAt": "2022-09-10T17:16:24.830514Z",
        "updatedAt": "2022-09-10T17:23:30.904423Z",
        "deletedAt": null,
        "description": "sale 1",
        "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
        "client": {
            "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "createdAt": "2022-09-02T18:45:30.471254Z",
            "updatedAt": "2022-09-10T17:23:30.90252Z",
            "deletedAt": null,
            "name": "client 1",
            "address": "client 1 address",
            "email": "client1@email.com"
        },
        "units": 65,
        "unitCost": 1.75,
        "total": 113.75
    },
    "error": {}
}
```

ok response with invoice:
```json
{
    "data": {
        "id": "6750e08a-2af8-11ed-a7ae-2ab0f0b3073f",
        "createdAt": "2022-09-02T19:49:49.198055Z",
        "updatedAt": "2022-09-09T18:43:21.072065Z",
        "deletedAt": null,
        "description": "sale 1",
        "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
        "client": {
            "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "createdAt": "2022-09-02T18:45:30.471254Z",
            "updatedAt": "2022-09-10T17:23:30.90252Z",
            "deletedAt": null,
            "name": "client 1",
            "address": "client 1 address",
            "email": "client1@email.com"
        },
        "invoiceId": "f8dd449f-2af8-11ed-a353-2ab0f0b3073f",
        "invoice": {
            "id": "f8dd449f-2af8-11ed-a353-2ab0f0b3073f",
            "createdAt": "2022-09-02T19:53:53.386146Z",
            "updatedAt": "2022-09-09T18:43:21.069586Z",
            "deletedAt": null,
            "description": "client 1: 2022-09-02",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-10T17:23:30.90252Z",
                "deletedAt": null,
                "name": "client 1",
                "address": "client 1 address",
                "email": "client1@email.com"
            },
            "sales": null,
            "salesTotal": 0,
            "labors": null,
            "laborsTotal": 0,
            "grandTotal": 0,
            "isPaid": true
        },
        "units": 12,
        "unitCost": 12.23,
        "total": 146.76
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

# Get Sales

`GET '/sale'`

ok response:
```json
{
    "data": [
        {
            "id": "5b04fb02-2f9e-11ed-a1a2-8e371d999111",
            "createdAt": "2022-09-08T17:47:49.68112Z",
            "updatedAt": "2022-09-09T19:44:19.836552Z",
            "deletedAt": null,
            "description": "sale 5",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-10T17:23:30.90252Z",
                "deletedAt": null,
                "name": "client 1",
                "address": "client 1 address",
                "email": "client1@email.com"
            },
            "invoiceId": "cbe3715f-3077-11ed-b75b-c62779d44d46",
            "invoice": {
                "id": "cbe3715f-3077-11ed-b75b-c62779d44d46",
                "createdAt": "2022-09-09T19:44:19.827806Z",
                "updatedAt": "2022-09-09T19:44:19.827806Z",
                "deletedAt": null,
                "description": "client 1: 2022-09-09",
                "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "client": {
                    "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                    "createdAt": "2022-09-02T18:45:30.471254Z",
                    "updatedAt": "2022-09-10T17:23:30.90252Z",
                    "deletedAt": null,
                    "name": "client 1",
                    "address": "client 1 address",
                    "email": "client1@email.com"
                },
                "sales": null,
                "salesTotal": 0,
                "labors": null,
                "laborsTotal": 0,
                "grandTotal": 0,
                "isPaid": false
            },
            "units": 32,
            "unitCost": 12.68,
            "total": 405.76
        },
        {
            "id": "4c63c313-312c-11ed-8176-328fd411f723",
            "createdAt": "2022-09-10T17:16:24.830514Z",
            "updatedAt": "2022-09-10T17:23:30.904423Z",
            "deletedAt": null,
            "description": "sale 1",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-10T17:23:30.90252Z",
                "deletedAt": null,
                "name": "client 1",
                "address": "client 1 address",
                "email": "client1@email.com"
            },
            "units": 65,
            "unitCost": 1.75,
            "total": 113.75
        }
    ],
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "error occurred retrieving sales"
    }
}
```

# Delete sale

`DELETE '/sale/:id'`

ok response:
```json
{
    "data": {
        "id": "4c63c313-312c-11ed-8176-328fd411f723",
        "createdAt": "0001-01-01T00:00:00Z",
        "updatedAt": "0001-01-01T00:00:00Z",
        "deletedAt": null,
        "description": "",
        "clientId": "00000000-0000-0000-0000-000000000000",
        "client": {
            "id": "00000000-0000-0000-0000-000000000000",
            "createdAt": "0001-01-01T00:00:00Z",
            "updatedAt": "0001-01-01T00:00:00Z",
            "deletedAt": null,
            "name": "",
            "address": "",
            "email": ""
        },
        "units": 0,
        "unitCost": 0,
        "total": 0
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
