START TRANSACTION;

-- Add enum type for discount_type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'discount_type_enum') THEN
        CREATE TYPE discount_type_enum AS ENUM ('percentage', 'fix_amount');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS "products"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL DEFAULT '',
    "category" TEXT NOT NULL DEFAULT '',
    "price" DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "images" (
    "id" SERIAL PRIMARY KEY,
    "product_id" INTEGER NOT NULL,
    "thumbnail" VARCHAR(255) NOT NULL DEFAULT '',
    "mobile" VARCHAR(255) NOT NULL DEFAULT '',
    "tablet" VARCHAR(255) NOT NULL DEFAULT '',
    "desktop" VARCHAR(255) NOT NULL DEFAULT '',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id" SERIAL PRIMARY KEY,
    "total_amount" DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    "status" VARCHAR(50) NOT NULL DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "order_items" (
    "id" SERIAL PRIMARY KEY,
    "order_id" INTEGER NOT NULL,
    "product_id" INTEGER NOT NULL,
    "quantity" INTEGER NOT NULL DEFAULT 1,
    "price" DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    "total_price" DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

COMMIT;