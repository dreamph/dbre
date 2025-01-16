CREATE TABLE country
(
    id          varchar NOT NULL,
    code        varchar NOT NULL,
    "name"      varchar NOT NULL,
    status      int4    NOT NULL,
    description varchar NULL,
    other_field varchar NULL,
    CONSTRAINT country_pk PRIMARY KEY (id)
);