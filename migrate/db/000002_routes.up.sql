CREATE TABLE routes (
  id                 serial,
  user_id            varchar(50),
  name               varchar(120),
  outward_trip       varchar(255),
  return_trip        varchar(255),
  update_time        timestamp default CURRENT_TIMESTAMP,
  create_time        timestamp default CURRENT_TIMESTAMP,
  CONSTRAINT routes_PK PRIMARY KEY (id)
);

INSERT INTO routes
VALUES
    (
        1,
        'user_id',
        'name',
        'outward_trip',
        'return_trip'
    );
