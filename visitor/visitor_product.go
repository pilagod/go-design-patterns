package visitor

type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type ProductVisitor interface {
	Visit(ProductInfoRetriever)
}

type ProductVisitable interface {
	Accept(ProductVisitor)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(pv ProductVisitor) {
	pv.Visit(p)
}

type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(pir ProductInfoRetriever) {
	pv.Sum += pir.GetPrice()
}

type ProductListVisitor struct {
	ProductList string
}

func (nv *ProductListVisitor) Visit(pir ProductInfoRetriever) {
	if nv.ProductList != "" {
		nv.ProductList += ", "
	}
	nv.ProductList += pir.GetName()
}
