-- name: CreatePost :one
INSERT INTO posts (
    id,
    title,
    description,
    url,
    published_at,
    feed_id,
    created_at,
    updated_at
    ) 
    VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
    RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.* FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY published_at DESC
LIMIT $2;
