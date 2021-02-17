drop table IF EXISTS constants;

create table constants (
  id      serial primary key,
  name    varchar(255)
);

INSERT INTO constants VALUES (1, 'test');



