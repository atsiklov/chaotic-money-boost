create table challenge_assignments (
    id                      serial primary key,
    user_id                 bigint not null references users(id),
    challenge_instance_id   bigint not null references challenge_instances(id) on delete cascade,
    status                  varchar(30) not null,    
    created_at              timestamptz not null default now(),
    updated_at              timestamptz not null default now(),
    submission              text,
    submitted_at            timestamptz
);