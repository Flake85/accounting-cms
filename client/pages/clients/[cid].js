function Client({ client }) {
    return (
        <div>
            <p>id: {client.id}</p>
            <p>name: {client.name}</p>
            <p>created: {client.createdAt}</p>
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const res = await fetch(`http://localhost:8080/client/${id}`)
    const client = await res.json()

    return {
        props: { client }
    }
}

export default Client