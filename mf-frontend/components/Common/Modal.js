export function Model() {
    return (
        <div
            className="modal bs-example-modal"
            tabIndex="-1"
            role="dialog"
        >
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title">Modal title</h5>
                        <button
                            type="button"
                            className="btn-close"
                            data-dismiss="modal"
                            aria-label="Close"
                        >
                        </button>
                    </div>
                    <div className="modal-body">
                        <p>One fine body</p>
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-primary">
                            Save changes
                        </button>
                        <button
                            type="button"
                            className="btn btn-light"
                            data-dismiss="modal"
                        >
                            Close
                        </button>
                    </div>
                </div>
            </div>
        </div>
    )
}