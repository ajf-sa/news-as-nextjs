-- name: AddNewPage :exec
INSERT INTO pages(title,descr,parent_id)values($1,$2,$3);

-- name: ListPages :many
SELECT * FROM pages;

-- name: EnablePage :exec
UPDATE pages SET is_active=true;

-- name: DisablePage :exec
UPDATE pages SET is_active=false;