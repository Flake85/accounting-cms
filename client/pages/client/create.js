import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"
import { useDispatch } from 'react-redux'
import { setAlertData, openAlertModal } from '../../slices/alertModalSlice'

export default function NewClient({ url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [clientName, setClientName] = useState('')
    const [clientEmail, setClientEmail] = useState('')
    const [clientAddress, setClientAddress] = useState('')

    function handleNameChange(event) { setClientName(event.target.value) }
    function handleEmailChange(event) { setClientEmail(event.target.value) }
    function handleAddressChange(event) { setClientAddress(event.target.value) }

    async function submitClient(event) {
        event.preventDefault();
        var newClient = {
            name: clientName,
            email: clientEmail,
            address: clientAddress
        }
        try {
            const res = await fetch(`${url}/client`, {
                method: 'POST',
                mode: 'cors',
                body: JSON.stringify(newClient)
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
            <h1>Create Client</h1>
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
                <Button href={`/client`}>Cancel</Button>
            </Form>
        </div>
    )
}

export async function getStaticProps() {
    return { props: { url: process.env.REACT_APP_BASEURL } }
}
