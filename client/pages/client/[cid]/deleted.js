import { useState } from 'react'
import { useRouter } from 'next/router'
import Alert from 'react-bootstrap/Alert'
import Button from 'react-bootstrap/Button'
// import Form from 'react-bootstrap/Form'

export default function Client({ client, url }) {
    const router = useRouter()
    // const [confirm, setConfirm] = useState("")
    const [showUndelete, setShowUndelete] = useState(false)
    // const [showPermDelete, setShowPermDelete] = useState(false)
    const undelete = () => setShowUndelete(true)
    const closeUndelete = () => setShowUndelete(false)
    // const permDelete = () => setShowPermDelete(true)
    // const closePermDelete = () => setShowPermDelete(false)

    // function handleConfirmChange(event) { setConfirm(event.target.value) }

    // function confirmed() {
    //     return confirm === "delete " + client.data.name ? false : true
    // }

    // function confirmPermaDelete() {
    //     fetch(`${url}/client/delete/${client.data.id}`, {
    //         method: 'DELETE',
    //         mode: 'cors'
    //     })
    //     .then(() => {
    //         alert("Successfully deleted ", client.data.name)
    //         router.push("/client")
    //     })
    //     .catch(err => alert(err))
    // }

    async function confirmUndelete() {
        try {
            const res = await fetch(`${url}/client_deleted/${client.data.id}`, {
                method: 'PUT',
                mode: 'cors'
            })
            const data = await res.json()
            if(!res.ok) throw new Error(data.error.message)
            alert("Successfully un-deleted client id: " + data.data.id)
            router.push('/client/' + data.data.id)
        } catch(err) { err => alert(err) }
    }

    return (
        <div>
            <Alert show={showUndelete} variant="warning" dismissible onClose={closeUndelete}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to un-delete "{ client.data.name }"?</p>
                <Button type="button" className='me-1' onClick={()=>confirmUndelete()}>Confirm</Button>
                <Button type="button" onClick={closeUndelete}>Cancel</Button>
            </Alert>
            {/* <Alert show={showPermDelete} variant="danger" dismissible onClose={closePermDelete}>
                <Alert.Heading>Warning</Alert.Heading>
                <p>Are you sure you want to permanently delete "{ client.data.name }"?</p>
                <p>
                    This action is irreversable and will deleted all associated client "{client.data.name}" data. <br />
                    That includes sales, labor, and invoice data associated with "{client.data.name}."
                </p>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Type "delete {client.data.name}" to delete client from the database</Form.Label>
                        <Form.Control type="email" onChange={handleConfirmChange} placeholder={"Enter text: 'delete " + client.data.name + "' to enable confirm button" }/>
                        <Form.Text className="text-muted">
                        Confirm button will work after entering correct text.
                        </Form.Text>
                    </Form.Group>
                </Form>
                <Button type="button" variant="danger" onClick={()=>confirmPermaDelete()} disabled={confirmed()} className='me-1' >Confirm</Button>
                <Button type="button" onClick={closePermDelete}>Cancel</Button>
            </Alert> */}
            <h1 className='text-danger'>Deleted Client</h1>
            <hr className='text-danger' />
            { client.data 
                ? <div>
                    <p>Name: {client.data.name}</p>
                    <p>Email: {client.data.email}</p>
                    <p>Address: {client.data.address}</p>
                    <p>Deleted At: {client.data.deletedAt}</p>
                </div>
                : <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ client.error.message }</p>
                </Alert>
            }
            <Button type='button' className='me-1' onClick={()=>undelete()}>Un-Delete</Button>
            {/* <Button variant="danger" onClick={()=>permDelete()}>Permanently Delete</Button> */}
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const url = process.env.REACT_APP_BASEURL
    const res = await fetch(`${url}/client_deleted/${id}`)
    const client = await res.json()
    return { props: { client, url } }
}
