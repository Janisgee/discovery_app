-- name: CreateUserBookmark :one
INSERT INTO users_bookmark (id, user_id, username, place_id, place_name, catagory, place_text, created_at)
VALUES(
  gen_random_uuid(),
  $1, $2, $3, $4,$5, $6, NOW()
)
RETURNING *;

-- name: GetUserBookmark :one
SELECT * FROM users_bookmark WHERE place_id = $1 AND user_id = $2 LIMIT 1;

-- name: GetAllUserBookmarkPlaceID :many
SELECT users_bookmark.place_id, users_bookmark.place_name,users_bookmark.catagory, users_bookmark.place_text, places.country, places.city  FROM users_bookmark
INNER JOIN places ON place_id = places.id 
WHERE user_id = $1;

-- name: DeleteUserBookmark :one
DELETE FROM users_bookmark WHERE place_id = $1 AND user_id = $2 
RETURNING *;
