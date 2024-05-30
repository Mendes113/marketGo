package graphql

import (
    "github.com/graphql-go/graphql"
)

var marketType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Market",
        Fields: graphql.Fields{
            "id":   &graphql.Field{Type: graphql.Int},
            "name": &graphql.Field{Type: graphql.String},
        },
    },
)

var productType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Product",
        Fields: graphql.Fields{
            "id":   &graphql.Field{Type: graphql.Int},
            "name": &graphql.Field{Type: graphql.String},
        },
    },
)

var queryType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Query",
        Fields: graphql.Fields{
            "markets": &graphql.Field{
                Type: graphql.NewList(marketType),
                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                    // Simular busca no banco de dados
                    markets := []map[string]interface{}{
                        {"id": 1, "name": "Market 1"},
                        {"id": 2, "name": "Market 2"},
                    }
                    return markets, nil
                },
            },
            "products": &graphql.Field{
                Type: graphql.NewList(productType),
                Resolve: func(params graphql.ResolveParams) (interface{}, error) {
                    // Simular busca no banco de dados
                    products := []map[string]interface{}{
                        {"id": 1, "name": "Product 1"},
                        {"id": 2, "name": "Product 2"},
                    }
                    return products, nil
                },
            },
        },
    },
)

var Schema, _ = graphql.NewSchema(
    graphql.SchemaConfig{
        Query: queryType,
    },
)
