
SELECT
    name AS Customers
FROM Customers
LEFT JOIN orders
    ON Orders.customerId = Customers.id
WHERE Orders.Id IS NULL;