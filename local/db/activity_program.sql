drop table IF EXISTS activity_program;

create table activity_program (
  id                 serial primary key,
  datetime           timestamp primary key,
  start_time         timestamp,
  end_time           timestamp,
  practice_section   varchar(255),
  practice_Contents  varchar(255),
  venue_id           serial,
  route_id           serial,
  contact_person_1   varchar(255),
  contact_abstract_1 varchar(255),
  contact_person_2   varchar(255),
  contact_abstract_2 varchar(255),
);



