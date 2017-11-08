DROP DATABASE IF EXISTS todo;
CREATE DATABASE todo;

USE todo;

CREATE TABLE todo (
	id bytes primary key,
	title string,
	description string,
	complete bool default false,
	completed_at timestamp,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp
);
