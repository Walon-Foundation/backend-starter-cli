from fastapi import FastAPI
from dotenv import load_dotenv
from fastapi.middleware.cors import CORSMiddleware

#loading in the .env
load_dotenv()

app = FastAPI()

#origins to be allowed
origins = [
    "add the url you want to allow"
]

#adding the cors middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/")
async def root():
    return {"message":"hello world"}
