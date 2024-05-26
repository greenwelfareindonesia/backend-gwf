package product

type ServiceProduct interface {
	CreateProduct(input ProductInput) (Products, error)
	GetProducts() ([]Products, error)
	GetProduct(ID int) (Products, error)
	DeleteProduct(ID int) (Products, error)
	GetProductByCategory(ID int) ([]Products, error)
}

type serviceProduct struct {
	repositoryProduct RepositoryProduct
}

func NewServiceProduct(repositoryProduct RepositoryProduct) *serviceProduct {
	return &serviceProduct{repositoryProduct}
}

func (s *serviceProduct) GetProductByCategory(ID int) ([]Products, error) {
	product, err := s.repositoryProduct.FindAllProductByCategory(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) GetProducts() ([]Products, error) {

	product, err := s.repositoryProduct.FindAll()
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) CreateProduct(input ProductInput) (Products, error) {
	product := Products{}

	product.Name = input.Title
	product.Price = input.Price
	product.Stock = input.Stock

	newProduct, err := s.repositoryProduct.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *serviceProduct) GetProduct(ID int) (Products, error) {

	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *serviceProduct) DeleteProduct(ID int) (Products, error) {

	product, err := s.repositoryProduct.FindById(ID)
	if err != nil {
		return product, err
	}
	productDel, err := s.repositoryProduct.Delete(product)

	if err != nil {
		return productDel, err
	}
	return productDel, nil

}
