import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"

export default function NewSale({ clients, url }) {
    const router = useRouter()
    const [saleDescription, setSaleDescription] = useState('')
    const [saleUnits, setUnits] = useState(0)
    const [saleUnitCost, setUnitCost] = useState(0)
    const [saleClient, setSaleClient] = useState('null')

    const handleDescriptionChange = event => setSaleDescription(event.target.value)
    const handleUnitsChange = event => setUnits(event.target.value)
    const handleUnitCostChange = event => setUnitCost(event.target.value)
    const handleClientChange = event => setSaleClient(event.target.value)

    const submitSale = async event => {
        event.preventDefault();
        var newSale = {
            description: saleDescription,
            units: parseFloat(saleUnits),
            unitCost: parseFloat(saleUnitCost),
            clientId: saleClient,
        }
        await fetch(`${url}/sale`, {
            method: 'POST',
            mode: 'cors',
            body: JSON.stringify(newSale)
        })
        .then(async (res) => {
            if(res.ok) return res.json()
            const json = await res.json();
            throw new Error(json.error.message);
        })
        .then(() => {
            alert("successfully submitted new sale: " + newSale.description)
            router.push('/sale')
        })
        .catch(err => alert(err))
    }

    return (
        <div>
            <h1>Create Sale</h1>
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
                                onChange={handleClientChange}>
                        <option value="null" disabled>Select a Client</option>
                        {clients.data.map((client) => (
                            <option key={client.id} value={client.id}>{ client.name }</option>
                        ))}
                    </Form.Select>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Unit Amount</Form.Label>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter Units"
                                value={saleUnits}
                                onChange={handleUnitsChange} />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Unit Cost</Form.Label>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter Unit Cost"
                                value={saleUnitCost}
                                onChange={handleUnitCostChange} />
                </Form.Group>
                <Button type="submit" className="me-1">Submit</Button>
                <Button href={`/sale`}>Cancel</Button> 
            </Form>
        </div>
    )
}

export async function getServerSideProps() {
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/client`)
    const clients = await res.json()
    return { props: { clients, url } }
}
