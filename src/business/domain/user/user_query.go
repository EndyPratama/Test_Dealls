package user

const (
	readUser = `
	SELECT 
	    user.id,
	    COALESCE(user.email, '') as email,
	    COALESCE(user.password, '') as password,
	    user.created_at,
	    COALESCE(user.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    user`

	detailUser = `
	SELECT
	    user.id,
	    COALESCE(user.email, '') as email,
	    COALESCE(user.password, '') as password,
	    user.created_at,
	    COALESCE(user.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    user
	WHERE 
		id = ?`

	login = `
	SELECT
	    user.id,
	    COALESCE(user.email, '') as email,
	    COALESCE(user.password, '') as password,
	    user.created_at,
	    COALESCE(user.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    user
	WHERE 
		email = ? AND password = ?`

	createUser = `
	INSERT INTO user (
	    email,
	    password,
	    created_at
	) VALUES (
	 	:email,
	    :password,
	    :created_at
	)`
)
