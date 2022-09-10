# Create Labor

    POST '/labor'

payload:

    {
        "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
        "invoiceId": null,
        "description": "labor 1",
        "hoursWorked": 25.5,
        "hourlyRate": 75.50
    }

ok response:

    {
        "data": {
            "id": "0cd67859-3097-11ed-9ebf-c62779d44d46",
            "createdAt": "2022-09-09T18:28:03.1952076-05:00",
            "updatedAt": "2022-09-09T18:28:03.1952076-05:00",
            "deletedAt": null,
            "description": "labor 1",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-09T18:28:03.1930972-05:00",
                "deletedAt": null,
                "name": "client 1",
                "address": "client 1 address",
                "email": "client1@email.com"
            },
            "hoursWorked": 25.5,
            "hourlyRate": 75.5,
            "total": 1925.25
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "labor decode malfunction"
        }
    }

# Update Labor

    PUT '/labor/:id'

payload:

    {
        "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
        "invoiceId": null,
        "description": "labor 1",
        "hoursWorked": 0.5,
        "hourlyRate": 60
    }

ok response:

    {
        "data": {
            "id": "0cd67859-3097-11ed-9ebf-c62779d44d46",
            "createdAt": "2022-09-09T23:28:03.195208Z",
            "updatedAt": "2022-09-09T18:38:06.0440079-05:00",
            "deletedAt": null,
            "description": "labor 1",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-09T18:38:06.0417558-05:00",
                "deletedAt": null,
                "name": "client 1",
                "address": "client 1 address",
                "email": "client1@email.com"
            },
            "hoursWorked": 0.5,
            "hourlyRate": 60,
            "total": 30
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "labor not found"
        }
    }

# Get labor

    GET '/labor/:id'

ok response without invoice:

    {
        "data": {
            "id": "0cd67859-3097-11ed-9ebf-c62779d44d46",
            "createdAt": "2022-09-09T23:28:03.195208Z",
            "updatedAt": "2022-09-09T23:38:06.044008Z",
            "deletedAt": null,
            "description": "labor 1",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-09T23:38:06.041756Z",
                "deletedAt": null,
                "name": "client 1",
                "address": "client 1 address",
                "email": "client1@email.com"
            },
            "hoursWorked": 0.5,
            "hourlyRate": 60,
            "total": 30
        },
        "error": {}
    }

ok response with invoice:

    {
        "data": {
            "id": "52ffb185-2af8-11ed-a7ae-2ab0f0b3073f",
            "createdAt": "2022-09-02T19:49:15.111666Z",
            "updatedAt": "2022-09-02T19:53:53.402799Z",
            "deletedAt": null,
            "description": "labor 1",
            "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
            "client": {
                "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "createdAt": "2022-09-02T18:45:30.471254Z",
                "updatedAt": "2022-09-09T23:38:06.041756Z",
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
                    "updatedAt": "2022-09-09T23:38:06.041756Z",
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
            "hoursWorked": 4,
            "hourlyRate": 32.43,
            "total": 129.72
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "labor not found"
        }
    }

# Get Labors

    GET '/labor'

ok response:

    {
        "data": [
            {
                "id": "52ffb185-2af8-11ed-a7ae-2ab0f0b3073f",
                "createdAt": "2022-09-02T19:49:15.111666Z",
                "updatedAt": "2022-09-02T19:53:53.402799Z",
                "deletedAt": null,
                "description": "labor 1",
                "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "client": {
                    "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                    "createdAt": "2022-09-02T18:45:30.471254Z",
                    "updatedAt": "2022-09-09T23:38:06.041756Z",
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
                        "updatedAt": "2022-09-09T23:38:06.041756Z",
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
                "hoursWorked": 4,
                "hourlyRate": 32.43,
                "total": 129.72
            },
            {
                "id": "0cd67859-3097-11ed-9ebf-c62779d44d46",
                "createdAt": "2022-09-09T23:28:03.195208Z",
                "updatedAt": "2022-09-09T23:38:06.044008Z",
                "deletedAt": null,
                "description": "labor 1",
                "clientId": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                "client": {
                    "id": "6b562268-2aef-11ed-b995-2ab0f0b3073f",
                    "createdAt": "2022-09-02T18:45:30.471254Z",
                    "updatedAt": "2022-09-09T23:38:06.041756Z",
                    "deletedAt": null,
                    "name": "client 1",
                    "address": "client 1 address",
                    "email": "client1@email.com"
                },
                "hoursWorked": 0.5,
                "hourlyRate": 60,
                "total": 30
            }
        ],
        "error": {}
    }

error response:

    {
        "error": {
            "message": "error occurred retrieving labors"
        }
    }

# Delete Labor

    DELETE '/labor/:id'

ok response:

    {
        "data": {
            "id": "0cd67859-3097-11ed-9ebf-c62779d44d46",
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
            "hoursWorked": 0,
            "hourlyRate": 0,
            "total": 0
        },
        "error": {}
    }

error response:

    {
        "error": {
            "message": "invalid uuid"
        }
    }