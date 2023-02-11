select 
    orders.id as order_id,
    users.fullname as fullname,
    users.email as email,
    orders.product_name as product_name,
    orders.unit_price as unit_price,
    orders.quantity as quantity,
    orders.order_date as order_date

    from orders
    join users on orders.user_id = users.id
    where users.status = 'active' and (orders.unit_price * orders.quantity) > 500000 or orders.quantity>20;