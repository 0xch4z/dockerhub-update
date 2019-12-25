package update

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/charliekenney23/dockerhub-go"
	"github.com/urfave/cli/v2"
)

// Entrypoint is the entrypoint of the CLI.
var Entrypoint = &cli.App{
	Name:  "dockerhub-update",
	Usage: "Update your Dockerhub repository",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "Dockerhub username",
			EnvVars: []string{"DOCKERHUB_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "Dockerhub password",
			EnvVars: []string{"DOCKERHUB_PASSWORD"},
		},
		&cli.StringFlag{
			Name:    "token",
			Aliases: []string{"t"},
			Usage:   "Dockerhub API token",
			EnvVars: []string{"DOCKERHUB_API_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "readme",
			Aliases: []string{"r"},
			Usage:   "Path to readme update",
			EnvVars: []string{"README_PATH"},
		},
		&cli.StringFlag{
			Name:    "description",
			Aliases: []string{"d"},
			Usage:   "Description update",
			EnvVars: []string{"README_PATH"},
		},
	},
	Action: entrypoint,
}

func getClient(ctx context.Context, c *cli.Context) (*dockerhub.Client, error) {
	tok := c.String("token")
	user := c.String("username")
	pass := c.String("password")

	client := dockerhub.NewClient(nil)
	if tok != "" {
		client.SetAuthToken(tok)
	} else if user != "" && pass != "" {
		client.Auth.Login(ctx, user, pass)
	} else {
		return nil, errors.New("Username and password or token is required for authentication")
	}

	return client, nil
}

func mustParseURI(uri string) (string, string, error) {
	comps := strings.Split(uri, "/")
	if len(comps) != 2 {
		return "", "", errors.New("Dockerhub repo URI must be of format:\n[USER|NAMESPACE]/[IMAGE]")
	}

	return comps[0], comps[1], nil
}

func entrypoint(c *cli.Context) error {
	ctx := context.Background()

	desc := c.String("description")
	readme := c.String("readme")
	uri := c.Args().First()

	if uri == "" {
		return errors.New("repo URI is required")
	}

	namespace, repo, err := mustParseURI(uri)
	if err != nil {
		return err
	}

	if readme != "" {
		d, err := ioutil.ReadFile(readme)
		if err != nil {
			return err
		}
		readme = string(d)
	}

	if readme == "" && desc == "" {
		fmt.Println("No readme or description update provided")
	}

	client, err := getClient(ctx, c)
	if err != nil {
		return err
	}

	_, err = client.Repositories.EditRepository(ctx, namespace, repo, &dockerhub.RepositoryPatch{
		FullDescription: readme,
		Description:     desc,
	})
	return err
}
