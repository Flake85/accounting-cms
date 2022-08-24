import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"

export default function NewClient({ url }) {
    const router = useRouter()
    const [clientName, setClientName] = useState('')
    const [clientEmail, setClientEmail] = useState('')
    const [clientAddress, setClientAddress] = useState('')

    function handleNameChange(event) { setClientName(event.target.value) }
    function handleEmailChange(event) { setClientEmail(event.target.value) }
    function handleAddressChange(event) { setClientAddress(event.target.value) }

    const submitClient = async event => {
        event.preventDefault();
        var newClient = {
            name: clientName,
            email: clientEmail,
            address: clientAddress,
        }
        const resp = await fetch(`${url}/client`, {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(newClient)
        })
        .then(() => {
            alert("successfully submitted new client: "+ newClient.name)
            router.push('/client')
        })
        .catch(err => alert(err))
    }

    return (
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
            <Button variant="primary" type="submit">
                Submit
            </Button>
        </Form>
    )
}

export async function getStaticProps() {
    return { props: { url: process.env.REACT_APP_BASEURL } }
}
