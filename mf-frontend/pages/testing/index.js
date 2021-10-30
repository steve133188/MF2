import {Dropzone} from "../../components/ImportContact";
import {NewContactForm} from "../../components/NewContactForm";
import {ContactList} from "../../components/ContactList"
import {EnhancedTable} from "../../components/Table";
import React from "react";

export default function Testing() {
    var React = require('react');
    var QRCode = require('qrcode.react');

    return (
        <QRCode value="http://facebook.github.io/react/" />
    )
}