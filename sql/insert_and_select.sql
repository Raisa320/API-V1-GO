INSERT INTO items (customer_name,order_date,product,quantity,price) 
VALUES  ('Juan','27-02-2017','Product 1',5,80),
        ('Jason','11-11-2011','Product 1',1,10.5),
        ('Steph','06-10-2022','Product 2',1,60),
        ('Cassie','25-03-2023','Product 3',6,10),
        ('Tim','01-04-2019','Product 4',7,85.3),
        ('Barbara','12-01-2018','Product 4',3,66);


SELECT * FROM items WHERE quantity > 2 AND price >50;