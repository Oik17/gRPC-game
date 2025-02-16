-- Create Questions Table
CREATE TABLE questions (
  id BIGSERIAL PRIMARY KEY,
  question TEXT NOT NULL,
  answer TEXT
);

-- Create Quiz Table
CREATE TABLE quiz (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  question_ids BIGINT[] REFERENCES questions(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- name: CreateQuestion :one
INSERT INTO questions (question, answer) 
VALUES ($1, $2) 
RETURNING *;

-- name: UpdateQuestion :one
UPDATE questions 
SET question = $2, answer = $3 
WHERE id = $1 
RETURNING *;

-- name: DeleteQuestion :one
DELETE FROM questions 
WHERE id = $1 
RETURNING *;

-- name: GetQuestionById :one
SELECT * FROM questions 
WHERE id = $1;

-- name: GetAllQuestions :many
SELECT * FROM questions 
ORDER BY id;

-- name: CreateQuiz :one
INSERT INTO quiz (title, question_ids) 
VALUES ($1, $2) 
RETURNING *;

-- name: UpdateQuiz :one
UPDATE quiz 
SET title = $2, question_ids = $3 
WHERE id = $1 
RETURNING *;

-- name: DeleteQuiz :one
DELETE FROM quiz 
WHERE id = $1 
RETURNING *;

-- name: GetQuizById :one
SELECT * FROM quiz 
WHERE id = $1;

-- name: GetAllQuizzes :many
SELECT * FROM quiz 
ORDER BY id;

-- Get a Quiz with Full Question Details
-- name: GetQuizWithQuestions :many
SELECT q.id AS quiz_id, q.title, qs.id AS question_id, qs.question, qs.answer
FROM quiz q
JOIN questions qs ON qs.id = ANY(q.question_ids)
WHERE q.id = $1;
