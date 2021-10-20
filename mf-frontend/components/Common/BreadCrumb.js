import React from "react"
import PropTypes from 'prop-types'
import Link from 'next/link'

const Breadcrumb = props => {
  return (
    <div className="row">
      <div className="col-12">
        <div className="page-title-box d-flex align-items-center justify-content-between">
          <h4 className="mb-0">{props.head}</h4>
          
          <div className="page-title-right">
            <ol className="breadcrumb m-0">
              <div>
                <Link href={props.url}>{props.title}</Link>
              </div>
              <div active>
                <Link href={props.url}>{props.title}</Link>
              </div>
            </ol>
          </div>
        </div>
      </div>
    </div>
  )
}

Breadcrumb.propTypes = {
  breadcrumbItem: PropTypes.string,
  title: PropTypes.string
}

export default Breadcrumb
