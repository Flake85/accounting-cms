import Link from 'next/link'

function Clients({ clients }) {
    return (
        <ul>
            {clients.data.map((client) => (
                <li key={client.id}>
                    <Link href={`/client/`+client.id}><a>{ client.name }</a></Link>
                </li>
            ))}
        </ul>
    )
}

export async function getServerSideProps() {
    const res = await fetch(`${process.env.BASEURL}/client`)
    const clients = await res.json()

    return { props: { clients } }
}

export default Clients