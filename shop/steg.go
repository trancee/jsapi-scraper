package shop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func XXX_steg() IShop {
	const _name = "Steg Electronics"
	const _url = "https://www.steg-electronics.ch/de/product/list/11853?sortKey=preisasc"

	type _Response struct {
		NewProductList string `json:"newProductList"`
	}

	var _result _Response

	resp, err := http.Post("https://www.steg-electronics.ch/de/product/list/11853?sortKey=preisasc", "application/json", bytes.NewBuffer(nil))
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))

	if err := json.Unmarshal(body, &_result); err != nil { // Parse []byte to go struct pointer
		panic(err)
	}
	// fmt.Println(_result.NewProductList)

	doc, err := html.Parse(strings.NewReader(_result.NewProductList))
	if err != nil {
		panic(err)
	}

	attr := func(n *html.Node, key string) (string, bool) {
		for _, attr := range n.Attr {
			if attr.Key == key {
				return attr.Val, true
			}
		}
		return "", false
	}

	var traverse func(n *html.Node) *html.Node
	traverse = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && len(n.Attr) == 2 {
			if s, ok := attr(n, "data-product-id"); ok {
				fmt.Println(s)

				product := Product{
					Code: s,
				}

				{
					c := n

					c = c.FirstChild.NextSibling
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "popularity" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
						// Skip
						c = c.NextSibling.NextSibling
					}
					if c.DataAtom.String() == "a" && c.Attr[0].Key == "class" && c.Attr[0].Val == "listItemImage" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)

						if s, ok := attr(c, "href"); ok {
							fmt.Println(s)

							product.URL = s
						}

						c = c.FirstChild.NextSibling
						if c.DataAtom.String() == "img" {
							// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)

							if s, ok := attr(c, "title"); ok {
								fmt.Println(s)

								product.Name = s
							}
						} else {
							panic("img")
						}

						c = c.Parent
					} else {
						panic("listItemImage")
					}

					c = c.NextSibling.NextSibling
					if c.DataAtom.String() == "h2" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
					} else {
						panic("h2")
					}

					c = c.NextSibling.NextSibling
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "rating" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
						// Skip
						c = c.NextSibling.NextSibling
					}
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "priceAndActionButtons" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
					} else {
						panic("priceAndActionButtons")
					}

					c = c.FirstChild.NextSibling
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "sm-flex lg-flex lg-dir-row lg-j-spaceBetween" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
					} else {
						panic("sm-flex")
					}

					c = c.FirstChild.NextSibling
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
					} else {
						panic("div class empty")
					}

					c = c.NextSibling.NextSibling
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "price" {
						// fmt.Printf("\t[%v]{%v} %v\n", c.Type, c.DataAtom, c.Attr)
					} else {
						panic("price")
					}

					c = c.FirstChild.NextSibling
					if c.DataAtom.String() == "div" && c.Attr[0].Key == "class" && c.Attr[0].Val == "generalPrice" {
						// fmt.Printf("\t[%v]{%v/%v} %v\n", c.Type, c.DataAtom, strings.TrimSpace(c.Data), c.Attr)

						if price, err := strconv.ParseFloat(strings.TrimSpace(c.Data), 32); err == nil {
							product.RetailPrice = float32(price)
						}
					} else {
						panic("generalPrice")
					}
				}

				// for c := n.FirstChild; c != nil; c = c.NextSibling {
				// 	fmt.Printf("{%v}\n", c.Attr)
				// }

				fmt.Println(product)

				// return n
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if result := traverse(c); result != nil {
				return result
			}
		}

		// for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 	if c.Type == html.TextNode && c.Parent.Data == tag {
		// 		fmt.Println(c.Parent.Attr)
		// 		fmt.Println(c.Data)
		// 		// *data = append(*data, c.Data)
		// 	}

		// 	res := traverse(c, tag)
		// 	if res != nil {
		// 		return res
		// 	}
		// }

		return nil
	}

	traverse(doc)

	return shop{}
}
