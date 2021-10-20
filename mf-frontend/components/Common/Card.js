import Link from 'next/link'

<<<<<<< HEAD
//Import Breadcrumb
import Breadcrumb from "./Breadcrumb";



export function CardA(props) {

    const [isOpen, setIsOpen] = useState(false);

    function toggle() {
        setIsOpen(!isOpen);
    }



=======
export function Card(props) {
    return (
        <div className="card">
            <img src={props.url} className="card-img-top" alt="" />
            <div className="card-body">
                <h5 className="card-title">{props.title}</h5>
                <p className="card-text">{props.children}</p>
                <Link href={props.link}><a className="btn btn-primary">{props.btnText}</a></Link>
            </div>
        </div>
    )
}

export function Card_horizontal(props) {
    return(
        <div className="card">
            <div className="g-0 align-items-center row">
                <div className="col-md-4">
                    <img src={props.url} alt="" className="img-fluid card-img" />
                </div>
                <div className="col-md-8">
                    <div className="card-body">
                        <div className="h5 card-title">{props.title}</div>
                        <p className="card-text">{props.children}</p><p className="card-text"><small className="text-muted">{props.lastUpdateTime}</small></p></div>
                </div>
            </div>
        </div>
    )
}

export function Card_colored(props) {
    const name = "text-white-50 card bg-" + props.color;
>>>>>>> 4a5bdd37ec6f1b622aee72be57823565cc3d9b45

    return (
<<<<<<< HEAD
        <>
            <div className="card">
                    <div className="p-4" onClick={toggle} >
                        <div className="d-flex align-items-center">
                            <div className="flex-shrink-0 me-3">
                                <div className="avatar-xs">
                                    <div className="avatar-title rounded-circle bg-soft-primary text-primary">
                                        01
                                    </div>
                                </div>
                            </div>
                            <div className="flex-grow-1 overflow-hidden">
                                <h5 className="font-size-16 mb-1">Billing Info</h5>
                                <p className="text-muted text-truncate mb-0">
                                    Fill all information below
                                </p>
                            </div>
                            <div className="flex-shrink-0">
                                <i className="mdi mdi-chevron-up accor-down-icon font-size-24"></i>
                            </div>
                        </div>
                    </div>
                <div className="collapse" >
                    <div className="p-4 border-top">
                        <form>
                            <div className="mb-3">
                                <label htmlFor="productname">Product Name</label>
                                <input
                                    id="productname"
                                    name="productname"
                                    type="text"
                                    className="form-control"
                                />
                            </div>
                            <div className="row">
                                <div className="col" lg="4">
                                    <div className="mb-3">
                                        <label htmlFor="manufacturername">Manufacturer Name</label>
                                        <input
                                            id="manufacturername"
                                            name="manufacturername"
                                            type="text"
                                            className="form-control"
                                        />
                                    </div>
                                </div>
                                <div className="col" lg="4">
                                    <div className="mb-3">
                                        <label htmlFor="manufacturerbrand">
                                            Manufacturer Brand
                                        </label>
                                        <input
                                            id="manufacturerbrand"
                                            name="manufacturerbrand"
                                            type="text"
                                            className="form-control"
                                        />
                                    </div>
                                </div>

                                <div className="col" lg="4">
                                    <div className="mb-3">
                                        <label htmlFor="price">Price</label>
                                        <input
                                            id="price"
                                            name="price"
                                            type="text"
                                            className="form-control"
                                        />
                                    </div>
                                </div>
                            </div>
                            <div className="row">
                                <div className="col" md="6">
                                    <div className="mb-3">
                                        <label className="control-label">Category</label>
                                        <select className="form-control select2">
                                            <option>Select</option>
                                            <option value="AK">Alaska</option>
                                            <option value="HI">Hawaii</option>
                                        </select>
                                    </div>
                                </div>
                                <div className="col" md="6">
                                    <div className="mb-3">
                                        <label className="control-label">Specifications</label>
                                        <select
                                            classNamePrefix="select2-selection"
                                            placeholder="Choose..."
                                            title="Country"
                                            options={options}
                                            isMulti
                                        />
                                    </div>
                                </div>
                            </div>
                            <div className="mb-0">
                                <label htmlFor="productdesc">Product Description</label>
                                <textarea className="form-control" id="productdesc" rows="4" />
                            </div>
                        </form>
                    </div>
=======
        <div className="col-lg-4">
            <div className={name}>
                <div className="card-body">
                    <h5 className="mt-0 mb-4 text-white">
                        <i className="uil uil-user me-3"></i> {props.title}
                    </h5>
                    <p className="card-text">
                        {props.content}
                    </p>
>>>>>>> 4a5bdd37ec6f1b622aee72be57823565cc3d9b45
                </div>
            </div>
        </div>
    )
}
<<<<<<< HEAD
export function CardB(props) {

    const [isOpenAddproduct, setIsOpenAddproduct] = useState(false);

    const [selectedFiles, setselectedFiles] = useState([]);
    function toggleAddproduct() {
        setIsOpenAddproduct(!isOpenAddproduct);
    }
    return (
        <>
            <div className="card">
                <Link
                    href={props.url}
                    className="text-dark collapsed"
                    onClick={toggleAddproduct}
                >
                    <div className="p-4">
                        <div className="d-flex align-items-center">
                            <div className="flex-shrink-0 me-3">
                                <div className="avatar-xs">
                                    <div className="avatar-title rounded-circle bg-soft-primary text-primary">
                                        02
                                    </div>
                                </div>
                            </div>
                            <div className="flex-grow-1 overflow-hidden">
                                <h5 className="font-size-16 mb-1">Product Image</h5>
                                <p className="text-muted text-truncate mb-0">
                                    Fill all information below
                                </p>
                            </div>
                            <div className="flex-shrink-0">
                                <i className="mdi mdi-chevron-up accor-down-icon font-size-24"></i>
                            </div>
                        </div>
                    </div>
                </Link>
                <div className="collapse" isOpen={isOpenAddproduct}>
                    <div className="p-4 border-top">
                        <form>
                            <div className="dropzone"></div>
                            <div className="dropzone-previews mt-3" id="file-previews"></div>
                        </form>
                    </div>
=======

export function Card_outline(props) {
    const borderColor = "border card border-" + props.borderColor;
    const textColor = "my-0 text-" + props.textColor;
    return(
        <div className="col-lg-4">
            <div className={borderColor}>
                <div className="bg-transparent card-header">
                    <h5 className={textColor}>
                        <i className="uil uil-user me-3"></i>
                        {props.cardName}
                    </h5>
                </div>
                <div className="card-body">
                    <div className="h5 mt-0 card-title">{props.cardTitle}</div>
                    <p className="card-text">{props.cardText}</p>
>>>>>>> 4a5bdd37ec6f1b622aee72be57823565cc3d9b45
                </div>
            </div>
        </div>
    )
}
<<<<<<< HEAD
export function CardC(props) {
    const [isOpenMetadata, setIsOpenMetadata] = useState(false);

    function toggleMetadata() {
        setIsOpenMetadata(!isOpenMetadata);
    }
    return (
        <>
            <div className="card">
                <Link
                    href={props.url}
                    className="text-dark collapsed"
                    onClick={toggleMetadata}
                >
                    <div className="p-4">
                        <div className="d-flex align-items-center">
                            <div className="flex-shrink-0 me-3">
                                <div className="avatar-xs">
                                    <div className="avatar-title rounded-circle bg-soft-primary text-primary">
                                        03
                                    </div>
                                </div>
                            </div>
                            <div className="flex-grow-1 overflow-hidden">
                                <h5 className="font-size-16 mb-1">Meta Data</h5>
                                <p className="text-muted text-truncate mb-0">
                                    Fill all information below
=======

export function Card_groups() {
    return(
        <div className={"row"}>
            <div className="col-12 col">
                <h4 className="my-3">
                    Card groups
                </h4>
                <div className="card-deck-wrapper card-deck">
                    <div className="card-group">
                        <div className="mb-4 card">
                            <img src="https://lh3.googleusercontent.com/Teg9v7SxByMdqGuwDV7ssdEhVYk_3XjG2OFtYK2a9p6Xj3NL3TxkWVVAPcCVis5hTEIDL1MFQ0Bvw3LD4iLrZtzWmA=w1400-k" alt="Card image cap"
                                 className="img-fluid card-img-top" />
                            <div className="card-body">
                                <div className="h4 mt-0 card-title"></div>
                                <p className="card-text">This is a longer card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.</p>
                                <p className="card-text">
                                    <small className="text-muted">
                                        Last updated 3 mins ago
                                    </small>
>>>>>>> 4a5bdd37ec6f1b622aee72be57823565cc3d9b45
                                </p>
                            </div>
                        </div>
                        <div className="mb-4 card">

                            <img src="https://lh3.googleusercontent.com/Teg9v7SxByMdqGuwDV7ssdEhVYk_3XjG2OFtYK2a9p6Xj3NL3TxkWVVAPcCVis5hTEIDL1MFQ0Bvw3LD4iLrZtzWmA=w1400-k" alt="Card image cap"
                                 className="img-fluid card-img-top" />
                            <div className="card-body">
                                <div className="h4 mt-0 card-title">Card title</div>
                                <p className="card-text">
                                    This card has supporting text below as a natural lead-in to
                                    additional content.
                                </p>
                                <p className="card-text">
                                    <small className="text-muted">Last
                                        updated 3 mins ago
                                    </small>
                                </p>
                            </div>

                        </div>
                        <div className="mb-4 card">

                            <img src="https://lh3.googleusercontent.com/Teg9v7SxByMdqGuwDV7ssdEhVYk_3XjG2OFtYK2a9p6Xj3NL3TxkWVVAPcCVis5hTEIDL1MFQ0Bvw3LD4iLrZtzWmA=w1400-k" alt="Card image cap"
                                 className="img-fluid card-img-top" />
                            <div className="card-body">
                                <div className="h4 mt-0 card-title">Card title</div>
                                <p className="card-text">This is a wider card with supporting text below as a
                                    natural lead-in to additional content. This card has even longer content than
                                    the first to show that equal height action.
                                </p>
                                <p className="card-text">
                                    <small
                                        className="text-muted">Last updated 3 mins ago
                                    </small>
                                </p>
                            </div>

                        </div>
                    </div>
                </div>
            </div>
<<<<<<< HEAD
        </>
    );
}

=======
        </div>
    )
}
>>>>>>> 4a5bdd37ec6f1b622aee72be57823565cc3d9b45
