# Create Expense

    POST '/expense'

Payload:

    {
        "description": "expense 1",
        "cost": 15.55
    }

ok response: 

    {
        "data": {
            "id": "1f96043f-308b-11ed-9ebf-c62779d44d46",
            "createdAt": "2022-09-09T17:02:40.6867731-05:00",
            "updatedAt": "2022-09-09T17:02:40.6867731-05:00",
            "deletedAt": null,
            "description": "expense 1",
            "cost": 15.55
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "expense validation error"
        }
    }

# Update Expense

    PUT 'expense/:id'

payload:

    {
        "description": "expense 1",
        "cost": 21.23
    }

ok response: 

    {
        "data": {
            "id": "1f96043f-308b-11ed-9ebf-c62779d44d46",
            "createdAt": "2022-09-09T22:02:40.686773Z",
            "updatedAt": "2022-09-09T17:04:06.0606121-05:00",
            "deletedAt": null,
            "description": "expense 1",
            "cost": 21.23
        },
        "error": {}
    }

# Get Expense

    GET '/expense/:id'

ok response:

    {
        "data": {
            "id": "1f96043f-308b-11ed-9ebf-c62779d44d46",
            "createdAt": "2022-09-09T22:02:40.686773Z",
            "updatedAt": "2022-09-09T22:04:06.060612Z",
            "deletedAt": null,
            "description": "expense 1",
            "cost": 21.23
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "invalid uuid"
        }
    }

# Get Expenses

    GET '/expense'

ok response:

    {
        "data": [
            {
                "id": "a7d9fa3e-2b15-11ed-96ce-2ab0f0b3073f",
                "createdAt": "2022-09-02T23:19:12.874191Z",
                "updatedAt": "2022-09-02T23:19:12.874191Z",
                "deletedAt": null,
                "description": "expense 2",
                "cost": 234.43
            },
            {
                "id": "1f96043f-308b-11ed-9ebf-c62779d44d46",
                "createdAt": "2022-09-09T22:02:40.686773Z",
                "updatedAt": "2022-09-09T22:04:06.060612Z",
                "deletedAt": null,
                "description": "expense 1",
                "cost": 21.23
            }
        ],
        "error": {}
    }

error response:

    {
        "error": {
            "message": "error occurred retrieving clients"
        }
    }

# Delete Expense

    DELETE '/expense/:id'

ok response:

    {
        "data": {
            "id": "1f96043f-308b-11ed-9ebf-c62779d44d46",
            "createdAt": "0001-01-01T00:00:00Z",
            "updatedAt": "0001-01-01T00:00:00Z",
            "deletedAt": null,
            "description": "",
            "cost": 0
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "invalid uuid"
        }
    }