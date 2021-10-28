import React, {useState} from 'react';
import {ChatBubble} from "./ChatBubble";
import TextField from '@mui/material/TextField';
import {IconButton} from "./Button";
import {LabelSelect} from "./Select";
import {ContactType} from "./ContactType";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import Button from "@mui/material/Button";

export function ChatArea() {
    const [value, setValue] = React.useState('Controlled');
    const handleChange = (event) => {
        setValue(event.target.value);
    };

    const theme = createTheme({
        palette: {
            neutral: {
                main: '#DEF0FF',
                contrastText: '#2198FA',
            },
            primary: {
                main: '#2198FA',
                contrastText: 'white',
            },
            secondary: {
                main: '#F1B44C',
                contrastText: 'white',
            },
            cancel: {
                main: '#F5F6F8',
                contrastText: '#444444',
            }
        },
    });

    const isSent = false;

    return (

        <div className="chatArea">
            <div className="chatAreaTop">
                <div className="targetContact">
                    <img className="targetIcon"
                         src="https://ath2.unileverservices.com/wp-content/uploads/sites/4/2020/02/IG-annvmariv-1024x1016.jpg"
                         alt=""/>
                    <span className="targetName">Debra Patel</span>
                    <ContactType/>
                </div>
                <div className="buttonGrp">
                    <IconButton/>
                    <IconButton/>
                    <LabelSelect/>
                </div>
            </div>
            <div className="chatBubbleContainer">
                <div className="messageDate">3 Jun 2021, 08:56 PM</div>
                <div className="chatBubbleGroup">
                    <ChatBubble type={isSent ? "sent" : "received"} messageTime="02:00 PM">Hello there. I am Ben Chow,
                        nice to
                        meet</ChatBubble>
                    <ChatBubble type={isSent ? "sent" : "received"} messageTime="">Lorem ipsum dolor sit amet,
                        consetetur sadipscing elitr,
                        sed diam nonumy</ChatBubble>
                    <ChatBubble type={isSent ? "sent" : "received"} messageTime="Mary Foster_02:05 PM">Lorem ipsum dolor
                        sit amet,
                        consetetur sadipscing elitr, sed diam nonumy</ChatBubble>
                    <ChatBubble type={isSent ? "sent" : "received"} messageTime="02:10 PM">Lorem ipsum dolor sit
                        amet</ChatBubble>
                    <ChatBubble type="sent" messageTime="Mary Foster_02:13 PM">Lorem ipsum dolor sit amet,
                        consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore.</ChatBubble>
                    <ChatBubble type="sent" messageTime="Mary Foster_02:13 PM">Lorem ipsum dolor sit amet,
                        consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore.</ChatBubble>
                    <ChatBubble type="sent" messageTime="Mary Foster_02:13 PM">Lorem ipsum dolor sit amet,
                        consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore.</ChatBubble>
                    <ChatBubble type="sent" messageTime="Mary Foster_02:13 PM">Lorem ipsum dolor sit amet,
                        consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore.</ChatBubble>
                </div>
            </div>
            <div className="messageInputFieldContainer">
                <div className="messageInputField">
                    <TextField
                        id="outlined-multiline-static"
                        label="Type something..."
                        multiline
                        rows={3}
                    />

                    <div className="buttonGroup">
                        <span className="left">
                            <span><img src="/" alt=""/></span>
                            <span><img src="logo-emoji.svg" alt=""/></span>
                            <span><img src="logo-tenthou.svg" alt=""/></span>
                            <span><img src="logo-lightingMsg.svg" alt=""/></span>
                            <span><img src="logo-money.svg" alt=""/></span>
                        </span>
                        <span className="right">
                            <span><img src="logo-mic.svg" alt=""/></span>
                            <span><img src="logo-messageSubmit.svg" width="34px" height="34px" alt=""/></span>
                        </span>
                    </div>
                </div>
            </div>
        </div>

    )
}