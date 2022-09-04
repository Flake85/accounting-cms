import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

export default function NavigationBar() {
  return (
    <Navbar bg="light" expand="lg" sticky="top">
      <Container>
        <Navbar.Brand href="/">Accounting CMS</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="ms-auto">
            <Nav.Link href="/">Dashboard</Nav.Link>
            <Nav.Link href="/client">Clients</Nav.Link>
            <Nav.Link href="/invoice">Invoices</Nav.Link>
            <Nav.Link href="/labor">Labor</Nav.Link>
            <Nav.Link href="/sale">Sales</Nav.Link>
            <Nav.Link href="/expense">Expenses</Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
