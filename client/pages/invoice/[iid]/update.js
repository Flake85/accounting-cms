import { useState } from "react"
import { useRouter } from "next/router"
import Form from 'react-bootstrap/Form'
import Button from "react-bootstrap/Button"

export default function UpdateInvoice({ invoice, url }) {
const router = useRouter()
const [invoiceIsPaid, setInvoiceIsPaid] = useState(invoice.data.isPaid)
console.log(invoiceIsPaid)

const handleIsPaidChange = event => setInvoiceIsPaid(event.target.checked)

const submitInvoice = async event => {
    event.preventDefault();
    var updatedInvoice = {
        clientId: invoice.data.clientId,
        isPaid: invoiceIsPaid,
    }
    await fetch(`${url}/invoice/${invoice.data.id}`, {
        method: 'PUT',
        mode: 'cors',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(updatedInvoice)
    })
    .then(async (res) => {
        if(res.ok) return res.json()
        const json = await res.json();
        throw new Error(json.error.message);
    })
    .then(() => {
        alert("successfully updated invoice: " + invoice.data.description)
        router.push(`/invoice/${invoice.data.id}`)
    })
    .catch(err => alert(err))
}

return (
    <div>
        <h1>Update Invoice</h1> 
        <hr />
        <Form onSubmit={submitInvoice}>
            <Form.Group className="mb-3" controlId="formBasicCheckbox">
                    <Form.Check value={invoiceIsPaid} checked={invoiceIsPaid} type="checkbox" label="Is Paid in Full" onChange={handleIsPaidChange} />
                </Form.Group>
            <Button variant="primary" type="submit" className="me-1">Submit</Button>
            <Button href={`/invoice`}>Cancel</Button>
        </Form>
    </div>
)
}

export async function getServerSideProps(context) {
const id = context.query.iid
const res = await fetch(`${process.env.REACT_APP_BASEURL}/invoice/${id}`)
const invoice = await res.json()
const url = process.env.REACT_APP_BASEURL
return { props: { invoice, url } }
}
