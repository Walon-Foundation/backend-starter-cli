package extras

var DatabaseList = map[string][]string{
	"node":    {"prisma", "drizzle", "mongodb", "postgresql"},
	"go":      {"sqlite", "postgresql", "mongodb"},
	"fastapi": {"sqlite3", "postgresql", "mongodb"},
}

var DatabaseListInfo = map[string]map[string][]string{
	// Node.js (JS/TS)
	"drizzle": {
		"js":  {"npm", "install", "drizzle-orm", "pg", "drizzle-kit", "tsx"},
		"ts":  {"npm", "install", "drizzle-orm", "pg", "drizzle-kit", "tsx", "@types/pg"},
	},
	"prisma": {
		"js": {"npx", "prisma", "init", "--datasource-provider", "postgresql"},
		"ts": {"npx", "prisma", "init", "--datasource-provider", "postgresql"},
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
