-- name: CreateFeed :exec
insert into
    feeds (
        id,
        created_at,
        updated_at,
        name,
        url,
        user_id
    )
values ($1, $2, $3, $4, $5, $6);

-- name: ShowFeeds :many
SELECT f.*, u.name as username
FROM feeds f
    LEFT JOIN users u on f.user_id = u.id
ORDER BY f.name;