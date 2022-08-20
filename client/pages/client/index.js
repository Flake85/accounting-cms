import ListGroup from 'react-bootstrap/ListGroup';
import Link from 'next/link'

export default function Clients({ clients }) {
  return (
        <ListGroup as="ul" numbered>
            {clients.data.map((client) => (
                <ListGroup.Item as="li" action key="client.id">
                    <Link href={`/client/`+client.id}><a>{ client.name }</a></Link>
                </ListGroup.Item>
            ))}
        </ListGroup>
  );
}

export async function getServerSideProps() {
    const res = await fetch(`${process.env.BASEURL}/client`)
    const clients = await res.json()

    return { props: { clients } }
}
