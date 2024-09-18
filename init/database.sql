CREATE DATABASE microservices;

\c microservices

CREATE TABLE IF NOT EXISTS public.url (
		id SERIAL PRIMARY KEY,
		shortUrl TEXT NOT NULL,
		longUrl TEXT NOT NULL
);