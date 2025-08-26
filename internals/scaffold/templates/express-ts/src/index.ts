import express from "express";

const app = express()

app.use(express.json())

app.get("/", async(req, res) => {
    return res.send("hello world")
})

app.listen(3000, () => {
    console.log(`server is running on http://localhost:${3000}`)
})