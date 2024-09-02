package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

// / @dev global variables
var (
	imageFilePath   string
	imageFileFormat string
)

/// imageCmd hold Command struct from cobra library
/// this struct contains the detail of the command
/// Run is the anonymous function that will be called when the command is executed, it's param are command and args(argument passed to the command)

var imageCmd = &cobra.Command{
	Use:     "image [your question] --path [image_path] --format [image_format]",
	Example: "gemini-cli image 'what is this image about?' --path image.png --format png",
	Args:    cobra.MinimumNArgs(1),
	Short:   "Know details about an image",
	Run: func(cmd *cobra.Command, args []string) {
		res := imageFunc(args)
		fmt.Println(res)
	},
}

// / @dev imageFunc is the function that will be called when the command is executed (this function is called in the Run field of the imageCmd)
// / @param args is the argument passed to the command
// / @return string is the response from the imageFunc
// / using join it combines all argument into a single string
// / Creating ctx that controls the lifecycle of the request
// / Creating new client to interact with Gemini API, fetched using gemini_api_key
// / if error occurs, it will be logged and the program will exit
// / creating model from the client, using gemini-1.5-flash model
// / reading the image file from the imageFilePath and storing it in imgData, if it cannot be read, it will be logged and the program will exit
// / prompt is the data sent to the model, it contains the image and the text
// / calling GenerateContent method with the prompt and model, it returns a response and error, if error occurs, it will be logged and the program will exit
// / Response Processing: extracting the final response from the response object and converting it to string that is returned
func imageFunc(args []string) string {
	userArgs := strings.Join(args[0:], " ")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	imgData, err := os.ReadFile(imageFilePath)
	if err != nil {
		log.Fatalf("Failed to read image file: %v", err)
	}

	// Supports multiple image inputs
	prompt := []genai.Part{
		genai.ImageData(imageFileFormat, imgData),
		genai.Text(userArgs),
	}
	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Panic(err)
	}
	finalResponse := resp.Candidates[0].Content.Parts[0]
	return fmt.Sprint(finalResponse)
}

// Flag definition: accesses the flag set associated with the imageCmd
// StringVarP is used to define a string flag with a shorthand and a longhand name, it takes 4 params,
// the first is the address of the variable to store the value, the second is the flag name, the third is the default value, and the fourth is the usage of the flag
// Ensuring that path and format flags are required, if user doesnot provide it, the command will fail
// check if there was an error in marking the flags as required, if there is, it will log the error and exit
func init() {
	imageCmd.Flags().StringVarP(&imageFilePath, "path", "p", "", "Enter the image path")
	imageCmd.Flags().StringVarP(&imageFileFormat, "format", "f", "", "Enter the image format (jpeg, png, etc.)")
	errPathF := imageCmd.MarkFlagRequired("path")
	errFormatF := imageCmd.MarkFlagRequired("format")
	if errPathF != nil || errFormatF != nil {
		log.Panic(errPathF, errFormatF)
	}
}
