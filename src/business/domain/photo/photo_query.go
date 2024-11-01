package photo

const (
	readPhoto = `
	SELECT 
	    photo.id,
	    photo.profile_id,
	    photo.photo_url,
	    photo.created_at,
	    COALESCE(photo.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    photo`

	detailPhoto = `
	SELECT
	    photo.id,
		photo.profile_id,
	    photo.photo_url,
	    photo.created_at,
	    COALESCE(photo.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    photo
	WHERE 
		user_id = ?`

	createPhoto = `
	INSERT INTO photo (
	    profile_id,
	    photo_url,
	    created_at
	) VALUES (
	    :profile_id,
	    :photo_url,
	    :created_at
	)`

	updatePhoto = `
	UPDATE 
	    photo 
	SET
	    profile_id    = :profile_id,
	    photo_url     = :photo_url,
	    updated_at    = :updated_at
	WHERE 
	    id = :id`
)
