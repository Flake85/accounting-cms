## Clients:
- id number
- name string
- address string
- email string

## Invoices:
- id number
- clientId number
- isInvoiced boolean

## Labor:
- id number
- description string
- clientId number
- invoiceId number
- hours_worked number
- hourly_rate number

## Sales:
- id number
- description string
- clientId number
- invoiceId number
- units number
- unit_cost number

## Expenses:
- id number
- description string
- cost number