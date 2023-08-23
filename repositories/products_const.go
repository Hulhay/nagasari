package repositories

var (
	GetProductsByStoreIDQuery = `
		select p.id, p.uuid, p.name, p.price, p.store_id, p.is_sold_out, p.product_photo_url, p.created_at
		from products p 
		join stores s on s.id = p.store_id
		where s.id = $1
	`
)
