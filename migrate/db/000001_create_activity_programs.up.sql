CREATE TABLE activity_programs (
  id                 serial,
  user_id            varchar(255),
  datetime           timestamp,
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
  create_time        timestamp default CURRENT_TIMESTAMP,
  update_time        timestamp default CURRENT_TIMESTAMP,
  CONSTRAINT activity_programs_PK PRIMARY KEY (id)
);

INSERT INTO activity_programs
VALUES
    (
        1,
        'user_id',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        'practice_section',
        'practice_Contents',
        1,
        1,
        'contact_person_1',
        'contact_abstract_1',
        'contact_person_2',
        'contact_abstract_2'
    );
