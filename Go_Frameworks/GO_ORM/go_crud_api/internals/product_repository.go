package internals

import ("go_crud_api/internals/models"
        "gorm.io/gorm")


type ProductRepository interface {
	FindAll(page, limit int) ([]models.Product, int64, error)
	FindByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}


func (r *productRepository) FindAll(page , limit int) ([]models.Product , int64 , error) {

	var products []models.Product

	var total int64

	offset := (page -1 ) * limit

	if err := r.db.Model(&models.Product{}).Count(&total).Error;  err != nil {
		return  nil , 0 , err
	}

	result := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&products)
	return products, total, result.Error
}


	

func (r *productRepository) FindById(id uint) (*models.Product , error) {
	var product models.Product

	result := r.db.Find(&product , id)

	if result.Error != nil {
		return  nil , result.Error
	   
	 }

	return &product , nil
}

func (r *productRepository) Create(product *models.Product) error {
	return  r.db.Create(product).Error
}


func (r *productRepository) Update(product *models.Product) error {
	return  r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}