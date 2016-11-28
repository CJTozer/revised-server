package backend

import (
    "revised-server/models"
)

func DummyResourcesList() []*models.Resource {
return []*models.Resource{
        &models.Resource{
            ID: 1,
            ResourceType: "Text",
            Text: "Resource A",
        },
        &models.Resource{
            ID: 2,
            ResourceType: "Link",
            Text: "https://en.wikipedia.org/wiki/Hyperlink",
        },
        &models.Resource{
            ID: 3,
            ResourceType: "Image",
            Text: "https://upload.wikimedia.org/wikipedia/commons/a/a8/TEIDE.JPG",
        },
        &models.Resource{
            ID: 4,
            ResourceType: "Text",
            Text: "Resource D",
        },
        &models.Resource{
            ID: 5,
            ResourceType: "Text",
            Text: "Resource E",
        },
        &models.Resource{
            ID: 6,
            ResourceType: "Text",
            Text: "Resource F",
        },
        &models.Resource{
            ID: 7,
            ResourceType: "Text",
            Text: "Resource G",
        },
        &models.Resource{
            ID: 8,
            ResourceType: "Text",
            Text: "Resource H",
        },
        &models.Resource{
            ID: 9,
            ResourceType: "Text",
            Text: "Resource I",
        },
        &models.Resource{
            ID: 10,
            ResourceType: "Text",
            Text: "Resource J",
        },
    }
}
