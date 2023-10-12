CREATE TABLE IF NOT EXISTS transactions(
    id integer primary key autoincrement,
    description string not null,
    operation string not null,
    value number not null,
    cost_center string not null,
    day int not null,
    month string not null,
    payment_method string not null,
    ignore boolean default false
);