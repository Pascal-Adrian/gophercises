-- name: GetPathByPath :one
SELECT *
FROM paths
WHERE path = ?1;