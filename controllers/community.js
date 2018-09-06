'use strict'

const Community = require('../models/communinty');

var getCommunity = (req, res) => {
    Community.findById(req.params.id, (err, community) => {
        if (err) {
            res.status(500).send({message: "Interval server error"});
        } else {
            if (!community) {
                res.status(404).send({message: "The community has not been found"});
            } else {
                res.status(200).send({community});
            }
        }
    });
};

var addCommunity = (req, res) => {
    var community = new Community();

    community.title = "first";
    community.description = "lorem ipsum tu gfa";

    community.save((err, community_save) => {
        if (err) {
            res.status(500).send({message: "Internal server error"});
        } else {
            if (!community_save) {
                res.status(404).send({message: "The user have not been saved"});
            } else {
                res.status(200).send({community: community_save});
            }
        }
    })
};

module.exports = {
    getCommunity,
    addCommunity
};
