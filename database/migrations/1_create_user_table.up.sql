create table users (
  id bigint(20) unsigned primary key not null auto_increment,
  email varchar(255) not null unique,
  password varchar(255) not null,
  api_key varchar(255) not null unique,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);