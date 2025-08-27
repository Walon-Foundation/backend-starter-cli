package extras

var DatabaseList = map[string][]string{
	"node":    {"prisma", "drizzle", "mongodb", "postgresql", "none"},
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


