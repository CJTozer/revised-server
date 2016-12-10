package backend

import (
    "revised-server/models"
)

var DummyBooksList = []*models.Book{
    &models.Book{
        Author: "Dava Sobel",
        ID: 1,
        Title:  "Longitude",
        Resources: []int64 {1, 3, 4, 5},
    },
    &models.Book{
        Author: "Robert K. Massie",
        ID: 2,
        Title:  "Peter The Great",
        Resources: []int64 {2, 4, 6, 7, 8},
    },
    &models.Book{
        Author: "Don Oberdorfer",
        ID: 3,
        Title:  "The Two Koreas",
        Resources: []int64 {1, 2, 3, 5, 7, 9, 10},
    },
}

func GetNextID() int64 {
    var nextID int64 = 1
    for _, book := range DummyBooksList {
        if book.ID >= nextID {
            nextID = book.ID + 1
        }
    }
    return nextID
}
