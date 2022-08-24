import Alert from 'react-bootstrap/Alert'

export default function Expense({ expense }) {
    return (
        <div>
            { expense.data &&
                <div>
                    <p>description: {expense.data.description}</p>
                    <p>cost: {expense.data.cost}</p>
                </div>
            }
            { expense.error.message &&
                <Alert variant="danger">
                    <Alert.Heading>Error</Alert.Heading>
                    <hr />
                    <p>{ expense.error.message }</p>
                </Alert>
            }
        </div>
    )
}

export async function getServerSideProps(context) {
    const id = context.query.eid
    const res = await fetch(`${process.env.REACT_APP_BASEURL}/expense/${id}`)
    const expense = await res.json()
    return { props: { expense } }
}
