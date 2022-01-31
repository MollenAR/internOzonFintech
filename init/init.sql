\c shorturl

CREATE TABLE IF NOT EXISTS urls
(
    id           serial        not null primary key,
    short_url    varchar(10)   not null,
    original_url varchar(2048) not null
);

CREATE UNIQUE INDEX IF NOT EXISTS shorturl_indx ON urls(short_url);
CREATE UNIQUE INDEX IF NOT EXISTS originalurl_indx ON urls(original_url);