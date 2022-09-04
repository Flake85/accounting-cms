import Form from "react-bootstrap/Form"
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"

export default function updatedLabor({ labor, clients, url }) {
    const router = useRouter()
    const [laborDescription, setLaborDescription] = useState(labor.data.description)
    const [laborHoursWorked, setHoursWorked] = useState(labor.data.hoursWorked)
    const [laborHourlyRate, setHourlyRate] = useState(labor.data.hourlyRate)
    const [laborClient, setLaborClient] = useState(labor.data.clientId)

    const handleDescriptionChange = event => setLaborDescription(event.target.value)
    const handleHoursWorkedChange = event => setHoursWorked(event.target.value)
    const handleHourlyRateChange = event => setHourlyRate(event.target.value)
    const handleClientChange = event => setLaborClient(event.target.value)

    async function submitLabor(event) {
        event.preventDefault();
        var updatedLabor = {
            description: laborDescription,
            hoursWorked: parseFloat(laborHoursWorked),
            hourlyRate: parseFloat(laborHourlyRate),
            clientId: laborClient,
        }
        try {
            const res = await fetch(`${url}/labor/${labor.data.id}`, {
                method: 'PUT',
                mode: 'cors',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(updatedLabor)
            })
            const data = await res.json()
            if(!res.ok) throw new Error(data.error.message)
            alert("Successfully updated labor: " + data.data.description)
            router.push(`/labor/${data.data.id}`)
        } catch(err) { err => alert(err) }
    }

    return (
        <div>
            <h1>Update Labor</h1>
            <hr />
            <Form onSubmit={submitLabor}>
                <Form.Group className="mb-3">
                    <Form.Label>Description</Form.Label>
                    <Form.Control placeholder="Enter labor description" 
                                value={laborDescription}
                                onChange={handleDescriptionChange}/>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Client</Form.Label>
                    <Form.Select value={laborClient} 
                                placeholder="select a client"
                                onChange={handleClientChange}>
                        <option value="null"></option>
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
                                value={laborHoursWorked}
                                onChange={handleHoursWorkedChange} />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Hourly Rate</Form.Label>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter hourly rate"
                                value={laborHourlyRate}
                                onChange={handleHourlyRateChange} />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.lid
    const url = process.env.REACT_APP_BASEURL
    const lRes = await fetch(`${url}/labor/${id}`)
    const labor = await lRes.json()
    const cRes = await fetch(`${url}/client`)
    const clients = await cRes.json()
    return { props: { labor, clients, url } }
}
