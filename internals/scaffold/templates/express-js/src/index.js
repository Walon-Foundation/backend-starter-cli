import express  from "express"
import cors from "cors"
import { corsOptions } from "./configs/corsOptions.js"
import { config } from "dotenv"

config()

const app = express()

//middleware 
app.use(express.json())
app.use(cors(corsOptions))


//test route
app.get("/", async(req,res) => {
    return res.send("Hello world")
})

//starting point
const PORT = process.env.PORT
app.listen(PORT, () => {
    console.log(`server in running on http://localhost:${PORT}`)
})