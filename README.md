# synthesized-bustabit-dataset

`gameResult` function logic sourced from [bustabit 2025 seeding event](https://bitcointalk.org/index.php?topic=5560454)

## dataset description

datasets 1k.csv, 10k.csv, 100k.csv, and 1m.csv contain a list of multipliers in a simulated bustabit crash game with random hash and salt

## generation

it is possible to generate bustabit history beyond a specific hash in the chain given that hash as the gameHash and the bitcoin block hash as the salt, uses legacy logic so is not compatible with 2024 seeding event logic
