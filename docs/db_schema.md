# DB スキーマ

## `artifact`

| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `id` | `VARCHAR(36)` | 聖遺物のID。 |
| `type` | `ENUM('flowerOfLife', 'plumeOfDeath', 'sandsOfEon', 'gobletOfEonothem', 'circletOfLogos')` | 聖遺物の種類。 |
| `mainStat` | `ENUM('flatHp', 'hpPercent', 'flatAtk', 'atkPercent', 'defPercent', 'energyRecharge', 'physicalDamage', 'hydroDamage', 'cryoDamage', 'electroDamage', 'anemoDamage', 'geoDamage', 'pyroDamage', 'dendroDamage', 'critRate', 'critDamage', 'healingBonus', 'elementalMastery')` | 聖遺物のメインステータスの種類。 |
| `mainStatValue` | `FLOAT` | 聖遺物のメインステータスの値。 |
| `subStats` | `JSON` | 聖遺物のサブステータス。キーはステータス名、値はその値。例: `{"flatHp": 100, "critRate": 0.05}`。 ステータス名は以下の通り。 |
|             |      | - `flatHp` |
|             |      | - `hpPercent` |
|             |      | - `flatAtk` |
|             |      | - `atkPercent` |
|             |      | - `flatDef` |
|             |      | - `defPercent` |
|             |      | - `energyRecharge` |
|             |      | - `healingBonus` |
|             |      | - `elementalMastery` |

### 元素(element)の種類
|日本語名|英語名|読み方|
|----|----|----|
|炎元素|Pyro|パイロ|
|水元素|Hydro|ハイドロ|
|氷元素|Cryo|クライオ|
|風元素|Anemo|アネモ|
|岩元素|Geo|ジオ|
|雷元素|Electro|エレクトロ|
|草元素|Dendro|デンドロ|



## `weapon`

| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `id` | `VARCHAR(36)` | 武器のID。 |
| `name` | `VARCHAR(255)` | 武器の名前。 |
| `type` | `ENUM('sword', 'claymore', 'polearm', 'bow', 'catalyst')` | 武器の種類。 |
| `baseAttack` | `INT` | 武器の基礎攻撃力。 |
| `subStat` | `ENUM('hpPercent', 'atkPercent', 'defPercent', 'energyRecharge', 'physicalDamage', 'critRate', 'critDamage', 'elementalMastery')` | 武器のサブステータスの種類。 |
| `subStatValue` | `FLOAT` | 武器のサブステータスの値。 |
| `additionalEffect` | `JSON` | 武器の効果。キーは効果の種類、値はその値。例: `{"effectType": "atkPercent", "effectValue": 0.2}`。効果の種類は以下の通り。 |
|             |      | - `atkPercent` |
|             |      | - `damageBonus` |
|             |      | - `elementalMastery` |
|             |      | - `critRate` |
|             |      | - `critDamage` |

## `character`
| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `characterID` | `VARCHAR(36)` | キャラクターのID。 |
| `name` | `VARCHAR(255)` | キャラクターの名前。 |
| `element` | `ENUM('pyro', 'hydro', 'electro', 'anemo', 'geo', 'dendro')` | キャラクターの属性。 |
| `weaponType` | `ENUM('sword', 'claymore', 'polearm', 'bow', 'catalyst')` | キャラクターの武器タイプ。 |
| `baseStats` | `JSON` | キャラクターの基礎ステータス(計算に必要な値のみを保持)。キーはステータス名、値はその値。例: `{"hp": 1000, "atk": 200, "def": 150}`。 |