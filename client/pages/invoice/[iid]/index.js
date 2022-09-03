import Alert from 'react-bootstrap/Alert'
import Row from 'react-bootstrap/Row'
import Col from 'react-bootstrap/Col'
import Table from 'react-bootstrap/Table'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useRouter } from 'next/router'

export default function Invoice({ url, invoice }) {
    const router = useRouter()
    const submitInvoicePaid = async event => {
        event.preventDefault();
        var updatedPaidInvoice = {
            isPaid: true,
        }
        await fetch(`${url}/invoice/${invoice.data.id}`, {
            method: 'PUT',
            mode: 'cors',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(updatedPaidInvoice)
        })
        .then(async (res) => {
            if(res.ok) return res.json()
            const json = await res.json();
            throw new Error(json.error.message);
        })
        .then(() => {
            alert("successfully updated invoice: " + invoice.data.description + " as paid")
            router.reload(window.location.pathname)
        })
        .catch(err => alert(err))
    }
    return (
        <div>
            <h1>Invoice</h1>
            <hr />
            { invoice.data
                ? <div>
                    <p><strong>Invoice Description: </strong>{invoice.data.description}</p>
                    { invoice.data.client.name
                        ? <p><strong>Client: </strong><Link href={`/client/${invoice.data.clientId}`}><a>{invoice.data.client.name}</a></Link></p>
                        : <p><strong>Client ID: </strong><Link href={`/client/${invoice.data.clientId}/deleted`}><a className='text-danger'>{invoice.data.clientId} (deleted)</a></Link></p>
                    }
                    <p><strong>Is Paid: </strong>{invoice.data.isPaid.toString()}</p>
                    <hr />
                    <Row>
                        <Col>
                            <h3>Sales</h3>
                                { invoice.data.sales.length 
                                    ? <div>
                                        <Table striped bordered hover>
                                            <thead>
                                                <tr>
                                                    <th>#</th>
                                                    <th>Description</th>
                                                    <th>Units</th>
                                                    <th>Unit Cost</th>
                                                    <th>Total</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {invoice.data.sales.map((sale, i) => (
                                                    <tr key={sale.id}>
                                                        <td>{ i + 1 }</td>
                                                        <td><Link href={`/sale/${sale.id}`}><a>{ sale.description }</a></Link></td>
                                                        <td>{ sale.units }</td>
                                                        <td>${ sale.unitCost }</td>
                                                        <td>${ sale.total }</td>
                                                    </tr>
                                                ))}
                                                <tr>
                                                    <td>Total Sales:</td>
                                                    <td></td>
                                                    <td></td>
                                                    <td></td>
                                                    <td>${invoice.data.salesTotal}</td>
                                                </tr>
                                            </tbody>
                                        </Table>
                                    </div>
                                    : <p>No Sales Added</p>
                                }
                        </Col>
                        <Col>
                            <h3>Labor</h3>
                            { invoice.data.labors.length 
                                    ? <div>
                                        <Table striped bordered hover>
                                            <thead>
                                                <tr>
                                                    <th>#</th>
                                                    <th>Description</th>
                                                    <th>Hours Worked</th>
                                                    <th>Hourly Rate</th>
                                                    <th>Total</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {invoice.data.labors.map((labor, i) => (
                                                    <tr key={labor.id}>
                                                        <td>{ i + 1 }</td>
                                                        <td><Link href={`/labor/${labor.id}`}><a>{ labor.description }</a></Link></td>
                                                        <td>{ labor.hoursWorked }</td>
                                                        <td>${ labor.hourlyRate }</td>
                                                        <td>${ labor.total }</td>
                                                    </tr>
                                                ))}
                                                <tr>
                                                    <td>Total Labor:</td>
                                                    <td></td>
                                                    <td></td>
                                                    <td></td>
                                                    <td>${invoice.data.laborsTotal}</td>
                                                </tr>
                                            </tbody>
                                        </Table>
                                    </div>
                                    : <p>No Sales Added</p>
                                }
                        </Col>
                    </Row>
                    <p><strong>Total billed: </strong>${invoice.data.grandTotal}</p>
                    {!invoice.data.isPaid && <Button onClick={submitInvoicePaid}>Invoice Paid</Button>}
                </div>
                : <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ invoice.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.iid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/invoice/${id}`)
    const invoice = await res.json()
    return { props: { url, invoice } }
}
