

CREATE TABLE IF NOT EXISTS items 
(
id SERIAL PRIMARY KEY,
customer_name VARCHAR(85) NOT NULL, 
order_date DATE NOT NULL, 
product VARCHAR(255) NOT NULL, 
quantity INTEGER NOT NULL CHECK(quantity > 0), 
price NUMERIC NOT NULL CHECK(price >= 0)
);