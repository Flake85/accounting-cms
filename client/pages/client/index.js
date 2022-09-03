import Table from 'react-bootstrap/Table';
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useState } from 'react';
import { useRouter } from 'next/router';

export default function Clients({ clients, url }) {
    const router = useRouter()
    const [show, setShow] = useState(false)
    const [target, setTarget] = useState('')

    function confirmDelete(client) {
        setShow(true)
        setTarget(client)
    }
    function closeAlert() {
        setShow(false)
        setTarget("")
    }
    function deleteClient() {
        fetch(`${url}/client/${target.id}`, {
            method: 'DELETE',
            mode: 'cors'
        })
        .then(async (res) => {
            if(res.ok) return res.json()
            const json = await res.json();
            throw new Error(json.error.message);
        })
        .then(() => {
            alert("Successfully deleted ", target.name)
            router.reload(window.location.pathname)
        })
        .catch(err => alert(err))
    }
    return (
        <div>
            <h1>Clients</h1>
            { clients.data.length
                ? <div>
                    <Alert show={show} variant="warning" dismissible onClose={closeAlert}>
                        <Alert.Heading>Warning</Alert.Heading>
                        <p>Are you sure you want to delete "{ target.name }"?</p>
                        <Button className="me-1" onClick={deleteClient}>Confirm</Button>
                        <Button onClick={closeAlert}>Cancel</Button>
                    </Alert>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th>#</th>
                                <th>Client</th>
                                <th>Address</th>
                                <th>Email</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {clients.data.map((client, i) => (
                                <tr key={client.id}>
                                    <td>{ i + 1 }</td>
                                    <td><Link href={`/client/${client.id}`}><a>{ client.name }</a></Link></td>
                                    <td>{ client.address }</td>
                                    <td>{ client.email }</td>
                                    <td>
                                        <Link href={`/client/${client.id}/update`}><a><i className="bi-pencil-square text-success"></i></a></Link>
                                        <Link href={`#`}><a onClick={() => confirmDelete(client)}><i className="bi-trash text-danger"></i></a></Link>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </Table>
                </div>
                : <div>
                    <hr />
                    <p>No Clients added.</p>
                </div>
            }
            <Button href="/client/create">Add Client</Button>
        </div>
    );
}

export async function getServerSideProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/client`)
    const clients = await res.json()
    return { props: { clients, url } }
}
