CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
	first_name TEXT,
	last_name TEXT,
	telegram_id BIGINT NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	first_query_time_stamp TEXT,
	query_counter BIGINT NOT NULL DEFAULT 0
);
