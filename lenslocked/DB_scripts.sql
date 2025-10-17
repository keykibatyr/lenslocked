CREATE TABLE users(
id SERIAL PRIMARY KEY,
age INT,
first_name TEXT,
last_name TEXT,
email Text Unique not null
);


CREATE TABLE users(
id SERIAL PRIMARY KEY
age INT,
first_name TEXT,
last_name TEXT,
pass_word TEXT,
email Text Unique not null;
)

CREATE TABLE tweets(
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id),
time_stamp DATE,
comments INT,
format TEXT CHECK (format IN ('video', 'image', 'gif'));
)

CREATE TABLE likes(
id SERIAL PRIMARY KEY,
tweet_id INT REFERENCES tweets(id),
amount INT;
)
