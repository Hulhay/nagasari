package repositories

var (
	insertUser = `
		insert into users (uuid, name, email, phone_number, password, created_at)
		values (uuid_generate_v4(), $1, $2, $3, $4, now())
	`

	getUserByEmail = `
		select u.id
		from users u
		where u.email = $1
	`

	getUserByPhoneNumber = `
		select u.id
		from users u
		where u.phone_number = $1
	`
)
