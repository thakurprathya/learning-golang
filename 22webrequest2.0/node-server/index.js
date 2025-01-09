const express = require('express');

const app = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true })); // 'extended' enables parsing of nested objects and arrays in request bodies, when set to true, it allows handling complex/nested JSON payloads

app.get('/get', (req, res) => {
    res.status(200).json({message: 'Hello World!'});
});

app.post('/post', (req, res) => {
    res.status(200).json(req.body);
});

app.post('/form', (req, res) => {  //endpoint for formdata
    res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
    console.log(`Server is running at http://localhost:${port}`);
});