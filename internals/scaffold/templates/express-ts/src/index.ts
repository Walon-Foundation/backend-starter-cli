import express from "express";
import cors from "cors";
import { config } from "dotenv";
import { corsOptions } from "./configs/corsOption";

config()

//making the server
const app = express()

//middleware
app.use(express.json())
app.use(cors(corsOptions))


//test route
app.get("/", async(req, res) => {
    return res.send("hello world")
})


//server starting point
const PORT = process.env.PORT
app.listen(PORT, () => {
    console.log(`server is running on http://localhost:${PORT}`)
})