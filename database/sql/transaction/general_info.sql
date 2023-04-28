with total_spend as (select sum(transaction.price * transaction.amount) total_spend
                     from transaction
                     where transaction.telegram_user_id = cast($1 as bigint)),
     report as (select transaction.asset                                                     "asset",
                       sum(transaction.amount)                                               "amount",
                       sum(transaction.price * transaction.amount) / sum(transaction.amount) "avg_price",
                       sum(transaction.price * transaction.amount)                           "spend",
                       sum(transaction.price * transaction.amount)
                from transaction
                where transaction.telegram_user_id = cast($1 as bigint)
                group by "asset")
select "asset",
       "amount",
       "avg_price",
       "spend",
       "spend" * 100 / "total_spend" "percent"
from report,
     total_spend
order by asset asc
