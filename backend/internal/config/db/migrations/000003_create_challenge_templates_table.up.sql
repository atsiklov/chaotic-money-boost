create table challenge_templates (
    id bigserial primary key,
    challenge_category_id int not null references challenge_categories(id) on delete cascade,
    description text not null
);

insert into challenge_templates (challenge_category_id, description) values (1, 'Отжаться 10 раз');