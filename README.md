# news-ajf-sa
موجز اخباري

# users table
    - id
    - showname
    - username
    - password
    - email
    - phone

# catogres table
    - id
    - name
    - description
    - parent_id
    - is_active


# post table
    - id
    - user_id
    - cat_id
    - name
    - description
    - body
    - is_active
    - create_at
    - publish_at

# tags table
    - id
    - post_id
    - name
    - slug