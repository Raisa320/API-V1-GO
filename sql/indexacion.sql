--Time sin INDEX: Duración: 0.945 ms
SELECT * FROM items WHERE product = 'Product 1';

-- CREATE INDEX : product
CREATE INDEX  idx_items_product ON items (product);

--Time con INDEX: Duración: (primera vez = 4.270 ms),(segunda vez=0.852 ms)
SELECT * FROM items WHERE product = 'Product 1';



