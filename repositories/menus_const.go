package repositories

var (
	GetMenusByStoreIDQuery = `
		select m.id, m.uuid, m.name, m.price, m.store_id, m.is_sold_out, m.created_at
		from menus m 
		join stores s on s.id = m.store_id
		where s.id = $1
	`
)
