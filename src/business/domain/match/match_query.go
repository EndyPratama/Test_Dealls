package match

const (
	readMatch = `
	SELECT 
	    matches.id,
	    matches.profile1_id,
	    matches.profile2_id,
	    matches.matched_at,
	    matches.created_at,
	    COALESCE(matches.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    matches`

	readMatchByProfile = `
	SELECT 
	    matches.id,
	    matches.profile1_id,
	    matches.profile2_id,
	    matches.matched_at,
	    matches.created_at,
	    COALESCE(matches.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    matches
	WHERE profile1_id = ? OR profile2_id = ?`

	detailMatch = `
	SELECT
	    matches.id,
		matches.profile1_id,
		matches.profile2_id,
	    matches.matched_at,
	    matches.created_at,
	    COALESCE(matches.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    matches
	WHERE 
		id = ?`

	createMatch = `
	INSERT INTO matches (
	    profile1_id,
	    profile2_id,
	    matched_at,
	    created_at
	) VALUES (
	    :profile1_id,
	    :profile2_id,
	    :matched_at,
	    :created_at
	)`

	delete = `
	DELETE FROM matches WHERE id = :id`
)
