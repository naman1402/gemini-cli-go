package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var numWords string = "150"

var searchCmd = &cobra.Command{
	Use:   "search [question]",
	Short: "Ask a question and get a response",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := getApiResponse(args)
		fmt.Println(res)
	},
}

// @dev getApiResponse is the function that will be called when the command is executed
// @param args is the argument passed to the command
// @return string is the response from the getApiResponse
// first of all, it joins all the arguments into a single string into userArgs
// creating a new client to interact with Gemini API, fetched using gemini_api_key
// if error occurs, it will be logged and the program will exit
// validating numWords to make sure it's a number, if not, it will be logged and the program will exit
// creating a new model, sending users query along with the specified number of words to Gemini API
// if content generation fails, it will be logged and the program will exit
// extracting the final response from the API response
// formatting the response using the formatAsText function
// and returning the formatted response
func getApiResponse(args []string) string {
	userArgs := strings.Join(args[0:], " ")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	_, err = strconv.Atoi(numWords)
	if err != nil {
		log.Panic(err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userArgs+"in"+numWords+" words."))
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	finalResponse := resp.Candidates[0].Content.Parts[0]
	return formatAsText(fmt.Sprint(finalResponse))
}

func formatAsText(input string) string {

	// removing headers
	re := regexp.MustCompile(`^#{1,6}\s+(.*)`)
	input = re.ReplaceAllString(input, "$1")
	// removing horizontal rules
	re = regexp.MustCompile(`\n---\n`)
	input = re.ReplaceAllString(input, "\n")
	// removing italic text
	re = regexp.MustCompile(`_([^_]+)_`)
	input = re.ReplaceAllString(input, "$1")
	// removing bold text
	re = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	input = re.ReplaceAllString(input, "$1")
	// removing bullet points
	re = regexp.MustCompile(`\n\* (.*)`)
	input = re.ReplaceAllString(input, "\n- $1")
	// returning the formatted response (cleaned)
	return input
}

// accessing the flags of the search command
// initializing the numWords flag with a default value of 150
func init() {
	searchCmd.Flags().StringVarP(&numWords, "words", "w", "150", "number of words")
}
