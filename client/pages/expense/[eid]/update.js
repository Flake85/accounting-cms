import { useState } from "react"
import { useRouter } from "next/router"
import Form from 'react-bootstrap/Form'
import Button from "react-bootstrap/Button"

export default function UpdateExpense({ expense, url }) {
    const router = useRouter()
    const [expenseDescription, setExpenseDescription] = useState(expense.data.description)
    const [expenseCost, setExpenseCost] = useState(expense.data.cost)

    const handleDescriptionChange = event => setExpenseDescription(event.target.value)
    const handleCostChange = event => setExpenseCost(event.target.value)
    
    const submitExpense = async event => {
        event.preventDefault();
        var newExpense = {
            description: expenseDescription,
            cost: parseFloat(expenseCost)
        }
        await fetch(`${url}/expense/${expense.data.id}`, {
            method: 'PUT',
            mode: 'cors',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(newExpense)
        })
        .then(() => {
            alert("successfully updated expense: " + newExpense.description)
            router.push(`/expense/${expense.data.id}`)
        })
        .catch(err => alert(err))
    }

    return (
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
    )
}

export async function getServerSideProps(context) {
    const id = context.query.eid
    const res = await fetch(`${process.env.REACT_APP_BASEURL}/expense/${id}`)
    const expense = await res.json()
    const url = process.env.REACT_APP_BASEURL
    return { props: { expense, url } }
}
