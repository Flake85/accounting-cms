import React from 'react'
import Table from 'react-bootstrap/Table';
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useState } from 'react';
import { useRouter } from 'next/router';
import { useDispatch } from 'react-redux';
import { openAlertModal, setAlertData } from '../../slices/alertModalSlice';

export default function Expenses({ expenses, url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [show, setShow] = useState(false)
    const [target, setTarget] = useState('')

    function confirmDelete(expense) { setShow(true); setTarget(expense) }
    function closeAlert() { setShow(false); setTarget("") }
    
    async function deleteExpense() {
        try {
            const res = await fetch(`${url}/expense/${target.id}`, {
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
            router.reload(window.location.pathname)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Expenses</h1>
            <Alert show={show} variant="warning" dismissible onClose={closeAlert}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ target.description }"?</p>
                <p>This action cannot be undone.</p>
                <Button onClick={deleteExpense}>Confirm</Button>
                <Button onClick={closeAlert}>Cancel</Button>
            </Alert>
            { expenses.data.length
                ? <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Description</th>
                            <th>Cost</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {expenses.data.map((expense, i) => (
                            <tr key={expense.id}>
                                <td>{ i + 1 }</td>
                                <td><Link href={`/expense/${expense.id}`}><a>{ expense.description }</a></Link></td>
                                <td>${ expense.cost }</td>
                                <td>
                                    <Link href={`/expense/${expense.id}/update`}><a><i className="bi-pencil-square text-success"></i></a></Link>
                                    <Link href={`#`}><a onClick={() => confirmDelete(expense)}><i className="bi-trash text-danger"></i></a></Link>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </Table>
                : <div>
                    <hr />
                    <p>No Expenses Added.</p>
                </div>
                }
            <Button href="/expense/create">Add Expense</Button>
        </div>
    );
}

export async function getServerSideProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/expense`)
    const expenses = await res.json()
    return { props: { expenses, url } }
}
