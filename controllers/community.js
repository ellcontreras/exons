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

    community.title = req.body.title;
    community.description = req.body.description;

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

var updateCommunity = (req, res) => {
    var body = req.body;

    Community.findByIdAndUpdate(body._id, body, (err, community) => {
        if (err) {
            res.ststua(500).send({message: "Internal server error!"});
        } else {
            if (!communityUpdated) {
                res.status(404).send({message: "The commnity has not been found!"});
            } else {
                res.status(200).send({community});
            }
        }
    });
};

var getAll = (req, res) => {
    Community.find((err, communities) => {
        if (err) {
            res.status(500).send({message: "Internal server errror"});
        } else {
            if (!communities) {
                res.status(404).send({message: "There is not communitites"});
            } else {
                res.status(200).send({communities});
            }
        }
    }).sort({_id:-1});
}

module.exports = {
    getCommunity,
    addCommunity,
    updateCommunity,
    getAll
};
