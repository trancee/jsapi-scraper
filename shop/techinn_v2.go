package shop

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sugawarayuuta/sonnet"

	"golang.org/x/net/html"

	helpers "jsapi-scraper/helpers"
)

var TechInnV2Regex = regexp.MustCompile(`(?i)(\d{1,2}/)?\d{1,3}[GT]b|(\d\.)?\d{1,2}´´|[345]G|Grade [ABC]| LTE| Dual Sim| Refurbished`)

var TechInnV2CleanFn = func(name string) string {
	name = strings.NewReplacer(" ", " ", "Enterprise Edition", "EE").Replace(name)
	// name = regexp.MustCompile(`^Renewd | \(?(SM-)?[AGMS]\d{3}[A-Z]*(/DSN)?\)?| XT\d{4}-\d+`).ReplaceAllString(name, "")

	if loc := TechInnV2Regex.FindStringSubmatchIndex(name); loc != nil {
		// fmt.Printf("%v\t%-30s %s\n", loc, name[:loc[0]], name)
		name = name[:loc[0]]
	}

	name = strings.NewReplacer(" (Product)Red", "", "Smartphone ", "", "Sonstige ", "").Replace(name)

	s := strings.Split(name, " ")

	if s[0] == "Apple" {
		name = strings.NewReplacer("SE2", "SE (2020)").Replace(name)
	}

	if s[0] == "Gigaset" {
		name = regexp.MustCompile(`(?i)IP\d{2}`).ReplaceAllString(name, "")
	}

	if s[0] == "Huawei" || s[0] == "HUAWEI" {
		if s[1] == "Honor" || s[1] == "HONOR" {
			name = regexp.MustCompile(`(?i)^HUAWEI\s*`).ReplaceAllString(name, "")
		}
	}

	if s[0] == "Realme" {
		name = regexp.MustCompile(`(?i)RMX\s+\d{4}`).ReplaceAllString(name, "")
	}

	if s[0] == "Samsung" {
		name = strings.ReplaceAll(name, " DS", "")

		if s[1] == "Z" {
			name = strings.ReplaceAll(name, "Samsung ", "Samsung Galaxy ")
		}
	}

	if s[0] == "Xiaomi" {
		name = strings.ReplaceAll(name, " Plus", "+")

		name = strings.ReplaceAll(name, "Xiaomi Redmi 12 Note 11S", "Xiaomi Redmi Note 11S")
	}

	if s[0] == "Zte" {
		name = strings.ReplaceAll(name, "I9", "L9")
	}

	return helpers.Lint(name)
}

func XXX_techinn_v2(isDryRun bool) IShop {
	const _name = "TechInn"
	const _count = 96
	const _url = "https://sr.tradeinn.com/"

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Model struct {
		ID string `json:"id_modelo"`

		Brand string `json:"marca"`

		Model struct {
			Name string `json:"eng"`
		} `json:"model"`

		Price float32 `json:"precio_win_192"`
	}

	type _Body struct {
		Hits struct {
			Total struct {
				Value int `json:""`
			} `json:"total"`

			Hits []struct {
				Source _Model `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	var _result []_Model
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/techinn.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			var jsonData = StringToBytes(
				fmt.Sprintf(
					`
					{
						"from": "%d",
						"size": "%d",
						"query": {
							"bool": {
								"filter": [
									{
										"range": {
											"image_created": {
												"gt": 0
											}
										}
									},
									{
										"nested": {
											"path": "productes.sellers",
											"query": {
												"bool": {
													"must": [
														{
															"range": {
																"productes.sellers.stock": {
																	"gt": 0
																}
															}
														},
														{
															"nested": {
																"path": "productes.sellers.precios_paises",
																"query": {
																	"match": {
																		"productes.sellers.precios_paises.id_pais": "192"
																	}
																}
															}
														}
													]
												}
											}
										}
									},
									{
										"nested": {
											"path": "productes",
											"query": {
												"term": {
													"productes.baja": {
														"value": 0
													}
												}
											}
										}
									}
								],
								"must": [
									{
										"nested": {
											"path": "familias",
											"query": {
												"term": {
													"familias.id_familia": {
														"value": "11488"
													}
												}
											}
										}
									}
								],
								"must_not": [
									{
										"terms": {
											"id_marca": [
												"28",
												"183",
												"189",
												"482",
												"550",
												"1699",
												"2090",
												"2426",
												"4878",
												"8513",
												"16190"
											]
										}
									},
									{
										"term": {
											"paises_prohibidos": "192"
										}
									}
								]
							}
						},
						"_source": {
							"includes": [
								"model.eng",
								"id_modelo",
								"precio_win_192",
								"marca"
							]
						},
						"sort": [
							{
								"precio_win_192": {
									"order": "asc"
								}
							}
						],
						"post_filter": {
							"bool": {
								"filter": [
									{
										"nested": {
											"path": "familias.subfamilias",
											"query": {
												"terms": {
													"familias.subfamilias.id_subfamilia": [
														"15806"
													]
												}
											}
										}
									}
								]
							}
						},
						"aggregations": {
							"group_by_marca": {
								"filter": {
									"bool": {
										"must": [
											{
												"nested": {
													"path": "familias.subfamilias",
													"query": {
														"terms": {
															"familias.subfamilias.id_subfamilia": [
																"15806"
															]
														}
													}
												}
											}
										]
									}
								},
								"aggs": {
									"marcas": {
										"terms": {
											"field": "marca.keyword",
											"size": 1000,
											"order": {
												"_key": "asc"
											}
										},
										"aggs": {
											"id_marca": {
												"terms": {
													"field": "id_marca"
												}
											}
										}
									}
								}
							},
							"group_by_categorias": {
								"nested": {
									"path": "familias"
								},
								"aggs": {
									"filter_id_familia": {
										"filter": {
											"term": {
												"familias.id_familia": "11488"
											}
										},
										"aggs": {
											"subfamilias": {
												"nested": {
													"path": "familias.subfamilias"
												},
												"aggs": {
													"id_subfamilia": {
														"terms": {
															"field": "familias.subfamilias.id_subfamilia",
															"size": 1000
														}
													}
												}
											}
										}
									}
								}
							},
							"group_by_tallas": {
								"filter": {
									"bool": {
										"must": [
											{
												"nested": {
													"path": "familias.subfamilias",
													"query": {
														"terms": {
															"familias.subfamilias.id_subfamilia": [
																"15806"
															]
														}
													}
												}
											}
										]
									}
								},
								"aggs": {
									"tallas": {
										"nested": {
											"path": "productes"
										},
										"aggs": {
											"talla_filter": {
												"filter": {
													"term": {
														"productes.baja": "0"
													}
												},
												"aggs": {
													"talla": {
														"terms": {
															"field": "productes.talla_filtro.keyword",
															"order": {
																"_key": "asc"
															},
															"size": 1000
														}
													}
												}
											}
										}
									}
								}
							},
							"group_by_tallas_2": {
								"filter": {
									"bool": {
										"must": [
											{
												"nested": {
													"path": "familias.subfamilias",
													"query": {
														"terms": {
															"familias.subfamilias.id_subfamilia": [
																"15806"
															]
														}
													}
												}
											}
										]
									}
								},
								"aggs": {
									"tallas": {
										"nested": {
											"path": "productes"
										},
										"aggs": {
											"talla_filter": {
												"filter": {
													"term": {
														"productes.baja": "0"
													}
												},
												"aggs": {
													"talla": {
														"terms": {
															"field": "productes.talla_filtro2.keyword",
															"order": {
																"_key": "asc"
															},
															"size": 1000
														}
													}
												}
											}
										}
									}
								}
							},
							"group_by_atributos": {
								"filter": {
									"bool": {
										"must": [
											{
												"nested": {
													"path": "familias.subfamilias",
													"query": {
														"terms": {
															"familias.subfamilias.id_subfamilia": [
																"15806"
															]
														}
													}
												}
											}
										]
									}
								},
								"aggs": {
									"atributos": {
										"nested": {
											"path": "atributos_padre"
										},
										"aggs": {
											"id_atributo": {
												"terms": {
													"field": "atributos_padre.id_atribut_pare",
													"size": 1000
												},
												"aggs": {
													"valor_atributos": {
														"nested": {
															"path": "atributos_padre.atributos"
														},
														"aggs": {
															"ids_atributos_valor": {
																"terms": {
																	"field": "atributos_padre.atributos.id_atribut_valor",
																	"size": 1000
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
					`,

					(p-1)*_count,
					_count,
				),
			)

			req, err := http.NewRequest("POST", _url, bytes.NewBuffer(jsonData))
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
			req.Header.Set("Origin", "https://www.tradeinn.com")
			req.Header.Set("Referer", "https://www.tradeinn.com/techinn/en/phones-smartphones/15806/s")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, req.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				// panic(resp.StatusCode)
				fmt.Printf("[%s] %d: %s (%s)\n", _name, resp.StatusCode, resp.Status, resp.Request.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			}

			if body, err := io.ReadAll(resp.Body); err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, resp.Request.URL)
				return NewShop(
					_name,
					_url,

					nil,
				)
			} else {
				_body = body
			}

			os.WriteFile(path+fn, _body, 0664)
		}
		// fmt.Println(BytesToString(_body))

		var body _Body
		if err := sonnet.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
			panic(err)
		}
		// fmt.Println(body.Model)

		for _, v := range body.Hits.Hits {
			_result = append(_result, v.Source)
		}

		if body.Hits.Total.Value <= (p-1)*_count {
			break
		}
	}

	r := regexp.MustCompile("[^a-z0-9 .-/]")

	_parseFn := func(s IShop) *[]*Product {
		products := []*Product{}

		fmt.Printf("-- %s (%d)\n", _name, len(_result))
		for _, _product := range _result {
			// fmt.Println(_product)

			_product.Brand = html.UnescapeString(_product.Brand)
			_product.Model.Name = html.UnescapeString(_product.Model.Name)

			if !strings.HasPrefix(strings.ToUpper(_product.Model.Name), strings.ToUpper(_product.Brand)) {
				_product.Model.Name = _product.Brand + " " + _product.Model.Name
			}

			_title := strings.TrimSpace(_product.Model.Name)
			_model := TechInnV2CleanFn(_title)

			if Skip(_model) {
				continue
			}
			if _debug {
				// fmt.Println(_title)
				fmt.Println(_model)
			}

			if _tests {
				testCases[_title] = _model
			}

			_retailPrice := _product.Price
			_price := _product.Price

			var _discount float32

			if _debug {
				fmt.Println(_retailPrice)
				fmt.Println(_price)
			}

			_savings := _price - _retailPrice
			if _debug {
				fmt.Println(_savings)
				fmt.Println(_discount)
			}

			_productName := strings.NewReplacer(" ", "-", "/", "-").Replace(r.ReplaceAllString(strings.ToLower(_title), "$1"))
			_link := fmt.Sprintf("https://www.tradeinn.com/techinn/en/%s/%s/p", _productName, _product.ID)
			if _debug {
				fmt.Println(_link)
				fmt.Println()
			}

			product := &Product{
				Code:  _name + "//" + _product.ID,
				Name:  _title,
				Model: _model,

				RetailPrice: _retailPrice,
				Price:       _price,
				Savings:     _savings,
				Discount:    _discount,

				URL: _link,
			}

			if s.IsWorth(product) {
				products = append(products, product)
			}
		}

		if _tests {
			keys := make([]string, 0, len(testCases))

			for k := range testCases {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })

			for _, k := range keys {
				fmt.Println("\"" + strings.ReplaceAll(k, "\"", "\\\"") + "\",")
			}
			fmt.Println()
			for _, k := range keys {
				fmt.Println("\"" + strings.ReplaceAll(testCases[k], "\"", "\\\"") + "\",")
			}
		}

		// fmt.Printf("%#v\n", products)
		return &products
	}

	return NewShop(
		_name,
		_url,

		_parseFn,
	)
}
