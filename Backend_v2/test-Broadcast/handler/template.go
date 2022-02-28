package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
	"log"
	"mf2-broadcast/config"
	"mf2-broadcast/model"
	"strings"
)

func CreateWtsTemplate(c *fiber.Ctx) error {
	template := new(model.Template)

	err := c.BodyParser(&template)
	if err != nil {
		log.Println(err)
	}

	component := make([]interface{}, 0)
	// header
	if len(template.Header.Type) != 0 {
		hType := strings.ToUpper(template.Header.Type)
		item := map[string]interface{}{
			"type":   graphql.String("HEADER"),
			"format": graphql.String(hType),
		}
		if hType == "TEXT" {
			item["text"] = graphql.String(template.Header.Body)
		} else {
			if len(template.Header.Example) != 0 {
				var hex struct {
					Example []graphql.String `json:"header_handler"`
				}
				hex.Example = append(hex.Example, graphql.String(template.Header.Example))
				item["example"] = hex
			}
		}
		component = append(component, item)
	}

	// body
	body := map[string]interface{}{
		"type": graphql.String("BODY"),
		"text": graphql.String(template.Body),
	}
	if len(template.BodyEX) != 0 {
		var bex struct {
			Example []graphql.String `json:"body_text"`
		}
		bex.Example = append(bex.Example, graphql.String(template.BodyEX))
		body["example"] = bex
	}
	component = append(component, body)

	// footer
	if template.Footer != "" {
		footer := map[string]interface{}{
			"type": graphql.String("FOOTER"),
			"test": graphql.String(template.Footer),
		}
		component = append(component, footer)
	}

	// buttons
	if len(template.BType) != 0 {
		buttons := map[string]interface{}{
			"type": graphql.String("BUTTONS"),
		}
		bs := make([]map[string]interface{}, 0)
		for _, v := range template.BBody {
			b := map[string]interface{}{
				"text": graphql.String(v.Text),
			}
			if template.BType == "url" {
				b["type"] = graphql.String("URL")
				b["url"] = graphql.String(v.Url)
			} else {
				b["type"] = graphql.String("QUICK_REPLY")
			}
			bs = append(bs, b)
		}
		buttons["buttons"] = bs
		component = append(component, buttons)
	}

	input := map[string]interface{}{
		"name":       graphql.String("test"),
		"category":   graphql.String(template.Category),
		"channelId":  graphql.String(template.ChannelId),
		"language":   graphql.String(template.Language),
		"components": component,
	}

	variable := map[string]interface{}{
		"input": input,
	}

	fmt.Println(variable)
	var mutation struct {
		CreateWhatsappMessageTemplate struct {
		} `graphql:"createWhatsappMessageTemplate(input: $input)"`
	}

	urlStr := config.GoDotEnvVariable("STELLAAPI")

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GoDotEnvVariable("ACCESSTOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient(urlStr, httpClient)
	err = client.Mutate(context.Background(), &mutation, variable)
	if err != nil {
		log.Println("error in mutation", err)
	}

	return c.JSON(variable)
}
