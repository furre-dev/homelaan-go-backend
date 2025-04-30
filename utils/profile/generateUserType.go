package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/furre-dev/homelaan-go-backend/graph/model"
	"github.com/furre-dev/homelaan-go-backend/utils"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

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
	messages = append(messages, openai.AssistantMessage(utils.UserTypeInstructions))

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
