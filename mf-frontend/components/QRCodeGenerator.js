export function QRCodeGenerator (props) {
    var React = require('react');
    var QRCode = require('qrcode.react');

    return (
        <QRCode value={props.children} />

    )
}