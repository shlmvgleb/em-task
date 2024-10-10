create table song (
  id bigserial primary key,
  song text not null,
  "group" text not null,
  "text" text not null,
  release_date datetime not null,
  "link" text not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
