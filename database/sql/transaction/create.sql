insert into transaction (telegram_user_id, purchase_date, asset, price, amount)
values (:telegram_user_id, :purchase_date, :asset, :price, :amount)
returning id;