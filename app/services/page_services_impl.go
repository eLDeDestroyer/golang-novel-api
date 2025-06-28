package services

import (
	"e-novel/app/repositories"
	"e-novel/model"
	"e-novel/model/dto"
)

type PageServiceImpl struct {
	pageRepository repositories.PageRepository
}

func NewPageServiceImpl(pageRepository repositories.PageRepository) *PageServiceImpl {
	return &PageServiceImpl{
		pageRepository: pageRepository,
	}
}

func (service *PageServiceImpl) GetPageBook(bookId int, page int) (*dto.ResponsePage, error) {
	var pages []int

	for i := 0; i < 3; i++ {
		if i == 0 {
			pagePrev := page - 1
			pages = append(pages, pagePrev)
		}
		if i == 1 {
			pageNow := page
			pages = append(pages, pageNow)
		}
		if i == 2 {
			nextPrev := page + 1
			pages = append(pages, nextPrev)
		}
	}

	request := &dto.RequestPage{
		BookId: bookId,
		Page:   pages,
	}

	rawData, err := service.pageRepository.GetPageBook(request)
	if err != nil {
		return nil, err
	}

	var response dto.ResponsePage

	for _, row := range rawData {
		pageNow := int(row["page"].(int32))
		if pageNow == page {
			response.Page = pageNow
			response.Text = row["text"].(string)
		}
		if pageNow == page-1 {
			response.PrevPage = &pageNow
		}
		if pageNow == page+1 {
			response.NextPage = &pageNow
		}
	}

	return &response, nil

}


func (service *PageServiceImpl) AddPageBook(page *model.Page) error {
	pageCount, err :=  service.pageRepository.GetPageBookCount(page.BookId)
	if err != nil {
		return err
	}

	page.Page = pageCount + 1

	err = service.pageRepository.AddPageBook(page)
	if err != nil {
		return err
	}

	return nil
} 


func (service *PageServiceImpl) UpdatePageBook(page *model.Page) error {
	err := service.pageRepository.UpdatePageBook(page)
	if err != nil {
		return err
	}

	return nil
} 
