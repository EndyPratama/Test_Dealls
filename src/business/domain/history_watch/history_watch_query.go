package historywatch

const (
	readHistoryWatch = `
	SELECT 
	    history_watch.id,
	    history_watch.profile1_id,
	    history_watch.profile2_id,
	    history_watch.label,
	    history_watch.created_at,
	    COALESCE(history_watch.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    history_watch`

	readHistoryWatchByProfileID = `
	SELECT 
	    history_watch.id,
	    history_watch.profile1_id,
	    history_watch.profile2_id,
	    history_watch.label,
	    history_watch.created_at,
	    COALESCE(history_watch.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    history_watch
	WHERE
		profile1_id = :profile1_id`

	createHistoryWatch = `
	INSERT INTO history_watch (
	    profile1_id,
	    profile2_id,
	    label,
	    created_at
	) VALUES (
	    :profile1_id,
	    :profile2_id,
	    :label,
	    :created_at
	)`

	updateHistoryWatch = `
	UPDATE 
	    history_watch 
	SET
	    label        = :label,
	    updated_at   = :updated_at
	WHERE 
	    id = :id`

	delete = `
	DELETE FROM history_watch WHERE id = :id`
)
