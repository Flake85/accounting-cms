import Table from 'react-bootstrap/Table';
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useState } from 'react';
import { useRouter } from 'next/router';

export default function Labors({ labors, url }) {
    const router = useRouter()
    const [show, setShow] = useState(false)
    const [target, setTarget] = useState('')

    function confirmDelete(labor) { setShow(true); setTarget(labor) }
    function closeAlert() { setShow(false); setTarget("") }

    async function deleteLabor() {
        try {
            const res = await fetch(`${url}/labor/${target.id}`, {
                method: 'DELETE',
                mode: 'cors'
            })
            const data = await res.json()
            if(!res.ok) throw new Error(data.error.message)
            alert("Successfully deleted labor id: " + data.data.id)
            router.reload(window.location.pathname)
        } catch(err) { err => alert(err) }
    }
    return (
        <div>
            <h1>Labors</h1>
            <Alert show={show} variant="warning" dismissible onClose={closeAlert}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ target.description }"?</p>
                <Button onClick={deleteLabor}>Confirm</Button>
                <Button onClick={closeAlert}>Cancel</Button>
            </Alert>
            { labors.data.length
                ? <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Description</th>
                            <th>Client</th>
                            <th>Invoice ID</th>
                            <th>Invoice Name</th>
                            <th>Hours Worked</th>
                            <th>Hourly Rate</th>
                            <th>Total</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {labors.data.map((labor, i) => (
                            <tr key={labor.id}>
                                <td>{ i + 1 }</td>
                                <td><Link href={`/labor/${labor.id}`}><a>{ labor.description }</a></Link></td>
                                {labor.client.name 
                                    ? <td><Link href={`/client/${labor.client.id}`}><a>{ labor.client.name }</a></Link></td>
                                    : <td className="text-danger"><Link href={`/client/${labor.clientId}/deleted`}><a>{ labor.clientId }</a></Link> (inactive)</td>
                                }
                                <td><Link href={`/invoice/${labor.invoiceId}`}><a>{ labor.invoiceId }</a></Link></td>
                                {labor.invoice 
                                    ? <td><Link href={`/invoice/${labor.invoiceId}`}><a>{labor.invoice.description}</a></Link></td>
                                    : <td></td>
                                }
                                <td>{ labor.hoursWorked }</td>
                                <td>{ labor.hourlyRate }</td>
                                <td>{ labor.total }</td>
                                { !labor.invoiceId
                                    ? <td>
                                        <Link href={`/labor/${labor.id}/update`}><a><i className="bi-pencil-square text-success"></i></a></Link>
                                        <Link href={`#`}><a onClick={() => confirmDelete(labor)}><i className="bi-trash text-danger"></i></a></Link>
                                    </td>
                                    : <td></td>
                                }
                            </tr>
                        ))}
                    </tbody>
                </Table>
                : <div><hr /><p>Labors haven't been added yet.</p></div>
            }
            <Button href="/labor/create">Add labor</Button>
        </div>
    );
}

export async function getServerSideProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/labor`)
    const labors = await res.json()
    return { props: { labors, url } }
}
