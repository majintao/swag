package main

import (
	"encoding/json"
	"fmt"
	"github.com/majintao/swag"
	"github.com/majintao/swag/gen"
	"github.com/urfave/cli/v2"
	"log"
	"majintao/go-yapi"
	"os"
)

const (
	searchDirFlag        = "dir"
	excludeFlag          = "exclude"
	generalInfoFlag      = "generalInfo"
	propertyStrategyFlag = "propertyStrategy"
	outputFlag           = "output"
	parseVendorFlag      = "parseVendor"
	parseDependencyFlag  = "parseDependency"
	markdownFilesFlag    = "markdownFiles"
	parseInternal        = "parseInternal"
	generatedTimeFlag    = "generatedTime"
)

var initFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    generalInfoFlag,
		Aliases: []string{"g"},
		Value:   "main.go",
		Usage:   "Go file path in which 'swagger general API Info' is written",
	},
	&cli.StringFlag{
		Name:    searchDirFlag,
		Aliases: []string{"d"},
		Value:   "./",
		Usage:   "Directory you want to parse",
	},
	&cli.StringFlag{
		Name:  excludeFlag,
		Usage: "exclude directories and files when searching, comma separated",
	},
	&cli.StringFlag{
		Name:    propertyStrategyFlag,
		Aliases: []string{"p"},
		Value:   "camelcase",
		Usage:   "Property Naming Strategy like snakecase,camelcase,pascalcase",
	},
	&cli.StringFlag{
		Name:    outputFlag,
		Aliases: []string{"o"},
		Value:   "./docs",
		Usage:   "Output directory for all the generated files(swagger.json, swagger.yaml and doc.go)",
	},
	&cli.BoolFlag{
		Name:  parseVendorFlag,
		Usage: "Parse go files in 'vendor' folder, disabled by default",
	},
	&cli.BoolFlag{
		Name:  parseDependencyFlag,
		Usage: "Parse go files in outside dependency folder, disabled by default",
	},
	&cli.StringFlag{
		Name:    markdownFilesFlag,
		Aliases: []string{"md"},
		Value:   "",
		Usage:   "Parse folder containing markdown files to use as description, disabled by default",
	},
	&cli.BoolFlag{
		Name:  "parseInternal",
		Usage: "Parse go files in internal packages, disabled by default",
	},
	&cli.BoolFlag{
		Name:  "generatedTime",
		Usage: "Generate timestamp at the top of docs.go, true by default",
	},
}

const (
	testInstanceURL = "http://yapi.iacorn.cn"
	testToken       = "8cde6e3bcbdb1c7bc0d8a2992fe0d79b0f07d186bfe7ebc5c6723b403e6830d8"
)

func initAction(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case swag.CamelCase, swag.SnakeCase, swag.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	config := &gen.Config{
		SearchDir:          c.String(searchDirFlag),
		Excludes:           c.String(excludeFlag),
		MainAPIFile:        c.String(generalInfoFlag),
		PropNamingStrategy: strategy,
		OutputDir:          c.String(outputFlag),
		ParseVendor:        c.Bool(parseVendorFlag),
		ParseDependency:    c.Bool(parseDependencyFlag),
		MarkdownFilesDir:   c.String(markdownFilesFlag),
		ParseInternal:      c.Bool(parseInternal),
		GeneratedTime:      c.Bool(generatedTimeFlag),
	}

	parser, err := gen.New().GetParser(config)
	if err != nil {
		return fmt.Errorf("generate error.")
	}

	for path, _ := range parser.GetSwagger().Paths.Paths {
		fmt.Println(path)
	}

	swagger, err := gen.New().BuildSwagger(config)

	if err != nil {
		return fmt.Errorf("generate error.")
	}

	yapiClient, err1 := yapi.NewClient(nil, testInstanceURL, testToken)
	if err1 != nil {
		return fmt.Errorf("Client creation -> Got an error: %s", err)
	}
	if yapiClient == nil {
		return fmt.Errorf("Expected a client. Got none")
	}

	swaggerJson := string(swagger)
	result, _, _ := yapiClient.Interface.UploadSwagger(&swaggerJson)
	marshal, _ := json.Marshal(result)
	api := string(marshal)
	fmt.Println(api)

	return nil
}

func listMenuAction(c *cli.Context) error {
	yapiClient, err1 := yapi.NewClient(nil, testInstanceURL, testToken)
	if err1 != nil {
		return fmt.Errorf("Client creation -> Got an error: %s", err1)
	}
	if yapiClient == nil {
		return fmt.Errorf("Expected a client. Got none")
	}

	projectResp, _, _ := yapiClient.Project.Get()
	result, _, _ := yapiClient.CatMenu.Get(projectResp.Data.ID)
	marshal, _ := json.Marshal(result)
	api := string(marshal)
	fmt.Println(api)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = swag.Version
	app.Usage = "yapi"
	app.Commands = []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create docs.go",
			Action:  initAction,
			Flags:   initFlags,
		}, {
			Name:    "listMenu",
			Aliases: []string{"lm"},
			Usage:   "列举当前的目录",
			Action:  listMenuAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
