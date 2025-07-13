create table challenge_templates (
    id bigserial primary key,
    category varchar(32) not null,
    description text not null,
    duration interval null 
);

insert into challenge_templates (category, description, duration) values ('GOOD_DEED', 'Сделать комплимент кассирше в магазине', interval '60 seconds');
insert into challenge_templates (category, description, duration) values ('ART', 'Нарисовать сердечко на бумаге', interval '30 seconds');
insert into challenge_templates (category, description, duration) values ('SPORT', 'Отжаться 10 раз', interval '10 seconds');