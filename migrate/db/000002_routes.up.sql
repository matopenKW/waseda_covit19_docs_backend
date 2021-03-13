CREATE TABLE routes (
  user_id            varchar(50),
  seq_no             serial,
  name               varchar(120),
  outward_trip       varchar(255),
  return_trip        varchar(255),
  update_time        timestamp default CURRENT_TIMESTAMP,
  create_time        timestamp default CURRENT_TIMESTAMP,
  CONSTRAINT routes_PK PRIMARY KEY (user_id, seq_no)
);