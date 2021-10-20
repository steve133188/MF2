// Cannot import ResponsiveTable component.

export function Table() {
    return (
        <div className="table-responsive">
            <table className="table mb-0 table">
                <thead>
                <tr>
                    <th>#</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Username</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <th scope="row">1</th>
                    <td>Mark</td>
                    <td>Otto</td>
                    <td>@mdo</td>
                </tr>
                <tr>
                    <th scope="row">2</th>
                    <td>Jacob</td>
                    <td>Thornton</td>
                    <td>@fat</td>
                </tr>
                <tr>
                    <th scope="row">3</th>
                    <td>Larry</td>
                    <td>the Bird</td>
                    <td>@twitter</td>
                </tr>
                </tbody>
            </table>
        </div>
    )
}

export function DataTable() {

}

{/*
    export function EditableTable() {
        const products = [
            {id: 1, age: 25, type: "Male", name: "David McHenry"},
            {id: 2, age: 34, type: "Male", name: "Frank Kirk"},
            {id: 3, age: 67, type: "Male", name: "Rafael Morales"},
            {id: 4, age: 23, type: "Male", name: "Mark Ellison"},
            {id: 5, age: 78, type: "Female", name: "Minnie Walter"},
        ]

        const columns = [
            {
                dataField: "id",
                text: "ID",
            },
            {
                dataField: "name",
                text: "Name",
            },
            {
                dataField: "age",
                text: "Age(AutoFill)",
            },
            {
                dataField: "type",
                text: "Gender(AutoFill and Editable)",
                editor: {
                    type: Type.SELECT,
                    options: [{
                        value: 'Male',
                        label: 'Male'
                    }, {
                        value: 'Female',
                        label: 'Female'
                    }]
                }
            },
        ]
        return (
            <div className="table-responsive">
                <BootstrapTable
                    keyField="id"
                    data={products}
                    columns={columns}
                    cellEdit={cellEditFactory({mode: "click"})}
                />
            </div>
        )
    }
*/}

export function ResponsiveTable() {

    return (
        <div className="table-rep-plugin">
            <div
                className="table-responsive mb-0"
                data-pattern="priority-columns"
            >
                <Table
                    id="tech-companies-1"
                    className="table table-striped table-bordered"
                >
                    <thead>
                    <tr>
                        <th>Company</th>
                        <th data-priority="1">Last Trade</th>
                        <th data-priority="3">Trade Time</th>
                        <th data-priority="1">Change</th>
                        <th data-priority="3">Prev Close</th>
                        <th data-priority="3">Open</th>
                        <th data-priority="6">Bid</th>
                        <th data-priority="6">Ask</th>
                        <th data-priority="6">1y Target Est</th>
                    </tr>
                    </thead>
                    <tbody>
                    <tr>
                        <th>
                            GOOG <span className="co-name">Google Inc.</span>
                        </th>
                        <td>597.74</td>
                        <td>12:12PM</td>
                        <td>14.81 (2.54%)</td>
                        <td>582.93</td>
                        <td>597.95</td>
                        <td>597.73 x 100</td>
                        <td>597.91 x 300</td>
                        <td>731.10</td>
                    </tr>
                    <tr>
                        <th>
                            AAPL <span className="co-name">Apple Inc.</span>
                        </th>
                        <td>378.94</td>
                        <td>12:22PM</td>
                        <td>5.74 (1.54%)</td>
                        <td>373.20</td>
                        <td>381.02</td>
                        <td>378.92 x 300</td>
                        <td>378.99 x 100</td>
                        <td>505.94</td>
                    </tr>
                    <tr>
                        <th>
                            AMZN{" "}
                            <span className="co-name">Amazon.com Inc.</span>
                        </th>
                        <td>191.55</td>
                        <td>12:23PM</td>
                        <td>3.16 (1.68%)</td>
                        <td>188.39</td>
                        <td>194.99</td>
                        <td>191.52 x 300</td>
                        <td>191.58 x 100</td>
                        <td>240.32</td>
                    </tr>

                    </tbody>
                </Table>
            </div>
        </div>
    )
}
