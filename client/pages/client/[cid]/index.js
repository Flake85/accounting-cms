import React from 'react'
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import { useState } from 'react'
import { useRouter } from 'next/router'
import { useDispatch } from 'react-redux'
import { setAlertData, openAlertModal } from '../../../slices/alertModalSlice'

export default function Client({ url, client }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [show, setShow] = useState(false)
    
    async function deleteClient() {
        try {
            const res = await fetch(`${url}/client/${client.data.id}`, {
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
            router.push(`/client/${data.data.id}/deleted`)
        } catch(err) { err => console.log(err) }
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
