package cli

import "github.com/hellflame/argparse"

type ProjectorOptions struct {
	Args   []string
	Config string
	Pwd    string
}

func GetOpts() (*ProjectorOptions, error) {
	parser := argparse.NewParser("projector", "gets all the values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	config := parser.String("c", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	pwd := parser.String("p", "pwd", &argparse.Option{
		Required: false,
		Default:  "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &ProjectorOptions{
		Args:   *args,
		Config: *config,
		Pwd:    *pwd,
	}, nil
}
