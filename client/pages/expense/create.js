import React from 'react'
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"
import { openAlertModal, setAlertData } from '../../slices/alertModalSlice';
import { useDispatch } from 'react-redux';

export default function NewExpense({ url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [expenseDescription, setExpenseDescription] = useState('')
    const [expenseCost, setExpenseCost] = useState(0)

    const handleDescriptionChange = event => setExpenseDescription(event.target.value)
    const handleCostChange = event => setExpenseCost(event.target.value)

    async function submitExpense(event) {
        event.preventDefault();
        var newExpense = {
            description: expenseDescription,
            cost: parseFloat(expenseCost)
        }
        try {
            const res = await fetch(`${url}/expense`, {
                method: 'POST',
                mode: 'cors',
                body: JSON.stringify(newExpense)
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
        } catch(err) { console.log(err) }
    }

    return (
        <div>
            <h1>Expense</h1> 
            <hr />
            <Form onSubmit={submitExpense}>
                <Form.Group className="mb-3">
                    <Form.Label>Description</Form.Label>
                    <Form.Control placeholder="Enter expense description" 
                                value={expenseDescription}
                                onChange={handleDescriptionChange}/>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Cost</Form.Label>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter cost"
                                value={expenseCost}
                                onChange={handleCostChange} />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </div>
    )
}

export async function getStaticProps() {
    return { props: { url: process.env.REACT_APP_BASEURL } }
}