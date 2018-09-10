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

app.use((req, res, next,) => {
    res.setHeader("Access-Control-Allow-Headers", "content-type")
    res.setHeader("Content-type", "application/json");
    res.setHeader("Access-Control-Allow-Origin", "*");

    next();
});

// Route declaration
app.use('/api/community', communityRouter);

module.exports = app;