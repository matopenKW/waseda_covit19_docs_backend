CREATE TABLE activity_programs (
  user_id              varchar(255),
  seq_no               serial,
  datetime             varchar(8),
  start_time           varchar(4),
  end_time             varchar(4),
  practice_section_id  smallint,
  practice_Contents_id smallint,
  outward_trip         varchar(255),
  return_trip          varchar(255),
  contact_person1      integer,
  contact_abstract1    varchar(255),
  contact_person2      integer,
  contact_abstract2    varchar(255),
  create_time          timestamp default CURRENT_TIMESTAMP,
  update_time          timestamp default CURRENT_TIMESTAMP,
  CONSTRAINT activity_programs_PK PRIMARY KEY (user_id, seq_no)
);