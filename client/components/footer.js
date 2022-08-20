import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';

export default function Footer() {
    return (
        <Navbar expand="lg" variant="light" bg="light" fixed="bottom">
            <Container>
                <Navbar.Brand href="#" className="mx-auto">Bentley Accounting CMS { new Date().getFullYear() }</Navbar.Brand>
            </Container>
        </Navbar>
    );
}
