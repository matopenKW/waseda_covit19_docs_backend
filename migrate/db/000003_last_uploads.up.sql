CREATE TABLE last_uploads (
  drive_id    varchar(255) NOT NULL,
  update_time timestamp default CURRENT_TIMESTAMP
);

INSERT INTO last_uploads (drive_id) VALUES ('test_drive_id');
