CREATE TABLE "users"
(
    "id"            VARCHAR(255),
    "first_name"    VARCHAR(255) NOT NULL,
    "last_name"     VARCHAR(255) NOT NULL,
    "date_of_birth" VARCHAR(255) NOT NULL,
    "email"         VARCHAR(255) NOT NULL,
    "password"      VARCHAR(255) NOT NULL
);
ALTER TABLE "users"
    ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
ALTER TABLE "users"
    ADD CONSTRAINT "users_email_unique" UNIQUE ("email")
