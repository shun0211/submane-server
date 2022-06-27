CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	name text NULL,
	email text NULL,
	uid bytea NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
