package likes

const (
	readLikes = `
	SELECT 
	    likes.id,
	    likes.liker_id,
	    likes.liked_id,
	    likes.created_at,
	    COALESCE(likes.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    likes`

	readLikesByLikerID = `
	SELECT 
	    likes.id,
	    likes.liker_id,
	    likes.liked_id,
	    likes.created_at,
	    COALESCE(likes.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    likes
	Where
		liker_id = ?`

	detailLikes = `
	SELECT
	    likes.id,
		likes.liker_id,
	    likes.liked_id,
	    likes.created_at,
	    COALESCE(likes.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    likes
	WHERE 
		id = ?`

	checkLikes = `
	SELECT
	    likes.id,
		likes.liker_id,
	    likes.liked_id,
	    likes.created_at,
	    COALESCE(likes.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    likes
	WHERE 
		liker_id = ? AND liked_id = ?`

	createLikes = `
	INSERT INTO likes (
	    liker_id,
	    liked_id,
	    created_at
	) VALUES (
	    :liker_id,
	    :liked_id,
	    :created_at
	)`

	delete = `
	DELETE FROM likes WHERE likes.id = :id`
)
