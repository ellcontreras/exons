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
app.use('/api/community', communityRouter);

module.exports = app;