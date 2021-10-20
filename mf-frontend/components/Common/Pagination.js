import Link from 'next/link'

export function Pagination() {
    return (
        <nav aria-label="Page navigation example">
            <ul className="pagination">
                <li className="page-item">
                    <Link href="#"><a className="page-link">Previous</a></Link>
                </li>
                <li className="page-item">
                    <Link href="#"><a className="page-link">1</a></Link>
                </li>
                <li className="page-item">
                    <Link href="#"><a className="page-link">2</a></Link>
                </li>
                <li className="page-item">
                    <Link href="#"><a className="page-link">3</a></Link>
                </li>
                <li className="page-item">
                    <Link href="#"><a className="page-link">Next</a></Link>
                </li>
            </ul>
        </nav>
    )
}