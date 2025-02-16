CREATE TABLE IF NOT EXISTS questions (
  id   BIGSERIAL PRIMARY KEY,
  question text      NOT NULL,
  answer  text
);

CREATE TABLE IF NOT EXISTS quiz (
  id         BIGSERIAL PRIMARY KEY,
  title      TEXT NOT NULL,
  question_ids BIGINT[] REFERENCES questions(id) ON DELETE CASCADE ON UPDATE CASCADE 
);
