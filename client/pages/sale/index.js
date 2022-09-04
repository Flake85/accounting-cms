import Table from 'react-bootstrap/Table';
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useState } from 'react';
import { useRouter } from 'next/router';

export default function Sales({ sales, url }) {
    const router = useRouter()
    const [show, setShow] = useState(false)
    const [target, setTarget] = useState('')

    function confirmDelete(sale) { setShow(true); setTarget(sale) }
    function closeAlert() { setShow(false); setTarget("") }

    async function deleteSale() {
        try {
            const res = await fetch(`${url}/sale/${target.id}`, {
                method: 'DELETE',
                mode: 'cors'
            })
            const data = await res.json()
            if(!res.ok) throw new Error(data.error.message)
            alert("Successfully deleted sale id: ", data.data.id)
            router.reload(window.location.pathname)
        } catch(err) { err => alert(err) }
    }

    return (
        <div>
            <h1>Sales</h1>
            <Alert show={show} variant="warning" dismissible onClose={closeAlert}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ target.description }"?</p>
                <p>This action cannot be undone.</p>
                <Button onClick={deleteSale}>Confirm</Button>
                <Button onClick={closeAlert}>Cancel</Button>
            </Alert>
            { sales.data.length
                ? <Table striped bordered hover>
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Description</th>
                            <th>Client</th>
                            <th>Invoice ID</th>
                            <th>Invoice Name</th>
                            <th>Units</th>
                            <th>Unit Cost</th>
                            <th>Total</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {sales.data.map((sale, i) => (
                            <tr key={sale.id}>
                                <td>{ i + 1 }</td>
                                <td><Link href={`/sale/${sale.id}`}><a>{ sale.description }</a></Link></td>
                                {sale.client.name 
                                    ? <td><Link href={`/client/${sale.client.id}`}><a>{ sale.client.name }</a></Link></td>
                                    : <td className="text-danger"><Link href={`/client/${sale.clientId}/deleted`}><a>{ sale.clientId }</a></Link> (deleted)</td>
                                }
                                <td><Link href={`/invoice/${sale.invoiceId}`}><a>{ sale.invoiceId }</a></Link></td>
                                { sale.invoice 
                                    ? <td><Link href={`/invoice/${sale.invoiceId}`}><a>{ sale.invoice.description }</a></Link></td>
                                    : <td></td>
                                }
                                <td>{ sale.units }</td>
                                <td>{ sale.unitCost }</td>
                                <td>{ sale.total }</td>
                                { !sale.invoiceId
                                    ? <td>
                                        <Link href={`/sale/${sale.id}/update`}><a><i className="bi-pencil-square text-success"></i></a></Link>
                                        <Link href={`#`}><a onClick={() => confirmDelete(sale)}><i className="bi-trash text-danger"></i></a></Link>
                                    </td>
                                    : <td></td>
                                }
                            </tr>
                        ))}
                    </tbody>
                </Table>
                : <div><hr /><p>Labors haven't been added yet.</p></div>
            }
            <Button href="/sale/create">Add sale</Button>
        </div>
    );
}

export async function getServerSideProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/sale`)
    const sales = await res.json()
    return { props: { sales, url } }
}
