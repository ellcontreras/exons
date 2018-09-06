'use strict'

const bodyParser = require('body-parser');
const express = require('express');

const { exec } = require('child_process');

const app = express();

// Routes
const communityRouter = require('./routes/community');

// Settings
app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

// Route declaration
app.use('/api/community/', communityRouter);

app.listen(8080, () => {
    console.log("The server is running on :8080");

    // exec('browse http://localhost:8080');
});
