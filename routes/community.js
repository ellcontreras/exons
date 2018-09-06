const express = require('express');
const router = express.Router();

const community = require('../controllers/community');

// Route list
router.get('/', community.getCommunity);

module.exports = router;
