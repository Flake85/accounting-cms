import React from 'react'
import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"
import { useDispatch } from 'react-redux';
import { openAlertModal, setAlertData } from '../../slices/alertModalSlice';

export default function NewInvoice({ url, clients }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [invoiceClient, setInvoiceClient] = useState('null')

    function handleClientChange(event) { setInvoiceClient(event.target.value) }

    async function submitInvoice(event) {
        event.preventDefault();
        var newInvoice = { clientId: invoiceClient }
        try {
            const res = await fetch(`${url}/invoice`, {
                method: 'POST',
                mode: 'cors',
                body: JSON.stringify(newInvoice)
            })
            const data = await res.json()
            if(!res.ok) {
                dispatch(setAlertData({
                    title: 'Something went wrong',
                    body: 'Error: ' + data.error.message
                }))
                dispatch(openAlertModal())
                throw new Error(data)
            }
            router.push(`/invoice/${data.data.id}`)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Create Invoice</h1>
            <hr />
            <Form onSubmit={submitInvoice}>
                <Form.Group className="mb-3">
                    <Form.Label>Client</Form.Label>
                    <Form.Select value={invoiceClient} 
                                onChange={handleClientChange}>
                        <option value="null" disabled>Select a Client</option>
                        {clients.data.map((client) => (
                            <option key={client.id} value={client.id}>{ client.name }</option>
                        ))}
                    </Form.Select>
                </Form.Group>
                <Button type="submit" className="me-1">Submit</Button>
                <Button href={`/invoice`}>Cancel</Button>
            </Form>
        </div>
    )
}

export async function getStaticProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/client`)
    const clients = await res.json()
    return { props: { url, clients } }
}
