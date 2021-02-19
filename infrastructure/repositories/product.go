package repositories

import "github.com/rennanbadaro/discount-service/infrastructure/storage"

type IProductRepository interface {
	FetchByID(id string) (*Product, error)
}

type ProductRepository struct {
	dbClient *storage.PostgresClient
}

type Product struct {
	Id           string
	PriceInCents int
	Title        string
	Description  string
}

func (pr *ProductRepository) FetchByID(id string) (*Product, error) {
	query := "SELECT id, title, description, price_in_cents from products WHERE id=$1"
	row := pr.dbClient.Conn.QueryRow(query, id)

	p := &Product{}
	if err := row.Scan(&p.Id, &p.Title, &p.Description, &p.PriceInCents); err != nil {
		return nil, err
	}

	return p, nil
}

func NewProductRepository(dbClient *storage.PostgresClient) *ProductRepository {
	return &ProductRepository{dbClient: dbClient}
}
