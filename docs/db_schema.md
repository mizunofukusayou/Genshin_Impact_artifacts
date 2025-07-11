# DB スキーマ
## `status`
| 列名 | 型 | 制約 | 説明 |
|-------------|------|-------------|-------------|
| `id` | `INT` | PRIMARY KEY | ステータスのID |
| `type` | `VARCHAR(32)` | UNIQUE | ステータスの種類 |

## `artifacts`

| 列名 | 型 | 制約 | 説明 |
|------------|------|-------------|-------------|
| `id` | `VARCHAR(36)` | PRIMARY KEY | 聖遺物のUUID |
| `setName` | `VARCHAR(32)` | NOT NULL | 聖遺物セットの名前 |
| `setEffect` | `JSON` | NOT NULL | 聖遺物セットの効果。キーは`status`テーブルの`id`、値はその値。例: `{"1": "0.2", "2": "0.15"}`|
| `type` | `ENUM('flowerOfLife', 'plumeOfDeath', 'sandsOfEon', 'gobletOfEonothem', 'circletOfLogos')` | NOT NULL | 部位の種類。 |
| `mainStat` | `INT` | NOT NULL | 聖遺物のメインステータスの種類`status`テーブルの`id`。 |
| `mainStatValue` | `FLOAT` | NOT NULL | 聖遺物のメインステータスの値。 |
| `subStats` | `JSON` | NOT NULL | 聖遺物のサブステータス。キーは`status`テーブルの`id`、値はその値。例: `{1: 100, 2: 0.05}` |

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



## `weapons`

| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `id` | `VARCHAR(36)` | 武器のUUID |
| `name` | `VARCHAR(32)` | 武器の名前 |
| `type` | `ENUM('sword', 'claymore', 'polearm', 'bow', 'catalyst')` | 武器の種類 |
| `baseAttack` | `INT` | 武器の基礎攻撃力。 |
| `effect` | `JSON` | 武器の効果（サブステータス+追加効果）。キーは`status`テーブルの`id`、値はその値。例: `{"1": 100, "2": 0.05}` |

## `characters`
| 列名 | 型 | 説明 |
|-------------|------|-------------|
| `id` | `VARCHAR(36)` | キャラクターのID。 |
| `name` | `VARCHAR(32)` | キャラクターの名前。 |
| `element` | `ENUM('pyro', 'hydro', 'electro', 'anemo', 'geo', 'dendro')` | キャラクターの属性。 |
| `weaponType` | `ENUM('sword', 'claymore', 'polearm', 'bow', 'catalyst')` | キャラクターの武器タイプ。 |
| `baseStatus` | `JSON` | キャラクターの基礎ステータス(計算に必要な値のみを保持)。キーはステータスID、値はその値。例: `{"hp": 1000, "attack": 200}`。 |
