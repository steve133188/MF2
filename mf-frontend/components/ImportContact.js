import React, {useMemo} from 'react';
import {useDropzone} from 'react-dropzone';
import {NormalButton, NormalButton2, CancelButton} from './Button';

export function Dropzone(props) {
    const {
        getRootProps,
        getInputProps,
        isDragActive,
        isDragAccept,
        isDragReject,
        acceptedFiles
    } = useDropzone();

    const baseStyle = {
        flex: 1,
        maxWidth: 592,
        maxHeight: 202,
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        paddingTop: '42px',
        paddingRight: '95.5px',
        paddingBottom: '48px',
        paddingLeft: '95.5px',
        borderWidth: 2,
        borderRadius: 10,
        borderColor: '#2385FC',
        borderStyle: 'dashed',
        backgroundColor: '#ECF2F8',
        color: '#bdbdbd',
        outline: 'none',
        transition: 'border .24s ease-in-out'
    };

    const files = acceptedFiles.map(file => (
        <li key={file.path}>
            {file.path} - {file.size} bytes
        </li>
    ));

    const activeStyle = {
        borderColor: '#2385FC'
    };

    const acceptStyle = {
        borderColor: '#48CEFF'
    };

    const rejectStyle = {
        borderColor: '#ff1744'
    };


    const style = useMemo(() => ({
        ...baseStyle,
        ...(isDragActive ? activeStyle : {}),
        ...(isDragAccept ? acceptStyle : {}),
        ...(isDragReject ? rejectStyle : {})
    }), [
        isDragActive,
        isDragReject,
        isDragAccept
    ]);
    return (
        <div className={"container importContactCenter"}>
            <div className={"importContactContainer"}>
                <div className="header">
                    <span>Import Contacts</span>
                    <div className="buttonGrp">
                        <NormalButton2>Confirm</NormalButton2>
                        <span style={{marginLeft: "30px"}}><CancelButton>Cancel</CancelButton></span>
                    </div>
                </div>
                <div {...getRootProps({style})}>
                    <input {...getInputProps()} />
                    <svg xmlns="http://www.w3.org/2000/svg" width="48" height="60" fill="#2198FA"
                         className="bi bi-file-earmark-arrow-up" viewBox="0 0 16 16">
                        <path
                            d="M8.5 11.5a.5.5 0 0 1-1 0V7.707L6.354 8.854a.5.5 0 1 1-.708-.708l2-2a.5.5 0 0 1 .708 0l2 2a.5.5 0 0 1-.708.708L8.5 7.707V11.5z"/>
                        <path
                            d="M14 14V4.5L9.5 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zM9.5 3A1.5 1.5 0 0 0 11 4.5h2V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.5v2z"/>
                    </svg>
                    <p style={{marginTop: "30px", color: "#444444"}}>Drag and drop a file or <span
                        style={{fontWeight: "600", color: "#2198FA", cursor: "pointer"}}>browse</span> to upload your
                        contacts.</p>
                </div>
                <aside style={{marginTop: "38px"}}>
                    <h4 style={{fontSize: "12px", color: "#444444"}}>Uploaded</h4>
                    <ul style={{
                        paddingLeft: "0px",
                        listStyle: "none",
                        marginTop: "17px",
                        fontSize: "16px",
                        color: "#444444"
                    }}>{files}</ul>
                </aside>
                <NormalButton>Download Template</NormalButton>
            </div>
        </div>
    );
}