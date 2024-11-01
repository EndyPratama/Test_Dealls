package subscription

const (
	readSubscription = `
	SELECT 
	    subscription.id,
	    subscription.name,
	    subscription.value,
	    COALESCE(subscription.created_at, TIMESTAMP("0001-01-01")) as created_at,
	    COALESCE(subscription.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
	    subscription`

	detailSubscription = `
	SELECT 
		subscription.id,
		subscription.name,
		subscription.value,
		COALESCE(subscription.created_at, TIMESTAMP("0001-01-01")) as created_at,
		COALESCE(subscription.updated_at, TIMESTAMP("0001-01-01")) as updated_at
	FROM
		subscription
	WHERE
		id = ?`

	createSubscription = `
	INSERT INTO subscription (
	    name,
	    value,
	    created_at
	) VALUES (
	    :name,
	    :value,
	    :created_at
	)`

	updateSubscription = `
	UPDATE 
	    subscription 
	SET
	    name          = :name,
	    value        = :value,
	    updated_at    = :updated_at
	WHERE 
	    id = :id`
)
