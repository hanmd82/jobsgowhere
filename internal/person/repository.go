package person

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jobsgowhere/jobsgowhere/internal/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

const errSqlNoRows = "sql: no rows in result set"

// Repository interface for person
type Repository interface {
	GetProfile(ctx context.Context, iamID string) (*models.Person, error)
	CreateProfile(ctx context.Context, iamID string, params CreateProfileParams) (*models.Person, error)
}

// personRepository struct
type personRepository struct {
	executor boil.ContextExecutor
}

func (repo *personRepository) GetProfile(ctx context.Context, iamID string) (*models.Person, error) {
	person, err := models.People(
		qm.Load(models.PersonRels.PersonProfiles),
		qm.Load(models.PersonRels.JobProvider),
		models.PersonWhere.IamID.EQ(iamID)).One(ctx, repo.executor)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("profile_not_found")
		}

		return nil, err
	}

	return person, nil
}

func (repo *personRepository) CreateProfile(ctx context.Context, iamID string, params CreateProfileParams) (*models.Person, error) {
	u1, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	u2, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	person, err := models.People(
		qm.Load(models.PersonRels.PersonProfiles),
		qm.Load(models.PersonRels.JobProvider),
		models.PersonWhere.IamID.EQ(iamID)).One(ctx, repo.executor)

	if err != nil {
		if err == sql.ErrNoRows {
			person = &models.Person{
				IamID:          iamID,
				FirstName:      null.StringFrom(params.FirstName),
				LastName:       null.StringFrom(params.LastName),
				CurrentCompany: null.StringFrom(params.Company),
				Email:          params.Email,
				IamProvider:    "LinkedIn",
				AvatarURL:      null.StringFrom(params.AvartarURL),
				ID:             u2.String(),
			}

			err = person.Insert(ctx, repo.executor, boil.Infer())

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if person.R == nil || len(person.R.PersonProfiles) == 0 {
		personProfile := &models.PersonProfile{
			ID:         u1.String(),
			PersonID:   u2.String(),
			ProfileURL: "dfds", // TODO: This is to be changed later.
		}

		err = personProfile.Insert(ctx, repo.executor, boil.Infer())

		if err != nil {
			return nil, err
		}

		if params.ProfileType == Recruiter.String() {
			jobProvider := &models.JobProvider{
				HuntingMode: null.IntFrom(1),
				Title:       params.Headline,
				WebsiteURL:  null.StringFrom(params.CompanyWebsite),
				PersonID:    u2.String(),
			}

			err = jobProvider.Insert(ctx, repo.executor, boil.Infer())

			if err != nil {
				return nil, err
			}
		}

		person, err := models.People(
			qm.Load(models.PersonRels.PersonProfiles),
			qm.Load(models.PersonRels.JobProvider),
			models.PersonWhere.IamID.EQ(iamID)).One(ctx, repo.executor)

		if err != nil {
			return nil, err
		}

		return person, nil
	}

	return person, nil
}
