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
SELECT a.id, a.title, a.content, a.created_at, a.updated_at, u.name as user_name, u.photo as user_photo FROM articles a
LEFT JOIN users u on u.id = a.user_id
WHERE a.id = $1 LIMIT 1;

-- name: GetArticleByUserId :one
SELECT a.id, a.title, a.content, a.created_at, a.updated_at, u.name as user_name, u.photo as user_photo FROM articles a
LEFT JOIN users u on u.id = a.user_id
WHERE user_id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT a.id, a.title, a.content, a.created_at, a.updated_at, u.name as user_name, u.photo as user_photo FROM articles a
LEFT JOIN users u on u.id = a.user_id
ORDER BY a.id
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
