CREATE TABLE cryptos (
    id varchar(50) NOT NULL PRIMARY KEY,
    rank varchar(255) NOT NULL,
    symbol varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    supply varchar(255) NOT NULL,
    maxSupply varchar(255) NOT NULL,
    marketCapUsd varchar(255) NOT NULL,
    volumeUsd24Hr varchar(255) NOT NULL,
    priceUsd varchar(255) NOT NULL,
    changePercent24Hr varchar(255) NOT NULL,
    vwap24Hr varchar(255) NOT NULL
)