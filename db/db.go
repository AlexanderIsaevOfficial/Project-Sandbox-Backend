package db

//
import (
	"context"
	"gameback_v1/types"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	db *pgx.Conn
}

func (r *Store) NewStore(connStr string) error {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	r.db = conn

	log.Println("INFO DB connetion")
	return nil
}

func (r *Store) GetMainCategoryList(size int) (data []types.CategoryDBRequest, err error) {

	cxt, _ := context.WithTimeout(context.Background(), time.Second*3)

	rows, err := r.db.Query(cxt,
		"select catid , cattitle , itemid , itemtitle , itemlogo from mvcategory where group_nn < $1 order by catrang , group_nn ASC ", size)
	if err != nil {
		log.Println("DB Error", err)
		return nil, err
	}

	for rows.Next() {
		var p types.CategoryDBRequest
		if err := rows.Scan(&p.Id, &p.Title, &p.I.Id, &p.I.Title, &p.I.Logo); err != nil {
			return nil, err
		}
		data = append(data, p)
	}
	return data, nil
}

func (r *Store) GetCategoryItemsById(catId, packege, limit int) (items []types.Item, err error) {
	cxt, _ := context.WithTimeout(context.Background(), time.Second*3)

	rows, err := r.db.Query(cxt,
		"select itemid , itemtitle , itemlogo from mvcategory where catid = $1 and group_nn > $2 LIMIT $3", catId, (packege-1)*limit, limit)
	if err != nil {
		log.Println("DB Error", err)
		return nil, err
	}

	for rows.Next() {
		var p types.Item
		if err := rows.Scan(&p.Id, &p.Title, &p.Logo); err != nil {
			return nil, err
		}
		items = append(items, p)
	}
	return items, nil
}

func (r *Store) GetItemById(id int) (i types.Item, err error) {

	tn := time.Now()

	cxt, _ := context.WithTimeout(context.Background(), time.Second*3)

	err = r.db.QueryRow(cxt, `select "id" , "title" , "description" , "locationId" , "logoImg" , "screenshotsImg" from "Items" where "activeFlg" = 1 and "status" = 3 and "id" = $1`, id).
		Scan(&i.Id, &i.Title, &i.Description, &i.LocationId, &i.Logo, &i.ScreenShots)
	if err != nil {
		log.Println("DB Error", err)
		return i, err
	}

	log.Println("DB GetItemById, time =", time.Since(tn).Microseconds())
	return i, nil
}

func (r *Store) GetLocatonById(id int) (i types.Location, err error) {
	tn := time.Now()

	cxt, _ := context.WithTimeout(context.Background(), time.Second*3)

	err = r.db.QueryRow(cxt, `select "id", "locationFilePath" from "Location" where "activeFlg" = 1 and "id" = $1`, id).
		Scan(&i.Id, &i.FilePath)
	if err != nil {
		log.Println("DB Error", err)
		return i, err
	}

	log.Println("DB GetLocatonById, time =", time.Since(tn).Microseconds())
	return i, nil
}
