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
	"regexp"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/recoilme/pudge"
	"github.com/sugawarayuuta/sonnet"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	helpers "jsapi-scraper/helpers"
	"jsapi-scraper/shop"
)

// https://docs.google.com/spreadsheets/d/1x28A6zoXXKeo7wmeoiAECyIzl-nlRjUSh6CJHUVifvI/edit#gid=238356703

var exclusionRegex = regexp.MustCompile(`(?i)^(emporia|htc|siemens|sony ericsson)|apple iphone \d(gs|g|c|s)?\b|fairphone (1|2)|gigaset (gl|gs)|google pixel (2|3|4|5)a?\b|motorola moto g[1234567]?\b|samsung galaxy (zoom|young|rex|note(\s[1234567]\b|$)|j\d|gt|alpha|ace|s\d?\b|a\d?\b|advance|mini|duos)`)

const PRODUCT_EXPIRATION = 5 * 24 * 60 * 60
const PRICE_DIFFERENCE_PERCENTAGE = 5.0
const PRICE_DIFFERENCE_VALUE = 5

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

			stack := strings.Join(strings.Split(shop.BytesToString(debug.Stack()), "\n")[7:], "\n")
			fmt.Printf("%s\n", stack)

			if !isDryRun {
				if _, err := SendMessage(fmt.Sprintf("%v\n\n%s", err, stack)); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	shop.EUR_CHF = helpers.EUR_CHF_v2()
	fmt.Printf("EUR/CHF = %f\n", shop.EUR_CHF)

	sheetId, err := strconv.Atoi(os.Getenv("SHEET_ID"))
	if err != nil {
		panic(err)
	}
	spreadsheetId := os.Getenv("SPREADSHEET_ID")

	_shops := []string{}
	_items := []string{}
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
		} else if err := sonnet.Unmarshal(products, &_products); err != nil {
			panic(err)
		}
	} else {
		wg := sync.WaitGroup{}

		_mutex := &sync.RWMutex{}

		for _, _shop := range []shop.IShop{
			shop.XXX_ackermann(isDryRun),
			shop.XXX_alltron(isDryRun),
			shop.XXX_alternate(isDryRun),
			shop.XXX_amazon(isDryRun),
			// shop.XXX_bohnettrade(isDryRun),
			shop.XXX_brack(isDryRun),
			shop.XXX_cashconverters(isDryRun),
			shop.XXX_conrad(isDryRun),
			// shop.XXX_electronova(isDryRun),
			shop.XXX_foletti(isDryRun),
			shop.XXX_fust(isDryRun),
			shop.XXX_galaxus(isDryRun),
			shop.XXX_interdiscount(isDryRun),
			shop.XXX_manor(isDryRun),
			shop.XXX_mediamarkt(isDryRun),
			shop.XXX_mediamarkt_refurbished(isDryRun),
			shop.XXX_melectronics(isDryRun),
			shop.XXX_microspot(isDryRun),
			// shop.XXX_mistore(isDryRun),
			shop.XXX_mistore_v2(isDryRun),
			// shop.XXX_mobiledevice(isDryRun),
			shop.XXX_mobiledevice_v2(isDryRun),
			shop.XXX_mobilezero(isDryRun),
			shop.XXX_mobilezone(isDryRun),
			shop.XXX_orderflow(isDryRun),
			// shop.XXX_stegpc(isDryRun), // out of order
			shop.XXX_techinn(isDryRun),
			shop.XXX_tutti(isDryRun),
			shop.XXX_ultimus(isDryRun),
			shop.XXX_venova(isDryRun),
		} {
			_shops = append(_shops, _shop.Name())

			if _shop.CanFetch() {
				wg.Add(1)

				go func(_shop shop.IShop) {
					defer wg.Done()

					_mutex.Lock()
					_products[_shop.Name()] = _shop.Fetch()
					_mutex.Unlock()
				}(_shop)
			}
		}

		wg.Wait()

		if products, err := json.Marshal(_products); err != nil {
			panic(err)
		} else {
			os.WriteFile(path+fn, products, 0664)
		}
	}

	type Price struct {
		ID    string
		Name  string
		Price float32
		Link  string
	}
	matrix := map[string]map[int]Price{}

	sort.Slice(_shops, func(i, j int) bool { return strings.ToLower(_shops[i]) < strings.ToLower(_shops[j]) })

	for _, shop := range _shops {
		items := _products[shop]
		// fmt.Println()
		// fmt.Println(shop)
		// fmt.Println(strings.Repeat("=", len(shop)))

		shopIndex := indexOf(shop, _shops)
		if shopIndex == -1 {
			panic("unknown shop")
		}

		if items != nil {
			for _, item := range *items {
				product := item.Model
				productKey := strings.ToUpper(product)

				if len(strings.Split(product, " ")) == 1 {
					// Skip products with only brand name
					continue
				}

				if exclusionRegex.MatchString(product) {
					continue
				}

				func() {
					for _, item := range _items {
						if item == productKey {
							return
						}
					}

					_items = append(_items, productKey)
				}()

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
						ID:    item.Code,
						Name:  product,
						Price: item.Price,
						Link:  item.URL,
					}

					matrix[productKey][shopIndex] = price
				}
			}
		}
	}

	{
		fmt.Printf("\n%s\n%s\n", "Discounts", strings.Repeat("=", len("Discounts")))

		db, err := pudge.Open("discounts", nil)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		ids := map[string]bool{}

		if keys, err := db.Keys(nil, 0, 0, true); err != nil {
			panic(err)
		} else {
			for _, key := range keys {
				ids[shop.BytesToString(key)] = true
			}
		}
		// fmt.Println(ids)

		// Loop:
		for _, item := range _items {
			min := Price{}
			max := Price{}
			// fmt.Printf("%s", item)

			for _, price := range matrix[item] {
				// if _shops[shop] == "Amazon" {
				// 	// Do not consider discounts from Amazon.
				// 	continue Loop
				// }

				// fmt.Printf(" %v", price.Price)
				if min.Price > price.Price || min.Price == 0 {
					min = price
				}
				if max.Price < price.Price || max.Price == 0 {
					max = price
				}
			}

			// fmt.Printf(" [%v/%v]", min.Price, max.Price)
			// fmt.Println()

			delete(ids, min.ID)

			notify := false

			// if max.Price-min.Price >= shop.ValueWorth {
			// 	fmt.Printf("%-25s %7.2f %7.2f %3.f%% %s\n", min.Name, min.Price, max.Price-min.Price, 100-((100/max.Price)*min.Price), min.Link)

			// 	var oldPrice Price
			// 	if ok, _ := db.Has(min.Name); ok {
			// 		db.Get(min.Name, &oldPrice)
			// 	}

			// 	if oldPrice != min {
			// 		db.Set(min.Name, min)

			// 		notify = true
			// 	}
			// }
			if 100-((100/max.Price)*min.Price) >= shop.ValueDiscount {
				fmt.Printf("%-25s %7.2f %7.2f %3.f%% %s\n", min.Name, min.Price, max.Price-min.Price, 100-((100/max.Price)*min.Price), min.Link)

				var oldPrice float32
				if ok, _ := db.Has(min.ID); ok {
					db.Get(min.ID, &oldPrice)
				}

				if oldPrice != min.Price {
					db.Set(min.ID, min.Price)

					notify = true
				}
			}

			if notify {
				productLine := fmt.Sprintf("%s\n%-8.2f ±%.2f %5.f%%\n\n%s", min.Name, min.Price, max.Price-min.Price, 100-((100/max.Price)*min.Price), min.Link)

				if _, err := SendMessage(productLine); err != nil {
					panic(err)
				}
			}
		}

		// fmt.Println(ids)
		for id := range ids {
			db.Delete(id)
		}
	}

	if !isDryRun {
		ctx := context.Background()

		// get bytes from base64 encoded google service accounts key
		credBytes, err := base64.StdEncoding.DecodeString(os.Getenv("CREDENTIALS"))
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

			for _, shop := range _shops {
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

		// items := make([]string, 0, len(matrix))
		// for k := range matrix {
		// 	items = append(items, k)
		// }
		// sort.Strings(_items)
		Sort(_items)

		for _, item := range _items {
			cells := make([]*sheets.CellData, 1+len(_shops))

			for i := 0; i < 1+len(_shops); i++ {
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
		colorDarkGray := &sheets.Color{Red: color(0x28), Green: color(0x28), Blue: color(0x28)}
		colorLightGray := &sheets.Color{Red: color(0xe4), Green: color(0xe4), Blue: color(0xe4)}

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
								EndRowIndex:      int64(1 + len(_items)),
								EndColumnIndex:   int64(1 + len(_shops)),
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
								EndRowIndex:      int64(1 + len(_items)),
								EndColumnIndex:   int64(1 + len(_shops)),
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
								EndRowIndex:      int64(1 + len(_items)),
								EndColumnIndex:   int64(1 + len(_shops)),
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
								EndRowIndex:      int64(1 + len(_items)),
								EndColumnIndex:   int64(1 + len(_shops)),
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
								EndRowIndex:      int64(1 + len(_items)),
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
								EndRowIndex:      int64(1 + len(_items)),
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
			&sheets.Request{
				AddConditionalFormatRule: &sheets.AddConditionalFormatRuleRequest{
					Index: 6,
					Rule: &sheets.ConditionalFormatRule{
						Ranges: []*sheets.GridRange{
							{
								SheetId:          int64(sheetId),
								StartRowIndex:    0,
								StartColumnIndex: 1,
								EndRowIndex:      1,
								EndColumnIndex:   int64(1 + len(_shops)),
							},
						},
						BooleanRule: &sheets.BooleanRule{
							Condition: &sheets.BooleanCondition{
								Type: "CUSTOM_FORMULA",
								Values: []*sheets.ConditionValue{
									{
										UserEnteredValue: `=COUNT(B2:B1000)=0`,
									},
								},
							},
							Format: &sheets.CellFormat{
								TextFormat:      &sheets.TextFormat{Bold: true, ForegroundColor: colorLightGray},
								BackgroundColor: colorDarkGray,
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
							EndRowIndex:      int64(1 + len(_items)),
							EndColumnIndex:   int64(1 + len(_shops)),
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

		if res, err := service.Spreadsheets.Values.Clear(spreadsheetId, fmt.Sprintf("Scraper!%d:%d", 1+len(_items)+1, 1+len(_items)+100), &sheets.ClearValuesRequest{}).Context(ctx).Do(); err != nil || res.HTTPStatusCode != 200 {
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
			ids[shop.BytesToString(key)] = true
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

				_name := product.Model
				if product.Quantity > 0 {
					_name += fmt.Sprintf(" (%d)", product.Quantity)
				}
				priceLine := ""
				if priceDiff(product.Price, product.RetailPrice) {
					priceLine = fmt.Sprintf("%8.2f %8.2f %3d%%", product.Price, product.Savings, int(product.Discount))
				}

				fmt.Printf("%-30s %8.2f %22s %s\n", _name, product.RetailPrice, priceLine, product.URL)

				notify := false

				var oldProduct shop.Product
				if ok, _ := db.Has(id); ok {
					db.Get(id, &oldProduct)

					product.Counter = oldProduct.Counter
					product.CreationDate = oldProduct.CreationDate
				} else {
					product.CreationDate = time.Now().Unix()
				}

				if ((product.EURPrice > 0 && priceDiff(oldProduct.EURPrice, product.EURPrice)) ||
					(oldProduct.RetailPrice > 0 && priceDiff(oldProduct.RetailPrice, product.RetailPrice)) ||
					(oldProduct.Price > 0 && priceDiff(oldProduct.Price, product.Price))) &&
					(product.RetailPrice <= shop.ValueWorth || product.Price <= shop.ValueWorth || product.Discount >= shop.ValueDiscount) {
					notify = true

					product.NotificationDate = time.Now().Unix()
				}

				product.Counter++
				product.ModificationDate = time.Now().Unix()

				db.Set(id, product)

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
		var oldProduct shop.Product
		if ok, _ := db.Has(id); ok {
			db.Get(id, &oldProduct)

			// Do not delete data if it is not older than 5 days.
			if time.Now().Unix()-oldProduct.ModificationDate < PRODUCT_EXPIRATION {
				continue
			}
		}

		db.Delete(id)
	}
}

func priceDiff(a float32, b float32) bool {
	_max := max(a, b)
	_min := min(a, b)
	_diff := _max - _min
	return (100/_max*_diff) > PRICE_DIFFERENCE_PERCENTAGE && _diff > PRICE_DIFFERENCE_VALUE
}

func color(v float64) float64 {
	return v / float64(0xff)
}

func getURL() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", os.Getenv("BOT_TOKEN"))
}

func SendMessage(text string) (bool, error) {
	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getURL())
	body, _ := json.Marshal(map[string]string{
		"chat_id": os.Getenv("CHAT_ID"),
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

// https://github.com/facette/natsort
type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(a, b int) bool {
	return Compare(s[a], s[b])
}

func (s stringSlice) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

// Sort sorts a list of strings in a natural order
func Sort(l []string) {
	sort.Sort(stringSlice(l))
}

// Compare returns true if the first string precedes the second one according to natural order
func Compare(a, b string) bool {
	ln_a := len(a)
	ln_b := len(b)
	posa := 0
	posb := 0

	for {
		if ln_a <= posa {
			if ln_b <= posb {
				// eof on both at the same time (equal)
				return false
			}
			return true
		} else if ln_b <= posb {
			// eof on b
			return false
		}

		av, bv := a[posa], b[posb]

		if av >= '0' && av <= '9' && bv >= '0' && bv <= '9' {
			// go into numeric mode
			intlna := 1
			intlnb := 1
			for {
				if posa+intlna >= ln_a {
					break
				}
				x := a[posa+intlna]
				if av == '0' {
					posa += 1
					av = x
					continue
				}
				if x >= '0' && x <= '9' {
					intlna += 1
				} else {
					break
				}
			}
			for {
				if posb+intlnb >= ln_b {
					break
				}
				x := b[posb+intlnb]
				if bv == '0' {
					posb += 1
					bv = x
					continue
				}
				if x >= '0' && x <= '9' {
					intlnb += 1
				} else {
					break
				}
			}
			if intlnb > intlna {
				// length of a value is longer, means it's a bigger number
				return true
			} else if intlna > intlnb {
				return false
			}
			// both have same length, let's compare as string
			v := strings.Compare(a[posa:posa+intlna], b[posb:posb+intlnb])
			if v < 0 {
				return true
			} else if v > 0 {
				return false
			}
			// equale
			posa += intlna
			posb += intlnb
			continue
		}

		if av == bv {
			posa += 1
			posb += 1
			continue
		}

		return av < bv
	}
}
