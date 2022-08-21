import Card from "react-bootstrap/Card"
import Row from "react-bootstrap/Row"
import Col from "react-bootstrap/Col"

export default function Home() {
    return (
      <Card className="text-center">
        <Card.Header>Dashboard</Card.Header>
        <Card.Body>
            <Row>
              <Col>
                <Card.Link href="/client">
                  <h3><i class="bi-people"></i></h3>
                  <h4>Clients</h4>
                </Card.Link>
              </Col>
              <Col>
                <Card.Link href="#">
                  <h3><i class="bi-receipt"></i></h3>
                  <h4>Invoices</h4>
                </Card.Link>
              </Col>
              <Col>
                <Card.Link href="#">
                  <h3><i class="bi-hammer"></i></h3>
                  <h4>Labor</h4>
                </Card.Link>
              </Col>
              <Col>
                <Card.Link href="#">
                  <h3><i class="bi-receipt-cutoff"></i></h3>
                  <h4>Sales</h4>
                </Card.Link>
              </Col>
              <Col>
                <Card.Link href="#">
                  <h3><i class="bi-cash-coin"></i></h3>
                  <h4>Expenses</h4>
                </Card.Link>
              </Col>
            </Row>
        </Card.Body>
      </Card>
    );
}
