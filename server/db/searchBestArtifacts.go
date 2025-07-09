package db

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

type Artifact struct {
	ID            uuid.UUID          `json:"id"`
	Type          string             `json:"type"`
	MainStat      string             `json:"mainStat"`
	MainStatValue float64            `json:"mainStatValue"`
	Substats      map[string]float64 `json:"substats"` // Key: ステータス名, Value: ステータス値
}

type ArtifactSet struct {
	FlowerOfLife     Artifact
	PlumeOfDeath     Artifact
	SandsOfEon       Artifact
	GobletOfEonothem Artifact
	CircletOfLogos   Artifact
	ExpectedDamage   float64
}

type Buff struct {
	AttackPercentage float64
	FlatAttack       float64
	CritDamage       float64
	CritRate         float64
}

type Character struct {
	ID         uuid.UUID          `json:"id"`
	Name       string             `json:"name"`
	Element    string             `json:"element"`
	WeaponType string             `json:"weaponType"`
	BaseStatus map[string]float64 `json:"baseStatus"` // Key: ステータス名, Value: ステータス値
	Buff       Buff               // キャラクターのバフ効果
}

func SearchBestArtifacts() echo.HandlerFunc {
	// 聖遺物の組み合わせを返す
	var combinationAllArtifacts []ArtifactSet
	combinationAllArtifacts, err := getAllCombinationArtifacts()
	if err != nil {
		return echo.NewHTTPError(500, "Failed to get artifacts")
	}

	// 聖遺物から期待値を計算する
	for i := range combinationAllArtifacts {
		// 聖遺物の効果をBuffでまとめる
		var buff Buff
		buff, err = sumUpBuff(&combinationAllArtifacts[i])
		if err != nil {
			return echo.NewHTTPError(500, "Failed to sum up buff")
		}

		// todo: ベネットのバフ、武器効果とかを計算してBuffと足し合わせたい
		var buff2 Buff
		buff2 = Buff{
			AttackPercentage: 0.2,
			FlatAttack:       100,
			CritDamage:       0.15,
			CritRate:         0.1,
		}
		buff += buff2

		// todo: 並列処理をさせたい
		value, err := calculateDamage(buff)
		if err != nil {
			return echo.NewHTTPError(500, "Failed to calculate damage")
		}
		combinationAllArtifacts[i].ExpectedDamage = value
	}
	return echo.NewHTTPError(200, combinationAllArtifacts)
}

func getAllCombinationArtifacts() ([]ArtifactSet, error) {
	var ret []ArtifactSet
	var flowerOfLifes []Artifact
	var plumeOfDeaths []Artifact
	var sandsOfEons []Artifact
	var gobletOfEonothem []Artifact
	var circletOfLogos []Artifact
	// todo: データベースの接続について勉強する
	// todo: データベースをGORMで扱う方法を調べる
	Artifact.Where("flowerOfLife").FindAll(&flowerOfLifes)
	Artifact.Where("plumeOfDeath").FindAll(&plumeOfDeaths)
	Artifact.Where("sandsOfEon").FindAll(&sandsOfEons)
	Artifact.Where("gobletOfEonothem").FindAll(&gobletOfEonothem)
	Artifact.Where("circletOfLogos").FindAll(&circletOfLogos)
	for _, flower := range flowerOfLifes {
		for _, plume := range plumeOfDeaths {
			for _, sands := range sandsOfEons {
				for _, goblet := range gobletOfEonothem {
					for _, circlet := range circletOfLogos {
						artifactSet := ArtifactSet{
							FlowerOfLife:     flower,
							PlumeOfDeath:     plume,
							SandsOfEon:       sands,
							GobletOfEonothem: goblet,
							CircletOfLogos:   circlet,
						}
						ret = append(ret, artifactSet)
					}
				}
			}
		}
	}
	if len(ret) == 0 {
		return nil, echo.NewHTTPError(404, "No artifacts found")
	}
	return ret, nil
}

func sumUpBuff(artifactSet *ArtifactSet) error {
	// Buffをまとめるロジックを実装する
	artifactSet.Buff = Buff{
		AttackPercentage: artifactSet.FlowerOfLife.Substats["atkPercent"] + artifactSet.PlumeOfDeath.Substats["atkPercent"] + artifactSet.SandsOfEon.Substats["atkPercent"] + artifactSet.GobletOfEonothem.Substats["atkPercent"] + artifactSet.CircletOfLogos.Substats["atkPercent"],
		FlatAttack:       artifactSet.FlowerOfLife.Substats["flatAtk"] + artifactSet.PlumeOfDeath.Substats["flatAtk"] + artifactSet.SandsOfEon.Substats["flatAtk"] + artifactSet.GobletOfEonothem.Substats["flatAtk"] + artifactSet.CircletOfLogos.Substats["flatAtk"],
		CritDamage:       artifactSet.FlowerOfLife.Substats["critDamage"] + artifactSet.PlumeOfDeath.Substats["critDMG"] + artifactSet.SandsOfEon.Substats["critDMG"] + artifactSet.GobletOfEonothem.Substats["critDMG"] + artifactSet.CircletOfLogos.Substats["critDMG"],
		CritRate:         artifactSet.FlowerOfLife.Substats["critRate"] + artifactSet.PlumeOfDeath.Substats["critRate"] + artifactSet.SandsOfEon.Substats["critRate"] + artifactSet.GobletOfEonothem.Substats["critRate"] + artifactSet.CircletOfLogos.Substats["critRate"],
		// todo: メインステータスも入れないと
	}
	return nil
}

func calculateDamage(artifactSet ArtifactSet) (float64, error) {
	// todo: 期待値を計算するロジックを実装する
	return 0, nil
}
