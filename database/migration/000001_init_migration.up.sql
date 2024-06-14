CREATE TABLE users
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255),
    username    VARCHAR(50) NOT NULL,
    password    TEXT NOT NULL,
    role        VARCHAR(50) NOT NULL,
    created_at  INT NOT NULL,
    updated_at  INT
);

CREATE TABLE products
(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    product_name    VARCHAR(255) NOT NULL,
    category        VARCHAR(255) NOT NULL,
    price           INT NOT NULL,
    description     TEXT,
    created_at      INT NOT NULL,
    updated_at      INT NOT NULL
);

CREATE TABLE carts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE cart_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cart_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cart_id) REFERENCES carts(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    total_amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unit_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
