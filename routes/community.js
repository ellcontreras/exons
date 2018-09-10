const express = require('express');
const router = express.Router();

const community = require('../controllers/community');

// Route list
router.get('/get/all', community.getAll);
router.get('/get/:id', community.getCommunity);
router.post('/add', community.addCommunity);
router.put('/update', community.updateCommunity);

module.exports = router;
