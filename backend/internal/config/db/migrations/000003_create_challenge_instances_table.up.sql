create table challenge_instances (
     id bigserial primary key,
     challenge_template_id bigint not null references challenge_templates(id),
     status varchar(20) not null,
     created_at timestamptz not null default now(),
     updated_at timestamptz not null default now(),
     started_at timestamptz,
     expires_at timestamptz
);