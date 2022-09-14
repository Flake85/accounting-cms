import { useState } from "react"
import { useRouter } from "next/router"
import Form from 'react-bootstrap/Form'
import Button from "react-bootstrap/Button"
import { useDispatch } from 'react-redux'
import { setAlertData, openAlertModal } from '../../../slices/alertModalSlice'

export default function UpdateClient({ client, url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [clientName, setClientName] = useState(client.data.name)
    const [clientEmail, setClientEmail] = useState(client.data.email)
    const [clientAddress, setClientAddress] = useState(client.data.address)

    const handleNameChange = event => setClientName(event.target.value)
    const handleEmailChange = event => setClientEmail(event.target.value)
    const handleAddressChange = event => setClientAddress(event.target.value)
    
    async function submitClient(event) {
        event.preventDefault();
        var updatedClient = {
            name: clientName,
            email: clientEmail,
            address: clientAddress
        }
        try {
            const res = await fetch(`${url}/client/${client.data.id}`, {
                method: 'PUT',
                mode: 'cors',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updatedClient)
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
            router.push(`/client/${data.data.id}`)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Update Client</h1>
            <hr />
            <Form onSubmit={submitClient}>
                <Form.Group className="mb-3">
                    <Form.Label>Name</Form.Label>
                    <Form.Control placeholder="Enter client name" 
                                value={clientName}
                                onChange={handleNameChange}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicEmail">
                    <Form.Label>Email</Form.Label>
                    <Form.Control type="email" 
                                placeholder="Enter client email"
                                value={clientEmail}
                                onChange={handleEmailChange} />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Address</Form.Label>
                    <Form.Control placeholder="Enter client address"
                                value={clientAddress}
                                onChange={handleAddressChange} />
                </Form.Group>
                <Button type="submit" className="me-1">Submit</Button>
                <Button href={'/client'}>Cancel</Button>
            </Form>
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/client/${id}`)
    const client = await res.json()
    return { props: { client, url } }
}
