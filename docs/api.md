# Open API
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