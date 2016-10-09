package backend

import (
    "revised-server/models"
)

func DummyResourcesList() []*models.Resource {
return []*models.Resource{
        &models.Resource{
            ResourceID: 1,
            ResourceType: "Link",
            Text: "Resource A",
        },
        &models.Resource{
            ResourceID: 2,
            ResourceType: "Link",
            Text: "Resource B",
        },
        &models.Resource{
            ResourceID: 3,
            ResourceType: "Link",
            Text: "Resource C",
        },
        &models.Resource{
            ResourceID: 4,
            ResourceType: "Link",
            Text: "Resource D",
        },
        &models.Resource{
            ResourceID: 5,
            ResourceType: "Link",
            Text: "Resource E",
        },
        &models.Resource{
            ResourceID: 6,
            ResourceType: "Link",
            Text: "Resource F",
        },
        &models.Resource{
            ResourceID: 7,
            ResourceType: "Link",
            Text: "Resource G",
        },
        &models.Resource{
            ResourceID: 8,
            ResourceType: "Link",
            Text: "Resource H",
        },
        &models.Resource{
            ResourceID: 9,
            ResourceType: "Link",
            Text: "Resource I",
        },
        &models.Resource{
            ResourceID: 10,
            ResourceType: "Link",
            Text: "Resource J",
        },
    }
}
