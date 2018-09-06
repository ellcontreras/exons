const express = require('express');
const router = express.Router();

const community = require('../controllers/community');

// Route list
router.get('/get/:id', community.getCommunity);
router.post('/add', community.addCommunity);

module.exports = router;
