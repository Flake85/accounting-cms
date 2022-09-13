# Create Invoice

`POST '/invoice'`

payload:
```json
{
    "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b"
}
```

ok response:
```json
{
    "data": {
        "id": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
        "createdAt": "2022-09-12T13:42:07.3690728-05:00",
        "updatedAt": "2022-09-12T13:42:07.3690728-05:00",
        "deletedAt": null,
        "description": "client 1: 2022-09-12",
        "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
        "client": {
            "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
            "createdAt": "2022-09-12T18:40:50.000956Z",
            "updatedAt": "2022-09-12T18:41:57.21796Z",
            "deletedAt": null,
            "name": "client 1",
            "address": "client address 1",
            "email": "client1@email.com"
        },
        "sales": [
            {
                "id": "945bc734-32ca-11ed-a08e-b28fcf2b227b",
                "createdAt": "2022-09-12T18:41:57.220248Z",
                "updatedAt": "2022-09-12T18:42:07.380271Z",
                "deletedAt": null,
                "description": "sale 3",
                "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                "client": {
                    "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                    "createdAt": "2022-09-12T18:40:50.000956Z",
                    "updatedAt": "2022-09-12T18:41:57.21796Z",
                    "deletedAt": null,
                    "name": "client 1",
                    "address": "client address 1",
                    "email": "client1@email.com"
                },
                "invoiceId": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
                "units": 65,
                "unitCost": 9.21,
                "total": 598.65
            }
        ],
        "salesTotal": 0,
        "labors": [
            {
                "id": "8c934981-32ca-11ed-a08e-b28fcf2b227b",
                "createdAt": "2022-09-12T18:41:44.162008Z",
                "updatedAt": "2022-09-12T18:42:07.388743Z",
                "deletedAt": null,
                "description": "labor 12",
                "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                "client": {
                    "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                    "createdAt": "2022-09-12T18:40:50.000956Z",
                    "updatedAt": "2022-09-12T18:41:57.21796Z",
                    "deletedAt": null,
                    "name": "client 1",
                    "address": "client address 1",
                    "email": "client1@email.com"
                },
                "invoiceId": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
                "hoursWorked": 25.5,
                "hourlyRate": 75.5,
                "total": 1925.25
            }
        ],
        "laborsTotal": 0,
        "grandTotal": 0,
        "isPaid": false
    },
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "invoice decode malfunction"
    }
}
```

# Update Invoice

`PUT '/invoice/:id'`

payload:
```json
{
    "isPaid": true
}
```

ok response:
```json
{
    "data": {
        "id": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
        "createdAt": "2022-09-12T18:42:07.369073Z",
        "updatedAt": "2022-09-12T13:46:37.0748988-05:00",
        "deletedAt": null,
        "description": "client 1: 2022-09-12",
        "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
        "client": {
            "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
            "createdAt": "2022-09-12T18:40:50.000956Z",
            "updatedAt": "2022-09-12T13:46:37.0730078-05:00",
            "deletedAt": null,
            "name": "client 1",
            "address": "client address 1",
            "email": "client1@email.com"
        },
        "sales": null,
        "salesTotal": 0,
        "labors": null,
        "laborsTotal": 0,
        "grandTotal": 0,
        "isPaid": true
    },
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "error occurred creating invoice"
    }
}
```

# Get Invoice

`GET '/invoice/:id'`

ok response:
```json
{
    "data": {
        "id": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
        "createdAt": "2022-09-12T18:42:07.369073Z",
        "updatedAt": "2022-09-12T18:46:37.074899Z",
        "deletedAt": null,
        "description": "client 1: 2022-09-12",
        "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
        "client": {
            "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
            "createdAt": "2022-09-12T18:40:50.000956Z",
            "updatedAt": "2022-09-12T18:46:37.073008Z",
            "deletedAt": null,
            "name": "client 1",
            "address": "client address 1",
            "email": "client1@email.com"
        },
        "sales": [
            {
                "id": "945bc734-32ca-11ed-a08e-b28fcf2b227b",
                "createdAt": "2022-09-12T18:41:57.220248Z",
                "updatedAt": "2022-09-12T18:42:07.380271Z",
                "deletedAt": null,
                "description": "sale 3",
                "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                "client": {
                    "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                    "createdAt": "2022-09-12T18:40:50.000956Z",
                    "updatedAt": "2022-09-12T18:46:37.073008Z",
                    "deletedAt": null,
                    "name": "client 1",
                    "address": "client address 1",
                    "email": "client1@email.com"
                },
                "invoiceId": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
                "units": 65,
                "unitCost": 9.21,
                "total": 598.65
            }
        ],
        "salesTotal": 598.65,
        "labors": [
            {
                "id": "8c934981-32ca-11ed-a08e-b28fcf2b227b",
                "createdAt": "2022-09-12T18:41:44.162008Z",
                "updatedAt": "2022-09-12T18:42:07.388743Z",
                "deletedAt": null,
                "description": "labor 12",
                "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                "client": {
                    "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                    "createdAt": "2022-09-12T18:40:50.000956Z",
                    "updatedAt": "2022-09-12T18:46:37.073008Z",
                    "deletedAt": null,
                    "name": "client 1",
                    "address": "client address 1",
                    "email": "client1@email.com"
                },
                "invoiceId": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
                "hoursWorked": 25.5,
                "hourlyRate": 75.5,
                "total": 1925.25
            }
        ],
        "laborsTotal": 1925.25,
        "grandTotal": 2523.9,
        "isPaid": true
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

# Get Invoices

`GET '/invoice'`

ok response:
```json
{
    "data": [
        {
            "id": "9a68b87e-32ca-11ed-a08e-b28fcf2b227b",
            "createdAt": "2022-09-12T18:42:07.369073Z",
            "updatedAt": "2022-09-12T18:42:07.369073Z",
            "deletedAt": null,
            "description": "client 1: 2022-09-12",
            "clientId": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
            "client": {
                "id": "6c4b3e2b-32ca-11ed-a08e-b28fcf2b227b",
                "createdAt": "2022-09-12T18:40:50.000956Z",
                "updatedAt": "2022-09-12T18:41:57.21796Z",
                "deletedAt": null,
                "name": "client 1",
                "address": "client address 1",
                "email": "client1@email.com"
            },
            "sales": null,
            "salesTotal": 0,
            "labors": null,
            "laborsTotal": 0,
            "grandTotal": 0,
            "isPaid": false
        }
    ],
    "error": {}
}
```

error response:
```json
{
    "error": {
        "message": "error occurred retrieving invoices"
    }
}
```

# Delete Invoice

`DELETE '/invoice/:id'`

ok response: 
```json
{
    "data": {
        "id": "9aa4726d-32cc-11ed-a08e-b28fcf2b227b",
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
        "sales": null,
        "salesTotal": 0,
        "labors": null,
        "laborsTotal": 0,
        "grandTotal": 0,
        "isPaid": false
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
