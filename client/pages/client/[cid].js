import Alert from 'react-bootstrap/Alert'

export default function Client({ client }) {
    // change to ternary 
    { if(client.data) {return (
        <div>
            <p>name: {client.data.name}</p>
            <p>email: {client.data.email}</p>
            <p>address: {client.data.address}</p>
        </div>
        )
    }}
    return (
        <Alert variant="danger">
            <Alert.Heading>Error</Alert.Heading>
            <hr />
            <p>{ client.error.message }</p>
        </Alert>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const res = await fetch(`${process.env.BASEURL}/client/${id}`)
    const client = await res.json()

    // change this up. its overkill
    return res.status === 404 ? { redirect: { destination: "/404" } }
         : res.status === 500 ? { redirect: { destination: "/500" } }
         : { props: {client} }
}
