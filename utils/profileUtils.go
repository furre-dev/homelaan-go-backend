package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/furre-dev/homelaan-go-backend/graph/model"
	"github.com/furre-dev/homelaan-go-backend/internal"
	"github.com/furre-dev/jsonschema-nullable"
	"github.com/jackc/pgx/v5"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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

func GenerateSchema[T any]() interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)

	return schema
}

// Generate the JSON schema at initialization time
var InvestorProfileJsonSchema = GenerateSchema[model.InvestorProfile]()

/* func AssignUserType(ctx context.Context, userType model.UserType) (*model.User, error) {

} */

func GenerateUserProfile(questionsWithAnswers []*model.QuestionInput) (*model.InvestorProfile, error) {
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "generate_user_profile",
		Description: openai.String("Generate user profile"),
		Schema:      InvestorProfileJsonSchema,
		Strict:      openai.Bool(true),
	}

	client := openai.NewClient(
		option.WithAPIKey(""), // defaults to os.LookupEnv("OPENAI_API_KEY")
	)

	var messages []openai.ChatCompletionMessageParamUnion
	messages = append(messages, openai.AssistantMessage(UserProfileInstructs))

	for _, qa := range questionsWithAnswers {
		messages = append(messages, openai.UserMessage(fmt.Sprintf(
			"\n------------------------------------------------------------------\nMessage: %s,\nProfileFieldInSchema: %s, \n------------------------------------------------------------------\n, \nUser answer: %s\n------------------------------------------------------------------\n",
			qa.QuestionTitle, qa.ProfileField, qa.UserInputAnswer,
		)))
	}

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    openai.ChatModelGPT4o,
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: schemaParam},
		},
	})

	if err != nil {
		panic(err.Error())
	}

	var investorProfile model.InvestorProfile
	//parse the string into the InvestorProfile object
	if err := json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &investorProfile); err != nil {
		panic(err.Error())
	}

	return &investorProfile, nil
}

type UserTypeResponseSchema struct {
	UserType model.UserType
}

var UserTypeResponseJsonSchema = GenerateSchema[UserTypeResponseSchema]()

func GenerateUserType(questionWithAnswers *model.QuestionInput) (*model.UserType, error) {
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "generate_user_type",
		Description: openai.String("Generate user type"),
		Schema:      UserTypeResponseJsonSchema,
		Strict:      openai.Bool(true),
	}

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY_GQL")),
	)

	var messages []openai.ChatCompletionMessageParamUnion
	messages = append(messages, openai.AssistantMessage(UserTypeInstructions))

	messages = append(messages, openai.UserMessage(fmt.Sprintf(`
		This was the question: %s,
		This was the user input answer: %s
	`, questionWithAnswers.QuestionTitle, questionWithAnswers.UserInputAnswer)))

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    openai.ChatModelGPT4o,
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{JSONSchema: schemaParam},
		},
	})

	if err != nil {
		panic(err.Error())
	}

	var userType model.UserType
	//parse the string into the InvestorProfile object
	if err := json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &userType); err != nil {
		panic(err.Error())
	}

	return &userType, nil
}

func AssignUserType(ctx context.Context, db *pgx.Conn, userType model.UserType) (*model.User, error) {
	userID, ok := internal.GetUserID(ctx)

	if !ok || userID == "" {
		return nil, fmt.Errorf("unauthenticated")
	}

	query := "UPDATE users SET user_type = $1 WHERE id = $2 RETURNING id, first_name, last_name, email, user_type"

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
