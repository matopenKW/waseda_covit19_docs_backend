CREATE TABLE places (
  id     SERIAL NOT NULL,
  name   VARCHAR(255) NOT NULL,
  hide   boolean,
  CONSTRAINT places_PK PRIMARY KEY (id)
);

INSERT INTO places (name, hide) VALUES
	 ('学生会館',NULL),
	 ('奉仕園',NULL),
	 ('四ツ木地区センター',NULL),
	 ('杉並公会堂',NULL),
	 ('堀切地区センター',NULL);