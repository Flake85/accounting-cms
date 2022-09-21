import React from 'react'
import { useState } from "react"
import { useRouter } from "next/router"
import Form from 'react-bootstrap/Form'
import InputGroup from 'react-bootstrap/InputGroup'
import Button from "react-bootstrap/Button"
import { useDispatch } from 'react-redux'
import { setAlertData } from '../../../slices/alertModalSlice'

export default function UpdateExpense({ expense, url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [expenseDescription, setExpenseDescription] = useState(expense.data.description)
    const [expenseCost, setExpenseCost] = useState(expense.data.cost)

    const handleDescriptionChange = event => setExpenseDescription(event.target.value)
    const handleCostChange = event => setExpenseCost(event.target.value)
    
    async function submitExpense(event) {
        event.preventDefault();
        var updatedExpense = {
            description: expenseDescription,
            cost: parseFloat(expenseCost)
        }
        try {
            const res = await fetch(`${url}/expense/${expense.data.id}`, {
                method: 'PUT',
                mode: 'cors',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updatedExpense)
            })
            const data = await res.json()
            if(!res.ok) {
                dispatch(setAlertData({
                    title: 'Something went wrong',
                    body: 'Error: ' + data
                }))
                dispatch(openAlertModal())
                throw new Error(data)
            }
            router.push(`/expense/${data.data.id}`)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Update Expense</h1> 
            <hr />
            <Form onSubmit={submitExpense}>
                <Form.Group className="mb-3">
                    <Form.Label>Description</Form.Label>
                    <Form.Control placeholder="Enter expense description" 
                                value={expenseDescription}
                                onChange={handleDescriptionChange}/>
                </Form.Group>

                <Form.Label>Cost</Form.Label>
                <InputGroup className="mb-3">
                    <InputGroup.Text>$</InputGroup.Text>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter cost"
                                value={expenseCost}
                                onChange={handleCostChange} />
                </InputGroup>
                <Button variant="primary" type="submit" className="me-1">Submit</Button>
                <Button href={`/expense`}>Cancel</Button>
            </Form>
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.eid
    const res = await fetch(`${process.env.REACT_APP_BASEURL}/expense/${id}`)
    const expense = await res.json()
    const url = process.env.REACT_APP_BASEURL
    return { props: { expense, url } }
}
