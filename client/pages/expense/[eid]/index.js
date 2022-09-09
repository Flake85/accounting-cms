import React from 'react'
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import { useState } from 'react'
import { useRouter } from 'next/router'
import { useDispatch } from 'react-redux'
import { openAlertModal, setAlertData } from '../../../slices/alertModalSlice'

export default function Expense({ url, expense }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [show, setShow] = useState(false)

    async function deleteExpense() {
        try {
            const res = await fetch(`${url}/expense/${expense.data.id}`, {
                method: 'DELETE',
                mode: 'cors'
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
            router.push(`/expense`)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Expense</h1>
            <hr />
            <Alert show={show} variant="warning" dismissible onClose={() => setShow(false)}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ expense.data.description }"?</p>
                <Button className="me-1" onClick={deleteExpense}>Confirm</Button>
                <Button onClick={() => setShow(false)}>Cancel</Button>
            </Alert>
            { expense.data
                ? <div>
                    <p><strong>Description: </strong>{expense.data.description}</p>
                    <p><strong>Cost: </strong>{expense.data.cost}</p>
                    <Button href={`/expense/${expense.data.id}/update`} className="me-1">Update Expense</Button>
                    <Button variant="danger" onClick={() => setShow(true)}>Delete</Button>
                </div>
                : <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ expense.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.eid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/expense/${id}`)
    const expense = await res.json()
    return { props: { url, expense } }
}
