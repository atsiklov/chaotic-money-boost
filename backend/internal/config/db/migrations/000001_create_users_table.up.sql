create table users (
    id bigserial primary key,
    nickname varchar(32) not null unique check (char_length(nickname) >= 3),
    email varchar(100) not null unique,
    created_at timestamptz not null default now()
);

insert into users(nickname, email) values ('atsiklov', 'atsiklov@gmail.com');
insert into users(nickname, email) values ('akomarov', 'akomarov@gmail.com');
