CREATE OR REPLACE FUNCTION calculate_total_price(item_id INTEGER)
RETURNS NUMERIC AS
$$
BEGIN
    RETURN (SELECT (price * quantity) FROM items WHERE id = item_id);
END;
$$
LANGUAGE plpgsql;

--- PRUBE DE LA FUNCION 
 SELECT calculate_total_price(4);