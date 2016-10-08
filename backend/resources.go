package backend

import (
    "revised-server/models"
)

func DummyResourcesList() []*models.Resource {
return []*models.Resource{
        &models.Resource{
            ResourceID: 1,
            ResourceType: "Link",
            Text: "Link to a photo",
        },
        &models.Resource{
            ResourceID: 2,
            ResourceType: "Link",
            Text: "Link to another photo",
        },
        &models.Resource{
            ResourceID: 3,
            ResourceType: "Link",
            Text: "Link to a map",
        },
    }
}
