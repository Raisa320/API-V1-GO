BEGIN; -- Inicia la transacción

-- Supón que deseas insertar un nuevo item en la tabla 'items'
INSERT INTO items (customer_name, order_date, product, quantity, price)
VALUES ('Jane Doe', '2023-05-17', 'View', 1, 1000);

-- Supón que también deseas actualizar el precio de otro item
UPDATE items
SET price = 2000
WHERE id = 2;

COMMIT; -- Finaliza la transacción

BEGIN;
DELETE FROM items
WHERE customer_name = 'Jane Doe' AND product = 'Apple Macbook Pro';
ROLLBACK;


