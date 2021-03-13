CREATE TABLE users (
  id              varchar(128) NOT NULL,
  email           varchar(255) NOT NULL,
  name            varchar(32),
  university_name varchar(32),
  faculty_name    varchar(32),
  student_no      varchar(32),
  cell_phon_no    varchar(11),
  create_time     timestamp default CURRENT_TIMESTAMP,

  CONSTRAINT users_PK PRIMARY KEY (id)
);
