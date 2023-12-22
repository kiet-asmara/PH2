export POSTGRESQL_URL='postgres://postgres:12345@localhost:5432/ngc10?sslmode=disable'

migrate create -ext sql -dir migrations -seq create_users_profile_table

migrate -database ${POSTGRESQL_URL} -path migrations up

migrate create -ext sql -dir migrations -seq add_role_to_users

migrate -database ${POSTGRESQL_URL} -path migrations up

migrate -database ${POSTGRESQL_URL} -path migrations down