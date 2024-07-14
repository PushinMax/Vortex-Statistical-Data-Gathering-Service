CREATE TABLE orderBook
(
    id       BIGINT not null unique,
    exchange VARCHAR(255),
    pair     VARCHAR(255)
);

CREATE TABLE asks
(
    id      serial primary key,
    id_book BIGINT,
    price   DOUBLE PRECISION,
    baseQty DOUBLE PRECISION
);

CREATE TABLE bids
(
    id      serial primary key,
    id_book BIGINT,
    price   DOUBLE PRECISION,
    baseQty DOUBLE PRECISION
);

CREATE TABLE orderHistory
(
    client_name           VARCHAR(255),
    exchange_name         VARCHAR(255),
    label                 VARCHAR(255),
    pair                  VARCHAR(255),
    side                  VARCHAR(255),
    type                  VARCHAR(255),
    base_qty              DOUBLE PRECISION,
    price                 DOUBLE PRECISION,
    algorithm_name_placed VARCHAR(255),
    lowest_sell_prc       DOUBLE PRECISION,
    highest_buy_prc       DOUBLE PRECISION,
    commission_quote_qty  DOUBLE PRECISION,
    time_placed           TIME
)