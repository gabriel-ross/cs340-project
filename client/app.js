"use strict";

const PORT = 8081;
const express = require("express");
const app = express();

app.use(express.urlencoded({extended: true}));
app.use(express.static("public"));

app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}...`);
});