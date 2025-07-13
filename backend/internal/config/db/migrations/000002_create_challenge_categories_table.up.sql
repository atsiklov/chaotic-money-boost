create table challenge_categories (
    id serial primary key,
    name varchar(50) not null unique
);

insert into challenge_categories (name) values ('sport');