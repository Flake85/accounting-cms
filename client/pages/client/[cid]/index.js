import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import { useState } from 'react'
import { useRouter } from 'next/router'

export default function Client({ url, client }) {
    const router = useRouter()
    const [show, setShow] = useState(false)
    
    function deleteClient() {
        fetch(`${url}/client/${client.data.id}`, {
            method: 'DELETE',
            mode: 'cors'
        })
        .then(async (res) => {
            if(res.ok) return res.json()
            const json = await res.json();
            throw new Error(json.error.message);
        })
        .then(() => {
            alert("Successfully deleted client")
            router.push("/client")
        })
        .catch(err => alert(err))
    }
    return (
        <div>
            <h1>Client</h1>
            <hr />
            <Alert show={show} variant="warning" dismissible onClose={() => setShow(false)}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ client.data.name }"?</p>
                <Button className="me-1" onClick={deleteClient}>Confirm</Button>
                <Button onClick={() => setShow(false)}>Cancel</Button>
            </Alert>
            { client.data
                ? <div>
                    <p><strong>Name: </strong>{client.data.name}</p>
                    <p><strong>Email: </strong>{client.data.email}</p>
                    <p><strong>Address: </strong>{client.data.address}</p>
                    <Button href={`/client/${client.data.id}/update`} className="me-1">Edit Client</Button>
                    <Button variant="danger" onClick={() => setShow(true)}>Delete</Button>
                </div>
                : <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ client.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/client/${id}`)
    const client = await res.json()
    return { props: { url, client } }
}
