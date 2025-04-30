package utils

import (
	"context"
	"fmt"

	"github.com/furre-dev/homelaan-go-backend/graph/model"
	"github.com/furre-dev/homelaan-go-backend/internal"
	"github.com/jackc/pgx/v5"
)

func MapProfileInfoToUser(userInput *model.InvestorProfileInput) *model.InvestorProfile {
	userProfile := &model.InvestorProfile{
		GeoInfo: &model.PersonalGeographicProfile{
			Location:                          (*model.Location)(userInput.GeoInfo.Location),
			YearsAbroad:                       userInput.GeoInfo.YearsAbroad,
			WillStayLongTermInCurrentLocation: userInput.GeoInfo.WillStayLongTermInCurrentLocation,
			HasFamilyAbroad:                   userInput.GeoInfo.HasFamilyAbroad,
			ResidencyStatus:                   userInput.GeoInfo.ResidencyStatus,
		},
		ProfessionalBackground: (*model.ProfessionalBackground)(userInput.ProfessionalBackground),
		NetworkMarketAccess:    (*model.NetworkMarketAccess)(userInput.NetworkMarketAccess),
		EngagementPreferences:  (*model.EngagementPreferences)(userInput.EngagementPreferences),
		InvestmentAppetite: &model.InvestmentAppetite{
			ActiveInvestorInStartupsOrSmallMidBusinesses: userInput.InvestmentAppetite.ActiveInvestorInStartupsOrSmallMidBusinesses,
			InvestmentRange: &model.InvestmentRange{
				FromUsd: userInput.InvestmentAppetite.InvestmentRange.FromUsd,
				ToUsd:   userInput.InvestmentAppetite.InvestmentRange.ToUsd,
			},
			AngelGroupMembership: userInput.InvestmentAppetite.AngelGroupMembership,
			InvolvementType:      userInput.InvestmentAppetite.InvolvementType,
		},
		AvailabilityAndCommitment: (*model.AvailabilityCommitment)(userInput.AvailabilityAndCommitment),
		SpousalInvolvement:        (*model.SpousalInvolvement)(userInput.SpousalInvolvement),
		AdditionalInfo:            userInput.AdditionalInfo,
	}

	return userProfile
}

func AssignUserType(ctx context.Context, db *pgx.Conn, userType model.UserType) (*model.User, error) {
	userID, ok := internal.GetUserID(ctx)

	if !ok || userID == "" {
		return nil, fmt.Errorf("unauthenticated")
	}

	query := "UPDATE users SET user_type = $1 WHERE id = $2 RETURNING *"

	// Assuming userType.UserID contains the ID of the user to update.
	// Also assuming that userType.Type contains the new user type to assign.
	var updatedUser model.User
	err := db.QueryRow(ctx, query, userType, userID).Scan(
		&updatedUser.ID,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.UserType,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating user type: %w", err)
	}

	return &updatedUser, nil
}

func GetCurrentUserFromPostgres(ctx context.Context, db *pgx.Conn) (*model.User, error) {
	userID, ok := internal.GetUserID(ctx)

	if !ok || userID == "" {
		return nil, fmt.Errorf("unauthenticated")
	}

	query := "SELECT * FROM users WHERE id = $1"

	row := db.QueryRow(ctx, query, userID)

	var user model.User

	rowError := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.UserType)
	if rowError != nil {
		return nil, rowError
	}

	return &user, nil
}
