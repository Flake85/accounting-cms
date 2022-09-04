import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useState } from 'react'

export default function Sale({ url, sale }) {
    const router = useRouter()
    const [show, setShow] = useState(false)

    async function deleteSale() {
        try {
            const res = await fetch(`${url}/sale/${sale.data.id}`, {
                method: 'DELETE',
                mode: 'cors'
            })
            const data = await res.json()
            if(!res.ok) throw new Error(data.error.message)
            alert("Successfully deleted sale id: " + data.data.id)
            router.push(`/sale`)
        } catch(err) { err => alert(err) }
    }

    return (
        <div>
            <h1>Sale</h1>
            <hr />
            <Alert show={show} variant="warning" dismissible onClose={() => setShow(false)}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to delete "{ sale.data.description }"?</p>
                <Button className="me-1" onClick={deleteSale}>Confirm</Button>
                <Button onClick={() => setShow(false)}>Cancel</Button>
            </Alert>
            { sale.data
                ? <div>
                    <p><strong>Description: </strong>{sale.data.description}</p>
                    { sale.data.client.name
                        ? <p><strong>Client: </strong><Link href={`/client/${sale.data.clientId}`}><a>{sale.data.client.name}</a></Link></p>
                        : <p><strong>Client ID: </strong><Link href={`/client/${sale.data.clientId}/deleted`}><a className='text-danger'>{sale.data.clientId} (deleted)</a></Link></p>
                    }
                    { sale.data.invoiceId &&
                        <div>
                            <p><strong>Invoice Id: </strong><Link href={`/invoice/${sale.data.invoiceId}`}><a>{sale.data.invoiceId}</a></Link></p>
                            <p><strong>Invoice Name: </strong><Link href={`/invoice/${sale.data.invoiceId}`}><a>{sale.data.invoice.description}</a></Link></p>
                        </div>
                    }
                    <p><strong>Units: </strong>{sale.data.units}</p>
                    <p><strong>Unit Cost: </strong>{sale.data.unitCost}</p>
                    <p><strong>Sale Total: </strong>{sale.data.total}</p>
                    { !sale.data.invoiceId &&
                        <div>
                            <Button href={`/sale/${sale.data.id}/update`} className="me-1">Update</Button>
                            <Button onClick={()=>setShow(true)} variant="danger">Delete</Button>
                        </div>
                    }
                </div>
                : <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ sale.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.sid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/sale/${id}`)
    const sale = await res.json()
    return { props: { url, sale } }
}
