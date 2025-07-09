# DB スキーマ

## `artifact`

| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `id` | `VARCHAR(36)` | 聖遺物のID。 |
| `setName` | `VARCHAR(255)` | 聖遺物セットの名前。 |
| `setEffect` | `JSON` | 聖遺物セットの効果。キーは効果の種類、値はその値。例: `{"hpPercent": "0.2", "damageBonus": "0.15"}`。効果の種類は以下の通り。 `ENUM('hpPercent', 'atkPercent', 'damageBonus')` |
| `type` | `ENUM('flowerOfLife', 'plumeOfDeath', 'sandsOfEon', 'gobletOfEonothem', 'circletOfLogos')` | 部位の種類。 |
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
| `additionalEffect` | `JSON` | 武器の効果。配列形式で複数の効果を定義。例: `[{"type": "atkPercent", "value": 0.2}, {"type": "damageBonus", "value": 0.12}]`。効果の種類は以下の通り。 |
|             |      | **効果タイプ (`type`):** |
|             |      | - `atkPercent` - 攻撃力％アップ |
|             |      | - `damageBonus` - ダメージボーナス |
|             |      | - `elementalMastery` - 元素熟知アップ |
|             |      | - `critRate` - 会心率アップ |
|             |      | - `critDamage` - 会心ダメージアップ |
|             |      | - `energyRecharge` - 元素チャージ効率アップ |

## `character`
| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `id` | `VARCHAR(36)` | キャラクターのID。 |
| `name` | `VARCHAR(255)` | キャラクターの名前。 |
| `element` | `ENUM('pyro', 'hydro', 'electro', 'anemo', 'geo', 'dendro')` | キャラクターの属性。 |
| `weaponType` | `ENUM('sword', 'claymore', 'polearm', 'bow', 'catalyst')` | キャラクターの武器タイプ。 |
| `baseStatus` | `JSON` | キャラクターの基礎ステータス(計算に必要な値のみを保持)。キーはステータス名、値はその値。例: `{"hp": 1000, "attack": 200}`。 |