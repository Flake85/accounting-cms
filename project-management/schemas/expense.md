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


