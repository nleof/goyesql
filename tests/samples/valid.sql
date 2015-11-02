-- name: simple
SELECT * FROM simple;

-- name: multiline
SELECT *
FROM multiline
WHERE line = 42;


-- name: comments
-- yoyo

SELECT *
-- inline
FROM comments;
