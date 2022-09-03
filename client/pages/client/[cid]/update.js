import { useState } from "react"
import { useRouter } from "next/router"
import Form from 'react-bootstrap/Form'
import Button from "react-bootstrap/Button"

export default function UpdateClient({ client, url }) {
    const router = useRouter()
    const [clientName, setClientName] = useState(client.data.name)
    const [clientEmail, setClientEmail] = useState(client.data.email)
    const [clientAddress, setClientAddress] = useState(client.data.address)

    const handleNameChange = event => setClientName(event.target.value)
    const handleEmailChange = event => setClientEmail(event.target.value)
    const handleAddressChange = event => setClientAddress(event.target.value)
    
    const submitClient = async event => {
        event.preventDefault();
        var newClient = {
            name: clientName,
            email: clientEmail,
            address: clientAddress,
        }
        await fetch(`${url}/client/${client.data.id}`, {
            method: 'PUT',
            mode: 'cors',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(newClient)
        })
        .then(async (res) => {
            if(res.ok) return res.json()
            const json = await res.json();
            throw new Error(json.error.message);
        })
        .then(() => {
            alert("successfully updated client: " + newClient.name)
            router.push(`/client/${client.data.id}`)
        })
        .catch(err => alert(err))
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
    const res = await fetch(`${process.env.REACT_APP_BASEURL}/client/${id}`)
    const client = await res.json()
    const url = process.env.REACT_APP_BASEURL
    return { props: { client, url } }
}
