CREATE TABLE IF NOT EXISTS subscriptions (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	name text NULL,
	price int8 NULL,
	contracted_at timestamptz NULL,
	user_id int8 NULL,
	CONSTRAINT subscriptions_pkey PRIMARY KEY (id),
	CONSTRAINT fk_subscriptions_user FOREIGN KEY (user_id) REFERENCES users(id)
);
