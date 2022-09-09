import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"
import { useDispatch } from 'react-redux'
import { setAlertData, openAlertModal } from '../../../slices/alertModalSlice'

export default function updatedSale({ sale, clients, url }) {
    const router = useRouter()
    const dispatch = useDispatch()

    const [saleDescription, setSaleDescription] = useState(sale.data.description)
    const [saleUnits, setUnits] = useState(sale.data.units)
    const [saleUnitCost, setUnitCost] = useState(sale.data.unitCost)
    const [saleClient, setSaleClient] = useState(sale.data.clientId)

    const handleDescriptionChange = event => setSaleDescription(event.target.value)
    const handleUnitsChange = event => setUnits(event.target.value)
    const handleUnitCostChange = event => setUnitCost(event.target.value)
    const handleClientChange = event => setSaleClient(event.target.value)



    async function submitSale(event) {
        event.preventDefault();
        var updatedSale = {
            description: saleDescription,
            units: parseFloat(saleUnits),
            unitCost: parseFloat(saleUnitCost),
            clientId: saleClient,
        }
        try {
            const res = await fetch(`${url}/sale/${sale.data.id}`, {
                method: 'PUT',
                mode: 'cors',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updatedSale)
            })
            const data = await res.json()
            if(!res.ok) {
                dispatch(setAlertData({ 
                    title: 'Something went wrong', 
                    body: 'Error: ' + data.error.message
                }))
                dispatch(openAlertModal())
                throw new Error(data)
            }
            router.push(`/sale/${data.data.id}`)
        } catch(err) { err => console.log(err) }
    }

    return (
        <div>
            <h1>Update Sale</h1>
            <hr />
            <Form onSubmit={submitSale}>
                <Form.Group className="mb-3">
                    <Form.Label>Description</Form.Label>
                    <Form.Control placeholder="Enter sale description" 
                                value={saleDescription}
                                onChange={handleDescriptionChange}/>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Client</Form.Label>
                    <Form.Select value={saleClient} 
                                placeholder="select a client"
                                onChange={handleClientChange}>
                        <option value="null">Select a Client</option>
                        {clients.data.map((client) => (
                            <option key={client.id} value={client.id}>{client.name}</option>
                        ))}
                    </Form.Select>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Hours Worked</Form.Label>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter hours worked"
                                value={saleUnits}
                                onChange={handleUnitsChange} />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Hourly Rate</Form.Label>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter hourly rate"
                                value={saleUnitCost}
                                onChange={handleUnitCostChange} />
                </Form.Group>
                <Button type="submit" className="me-1">Submit</Button>
                <Button href={`/sale/${sale.data.id}`}>Cancel</Button>
            </Form>
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.sid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/sale/${id}`)
    const sale = await res.json()
    const cRes = await fetch(`${url}/client`)
    const clients = await cRes.json()
    return { props: { sale, clients, url } }
}
