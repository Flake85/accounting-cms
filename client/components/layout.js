import Navbar from "./navbar";
import Footer from "./footer";

export default function Layout({ children }) {
    return (
        <div>
            <Navbar />
            <main className="container">{ children }</main>
            <Footer />
        </div>
    )
}