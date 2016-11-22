package backend

import (
    "revised-server/models"
)

func DummyResourcesList() []*models.Resource {
return []*models.Resource{
        &models.Resource{
            ID: 1,
            ResourceType: "Link",
            Text: "Resource A",
        },
        &models.Resource{
            ID: 2,
            ResourceType: "Link",
            Text: "Resource B",
        },
        &models.Resource{
            ID: 3,
            ResourceType: "Link",
            Text: "Resource C",
        },
        &models.Resource{
            ID: 4,
            ResourceType: "Link",
            Text: "Resource D",
        },
        &models.Resource{
            ID: 5,
            ResourceType: "Link",
            Text: "Resource E",
        },
        &models.Resource{
            ID: 6,
            ResourceType: "Link",
            Text: "Resource F",
        },
        &models.Resource{
            ID: 7,
            ResourceType: "Link",
            Text: "Resource G",
        },
        &models.Resource{
            ID: 8,
            ResourceType: "Link",
            Text: "Resource H",
        },
        &models.Resource{
            ID: 9,
            ResourceType: "Link",
            Text: "Resource I",
        },
        &models.Resource{
            ID: 10,
            ResourceType: "Link",
            Text: "Resource J",
        },
    }
}
