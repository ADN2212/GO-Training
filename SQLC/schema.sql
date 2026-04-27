CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    language smallint,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT uni_users_username UNIQUE (username)
)
