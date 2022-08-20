import Navbar from "./navbar";
import Footer from "./footer";
import Container from "react-bootstrap/Container";

export default function Layout({ children }) {
    return (
        <div>
            <Navbar />
            <Container>
                <main>{ children }</main>
            </Container>
            <Footer />
        </div>
    )
}
