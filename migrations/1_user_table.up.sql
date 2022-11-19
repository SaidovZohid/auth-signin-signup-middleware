CREATE TABLE IF NOT EXISTS "users" (
    "id" serial not null primary key,
    "first_name" varchar(50) not null,
    "last_name" varchar(50) not null,
    "email" varchar not null UNIQUE,
    "phone_number" varchar(30) UNIQUE,
    "password" varchar not NULL,
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    "last_login" TIMESTAMP,
    "deleted_at" TIMESTAMP
);