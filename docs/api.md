# Open API
## `GET /calculate/expectedDamage`
operationID: `calculateExpectedDamage`

聖遺物の全ての組み合わせにおいて期待値を計算し、最適な組み合わせを検索する。

### Request

```json
{
    "characterID": "<character_id>"
}
```
// todo: 聖遺物のidを送信して、その中での最適を検索するようにする

### Response

```json
[
    {
        "expectedDamage": 123456.78,
        "artifacts": {
            "plumeOfDeath":     "107C367C-9DBC-450C-90E0-61C7AE13AB5A",
            "flowerOfLife":     "1230c3c7-3256-9595-9553-5a314866ff15",
            "sandsOfEon":       "367ff006-6b25-8c3c-16a5-30e11fca8744",
            "gobletOfEonothem": "85c2db39-90ca-3ad2-0823-b16aeef654b1",
            "circletOfLogos":   "059f58b3-96cd-ee88-d468-e2143cda8779"
            }
    },
    {
        "expectedDamage": 234567.89,
        "artifacts": {
            "plumeOfDeath":     "78ab3b42-daa4-d139-0dd9-2e01ca38ad59",
            "flowerOfLife":     "81dbba18-1fd7-8d62-4fb7-bea36a201f0d",
            "sandsOfEon":       "0ac3a027-17e8-86d6-63e8-47ec2ae9d3a7",
            "gobletOfEonothem": "c4e2f7f3-7592-a94d-0e4e-80848b4a1eab",
            "circletOfLogos":   "1099d2d0-46d9-3973-0b8a-82ed45758609"
            }
    }
]
```

## `POST /data/add/artifact`

operationID: `addArtifact`

聖遺物のデータを新規登録する。

### Request

```json
{
    "artifactID": "<artifact_id>",
    "type": "plumeOfDeath",
    "mainStat": "flatAtk",
    "mainStatValue": 100.0,
    "subStats": {
        "flatHp": 100,
        "critRate": 0.05
    }
}
```

### Response

成功 204

既存　409

## `POST /data/add/weapon`

operationID: `addWeapon`

武器のデータを新規登録する。

### Request

```json
{
    "weaponID": "<weapon_id>",
    "name": "天空剣",
    "type": "sword",
    "baseAttack": 608,
    "subStat": "energyRecharge",
    "subStatValue": 0.551,
    "additionalEffect": {
        "effectType": "critRate",
        "effectValue": 0.04
    }
}
```

### Response

成功 204

既存　409


## `POST /data/add/character`

operationID: `addCharacter`

キャラクターのデータを新規登録する。

### Request

```json
{
    "characterID": "<character_id>",
    "name": "リネ",
    "element": "pyro",
    "weaponType": "bow",
    "baseStats": {
        "atk": 318,
        "critRate": 0.192
    }
}
```

### Response

成功 204

既存　409