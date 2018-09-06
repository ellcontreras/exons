'use strict'

const app = require('./app');
const mongoose = require('mongoose');

mongoose.connect('mongodb://localhost:27017/comm', {useNewUrlParser: true}, (err, res) => {
    if (err) {
        throw err;
    } else {
        app.listen(8080, () => {
            console.log("Server is running on localhost:8080");
        });
    }
});
