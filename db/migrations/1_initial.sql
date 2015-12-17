-- +migrate Up
CREATE TABLE users
(
  id      CHAR(36)      NOT NULL PRIMARY KEY,
  name    VARCHAR(128)  NOT NULL,
  email   VARCHAR(128)  NOT NULL UNIQUE,
  hash    CHAR(64)      NOT NULL,
  salt    CHAR(64)      NOT NULL UNIQUE
);

CREATE TABLE notes
(
  id      BIGINT        NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id CHAR(36)      NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE revisions
(
  id      BIGINT        NOT NULL,
  note_id BIGINT        NOT NULL,
  content BLOB,
  time    TIMESTAMP     NOT NULL,

  PRIMARY KEY (id, note_id),
  FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE revisions;
DROP TABLE notes;
DROP TABLE users;
