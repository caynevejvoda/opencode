package models

const (
        ProviderMistral ModelProvider = "mistral"

        MistralSmallLatest         ModelID = "mistral-small-latest"
        MistralMediumLatest        ModelID = "mistral-medium-latest"
        MistralLargeLatest         ModelID = "mistral-large-latest"
        MistralCodestralLatest     ModelID = "codestral-latest"
        MistralDevstralSmallLatest ModelID = "devstral-small-latest"
)

var MistralModels = map[ModelID]Model{
        MistralSmallLatest: {
                ID:                 MistralSmallLatest,
                Name:               "Mistral Small Latest",
                Provider:           ProviderMistral,
                APIModel:           "mistral-small-latest",
                CostPer1MIn:        0.1,
                CostPer1MInCached:  0,
                CostPer1MOut:       0.3,
                CostPer1MOutCached: 0,
                ContextWindow:      40_000,
                DefaultMaxTokens:   20_000,
        },
        MistralMediumLatest: {
                ID:                 MistralMediumLatest,
                Name:               "Mistral Medium Latest",
                Provider:           ProviderMistral,
                APIModel:           "mistral-medium-latest",
                CostPer1MIn:        0.4,
                CostPer1MInCached:  0,
                CostPer1MOut:       2,
                CostPer1MOutCached: 0,
                ContextWindow:      128_000,
                DefaultMaxTokens:   20_000,
        },
        MistralLargeLatest: {
                ID:                 MistralLargeLatest,
                Name:               "Mistral Large Latest",
                Provider:           ProviderMistral,
                APIModel:           "mistral-large-latest",
                CostPer1MIn:        2,
                CostPer1MInCached:  0,
                CostPer1MOut:       6,
                CostPer1MOutCached: 0,
                ContextWindow:      128_000,
                DefaultMaxTokens:   20_000,
        },
        MistralCodestralLatest: {
                ID:                 MistralCodestralLatest,
                Name:               "Codestral Latest",
                Provider:           ProviderMistral,
                APIModel:           "codestral-latest",
                CostPer1MIn:        0.3,
                CostPer1MInCached:  0,
                CostPer1MOut:       0.9,
                CostPer1MOutCached: 0,
                ContextWindow:      256_000,
                DefaultMaxTokens:   20_000,
        },
        MistralDevstralSmallLatest: {
                ID:                 MistralDevstralSmallLatest,
                Name:               "Devstral Small Latest",
                Provider:           ProviderMistral,
                APIModel:           "devstral-small-latest",
                CostPer1MIn:        0.1,
                CostPer1MInCached:  0,
                CostPer1MOut:       0.3,
                CostPer1MOutCached: 0,
                ContextWindow:      128_000,
                DefaultMaxTokens:   20_000,
        },
}
