CREATE SCHEMA IF NOT EXISTS "production";

CREATE TABLE IF NOT EXISTS "production"."category" (
    category_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parent_category_id UUID,
    name VARCHAR(255) NOT NULL CHECK (name <> ''),
    description TEXT CHECK (description <> ''),
    tags VARCHAR(255)[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "production"."brand" (
    brand_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL CHECK (name <> ''),
    description TEXT CHECK (description <> ''),
    tags VARCHAR(255)[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "production"."product" (
    product_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_id UUID NOT NULL,
    brand_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL CHECK (name <> ''),
    description TEXT CHECK (description <> ''),
    tags VARCHAR(255)[] NOT NULL DEFAULT '{}',
    list_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    sell_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    tax_rate DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (category_id) REFERENCES "production"."category" (category_id) ON DELETE CASCADE,
    FOREIGN KEY (brand_id) REFERENCES "production"."brand" (brand_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "production"."product_image" (
    product_image_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    image_url TEXT NOT NULL CHECK (image_url <> ''),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (product_id) REFERENCES "production"."product" (product_id) ON DELETE CASCADE
);