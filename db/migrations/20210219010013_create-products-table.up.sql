CREATE TABLE "products"
(
    "id"             VARCHAR(255),
    "title"          VARCHAR(255) NOT NULL,
    "description"    VARCHAR(500) NOT NULL,
    "price_in_cents" INTEGER      NOT NULL
);
ALTER TABLE "products"
    ADD CONSTRAINT "products_pkey" PRIMARY KEY ("id")
