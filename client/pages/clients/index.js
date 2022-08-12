import Link from 'next/link'

function Clients({ clients }) {
    return (
        <ul>
            {clients.data.map((client) => (
                <li key={client.id}>
                    <Link href={`/clients/`+client.id}><a>{ client.name }</a></Link>
                </li>
            ))}
        </ul>
    )
}

export async function getServerSideProps() {
    const res = await fetch('http://localhost:8080/client')
    const clients = await res.json()

    return {
        props: { clients }
    }
}

export default Clients