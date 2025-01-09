package maritaca

import (
	"log"
	"net/http"
	"net/url"

	"github.com/silverfox07/langchaingo/llms/maritaca/internal/maritacaclient"
)

type options struct {
	maritacaServerURL   *url.URL
	httpClient          *http.Client
	model               string
	maritacaOptions     maritacaclient.Options
	customModelTemplate string
	system              string
	format              string
}

type Option func(*options)

// WithModel Set the model to use.
func WithModel(model string) Option {
	return func(opts *options) {
		opts.model = model
	}
}

// WithFormat Sets the maritaca output format (currently maritaca only supports "json").
func WithFormat(format string) Option {
	return func(opts *options) {
		opts.format = format
	}
}

// WithSystem Set the system prompt. This is only valid if
// WithCustomTemplate is not set and the maritaca model use
// .System in its model template OR if WithCustomTemplate
// is set using {{.System}}.
func WithSystemPrompt(p string) Option {
	return func(opts *options) {
		opts.system = p
	}
}

// WithCustomTemplate To override the templating done on maritaca model side.
func WithCustomTemplate(template string) Option {
	return func(opts *options) {
		opts.customModelTemplate = template
	}
}

// WithServerURL Set the URL of the maritaca instance to use.
func WithServerURL(rawURL string) Option {
	return func(opts *options) {
		var err error
		opts.maritacaServerURL, err = url.Parse(rawURL)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// WithHTTPClient Set custom http client.
func WithHTTPClient(client *http.Client) Option {
	return func(opts *options) {
		opts.httpClient = client
	}
}

// WithChatMode Set the chat mode.
// default: true
// If True, the model will run in chat mode, where messages is a string containing the
// user's message or a list of messages containing the iterations of the conversation
// between user and assistant. If False, messages must be a string containing the desired prompt.
func WithChatMode(chatMode bool) Option {
	return func(opts *options) {
		opts.maritacaOptions.ChatMode = chatMode
	}
}

// WithMaxTokens Set the maximum number of tokens that will be generated by the mode.
// minimum: 1
// Maximum number of tokens that will be generated by the mode.
func WithMaxTokens(maxTokens int) Option {
	return func(opts *options) {
		opts.maritacaOptions.MaxTokens = maxTokens
	}
}

// WithDoSample Set the model's generation will be sampled via top-k sampling.
// Default: true
// If True, the model's generation will be sampled via top-k sampling.
// Otherwise, the generation will always select the token with the highest probability.
// Using do_sample=False leads to a deterministic result, but with less diversity.
func WithDoSample(doSample bool) Option {
	return func(opts *options) {
		opts.maritacaOptions.DoSample = doSample
	}
}

// WithTemperature Set the sampling temperature.
// minimum: 0
// default: 0.7
// Sampling temperature (greater than or equal to zero).
// Higher values lead to greater diversity in generation but also increase the likelihood of generating nonsensical texts.
// Values closer to zero result in more plausible texts but increase the chances of generating repetitive texts.
func WithTemperature(temperature float64) Option {
	return func(opts *options) {
		opts.maritacaOptions.Temperature = temperature
	}
}

// WithTopK Set the number of top tokens to consider for sampling.
// exclusiveMaximum: 1
// exclusiveMinimum: 0
// default: 0.95
// If less than 1, it retains only the top tokens with cumulative probability >= top_p (nucleus filtering).
// For example, 0.95 means that only the tokens that make up the top 95% of the probability mass are considered when predicting the next token.
//
//	Nucleus filtering is described in Holtzman et al. (http://arxiv.org/abs/1904.09751).
func WithTopP(topP float64) Option {
	return func(opts *options) {
		opts.maritacaOptions.TopP = topP
	}
}

// WithFrequencyPenalty Set the frequency penalty.
//
//	minimum: 0
//
// default: 1
// Repetition penalty. Positive values encourage the model not to repeat previously generated tokens.
func WithRepetitionPenalty(repetitionPenalty float64) Option {
	return func(opts *options) {
		opts.maritacaOptions.RepetitionPenalty = repetitionPenalty
	}
}

// WithFrequencyPenalty Set the frequency penalty.
// List of tokens that, when generated, indicate that the model should stop generating tokens.
func WithStoppingTokens(tokens []string) Option {
	return func(opts *options) {
		opts.maritacaOptions.StoppingTokens = tokens
	}
}

// WithStream Set the model will run in streaming mode.
// default: false
// If True, the model will run in streaming mode,
// where tokens will be generated and returned to the client as they are produced.
// If False, the model will run in batch mode, where all tokens will be generated before being returned to the client.
func WithStream(stream bool) Option {
	return func(opts *options) {
		opts.maritacaOptions.Stream = stream
	}
}

// WithTokensPerMessage Set the number of tokens that will be returned per message.
//
//	minimum: 1
//
// default: 4
// Number of tokens that will be returned per message. This field is ignored if stream=False.
func WithTokensPerMessage(tokensPerMessage int) Option {
	return func(opts *options) {
		opts.maritacaOptions.NumTokensPerMessage = tokensPerMessage
	}
}

// WithToken Set the token to use.
// token use as key.
func WithToken(token string) Option {
	return func(opts *options) {
		opts.maritacaOptions.Token = token
	}
}
