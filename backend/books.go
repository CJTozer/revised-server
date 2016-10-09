package backend

import (
    "revised-server/models"
)

func DummyBooksList() []*models.Book {
    return []*models.Book{
        &models.Book{
            Author: "Dava Sobel",
            BookID: 1,
            Title:  "Longitude",
            Resources: []int64 {1, 3, 4, 5},
        },
        &models.Book{
            Author: "Robert K. Massie",
            BookID: 2,
            Title:  "Peter The Great",
            Resources: []int64 {2, 4, 6, 7, 8},
        },
        &models.Book{
            Author: "Don Oberdorfer",
            BookID: 3,
            Title:  "The Two Koreas",
            Resources: []int64 {1, 2, 3, 5, 7, 9, 10},
        },
    }
}
