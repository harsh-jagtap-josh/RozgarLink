package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// PostgreSQL Queries
const (
	createAddressQuery            = `INSERT INTO address (details, street, city, state, pincode) VALUES (:details, :street, :city, :state, :pincode) RETURNING *`
	updateAddressQuery            = "UPDATE address SET details=:details, street=:street, city=:city, state=:state, pincode=:pincode WHERE id=:id RETURNING *;"
	fetchAddressByIdQuery         = "SELECT * FROM address where id=$1;"
	fetchAddressByWorkerIdQuery   = "SELECT address.* FROM address inner join workers on address.id = workers.location where workers.id=$1;"
	fetchAddressByEmployerIdQuery = "SELECT address.* FROM address inner join employers on address.id = employers.location where employers.id=$1;"
	fetchAddressByJobIdQuery      = "SELECT address.* FROM address inner join jobs on address.id = jobs.location where jobs.id=$1;"
	deleteAddressByIdQuery        = "DELETE FROM address WHERE id=$1;"
)

// create a new address and return newly created address object, and error
func CreateAddress(ctx context.Context, sqlxDb *sqlx.DB, addressData Address) (Address, error) {

	var newAddress Address
	rows, err := sqlxDb.NamedQuery(createAddressQuery, addressData)
	if err != nil {
		return Address{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&newAddress)
		if err != nil {
			return Address{}, err
		}
	}
	return newAddress, nil
}

// update address based on ID, and return updated address, and error
func UpdateAddress(ctx context.Context, sqlxDb *sqlx.DB, addressData Address) (Address, error) {

	var address Address
	rows, err := sqlxDb.NamedQuery(updateAddressQuery, addressData)
	if err != nil {
		return Address{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&address)
		if err != nil {
			return Address{}, err
		}
	}

	return address, nil
}

// delete address and return ID of address obj deleted, and error
func DeleteAddress(ctx context.Context, sqlxDb *sqlx.DB, addressId int) error {
	_, err := sqlxDb.Exec(deleteAddressByIdQuery, addressId)
	if err != nil {
		return err
	}

	return nil
}

// fetch address by id
func GetAddressById(ctx context.Context, sqlxDb *sqlx.DB, addressId int) (Address, error) {

	var address Address

	err := sqlxDb.Get(&address, fetchAddressByIdQuery, addressId)

	if err != nil {
		return Address{}, err
	}

	return address, nil
}

// fetch address by worker id
func GetAddressByWorkerId(ctx context.Context, sqlxDb *sqlx.DB, workerId int) (Address, error) {
	var address Address

	err := sqlxDb.Get(&address, fetchAddressByWorkerIdQuery, workerId)

	if err != nil {
		fmt.Println(err.Error())
		return Address{}, err
	}
	return address, nil
}

func GetAddressByEmployerId(ctx context.Context, sqlxDb *sqlx.DB, employerId int) (Address, error) {
	var address Address

	err := sqlxDb.Get(&address, fetchAddressByWorkerIdQuery, employerId)

	if err != nil {
		return Address{}, err
	}

	return address, nil
}

func GetAddressByJobId(ctx context.Context, sqlxDb *sqlx.DB, jobId int) (Address, error) {
	var address Address

	err := sqlxDb.Get(&address, fetchAddressByJobIdQuery, jobId)

	if err != nil {
		return Address{}, err
	}

	return address, nil
}
