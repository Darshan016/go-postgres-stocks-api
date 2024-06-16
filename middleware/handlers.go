package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres-stocksAPI/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func creatConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
	return db

}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	// CreateStock := &models.Stock{}
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Error while parsing request body %v", err)
	}

	insertID := insertStock(stock)
	res := response{
		ID:      insertID,
		Message: "Stock created successfuly.",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Unable to convert string Id")
	}
	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatal("unable to get stock %v", err)
	}
	json.NewEncoder(w).Encode(stock)
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}
	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		panic(err)
	}

	updatedStock := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated successfully. updated: %v", updatedStock)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	deletedStocks := deleteStock(int64(id))
	msg := fmt.Sprintf("Stock deleted successfully. Deleted: %v", deletedStocks)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := creatConnection()
	defer db.Close()
	sqlStatement := `insert into stocks(name,price,company) values($1,$2,$3) returning stockid`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("unable to execute query: %v", err)
	}

	fmt.Printf("Inserted a record: %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {
	db := creatConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `select * from stocks where stockid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return stock, err
}

func getAllStocks() ([]models.Stock, error) {
	db := creatConnection()
	defer db.Close()

	var stocks []models.Stock
	sqlStatement := `select * from stocks`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Ubable to execute the query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var stock models.Stock

		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the rows %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, err

}

func updateStock(id int64, stock models.Stock) int64 {
	db := creatConnection()

	defer db.Close()

	sqlStatement := `update stocks set name=$2, price=$3, company=$4 where stockid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows affected: %v", rowsAffected)
	return rowsAffected
}

func deleteStock(id int64) int64 {
	db := creatConnection()
	defer db.Close()

	sqlStatement := `delete from stocks where stockid=$1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	return rowsAffected
}
