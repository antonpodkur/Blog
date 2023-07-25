CREATE TABLE "users" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "photo" VARCHAR NOT NULL,
    "verified" BOOLEAN NOT NULL,
    "password" VARCHAR NOT NULL,
    "role" VARCHAR NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

CREATE TABLE "articles" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "title" VARCHAR(500) NOT NULL,
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,
    "user_id" UUID NOT NULL, 

    CONSTRAINT "articles_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "users"("id") 
);
