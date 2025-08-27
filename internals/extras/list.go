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
