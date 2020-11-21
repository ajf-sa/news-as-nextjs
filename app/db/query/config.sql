-- name: SetConfig :exec
INSERT INTO config(seo_title,seo_descript)values($1,$2);