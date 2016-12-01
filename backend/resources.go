package backend

import (
    "revised-server/models"
)

func DummyResourcesList() []*models.Resource {
return []*models.Resource{
        &models.Resource{
            ID: 1,
            ResourceType: "Text",
            Text: "The letter 'E'",
            Description: "E is the fifth letter and the second vowel in the modern English alphabet and the ISO basic Latin alphabet. It is the most commonly used letter in many languages, including: Czech, Danish, Dutch, English, Latvian French, German, Hungarian, Latin, Norwegian, Spanish, and Swedish.",
        },
        &models.Resource{
            ID: 2,
            ResourceType: "Link",
            Text: "https://en.wikipedia.org/wiki/T",
            Description: "Wikipedia page on the letter T",
        },
        &models.Resource{
            ID: 3,
            ResourceType: "Image",
            Text: "http://2.bp.blogspot.com/-S8IeK4SBo5A/Vf_OtsUuPOI/AAAAAAAAF3s/E8iJVa-OlTA/s1600/file000514157119.jpg",
            Description: "A picture of a book",
        },
        &models.Resource{
            ID: 4,
            ResourceType: "Text",
            Text: "Resource D",
            Description: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 5,
            ResourceType: "Text",
            Text: "Resource E",
            Description: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 6,
            ResourceType: "Text",
            Text: "Resource F",
            Description: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 7,
            ResourceType: "Text",
            Text: "Resource G",
            Description: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 8,
            ResourceType: "Text",
            Text: "Resource H",
            Description: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 9,
            ResourceType: "Text",
            Text: "Resource I",
            Description: "Placeholder text for resource D",
        },
        &models.Resource{
            ID: 10,
            ResourceType: "Text",
            Text: "Resource J",
            Description: "Placeholder text for resource D",
        },
    }
}
