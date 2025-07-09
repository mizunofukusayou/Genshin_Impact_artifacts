package db

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
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

func SearchBestArtifacts(c echo.Context) error {
	characterIDStr := c.Param("characterID")
	characterID, err := uuid.FromString(characterIDStr)
	if err != nil {
		c.Logger().Error("Invalid character ID:", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid character ID")
	}
	// キャラクターの情報を取得
	character, err := GetCharacter(characterID)
	if err != nil {
		c.Logger().Error("Failed to get character:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get character")
	}
	// キャラクターが存在しない場合
	if character.ID == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, "Character not found")
	}

	// 聖遺物の組み合わせを返す
	var combinationAllArtifacts []ArtifactSet
	combinationAllArtifacts, err = getAllCombinationArtifacts()
	if err != nil {
		c.Logger().Error("Failed to get all combination artifacts:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get artifacts")
	}

	// 聖遺物から期待値を計算する
	for i := range combinationAllArtifacts {
		// 聖遺物の効果をBuffでまとめる
		buff, err := sumUpBuff(&combinationAllArtifacts[i])
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to sum up buff")
		}

		// todo: ベネットのバフ、武器効果とかを計算してBuffと足し合わせたい
		buff2 := Buff{
			AttackPercentage: 0.2,
			FlatAttack:       100,
			CritDamage:       0.15,
			CritRate:         0.1,
		}
		// Buffの合計を計算
		character.Buff = Buff{
			AttackPercentage: buff.AttackPercentage + buff2.AttackPercentage,
			FlatAttack:       buff.FlatAttack + buff2.FlatAttack,
			CritDamage:       buff.CritDamage + buff2.CritDamage,
			CritRate:         buff.CritRate + buff2.CritRate,
		}

		// todo: 並列処理をさせたい
		value, err := calculateDamage(character)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to calculate damage")
		}
		combinationAllArtifacts[i].ExpectedDamage = value
	}
	return echo.NewHTTPError(http.StatusOK, combinationAllArtifacts)
}

func GetCharacter(characterID uuid.UUID) (Character, error) {
	var character Character
	err := Character.Where("id", characterID).FindAll(&character)
	if err != nil {
		return Character{}, err
	}
	return character, nil
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

func sumUpBuff(artifactSet *ArtifactSet) (Buff, error) {
	// Buffをまとめるロジックを実装する
	return Buff{
		AttackPercentage: artifactSet.FlowerOfLife.Substats["atkPercent"] + artifactSet.PlumeOfDeath.Substats["atkPercent"] + artifactSet.SandsOfEon.Substats["atkPercent"] + artifactSet.GobletOfEonothem.Substats["atkPercent"] + artifactSet.CircletOfLogos.Substats["atkPercent"],
		FlatAttack:       artifactSet.FlowerOfLife.Substats["flatAtk"] + artifactSet.PlumeOfDeath.Substats["flatAtk"] + artifactSet.SandsOfEon.Substats["flatAtk"] + artifactSet.GobletOfEonothem.Substats["flatAtk"] + artifactSet.CircletOfLogos.Substats["flatAtk"],
		CritDamage:       artifactSet.FlowerOfLife.Substats["critDamage"] + artifactSet.PlumeOfDeath.Substats["critDMG"] + artifactSet.SandsOfEon.Substats["critDMG"] + artifactSet.GobletOfEonothem.Substats["critDMG"] + artifactSet.CircletOfLogos.Substats["critDMG"],
		CritRate:         artifactSet.FlowerOfLife.Substats["critRate"] + artifactSet.PlumeOfDeath.Substats["critRate"] + artifactSet.SandsOfEon.Substats["critRate"] + artifactSet.GobletOfEonothem.Substats["critRate"] + artifactSet.CircletOfLogos.Substats["critRate"],
		// todo: メインステータスも入れないと
	}, nil
}

func calculateDamage(character Character) (float64, error) {
	// todo: 期待値を計算するロジックを実装する
	return 0, nil
}
