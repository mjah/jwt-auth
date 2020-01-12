DROP TABLE IF EXISTS user;
CREATE TABLE user (
  user_id serial NOT NULL PRIMARY KEY,
  group_id serial NOT NULL,
  email varchar(128) NOT NULL,
  username varchar(40) NOT NULL,
  password_hash varchar(100) NOT NULL,
  password varchar(255) NOT NULL,
  first_name varchar(32) NOT NULL,
  last_name varchar(32) NOT NULL,
  created timestamp NOT NULL,
  modified timestamp NOT NULL,
  last_login timestamp,
  is_confirmed smallint(6) NOT NULL,
  is_active smallint(6) NOT NULL
  resetpass_token text,
  resetpass_token_created text,
  first_failure timestamp,
  lock_expires timestamp
);

DROP TABLE IF EXISTS

DROP TABLE IF EXISTS group;
CREATE TABLE group (
  group_id serial NOT NULL PRIMARY KEY,
  group_name varchar(32) NOT NULL,
  permissions text
);

DROP TABLE IF EXISTS token_revocation;
CREATE TABLE token_revocation (

);
INSERT INTO group (group_id, group_name) VALUES (1, 'admin');
INSERT INTO group (group_id, group_name) VALUES (2, 'guest');
