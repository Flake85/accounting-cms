## Clients:
- id uuid
- createdAt time
- updatedAt time
- deletedAt time
- name string
- address string
- email string

## Invoices:
- id uuid
- createdAt time
- updatedAt time
- deletedAt time
- clientId number
- isInvoiced boolean

## Labor:
- id uuid
- createdAt time
- updatedAt time
- deletedAt time
- description string
- clientId number
- invoiceId number
- hours_worked number
- hourly_rate number

## Sales:
- id uuid
- createdAt time
- updatedAt time
- deletedAt time
- description string
- clientId number
- invoiceId number
- units number
- unit_cost number

## Expenses:
- id uuid
- createdAt time
- updatedAt time
- deletedAt time
- description string
- cost number