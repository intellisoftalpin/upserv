package main

const CHTMLTemplate = `
    <html>
    <head><title>Uploaded Files</title></head>
    <body>
        <h1>Uploaded Files</h1>
        <ul>
            {{range .}}
                <li>{{.Name}} - {{.ModTime}}</li>
            {{end}}
        </ul>
    </body>
    </html>
`
