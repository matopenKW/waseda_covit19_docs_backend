CREATE TABLE last_uploads (
  week_id     smallint NOT NULL,
  drive_id    varchar(255) NOT NULL,
  update_time timestamp default CURRENT_TIMESTAMP,
  CONSTRAINT last_uploads_PK PRIMARY KEY (week_id)
);

INSERT INTO last_uploads (week_id, drive_id) VALUES (0, '');
INSERT INTO last_uploads (week_id, drive_id) VALUES (1, '');
INSERT INTO last_uploads (week_id, drive_id) VALUES (2, '');
INSERT INTO last_uploads (week_id, drive_id) VALUES (3, '');
INSERT INTO last_uploads (week_id, drive_id) VALUES (4, '');
INSERT INTO last_uploads (week_id, drive_id) VALUES (5, '');
INSERT INTO last_uploads (week_id, drive_id) VALUES (6, '');
