package extradependencies


var FrameworkDeps = map[string][]string{
	"express-js": {
		"dotenv", "drizzle", "prisma", "eslint", "nodemon", "mongodb", "cors", "zod",
		"axios", "lodash", "moment", "dayjs", "bcrypt", "bcryptjs", "uuid", 
		"mongoose", "redis", "ioredis", "helmet", "morgan", "winston", "joi", "yup",
		"socket.io", "passport", "multer", "sharp", "compression", "rate-limit", 
		"swagger", "mocha", "chai", "sinon", "prettier", "husky", "lint-staged",
		"concurrently", "cross-env",
	},
	"express-ts": {
		"dotenv", "jest", "prisma", "eslint", "drizzle", "nodemon", "cors", "zod",
		"axios", "lodash", "moment", "dayjs", "bcrypt", "bcryptjs", "uuid",
		"class-validator", "class-transformer", "typeorm", "mongoose", "redis", 
		"ioredis", "helmet", "morgan", "winston", "joi", "yup", "socket.io", 
		"passport", "multer", "sharp", "compression", "rate-limit", "swagger",
		"mocha", "chai", "sinon", "prettier", "husky", "lint-staged", "concurrently", 
		"cross-env",
	},
	"gin": {
		"gorm", "jwt-go", "godotenv", "validator",
		// Additional Go packages would be added here, but most are already covered
	},
	"fastapi": {
		"pydantic", "sqlalchemy", "alembic", "python-dotenv",
		// Additional Python packages would be: requests, pytest, black, flake8, etc.
	},
	"nextjs": {
		"drizzle", "prisma", "clerk", "supabase", "zod",
		"axios", "lodash", "moment", "dayjs", "bcrypt", "bcryptjs", "uuid",
		"mongoose", "redis", "ioredis", "helmet", "joi", "yup", "socket.io",
		"multer", "sharp", "compression", "rate-limit", "swagger", "jest",
		"prettier", "husky", "lint-staged", "concurrently", "cross-env",
	},
	"hono": {
		"dotenv", "mongodb", "drizzle", "prisma", "eslint", "zod",
		"axios", "lodash", "moment", "dayjs", "bcrypt", "bcryptjs", "uuid",
		"mongoose", "redis", "ioredis", "helmet", "morgan", "winston", "joi", "yup",
		"compression", "rate-limit", "swagger", "jest", "prettier", "husky", 
		"lint-staged", "concurrently", "cross-env",
	},
	"fiber": {
		"godotenv", "jwt-go", "gorm", "validator",
		// Additional Go packages would be added here
	},
	"nestjs": {
		"dotenv", "jsonwebtoken", "prisma", "drizzle", "eslint", "zod", "mongodb",
		"axios", "lodash", "moment", "dayjs", "bcrypt", "bcryptjs", "uuid",
		"class-validator", "class-transformer", "typeorm", "mongoose", "redis", 
		"ioredis", "helmet", "morgan", "winston", "joi", "yup", "socket.io", 
		"passport", "multer", "sharp", "compression", "rate-limit", "swagger",
		"jest", "prettier", "husky", "lint-staged", "concurrently", "cross-env",
	},
}