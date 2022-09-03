import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"

export default function NewExpense({ url }) {
    const router = useRouter()
    const [expenseDescription, setExpenseDescription] = useState('')
    const [expenseCost, setExpenseCost] = useState(0)

    const handleDescriptionChange = event => setExpenseDescription(event.target.value)
    const handleCostChange = event => setExpenseCost(event.target.value)

    const submitExpense = async event => {
        event.preventDefault();
        var newExpense = {
            description: expenseDescription,
            cost: parseFloat(expenseCost),
        }
        await fetch(`${url}/expense`, {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(newExpense)
        })
        .then(async (res) => {
            if(res.ok) return res.json()
            const json = await res.json();
            throw new Error(json.error.message);
        })
        .then(() => {
            alert("successfully submitted new expense: "+newExpense.description)
            router.push('/expense')
        })
        .catch(err => alert(err))
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