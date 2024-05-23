Table users {
  id serial [primary key]
  first_name varchar(50) [not null]
  last_name varchar(50) [not null]
  email varchar(100) [not null, unique]
  password varchar(100) [not null]
}

Table products {
  id serial [primary key]
  name varchar(100) [not null]
  description text
  image varchar(255)
  quantity int [not null]
  created_at timestamp [not null, default: `CURRENT_TIMESTAMP`]
}

Table orders {
  id serial [primary key]
  user_id int [not null, ref: > users.id]
  total int [not null]
  status varchar(50) [not null]
  address text [not null]
  created_at timestamp [not null, default: `CURRENT_TIMESTAMP`]
}

