const express = require('express')
const cors = require('cors')

const { GraphQLClient } = require('graphql-request')
const gql = require('graphql-tag');

const app = express()
const port = 8082


app.use(cors())
app.use(express.json())
// app.use(express.urlencoded({extended:true}))

const accessToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJBUEkiLCJhcHAiOiI1ZTQyMDgzNzllNDhjMDZiYzhiYzgzOGUiLCJhY2wiOlsiYXBpOmFkbWluIiwiYXR0YWNobWVudElkOmNyZWF0ZSIsIndoYXRzYXBwOmdldEZpbGUiLCJib3Q6YWRtaW4iXSwianRpIjoiZTk2ODNlZTAtOWQyNS01ZDliLTliNWYtYmZlMTVlMTUyZmY2IiwiaXNzIjoiNWU0MjA4Mzc5ZTQ4YzAzM2YyYmM4MzkwIiwiaWF0IjoxNjQxOTE3Nzc4Njc1fQ.mWJiW1yh4D7d3c03ut9eUGHQxiSS2eLeF8Gmn6zURSc'


app.get("/bc-api/", async (req, res) => {
    res.send({'testing': 'broadcast server is running'})
})

app.post("/bc-api/template", async (req, res) => {
    const body = req.body
    console.log(body)

    let components = []

    // header
    if (body.header_type != null) {
        let header = {}
        header.type = "HEADER"
        header.format = body.header_type.toUpperCase()
        if (body.header_type === 'text'){
            header.text = body.header_body[0]
        } else {
            if (body.header_body != null) {
                let ex = []
                for(let i = 0; i < body.header_body.length; i++){
                    ex.push(body.header_body[i])
                }
                header.example = {}
                header.example.header_handle = ex
            }
        }
        components.push(header)
    }

    // body
    if (body.body == null) {
        res.status(400).send({'error': 'missing body'})
    }
    const msgBody = {
        type: 'BODY',
        text: body.body,
    }
    components.push(msgBody)

    //footer
    if (body.footer != null){
        const footer = {
            type: 'FOOTER',
            text: body.footer,
        }
        components.push(footer)
    }

    //buttons
    if (body.buttons_type != null) {
        let butBody = {}
        butBody.type = 'BUTTONS'
        let buttons = []
        for (let i = 0; i < body.buttons_body.length; i++){
            let button = {}
            if (body.buttons_type === 'url') {
                button.type = 'URL'
                button.url =  body.buttons_body[i].url
                button.text = body.buttons_body[i].text
            } else if (body.buttons_type === 'quick_reply'){
                button.type = 'QUICK_REPLY'
                button.text = body.buttons_body[i].text
            }
            buttons.push(button)
        }
        butBody.buttons = buttons
        components.push(butBody)
    }

    const endpoint = 'https://openapi.stellabot.com/v2'
    const client = new GraphQLClient(endpoint)
    const requestHeaders = {
        "Authorization": "Bearer " + `${accessToken}`
    }

    const mutation = gql
        `mutation TestTextTemplate($input: CreateWhatsAppMessageTemplateInput!){
        createWhatsAppMessageTemplate(input: $input){
            result
            clientMutationId
        }
    }`

    const variables = {
        input: {
            name: body.name,
            category: body.category,
            channelId: body.channel_id,
            language: body.language,
            components: components
        }
    }
    console.log(JSON.stringify(variables))
    await client.request(mutation, variables, requestHeaders)
        .then(res => {
            if (res.createWhatsAppMessageTemplate.result.error) {
                res.status(400).send({'error': JSON.stringify(res.createWhatsAppMessageTemplate.result.error)})
            }
            console.log(JSON.stringify(res))
            res.sendStatus(200)
        })
        .catch(err => {
            console.log(err)
            res.status(500).send({'error': err})
        })

    res.sendStatus(200)
})

app.listen(port, () => {
    console.log('Broadcast server is running on port ', port)
})
