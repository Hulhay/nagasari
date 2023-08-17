package repositories

import (
	"context"
	"database/sql"

	"github.com/hulhay/nagasari/lib/utils"
	"github.com/hulhay/nagasari/models"
)

func (r *repository) GetUserByEmailFromDB(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.qry.Read().GetContext(ctx, &user, getUserByEmail, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetUserByPhoneNumberFromDB(ctx context.Context, phoneNumber string) (*models.User, error) {
	var user models.User

	err := r.qry.Read().GetContext(ctx, &user, getUserByPhoneNumber, phoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *repository) InsertUserToDB(ctx context.Context, req models.RegisterRequest) error {
	res, err := r.qry.Write().ExecContext(ctx,
		insertUser,
		req.Name,
		req.Email,
		req.PhoneNumber,
		req.EncryptedPassword,
	)
	if err != nil {
		return err
	}

	rowAffected, _ := res.RowsAffected()
	if rowAffected == 0 {
		return utils.ErrNoRowsAffected
	}

	return nil
}
