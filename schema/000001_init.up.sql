create table links (
    id serial not null,
    hcode varchar(10) not null,
    original text not null
);

create index hashcode on links (hcode);

create table customlinks (
    id serial not null,
    custom varchar(10) not null,
    original text not null
);

create index customcode on customlinks (custom);