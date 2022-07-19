## Clients:
- id number
- name string
- address string
- email string

## Invoices:
- id number
- clientId number
- laborId number
- salesId number
- job_completed boolean
- isInvoiced boolean

## Labor:
- id number
- clientId number
- transactionId number
- hours_worked number
- hourly_rate number

## Sales:
- id number
- clientId number
- transactionId number
- units number
- unit_cost number

## Expenses:
- id number
- description string
- cost number