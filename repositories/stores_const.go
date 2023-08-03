package repositories

var (
	getStoresByKeywordName = `(s.store_name ilike '%' || $1 || '%')`

	getStoresQuery = `
		select s.id, s.uuid, s.store_name, s.store_photo_url, u.uuid "owner_uuid", u.name "owner_name", u.phone_number "owner_phone_number", s.created_at
		from stores s
		join users u on s.owner_id = u.id
	`
)
