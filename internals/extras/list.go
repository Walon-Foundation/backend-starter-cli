package extras

var DatabaseList = map[string][]string{
	"node":    {"prisma", "drizzle", "mongodb", "none"},
	"go":      {"sqlite", "postgresql", "mongodb", "none"},
	"fastapi": {"sqlite3", "postgresql", "mongodb", "none"},
}

var DatabaseListInfo = map[string]map[string][]string{
	// Node.js (JS/TS)
	"drizzle": {
		"js": {"npm", "install", "drizzle-orm", "pg", "drizzle-kit", "tsx"},
		"ts": {"npm", "install", "drizzle-orm", "pg", "drizzle-kit", "tsx", "@types/pg"},
	},
	"prisma": {
		"js": {"npm", "install", "prisma", "@prisma/client"},
		"ts": {"npm", "install", "prisma", "@prisma/client"},
	},
	"mongodb": {
		"js": {"npm", "install", "mongodb"},
		"ts": {"npm", "install", "mongodb", "@types/mongodb"},
	},
	"postgresql": {
		"go":      {"go", "get", "github.com/lib/pq"},
		"fastapi": {"pip", "install", "psycopg2-binary"},
	},

	// Go only
	"sqlite": {
		"go": {"go", "get", "github.com/mattn/go-sqlite3"},
	},

	// FastAPI only
	"sqlite3": {
		"fastapi": {"pip", "install", "aiosqlite"},
	},
}


var ValidationList = map[string][]string{
	"node":    {"zod", "yup", "joi", "class-validator", "none"},
	"fastapi": {"pydantic", "none"},
	"go":      {"go-playground/validator", "none"},
}


var ValidationListInfo = map[string]map[string][]string{
	"zod": {
		"js": {"npm", "install", "zod"},
		"ts": {"npm", "install", "zod"}, 
	},
	"yup": {
		"js": {"npm", "install", "yup"},
		"ts": {"npm", "install", "yup", "@types/yup"},
	},
	"joi": {
		"js": {"npm", "install", "joi"},
		"ts": {"npm", "install", "joi", "@types/joi"},
	},
	"class-validator": {
		"js": {"npm", "install", "class-validator", "class-transformer"}, // works but not common in JS-only
		"ts": {"npm", "install", "class-validator", "class-transformer"},
	},
	"pydantic": {
		"fastapi": {"pip", "install", "pydantic"},
	},
	"go-playground/validator": {
		"go": {"go", "get", "github.com/go-playground/validator/v10"},
	},
}



var AuthProviderList = map[string][]string{
	"node":    {"jwt", "clerk", "auth0", "supabase", "firebase", "okta", "none"},
	"go":      {"jwt", "auth0", "supabase", "firebase", "okta", "none"},
	"fastapi": {"jwt", "auth0", "supabase", "firebase", "okta", "none"},
}



var AuthProviderListInfo = map[string]map[string][]string{
	"jwt": {
		"js":      {"npm", "install", "jsonwebtoken", "bcryptjs"},
		"ts":      {"npm", "install", "jsonwebtoken", "bcryptjs", "@types/jsonwebtoken", "@types/bcryptjs"},
		"go":      {"go", "get", "github.com/golang-jwt/jwt/v5"},
		"fastapi": {"pip", "install", "python-jose", "passlib[bcrypt]"},
	},
	"clerk": {
		"js":      {"npm", "install", "@clerk/clerk-sdk-node"},
		"ts":      {"npm", "install", "@clerk/clerk-sdk-node"},
	},
	"auth0": {
		"js":      {"npm", "install", "express-jwt", "jwks-rsa"},
		"ts":      {"npm", "install", "express-jwt", "jwks-rsa", "@types/express-jwt"},
		"go":      {"go", "get", "github.com/auth0/go-jwt-middleware"},
		"fastapi": {"pip", "install", "python-jose", "requests"},
	},
	"supabase": {
		"js":      {"npm", "install", "@supabase/supabase-js"},
		"ts":      {"npm", "install", "@supabase/supabase-js"},
		"go":      {"go", "get", "github.com/golang-jwt/jwt/v5"}, // verify Supabase JWTs
		"fastapi": {"pip", "install", "supabase-py", "python-jose"},
	},
	"firebase": {
		"js":      {"npm", "install", "firebase-admin"},
		"ts":      {"npm", "install", "firebase-admin", "@types/firebase-admin"},
		"go":      {"go", "get", "firebase.google.com/go"},
		"fastapi": {"pip", "install", "firebase-admin"},
	},
	"okta": {
		"js":      {"npm", "install", "@okta/okta-sdk-nodejs", "@okta/okta-auth-js"},
		"ts":      {"npm", "install", "@okta/okta-sdk-nodejs", "@okta/okta-auth-js"},
		"go":      {"go", "get", "github.com/okta/okta-jwt-verifier-golang"},
		"fastapi": {"pip", "install", "okta"},
	},
}
