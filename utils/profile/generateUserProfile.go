package profile

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/furre-dev/homelaan-go-backend/graph/model"
	"github.com/furre-dev/homelaan-go-backend/utils"
	"github.com/jackc/pgx/v5"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Generate the JSON schema at initialization time
var InvestorProfileJsonSchema = GenerateSchema[model.InvestorProfile]()

func GenerateUserProfile(ctx context.Context, db *pgx.Conn, questionsWithAnswers []*model.QuestionInput) (*model.InvestorProfile, error) {
	/* user, err := utils.GetCurrentUserFromPostgres(ctx, db)

	if err != nil {
		return nil, fmt.Errorf("unauthenticated")
	}

	userType := user.UserType */

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
	messages = append(messages, openai.AssistantMessage(utils.UserProfileInstructs))

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
