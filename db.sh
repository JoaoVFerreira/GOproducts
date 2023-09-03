#!/bin/bash

jsonFile="/db/db.json"
targetFolder="/db"

# Verify if the file already exists
if [ ! -e "$jsonFile" ]; then
    # if the file does not exists, create
    mkdir -p "$targetFolder"
    echo '[
        {
            "id": 1,
            "title": "Ragdoll Cat",
            "price": 5500.00,
            "description": "A very lovely animal",
            "category": "animals",
            "rating": {
                "rate": 4.9,
                "count": 3
            }
        }
    ]' > "$jsonFile"
    echo "Arquivo JSON criado em: $jsonFile"
else
    echo "O arquivo JSON já existe em: $jsonFile"
fi
