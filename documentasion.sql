-- Active: 1706099661876@@127.0.0.1@5432@assignment_project@public
CREATE table gmail (
    id SERIAL PRIMARY KEY NOT NULL, message VARCHAR(300), receiver VARCHAR(50), created_at TIMESTAMP
);

-- craete user
CREATE Table users (
    id SERIAL PRIMARY KEY NOT NULL, name VARCHAR(60), dob TIMESTAMP, gmail VARCHAR(60), username VARCHAR(500), password VARCHAR(500), created_at TIMESTAMP, updated_at TIMESTAMP
);

-- create table products
CREATE Table products (
    id SERIAL PRIMARY KEY NOT NULL, name VARCHAR(60), stock INT, price FLOAT, created_at TIMESTAMP, updated_at TIMESTAMP
);

-- create table promo
CREATE TABLE promo (
    id SERIAL PRIMARY KEY NOT NULL, kode_promo VARCHAR(50), token VARCHAR(500), created_at TIMESTAMP
);

-- create table Order
CREATE TABLE orders (
    id SERIAL PRIMARY KEY NOT NULL, user_id INT NOT NULL, product_id INT NOT NULL, qty INT, promo FLOAT, promo_code VARCHAR(60), total_price FLOAT, total_paid FLOAT, total_return FLOAT, created_at TIMESTAMP, CONSTRAINT fk_orders_users FOREIGN KEY (user_id) REFERENCES users (id), CONSTRAINT fk_orders_products FOREIGN KEY (product_id) REFERENCES products (id)
);