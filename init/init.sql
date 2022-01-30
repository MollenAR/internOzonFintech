\c shorturl

CREATE TABLE IF NOT EXISTS urls
(
    id           serial        not null primary key,
    short_url    varchar(10)   not null,
    original_url varchar(2048) not null
)