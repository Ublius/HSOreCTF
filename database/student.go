package database

import (
	"context"
	"database/sql"
	"fmt"
)

func (d *Database) GetStudentByEmail(ctx context.Context, email string) (*Student, error) {
	var student Student
	var parentEmail, signatory, dietaryRestrictions sql.NullString
	err := d.DB.QueryRow(ctx, `
		SELECT teamid, email, name, age, parentemail, signatory, ctfdpassword,
		EmailConfirmed, liabilitywaiver, computerusewaiver, dietaryrestrictions, qrcodesent, checkedin
		FROM students
		WHERE email = $1
	`, email).Scan(&student.TeamID, &student.Email, &student.Name, &student.Age,
		&parentEmail, &signatory, &student.CTFdPassword, &student.EmailConfirmed,
		&student.LiabilitySigned, &student.ComputerUseWaiverSigned,
		&dietaryRestrictions, &student.QRCodeSent, &student.CheckedIn)
	if err != nil {
		return nil, err
	}

	if parentEmail.Valid {
		student.ParentEmail = parentEmail.String
	}

	if signatory.Valid {
		student.Signatory = signatory.String
	}

	if dietaryRestrictions.Valid {
		student.DietaryRestrictions = dietaryRestrictions.String
	}

	// if campusTour.Valid {
	// 	student.CampusTour = campusTour.Bool
	// }

	return &student, err
}

func (d *Database) ConfirmStudent(ctx context.Context, email string, dietaryRestrictions, parentEmail string) error {
	_, err := d.DB.Exec(ctx, `
		UPDATE students
		SET emailconfirmed = true, dietaryrestrictions = $1, parentemail = $2
		WHERE email = $3
	`, dietaryRestrictions, parentEmail, email)
	return err
}

func (d *Database) SignFormsForStudent(ctx context.Context, email, signatory string) error {
	computerUseQuery := "computerusewaiver = true,"
	q := fmt.Sprintf(`
		UPDATE students
		SET liabilitywaiver = true, %s signatory = $1
		WHERE email = $2
	`, computerUseQuery)
	_, err := d.DB.Exec(ctx, q, signatory, email)
	return err
}

func (d *Database) GetAllDietaryRestrictions(ctx context.Context) ([]string, error) {
	rows, err := d.DB.Query(ctx, `
		SELECT dietaryrestrictions
		FROM students
		WHERE dietaryrestrictions != '' AND dietaryrestrictions IS NOT NULL
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dietaryRestrictions []string
	for rows.Next() {
		var restriction string
		if err = rows.Scan(&restriction); err != nil {
			return nil, err
		}
		dietaryRestrictions = append(dietaryRestrictions, restriction)
	}
	return dietaryRestrictions, nil
}

func (d *Database) MarkQRCodeSent(ctx context.Context, email string) error {
	_, err := d.DB.Exec(ctx, `
		UPDATE students
		SET qrcodesent = true
		WHERE email = $1
	`, email)
	return err
}

func (d *Database) CheckInStudent(ctx context.Context, email string) error {
	_, err := d.DB.Exec(ctx, `
		UPDATE students
		SET checkedin = true
		WHERE email = $1
	`, email)
	return err
}
