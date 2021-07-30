package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/leumas3003/beercatalogue-go/internal/pkg/db"

	"github.com/labstack/echo/v4"
)

type BeerData struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Graduation  float32 `json:"graduation"`
}

type Beer struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Graduation  float32 `json:"graduation"`
}

// AddBeer godoc
// @Summary AddBeer
// @Description AddBeer to catalogue
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Param beer body Beer true "beer"
// @Router /addbeer [post]
func AddBeer(c echo.Context) error {
	Pool, err := db.Db()
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("something went wrong with the DB. %s", err))
	}

	b := &Beer{}
	if err := c.Bind(b); err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("something went wrong unmarshalling the Beer data. %s", err))
	}
	qry := `insert into public.beer (beer_name, description, graduation) values ($1, $2, $3)`
	_, e := Pool.Exec(context.Background(), qry, strings.ToUpper(b.Name), b.Description, b.Graduation)
	if e != nil {
		log.Printf("beer %s already in the DB. %s", b.Name, e)
		return c.String(http.StatusBadRequest, fmt.Sprintf("beer %s already in the DB. %s", b.Name, e))
	}

	log.Printf("this is your beer %#v\n", b)
	return c.String(http.StatusOK, fmt.Sprintf("the beer %s was added successfully", b.Name))
}

// GetBeer godoc
// @Summary GetBeer
// @Description GetBeer from catalogue
// @Accept json
// @Produce json
// @Success 200 {object} Beer
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Param name path string true "name"
// @Router /beer/{name} [get]
func GetBeer(c echo.Context) error {
	Pool, err := db.Db()
	if err != nil {
		log.Println(err)
		return err
	}
	beername := c.Param("name")
	qry := `select * from public.beer where beer_name = $1`

	row, e := Pool.Query(context.Background(), qry, strings.ToUpper(beername))
	if e != nil {
		log.Printf("beer %s is not in the DB. %s", beername, e)
		return c.String(http.StatusBadRequest, fmt.Sprintf("beer %s is not in the DB. %s", beername, e))
	}
	defer row.Close()

	var beer BeerData
	for row.Next() {
		if err := row.Scan(&beer.Id, &beer.Name, &beer.Description, &beer.Graduation); err != nil {
			log.Println(err)
			return c.String(http.StatusInternalServerError, fmt.Sprintf("something went wrong. %s", e))
		}
		return c.JSON(http.StatusOK, map[string]string{
			"name":        beer.Name,
			"description": beer.Description,
			"graduation":  fmt.Sprintf("%.2f", beer.Graduation),
		})

	}
	return c.String(http.StatusBadRequest, fmt.Sprintf("beer %s not found", beername))

}
