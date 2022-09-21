import React from 'react'
import Form from "react-bootstrap/Form"
import InputGroup from 'react-bootstrap/InputGroup'
import Button from "react-bootstrap/Button"
import { useState } from "react";
import { useRouter } from "next/router"
import { useDispatch } from 'react-redux';
import { openAlertModal, setAlertData } from '../../slices/alertModalSlice';

export default function NewLabor({ clients, url }) {
    const router = useRouter()
    const dispatch = useDispatch()
    
    const [laborDescription, setLaborDescription] = useState('')
    const [laborHoursWorked, setHoursWorked] = useState(0)
    const [laborHourlyRate, setHourlyRate] = useState(0)
    const [laborClient, setLaborClient] = useState('null')

    const handleDescriptionChange = event => setLaborDescription(event.target.value)
    const handleHoursWorkedChange = event => setHoursWorked(event.target.value)
    const handleHourlyRateChange = event => setHourlyRate(event.target.value)
    const handleClientChange = event => setLaborClient(event.target.value)

    async function submitLabor(event) {
        event.preventDefault();
        var newLabor = {
            description: laborDescription,
            hoursWorked: parseFloat(laborHoursWorked),
            hourlyRate: parseFloat(laborHourlyRate),
            clientId: laborClient,
        }
        try {
            const res = await fetch(`${url}/labor`, {
                method: 'POST',
                mode: 'cors',
                body: JSON.stringify(newLabor)
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
            router.push(`/labor/${data.data.id}`)
        } catch(err) { err => alert(err) }
    }

    return (
        <div>
            <h1>Create Labor</h1>
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
                <Form.Label>Hourly Rate</Form.Label>
                <InputGroup className="mb-3">
                    <InputGroup.Text>$</InputGroup.Text>
                    <Form.Control type="number"
                                step={0.01}  
                                placeholder="Enter hourly rate"
                                value={laborHourlyRate}
                                onChange={handleHourlyRateChange} />
                </InputGroup>
                <Button type="submit" className="me-1">Submit</Button>
                <Button href={`/labor`}>Cancel</Button>
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
