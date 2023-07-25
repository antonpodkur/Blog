-- name: CreateArticle :one
INSERT INTO articles (
  title,
  content,
  updated_at,
  user_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetArticleById :one
SELECT * FROM articles
WHERE id = $1 LIMIT 1;

-- name: GetArticleByUserId :one
SELECT * FROM articles
WHERE user_id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateArticle :one
UPDATE articles
set title = $2,
content = $3,
updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;
