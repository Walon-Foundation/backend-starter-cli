import express  from "express"

const app = express()

//middleware 
app.use(express.json())

//test route
app.get("/", async(req,res) => {
    return res.send("Hello world")
})

//starting point
app.listen(3000, () => {
    console.log(`server in running on http://localhost:${3000}`)
})