package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/furre-dev/homelaan-go-backend/graph/model"
	"github.com/furre-dev/jsonschema-nullable"
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
