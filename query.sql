CREATE TABLE category(
	category_id SERIAL PRIMARY KEY ,
	name VARCHAR(100),
	description TEXT,
	created_at timestamp default current_timestamp,
    updated_at date
)
