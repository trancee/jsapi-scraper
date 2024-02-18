package shop

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/sugawarayuuta/sonnet"

	"golang.org/x/net/html"

	helpers "jsapi-scraper/helpers"
)

var TechInnRegex = regexp.MustCompile(`(?i)(\d{1,2}\/)?\d{1,3}[GT]b|(\d\.)?\d{1,2}´´|[345]G|Grade [ABC]| LTE| Dual Sim| Refurbished| Enterprise Edition| EE`)

var TechInnCleanFn = func(name string) string {
	// name = regexp.MustCompile(`^Renewd | \(?(SM-)?[AGMS]\d{3}[A-Z]*(\/DSN)?\)?| XT\d{4}-\d+`).ReplaceAllString(name, "")

	if loc := TechInnRegex.FindStringSubmatchIndex(name); loc != nil {
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

var conversion = EUR_CHF

func get_descuento_producto(precio float32, rec float32) float32 {
	precio_calculo := precio / float32(conversion)
	precio_calculo = precio_calculo * 1.21

	descuento_aux := 100 - (precio_calculo * 100 / rec)
	descuento := float32(math.Round(float64(descuento_aux)))

	return descuento
}

func get_precio_anterior(precio float32, descompte float32) float32 {
	precio1 := precio * 100
	precio2 := float32(100 - descompte)

	precio_anterior := (precio1 / precio2)

	return precio_anterior
}

func XXX_techinn(isDryRun bool) IShop {
	const _name = "TechInn"
	const _count = 96
	const _url = "https://www.tradeinn.com/index.php?action=get_info_elastic_listado&id_tienda=16&idioma=eng"

	const _debug = false
	const _tests = false

	testCases := map[string]string{}

	type _Model struct {
		ID    string  `json:"id_modelo"`
		Brand string  `json:"marca"`
		Name  string  `json:"nombre_modelo"`
		Price float32 `json:"precio_win"`

		Product []struct {
			Rec string `json:"rec"`
		} `json:"productes"`
	}

	type _Body struct {
		Model []_Model `json:"id_modelos"`

		Total struct {
			Value int `json:"value"`
		} `json:"total"`
	}

	var _result []_Model
	var _body []byte

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += "/"

	{
		type _Body struct {
			Currencies struct {
				CHF string `json:"CHF"`
			} `json:"currencies"`
		}

		resp, err := http.Get("https://www.tradeinn.com/?action=get_info_pais&id_tienda=16&id_pais=192")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		_body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var body _Body
		if err := sonnet.Unmarshal(_body, &body); err != nil { // Parse []byte to go struct pointer
			panic(err)
		}

		EUR_CHF, err := strconv.ParseFloat(body.Currencies.CHF, 32)
		if err != nil {
			panic(err)
		}

		conversion = EUR_CHF
	}

	for p := 1; p <= 10; p++ {
		fn := fmt.Sprintf("shop/techinn.%d.json", p)

		if isDryRun {
			if body, err := os.ReadFile(path + fn); err != nil {
				panic(err)
			} else {
				_body = body
			}
		} else {
			// _url := fmt.Sprintf(_url, (p-1)*_count)

			form := url.Values{
				"vars[]": []string{
					"id_familia=11488",
					"atributos_e=5091,6017",
					// "model.ger;model.eng;video_mp4;id_marca;precio_tachado;sostenible;productes.talla2;productes.talla_usa;productes.talla_jp;productes.talla_uk;tres_sesenta;atributos_padre.atributos.id_atribut_valor;productes.v360;productes.v180;productes.v90;productes.v30;productes.exist;productes.stock_reservat;productes.pmp;productes.id_producte;productes.color;productes.referencia;productes.brut;productes.desc_brand;image_created;id_modelo;familias.eng;familias.ger;familias.id_familia;familias.subfamilias.eng;familias.subfamilias.ger;familias.subfamilias.id_tienda;familias.subfamilias.id_subfamilia;productes.talla;productes.baja;productes.rec;precio_win_192;productes.sellers.id_seller;productes.sellers.precios_paises.precio;productes.sellers.precios_paises.id_pais;fecha_descatalogado;marca",
					"model.eng;id_marca;precio_tachado;tres_sesenta;productes.exist;productes.stock_reservat;productes.pmp;productes.id_producte;productes.referencia;productes.brut;productes.desc_brand;id_modelo;productes.baja;productes.rec;precio_win_192;fecha_descatalogado;marca",
					"precio_win_192;asc",
					"96",
					"productos",
					"search",
					"id_subfamilia=15806",
					fmt.Sprintf("%d", (p-1)*_count),
				},
			}
			// fmt.Println(form)

			resp, err := http.PostForm(_url, form)
			if err != nil {
				// panic(err)
				fmt.Printf("[%s] %s (%s)\n", _name, err, _url)
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

		_result = append(_result, body.Model...)

		if body.Total.Value <= (p-1)*_count {
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
			_product.Name = html.UnescapeString(_product.Name)

			if !strings.HasPrefix(strings.ToUpper(_product.Name), strings.ToUpper(_product.Brand)) {
				_product.Name = _product.Brand + " " + _product.Name
			}

			_title := strings.TrimSpace(_product.Name)
			_model := TechInnCleanFn(_title)

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

			_rec, err := strconv.ParseFloat(_product.Product[0].Rec, 32)
			if err != nil {
				panic(err)
			}

			_retailPrice := _product.Price
			_price := _product.Price

			var _discount float32

			if descuento := get_descuento_producto(_product.Price, float32(_rec)); descuento >= 5 {
				_retailPrice = get_precio_anterior(_product.Price, descuento)
				_discount = descuento
			}
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
