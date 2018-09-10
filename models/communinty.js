'use strict'

const mongoose = require('mongoose');

var CommunitySchema = mongoose.Schema({
    title: String,
    description: String
});

module.exports = mongoose.model("community", CommunitySchema);
