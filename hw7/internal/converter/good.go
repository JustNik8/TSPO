package converter

import (
	"ecom/internal/domain"
	"ecom/internal/transport/rest/dto"
)

type GoodConverter struct {
}

func (c *GoodConverter) MapDomainToDto(d domain.Good) dto.Good {
	return dto.Good{
		ID:            d.ID,
		Name:          d.Name,
		Price:         d.Price,
		Desc:          d.Desc,
		StockQuantity: d.StockQuantity,
		MeasureUnit: &dto.MeasureUnit{
			ID:   d.MeasureUnit.ID,
			Name: d.MeasureUnit.Name,
		},
	}
}

func (c *GoodConverter) MapDomainsToDtos(d []domain.Good) []dto.Good {
	dtos := make([]dto.Good, 0, len(d))

	for _, item := range d {
		dtos = append(dtos, c.MapDomainToDto(item))
	}

	return dtos
}

func (c *GoodConverter) MapRequestToDomain(dto dto.Good) domain.Good {
	return domain.Good{
		ID:            dto.ID,
		Name:          dto.Name,
		Price:         dto.Price,
		StockQuantity: dto.StockQuantity,
		MeasureUnit: domain.MeasureUnit{
			ID:   dto.MeasureUnit.ID,
			Name: dto.MeasureUnit.Name,
		},
	}
}
