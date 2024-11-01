package profile

const (
	readProfile = `
	SELECT 
	    profile.id,
	    profile.user_id,
	    profile.name,
	    profile.gender,
	    profile.bio,
	    profile.birthdate,
	    profile.location,
	    profile.subscription_id,
	    profile.created_at,
	    COALESCE(profile.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    profile`

	SearchProfilePeople = `
	SELECT 
	    profile.id,
	    profile.user_id,
	    profile.name,
	    profile.gender,
	    profile.bio,
	    profile.birthdate,
	    profile.location,
	    profile.created_at,
		profile.subscription_id,
	    COALESCE(profile.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    profile
	Where
		user_id != ? 
		AND gender != ?
	ORDER BY rand()`

	detailProfile = `
	SELECT
	    profile.id,
	    profile.user_id,
	    profile.name,
	    profile.gender,
	    profile.bio,
	    profile.birthdate,
	    profile.location,
		profile.subscription_id,
	    profile.created_at,
	    COALESCE(profile.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    profile
	WHERE 
		id = ?`

	detailProfileByUserID = `
	SELECT
	    profile.id,
	    profile.user_id,
	    profile.name,
	    profile.gender,
	    profile.bio,
	    profile.birthdate,
	    profile.location,
		profile.subscription_id,
	    profile.created_at,
	    COALESCE(profile.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    profile
	WHERE 
		user_id = ?`

	createProfile = `
	INSERT INTO profile (
	    user_id,
	    name,
	    gender,
	    bio,
	    birthdate,
	    location,
		subscription_id,
	    created_at
	) VALUES (
	 	:user_id,
	    :name,
	    :gender,
	    :bio,
	    :birthdate,
	    :location,
		:subscription_id,
	    :created_at
	)`

	updateProfile = `
	UPDATE 
	    profile 
	SET
	    user_id              = :user_id,
	    name                 = :name,
	    gender               = :gender,
	    bio                  = :bio,
	    birthdate            = :birthdate,
	    location             = :location,
	    subscription_id      = :subscription_id,
	    updated_at           = :updated_at
	WHERE 
	    id = :id`

	delete = `
	UPDATE
	    profile
	SET
	    profile.is_deleted = :is_deleted,
	    profile.deleted_at = :deleted_at,
	    profile.deleted_by = :deleted_by
	WHERE
	    profile.id = :id`
)
