function Client({ client }) {
    return (
        <div>
            <p>name: {client.data.name}</p>
            <p>email: {client.data.email}</p>
            <p>address: {client.data.address}</p>
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.cid
    const res = await fetch(`${process.env.HOST}/client/${id}`)
    const client = await res.json()

    if (client.error.message) {
        return {
            redirect: {
                destination: "/404"
            }
        }
    }
    return {
        props: { client }
    }
}

export default Client