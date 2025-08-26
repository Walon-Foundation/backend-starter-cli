package extradependencies

type PackageDetails struct {
	PackageName string
	HasType     bool
}

type Extra struct {
	Name   string
	Detail PackageDetails
}

var All = []Extra{
	// Core packages from FrameworkDeps
	{Name: "dotenv", Detail: PackageDetails{PackageName: "dotenv", HasType: false}},
	{Name: "drizzle", Detail: PackageDetails{PackageName: "drizzle-orm", HasType: true}},
	{Name: "prisma", Detail: PackageDetails{PackageName: "@prisma/client", HasType: true}},
	{Name: "eslint", Detail: PackageDetails{PackageName: "eslint", HasType: false}},
	{Name: "nodemon", Detail: PackageDetails{PackageName: "nodemon", HasType: false}},
	{Name: "mongodb", Detail: PackageDetails{PackageName: "mongodb", HasType: true}},
	{Name: "cors", Detail: PackageDetails{PackageName: "cors", HasType: true}},
	{Name: "zod", Detail: PackageDetails{PackageName: "zod", HasType: false}},
	{Name: "jest", Detail: PackageDetails{PackageName: "jest", HasType: true}},
	{Name: "jsonwebtoken", Detail: PackageDetails{PackageName: "@nestjs/jwt", HasType: true}},
	{Name: "clerk", Detail: PackageDetails{PackageName: "@clerk/nextjs", HasType: true}},
	{Name: "supabase", Detail: PackageDetails{PackageName: "@supabase/supabase-js", HasType: true}},

	// Go packages
	{Name: "gorm", Detail: PackageDetails{PackageName: "gorm.io/gorm", HasType: false}},
	{Name: "jwt-go", Detail: PackageDetails{PackageName: "github.com/golang-jwt/jwt/v4", HasType: false}},
	{Name: "godotenv", Detail: PackageDetails{PackageName: "github.com/joho/godotenv", HasType: false}},
	{Name: "validator", Detail: PackageDetails{PackageName: "github.com/go-playground/validator/v10", HasType: false}},

	// Python packages
	{Name: "pydantic", Detail: PackageDetails{PackageName: "pydantic", HasType: true}},
	{Name: "sqlalchemy", Detail: PackageDetails{PackageName: "sqlalchemy", HasType: true}},
	{Name: "alembic", Detail: PackageDetails{PackageName: "alembic", HasType: false}},
	{Name: "python-dotenv", Detail: PackageDetails{PackageName: "python-dotenv", HasType: false}},

	// Additional packages from FrameworkDeps
	{Name: "axios", Detail: PackageDetails{PackageName: "axios", HasType: true}},
	{Name: "lodash", Detail: PackageDetails{PackageName: "lodash", HasType: true}},
	{Name: "moment", Detail: PackageDetails{PackageName: "moment", HasType: true}},
	{Name: "dayjs", Detail: PackageDetails{PackageName: "dayjs", HasType: true}},
	{Name: "bcrypt", Detail: PackageDetails{PackageName: "bcrypt", HasType: true}},
	{Name: "bcryptjs", Detail: PackageDetails{PackageName: "bcryptjs", HasType: true}},
	{Name: "uuid", Detail: PackageDetails{PackageName: "uuid", HasType: true}},
	{Name: "class-validator", Detail: PackageDetails{PackageName: "class-validator", HasType: true}},
	{Name: "class-transformer", Detail: PackageDetails{PackageName: "class-transformer", HasType: true}},
	{Name: "typeorm", Detail: PackageDetails{PackageName: "typeorm", HasType: true}},
	{Name: "mongoose", Detail: PackageDetails{PackageName: "mongoose", HasType: true}},
	{Name: "redis", Detail: PackageDetails{PackageName: "redis", HasType: true}},
	{Name: "ioredis", Detail: PackageDetails{PackageName: "ioredis", HasType: true}},
	{Name: "helmet", Detail: PackageDetails{PackageName: "helmet", HasType: true}},
	{Name: "morgan", Detail: PackageDetails{PackageName: "morgan", HasType: true}},
	{Name: "winston", Detail: PackageDetails{PackageName: "winston", HasType: true}},
	{Name: "joi", Detail: PackageDetails{PackageName: "joi", HasType: true}},
	{Name: "yup", Detail: PackageDetails{PackageName: "yup", HasType: true}},
	{Name: "socket.io", Detail: PackageDetails{PackageName: "socket.io", HasType: true}},
	{Name: "passport", Detail: PackageDetails{PackageName: "passport", HasType: true}},
	{Name: "multer", Detail: PackageDetails{PackageName: "multer", HasType: true}},
	{Name: "sharp", Detail: PackageDetails{PackageName: "sharp", HasType: true}},
	{Name: "compression", Detail: PackageDetails{PackageName: "compression", HasType: true}},
	{Name: "rate-limit", Detail: PackageDetails{PackageName: "express-rate-limit", HasType: true}},
	{Name: "swagger", Detail: PackageDetails{PackageName: "swagger-ui-express", HasType: true}},
	{Name: "mocha", Detail: PackageDetails{PackageName: "mocha", HasType: true}},
	{Name: "chai", Detail: PackageDetails{PackageName: "chai", HasType: true}},
	{Name: "sinon", Detail: PackageDetails{PackageName: "sinon", HasType: true}},
	{Name: "prettier", Detail: PackageDetails{PackageName: "prettier", HasType: false}},
	{Name: "husky", Detail: PackageDetails{PackageName: "husky", HasType: false}},
	{Name: "lint-staged", Detail: PackageDetails{PackageName: "lint-staged", HasType: false}},
	{Name: "concurrently", Detail: PackageDetails{PackageName: "concurrently", HasType: false}},
	{Name: "cross-env", Detail: PackageDetails{PackageName: "cross-env", HasType: false}},
}