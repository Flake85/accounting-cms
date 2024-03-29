import React from 'react'
import Table from 'react-bootstrap/Table';
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useState } from 'react';
import { useRouter } from 'next/router';
import { useDispatch } from 'react-redux';
import { openAlertModal, setAlertData } from '../../slices/alertModalSlice';

export default function Invoices({ invoices, url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [show, setShow] = useState(false)
    const [target, setTarget] = useState('')

    function confirmDelete(invoice) {
        setShow(true); setTarget(invoice)
    }
    function closeAlert() {
        setShow(false); setTarget("")
    }

    async function deleteInvoice() {
        try {
            const res = await fetch(`${url}/invoice/${target.id}`, {
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
            <h1>Invoices</h1>
            <Alert show={show} variant="warning" dismissible onClose={closeAlert}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ target.description }"?</p>
                <Button className="me-1" onClick={deleteInvoice}>Confirm</Button>
                <Button onClick={closeAlert}>Cancel</Button>
            </Alert>
            { invoices.data.length
                ? <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Invoice #</th>
                            <th>Description</th>
                            <th>Client</th>
                            <th>Is Paid</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {invoices.data.map((invoice, i) => (
                            <tr key={invoice.id}>
                                <td>{ i + 1 }</td>
                                <td><Link href={`/invoice/${invoice.id}`}><a>{ invoice.id }</a></Link></td>
                                <td><Link href={`/invoice/${invoice.id}`}><a>{ invoice.description }</a></Link></td>
                                { invoice.client.name
                                    ? <td><Link href={`/client/${invoice.clientId}`}><a>{ invoice.client.name }</a></Link></td>
                                    : <td><Link href={`/client/${invoice.clientId}/deleted`}><a className="text-danger">{ invoice.clientId } (inactive)</a></Link></td>
                                }
                                <td>{ invoice.isPaid.toString() }</td>
                                { invoice.client.name
                                    ? <td>
                                        <Link href={`/invoice/${invoice.id}/update`}><a><i className="bi-pencil-square text-success"></i></a></Link>
                                        {!invoice.isPaid && <Link href={`#`}><a onClick={() => confirmDelete(invoice)}><i className="bi-trash text-danger"></i></a></Link>}
                                    </td>
                                    : <td></td>
                                }
                            </tr>
                        ))}
                    </tbody>
                </Table>
                : <div><hr /><p>Invoices haven't been added yet.</p></div>
            }
            <Button href="/invoice/create" className="mb-5">Add Invoice</Button>
        </div>
    );
}

export async function getServerSideProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/invoice`)
    const invoices = await res.json()
    return { props: { invoices, url } }
}
