package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/recoilme/pudge"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	helpers "jsapi-scraper/helpers"
	"jsapi-scraper/shop"
)

const ValueDiscount = 50
const ValueWorth = 100
const ValueMaximum = 300

const Token = "6219604147:AAERFP-_PfSELN3-gorzE9czM6WR-3Rum-Q"
const ChatID = "1912073977"

// https://docs.google.com/spreadsheets/d/1x28A6zoXXKeo7wmeoiAECyIzl-nlRjUSh6CJHUVifvI/edit#gid=238356703
const sheetId = 238356703
const spreadsheetId = "1x28A6zoXXKeo7wmeoiAECyIzl-nlRjUSh6CJHUVifvI"

func main() {
	isDryRun := false

	_isDryRun := flag.Bool("dryrun", isDryRun, "dry run (avoid making external calls)")
	flag.Parse()

	if _isDryRun != nil {
		isDryRun = *_isDryRun
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println()

			stack := strings.Join(strings.Split(string(debug.Stack()), "\n")[7:], "\n")
			fmt.Printf("%s\n", stack)

			if !isDryRun {
				if _, err := SendMessage(fmt.Sprintf("%v\n\n%s", err, stack)); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	_products := map[string]*[]*shop.Product{}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	fn := "products.json"

	if isDryRun {
		if products, err := os.ReadFile(path + fn); err != nil {
			panic(err)
		} else if err := json.Unmarshal(products, &_products); err != nil {
			panic(err)
		}
	} else {
		wg := sync.WaitGroup{}

		_mutex := &sync.RWMutex{}

		for _, _shop := range []shop.IShop{
			shop.XXX_alltron(isDryRun),
			shop.XXX_alternate(isDryRun),
			shop.XXX_amazon(isDryRun),
			// shop.XXX_bohnettrade(isDryRun),
			shop.XXX_brack(isDryRun),
			shop.XXX_conrad(isDryRun),
			// shop.XXX_electronova(isDryRun),
			shop.XXX_foletti(isDryRun),
			shop.XXX_fust(isDryRun),
			shop.XXX_galaxus(isDryRun),
			shop.XXX_interdiscount(isDryRun),
			shop.XXX_mediamarkt(isDryRun),
			shop.XXX_mediamarkt_refurbished(isDryRun),
			shop.XXX_melectronics(isDryRun),
			shop.XXX_microspot(isDryRun),
			shop.XXX_mistore(isDryRun),
			shop.XXX_mobiledevice(isDryRun),
			shop.XXX_mobilezone(isDryRun),
			shop.XXX_orderflow(isDryRun),
			shop.XXX_stegpc(isDryRun),
		} {
			wg.Add(1)

			go func(_shop shop.IShop) {
				defer wg.Done()

				_mutex.Lock()
				_products[_shop.Name()] = _shop.Fetch()
				_mutex.Unlock()
			}(_shop)
		}

		wg.Wait()

		if products, err := json.Marshal(_products); err != nil {
			panic(err)
		} else {
			os.WriteFile(path+fn, products, 0664)
		}
	}

	lint := func(text string) string {
		return helpers.Lint(helpers.Model(helpers.Title(strings.ToLower(strings.TrimSpace(text)))))
	}

	type Price struct {
		Name  string
		Price float32
		Link  string
	}
	matrix := map[string]map[int]Price{}

	shops := make([]string, 0, len(_products))
	for k := range _products {
		shops = append(shops, k)
	}
	// sort.Strings(shops)
	sort.Slice(shops, func(i, j int) bool { return strings.ToLower(shops[i]) < strings.ToLower(shops[j]) })

	for _, shop := range shops {
		items := _products[shop]
		// fmt.Println()
		// fmt.Println(shop)
		// fmt.Println(strings.Repeat("=", len(shop)))

		shopIndex := indexOf(shop, shops)
		if shopIndex == -1 {
			panic("unknown shop")
		}

		if items != nil {
			for _, item := range *items {
				product := lint(item.Model)
				productKey := strings.ToUpper(product)

				if len(strings.Split(product, " ")) == 1 {
					// Skip products with only brand name
					continue
				}

				if _, ok := matrix[productKey]; !ok {
					matrix[productKey] = map[int]Price{}
				}
				if price, ok := matrix[productKey][shopIndex]; ok {
					// fmt.Printf("%s %s %.2f %.2f\n", shop, product, price.Price, item.Price)
					if price.Price > item.Price {
						price.Price = item.Price
					}
				} else {
					price := Price{
						Name:  product,
						Price: item.Price,
						Link:  item.URL,
					}

					matrix[productKey][shopIndex] = price
				}
			}
		}
	}

	if !isDryRun {
		ctx := context.Background()

		// get bytes from base64 encoded google service accounts key
		credBytes, err := base64.StdEncoding.DecodeString(`ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAic2NyYXBlci0zODAzMjAiLAogICJwcml2YXRlX2tleV9pZCI6ICIxMWQ5MjFkZGExMzM1NGZlNTdiY2U4MDgyNTI0Yzg1YmRhYTM3ZmZhIiwKICAicHJpdmF0ZV9rZXkiOiAiLS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tXG5NSUlFdmdJQkFEQU5CZ2txaGtpRzl3MEJBUUVGQUFTQ0JLZ3dnZ1NrQWdFQUFvSUJBUUM4YW83OFRVRFhlOEYwXG5EbU13YnNZMTlHOTVxdHRjU0JyNGRZSkZ2TDhhSDBqYjJaS29vVlV4Vm9Va1NEbmJzbU8vMUpiaDlxbmhHMndDXG5nRVRyZVhFWDMxKzM2Uk44eGJPWUFMeEpVSkpSdDA0ZjMyZ1MxdmsySGlEYlhVRHNXNjBzVGpUcTAyU3BSY1kvXG5CQnJGV29MR2N3Z3ZJU0lrVGZKbkNuVEhJdnlRRG50dWhjaXZLNkFjNTF0dUw0ZndyOEg4ZFZ0TFA0Q1JxdHdsXG52bnJzVXRSNnZsNDN2UDg5Umk1V25ZK1A1aGp4OW45UEVBZlltTnlUYndoOURCb3Y4c3hFUUx6OWFKS3RTSGxTXG53VUdWVU5WamY5d2djUmlSdEtDSWY1ZDBraDE1RStZaHE2OXpNWjJuKytnR3hoM0YrREh3c09YbW43ODh6MVUxXG5qMWprZHNUdEFnTUJBQUVDZ2dFQU1vRi8zWnJad0VsZ3RIcnMxTDFFN1k2ZDJTZlhFRmdWdnJkRkdldDc4SVVsXG5VeVZ4M2prTTdLSk1JMHNuRTBDdzQybVpubTJ2NFBNb1UwMU43QzhNQlVHdjEwMG5sNkVwUUp3bDNLTTM3YWFzXG56dmRrWHZSNExpMEtVck1mSlp4M2diSmZGZmxmZU02RzB6cUc4Sk1RRGlFa3R2bHpQUGNWL00vOU9Lb2t1SHBxXG50a0hKL0ZMV0RaczNEL3BNNi9HREdpV0xsdld6MDY3VlQ2LzNLWGJsOHhWYUJ0Mm8wSFRIQTcrbW9qNjlIbXVuXG5aWG1uQTM3SWx0eitteWRFY2o2aWxWUHJTZDZiY0VJOENDU3FoV0VucmpwR3JUbkYwZFJOR0RhVThFUXMrM3J6XG5jL25ZWkFBVkVSZ2g4NVFUazlTbWowWWRCSHNidTU4bVlJQjFkMnhYUXdLQmdRRDRGajZMWi9UTmNSRG5RRTExXG5GalRvdGg2YWQ3Y3JUb1ZXRHVjNEp1MDNiU2owOERJc09jWllPa0Z2Z1hiWjExVGdweWpTMk40ZDRiQkxDN3JhXG50N0ZMRndicmVDdW12anI5dTEwVlZkRzJZYzF1elBvQmtpcER2d29HcDNYUFRyTUhoWFFwNWlMbWtYMm5CTWMxXG52WkVNeUg5UjBsT3NZWXE2TTVtbHZleGw2d0tCZ1FEQ2JSS3g5dmxaNlREVVJEOFRjS2dGdTVRa3BaeGJ6TGJ5XG40dDFQZ0VXa21oQnFPeFNIemRQNGFBMjRaNWp3cE9QVEhVL2RyQUhFOVZLOHhTbEwzUVhJUXQyanBsVlZSNFRhXG4rQ0prQzRIUTdmc3dNZWZoRi9LS21FKzJNalJscVVTcWJYcWRZdWhGdGxsZTRNanZjNVlyK0g3c25wMi9wdm9OXG5wVzlYMDhlU2h3S0JnR01CTFpDZ3VmZEt5ZjRma1VuS3hPNmh6M0RCbWQyMGhrMmp3TzZOeWxrMlBRUVMzMUw2XG44NGErS09NQS9aZE44ZGQ5bmpNV3pQMkwxYmo5UTJLSnNEMVJRVGV6UzJoTnZta0gzc3ZtNWJ3dEo3aXlJSXVEXG44MDM1N1Z4ZWRBdDVVc1VMb3lJZGI0d29QOGJwaHo2UkdsUEpwOVhWWkFNRklrSFEyZDVrL3ZSbEFvR0JBSWQ4XG5CSXdaWTdlQTVYTDF2OUtuTFo4WkVPbmNzakhTWFNheWFyQXMzZHNQTlNNaDJuT3NQZXNiYjN3eVRRUmNreG9aXG5rZjhTRHdXV1FycWkxZDAwdndQSGZMVytnalowS1NPQnlFMVpLM1JSY2pvcWZNQ0J0SlZhQUNvaG9CdTdzY3JsXG5rWTA5VUVqTUFrazRjUzFUcWJFb2NDSXBnaG44bk1HSHFDaFd2dnJmQW9HQkFLZ1JFRzVVVldrMXcwVWNsN1VFXG5YdnRsbUEwSFU4Mk12ZGNQV1VYeVpubHMvQXpsZ1BmUVVmNDU5RUl0SHYvcXpxVzJvSkNhaFhNZnYya08xMUZnXG54MGhrMlNZTXRYcitCb25oaHhXaEI4VDRGN3phVTlMNWVKMzJzKzJHbHp2dzgrYWIxQ2JDN1JuRmhxSmx0V1ByXG5oWExURUZvY1VjTVhoUTZqQzZQcHFZVGlcbi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS1cbiIsCiAgImNsaWVudF9lbWFpbCI6ICJzY3JhcGVyQHNjcmFwZXItMzgwMzIwLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwKICAiY2xpZW50X2lkIjogIjEwMzg2Njg5Mjk1MTM0MTY3MjQ0MCIsCiAgImF1dGhfdXJpIjogImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoIiwKICAidG9rZW5fdXJpIjogImh0dHBzOi8vb2F1dGgyLmdvb2dsZWFwaXMuY29tL3Rva2VuIiwKICAiYXV0aF9wcm92aWRlcl94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL29hdXRoMi92MS9jZXJ0cyIsCiAgImNsaWVudF94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL3JvYm90L3YxL21ldGFkYXRhL3g1MDkvc2NyYXBlciU0MHNjcmFwZXItMzgwMzIwLmlhbS5nc2VydmljZWFjY291bnQuY29tIgp9Cg==`) //os.ReadFile("scraper-380320-11d921dda133.json")
		if err != nil {
			panic(err)
		}

		// authenticate and get configuration
		config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
		if err != nil {
			panic(err)
		}

		// create client with config and context
		client := config.Client(ctx)

		// create new service using client
		service, err := sheets.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			panic(err)
		}

		rows := []*sheets.RowData{}

		{
			_date := strings.ReplaceAll(time.Now().Format(time.RFC3339), "T", "\n")
			_date = _date[:len(_date)-6]

			cells := []*sheets.CellData{
				{
					UserEnteredValue: &sheets.ExtendedValue{StringValue: &_date},
				},
			}

			for _, shop := range shops {
				_shop := strings.ReplaceAll(shop, " ", "\n")

				cells = append(
					cells,

					&sheets.CellData{
						UserEnteredValue: &sheets.ExtendedValue{StringValue: &_shop},
					},
				)
			}

			rows = append(rows, &sheets.RowData{Values: cells})
		}

		items := make([]string, 0, len(matrix))
		for k := range matrix {
			items = append(items, k)
		}
		sort.Strings(items)

		for _, item := range items {
			cells := make([]*sheets.CellData, 1+len(shops))

			for i := 0; i < 1+len(shops); i++ {
				cells[i] = &sheets.CellData{
					// UserEnteredValue: &sheets.ExtendedValue{},
				}
			}

			var cheapestPrice float32
			// Pre-process to find cheapest price
			for _, price := range matrix[item] {
				if cheapestPrice == 0 || cheapestPrice > price.Price {
					cheapestPrice = price.Price
				}
			}

			for shop, price := range matrix[item] {
				// fmt.Printf("%-50s %-25s %8.2f\n", item, shops[shop], price.Price)

				cells[0] = &sheets.CellData{
					UserEnteredValue: &sheets.ExtendedValue{StringValue: &price.Name},
				}

				_price := float64(price.Price) //fmt.Sprintf("%.2f", price.Price)

				cells[1+shop] = &sheets.CellData{
					UserEnteredValue:  &sheets.ExtendedValue{NumberValue: &_price},
					UserEnteredFormat: &sheets.CellFormat{TextFormat: &sheets.TextFormat{Link: &sheets.Link{Uri: price.Link}}},

					// TextFormatRuns: []*sheets.TextFormatRun{
					// 	{
					// 		Format: &sheets.TextFormat{
					// 			// Bold: cheapestPrice == price.Price,
					// 			Link: &sheets.Link{
					// 				Uri: price.Link,
					// 			},
					// 		},
					// 	},
					// },
				}
			}

			rows = append(rows, &sheets.RowData{Values: cells})
		}

		valueWorth := "100"
		valueMaximum := "200"

		colorWhite := &sheets.Color{Red: 1.0, Green: 1.0, Blue: 1.0}
		colorGray := &sheets.Color{Red: color(0x3f), Green: color(0x4d), Blue: color(0x59)}
		colorLight := &sheets.Color{Red: color(0xbd), Green: color(0xbd), Blue: color(0xbd)}
		colorGreen := &sheets.Color{Red: color(0x34), Green: color(0xa8), Blue: color(0x53)}
		colorYellow := &sheets.Color{Red: color(0xfb), Green: color(0xbc), Blue: color(0x04)}
		colorRed := &sheets.Color{Red: color(0xea), Green: color(0x43), Blue: color(0x35)}

		requests := []*sheets.Request{}

		if res, err := service.Spreadsheets.Get(spreadsheetId).Fields("sheets(properties(sheetId,title),conditionalFormats)").Do(); err != nil || res.HTTPStatusCode != 200 {
			panic(err)
		} else {
			for _, sheet := range res.Sheets {
				property := sheet.Properties

				if property.SheetId == int64(sheetId) {
					for i := 0; i < len(sheet.ConditionalFormats); i++ {
						requests = append(requests, &sheets.Request{
							DeleteConditionalFormatRule: &sheets.DeleteConditionalFormatRuleRequest{SheetId: int64(sheetId), Index: 0},
						})
					}
					break
				}
			}
		}

		requests = append(requests,
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 0,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    1,
								StartColumnIndex: 1,
								EndRowIndex:      int64(1 + len(items)),
								EndColumnIndex:   int64(1 + len(shops)),
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "CUSTOM_FORMULA",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: `=IF(AND(COUNT($B2:2)>1,B2<` + valueWorth + `),(B2=MIN($B2:2))*(B2<>""))`,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: true, ForegroundColor: colorWhite},
								BackgroundColor: colorGreen,
							},
						},
					},
				},
			},
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 1,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    1,
								StartColumnIndex: 1,
								EndRowIndex:      int64(1 + len(items)),
								EndColumnIndex:   int64(1 + len(shops)),
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "NUMBER_LESS_THAN_EQ",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: valueWorth,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: true, ForegroundColor: colorWhite},
								BackgroundColor: colorYellow,
							},
						},
					},
				},
			},
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 2,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    1,
								StartColumnIndex: 1,
								EndRowIndex:      int64(1 + len(items)),
								EndColumnIndex:   int64(1 + len(shops)),
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "NUMBER_BETWEEN",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: valueWorth,
									},
									{
										UserEnteredValue: valueMaximum,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: false, ForegroundColor: colorLight},
								BackgroundColor: colorGray,
							},
						},
					},
				},
			},
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 3,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    1,
								StartColumnIndex: 1,
								EndRowIndex:      int64(1 + len(items)),
								EndColumnIndex:   int64(1 + len(shops)),
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "NUMBER_GREATER_THAN_EQ",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: valueMaximum,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: true, ForegroundColor: colorWhite},
								BackgroundColor: colorRed,
							},
						},
					},
				},
			},
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 4,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    1,
								StartColumnIndex: 0,
								EndRowIndex:      int64(1 + len(items)),
								EndColumnIndex:   1,
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "CUSTOM_FORMULA",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: `=MIN($B2:2)<` + valueWorth + ``,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: true, ForegroundColor: colorWhite},
								BackgroundColor: colorGreen,
							},
						},
					},
				},
			},
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 5,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    1,
								StartColumnIndex: 0,
								EndRowIndex:      int64(1 + len(items)),
								EndColumnIndex:   1,
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "CUSTOM_FORMULA",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: `=MIN($B2:2)>` + valueMaximum + ``,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: true, ForegroundColor: colorWhite},
								BackgroundColor: colorRed,
							},
						},
					},
				},
			},
		)

		batchUpdateRequest := sheets.BatchUpdateSpreadsheetRequest{
			Requests: append(requests,
				&sheets.Request{
					UpdateCells: &sheets.UpdateCellsRequest{
						Fields: "userEnteredValue,userEnteredFormat.textFormat.link",
						Range: &sheets.GridRange{
							SheetId:          int64(sheetId),
							StartRowIndex:    0,
							StartColumnIndex: 0,
							EndRowIndex:      int64(1 + len(items)),
							EndColumnIndex:   int64(1 + len(shops)),
						},
						Rows: rows,
					},
				},
			),
		}

		// execute the request using spreadsheetId
		if res, err := service.Spreadsheets.BatchUpdate(spreadsheetId, &batchUpdateRequest).Context(ctx).Do(); err != nil || res.HTTPStatusCode != 200 {
			panic(err)
		}

		if res, err := service.Spreadsheets.Values.Clear(spreadsheetId, fmt.Sprintf("Scraper!%d:%d", 1+len(items)+1, 1+len(items)+100), &sheets.ClearValuesRequest{}).Context(ctx).Do(); err != nil || res.HTTPStatusCode != 200 {
			panic(err)
		}
	}

	db, err := pudge.Open("products", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ids := map[string]bool{}

	if keys, err := db.Keys(nil, 0, 0, true); err != nil {
		panic(err)
	} else {
		for _, key := range keys {
			ids[string(key)] = true
		}
	}
	// fmt.Println(ids)

	for name, products := range _products {
		fmt.Println()

		_num := 0
		if products != nil {
			_num = len(*products)
		}
		_name := fmt.Sprintf("%s (%d)", name, _num)

		fmt.Printf("%s\n%s\n", _name, strings.Repeat("=", len(_name)))

		if products != nil {
			for _, product := range *products {
				id := product.Code
				delete(ids, id)

				_name := product.Name
				if product.Quantity > 0 {
					_name += fmt.Sprintf(" (%d)", product.Quantity)
				}
				priceLine := ""
				if product.Price != product.RetailPrice {
					priceLine = fmt.Sprintf("%8.2f %8.2f %3d%%", product.Price, product.Savings, int(product.Discount))
				}

				fmt.Printf("%-69s %8.2f %22s %s\n", _name, product.RetailPrice, priceLine, product.URL)

				notify := false

				var oldProduct shop.Product
				if ok, _ := db.Has(id); ok {
					db.Get(id, &oldProduct)
				}

				if oldProduct != *product {
					db.Set(id, product)

					if oldProduct.RetailPrice != product.RetailPrice && (product.RetailPrice < ValueWorth || product.Discount > ValueDiscount) {
						notify = true
					}
				}

				if notify {
					if !isDryRun {
						// fmt.Println()
						// fmt.Println(oldProduct)
						// fmt.Println(product)

						priceDiff := ""
						priceLine := ""
						if product.Price != product.RetailPrice {
							priceLine = fmt.Sprintf("%-8.2f %-8.2f %3d%%", product.Price, product.Savings, int(product.Discount))
						} else {
							if oldProduct.RetailPrice > 0 {
								priceDiff = fmt.Sprintf("\n%-8.2f %+8.2f", oldProduct.RetailPrice, product.RetailPrice-oldProduct.RetailPrice)
							}
						}

						productLine := fmt.Sprintf("%s\n%-8.2f %s%s\n\n%s", _name, product.RetailPrice, priceLine, priceDiff, product.URL)

						if _, err := SendMessage(productLine); err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}

	// fmt.Println(ids)
	for id := range ids {
		db.Delete(id)
	}
}

func color(v float64) float64 {
	return v / float64(0xff)
}

func getURL() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessage(text string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getURL())
	body, _ := json.Marshal(map[string]string{
		"chat_id": ChatID,
		"text":    text,
	})
	if response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	); err != nil {
		return false, err
	} else {
		// Close the request at the end
		defer response.Body.Close()

		// Body
		if _, err := io.ReadAll(response.Body); err != nil {
			return false, err
		}
	}

	// Return
	return true, nil
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
