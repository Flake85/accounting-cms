import Alert from 'react-bootstrap/Alert'

export default function Client({ client }) {
    return (
        <div>
            { client.data &&
                <div>
                    <p>name: {client.data.name}</p>
                    <p>email: {client.data.email}</p>
                    <p>address: {client.data.address}</p>
                </div>
            }
            { client.error.message &&
                <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ client.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const res = await fetch(`${process.env.BASEURL}/client/${id}`)
    const client = await res.json()
    return { props: { client } }
}
