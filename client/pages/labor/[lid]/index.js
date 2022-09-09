import React from 'react'
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useState } from 'react'
import { useDispatch } from 'react-redux'
import { openAlertModal, setAlertData } from '../../../slices/alertModalSlice'

export default function Labor({ url, labor }) {
    const router = useRouter()
    const dispatch = useDispatch()
    const [show, setShow] = useState(false)

    async function deleteLabor() {
        try {
            const res = await fetch(`${url}/labor/${labor.data.id}`, {
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
            router.push(`/labor`)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Labor</h1>
            <hr />
            <Alert show={show} variant="warning" dismissible onClose={() => setShow(false)}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ labor.data.description }"?</p>
                <Button className="me-1" onClick={deleteLabor}>Confirm</Button>
                <Button onClick={() => setShow(false)}>Cancel</Button>
            </Alert>
            { labor.data
                ? <div>
                    <p><strong>Description: </strong>{labor.data.description}</p>
                    { labor.data.client.name
                        ? <p><strong>Client Name: </strong><Link href={`/client/${labor.data.clientId}`}><a>{labor.data.client.name}</a></Link></p>
                        : <p><strong>Client ID: </strong><Link href={`/client/${labor.data.clientId}/deleted`}><a className='text-danger'>{labor.data.clientId} (deleted)</a></Link></p>
                    }
                    { labor.data.invoiceId &&
                        <div>
                            <p><strong>Invoice Id: </strong><Link href={`/invoice/${labor.data.invoiceId}`}><a>{labor.data.invoiceId}</a></Link></p>
                            <p><strong>Invoice Name: </strong><Link href={`/invoice/${labor.data.invoiceId}`}><a>{labor.data.invoice.description}</a></Link></p>
                        </div>
                    }
                    <p><strong>Hourly Rate: </strong>{labor.data.hourlyRate}</p>
                    <p><strong>Hours Worked: </strong>{labor.data.hoursWorked}</p>
                    <p><strong>Labor Total: </strong>{labor.data.total}</p>
                    { !labor.data.invoiceId &&
                        <div>
                            <Button href={`/labor/${labor.data.id}/update`} className="me-1">Update</Button>
                            <Button onClick={()=>setShow(true)} variant="danger">Delete</Button>
                        </div>
                    }
                </div>
                : <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ labor.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.lid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/labor/${id}`)
    const labor = await res.json()
    return { props: { url, labor } }
}
