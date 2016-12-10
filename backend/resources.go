package backend

import (
    "revised-server/models"
)

func DummyResourcesList() []*models.Resource {
return []*models.Resource{
        &models.Resource{
            ID: 1,
            ResourceType: "Text",
            Description: "The letter 'E'",
            Text: "E is the fifth letter and the second vowel in the modern English alphabet and the ISO basic Latin alphabet. It is the most commonly used letter in many languages, including: Czech, Danish, Dutch, English, Latvian French, German, Hungarian, Latin, Norwegian, Spanish, and Swedish.",
        },
        &models.Resource{
            ID: 2,
            ResourceType: "Link",
            Description: "Wikipedia page on the letter T",
            Text: "https://en.wikipedia.org/wiki/T",
        },
        &models.Resource{
            ID: 3,
            ResourceType: "Image",
            Description: "A picture of a book",
            Text: "https://2.bp.blogspot.com/-S8IeK4SBo5A/Vf_OtsUuPOI/AAAAAAAAF3s/E8iJVa-OlTA/s1600/file000514157119.jpg",
        },
        &models.Resource{
            ID: 4,
            ResourceType: "Text",
            Description: "Resource D",
            Text: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 5,
            ResourceType: "Text",
            Description: "Resource E",
            Text: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 6,
            ResourceType: "Text",
            Description: "Resource F",
            Text: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 7,
            ResourceType: "Text",
            Description: "Resource G",
            Text: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 8,
            ResourceType: "Text",
            Description: "Resource H",
            Text: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 9,
            ResourceType: "Text",
            Description: "Resource I",
            Text: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 10,
            ResourceType: "Text",
            Description: "Resource J",
            Text: "Placeholder text for resource D",
        },
    }
}
