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
    }
}
