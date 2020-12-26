package fantasy

import (
	"fmt"
	"strconv"

	"github.com/erichaase/fantasy/internal/espn"
)

type GameLine struct {
	EspnId    int
	FirstName string
	LastName  string
	Min       int
	Fgm       int
	Fga       int
	Ftm       int
	Fta       int
	Tpm       int
	Tpa       int
	Pts       int
	Reb       int
	Ast       int
	Stl       int
	Blk       int
	To        int
	Zfg       float64
	Zft       float64
	Ztp       float64
	Zpts      float64
	Zreb      float64
	Zast      float64
	Zstl      float64
	Zblk      float64
	Zto       float64
	Zsum      float64
}

func BuildGameLine(egl espn.GameLine) GameLine {
	// populateStats()
	// populateZScores()
	min, _ := strconv.Atoi(egl.Minutes)
	pts, _ := strconv.Atoi(egl.Points)
	reb, _ := strconv.Atoi(egl.Rebounds)
	ast, _ := strconv.Atoi(egl.Assists)
	stl, _ := strconv.Atoi(egl.Steals)
	blk, _ := strconv.Atoi(egl.Blocks)
	to, _ := strconv.Atoi(egl.Turnovers)

	var fgm, fga int
	fmt.Sscanf(egl.Fg, "%d/%d", &fgm, &fga)
	var ftm, fta int
	fmt.Sscanf(egl.Ft, "%d/%d", &ftm, &fta)
	var tpm, tpa int
	fmt.Sscanf(egl.Threept, "%d/%d", &tpm, &tpa)

	// details: https://www.reddit.com/r/fantasybball/comments/71bdq0/how_to_calculate_weighted_zscore_for_fg/
	zfg := 0.0
	if fga != 0 {
		zfg = ((((float64(fgm) / float64(fga)) - 0.478) * float64(fga)) + 0.07) / 0.63
	}
	zft := 0.0
	if fta != 0 {
		zft = ((((float64(ftm) / float64(fta)) - 0.780) * float64(fta)) - 0.02) / 0.33
	}
	ztp := (float64(tpm) - 1.69) / 1.01
	zpts := (float64(pts) - 16.29) / 5.88
	zreb := (float64(reb) - 6.21) / 2.58
	zast := (float64(ast) - 3.52) / 2.17
	zstl := (float64(stl) - 1.00) / 0.36
	zblk := (float64(blk) - 0.71) / 0.52
	zto := -((float64(to) - 1.95) / 0.87)
	zsum := zfg + zft + ztp + zpts + zreb + zast + zstl + zblk + zto

	return GameLine{
		EspnId:    egl.Id,
		FirstName: egl.FirstName,
		LastName:  egl.LastName,
		Min:       min,
		Fgm:       fgm,
		Fga:       fga,
		Ftm:       ftm,
		Fta:       fta,
		Tpm:       tpm,
		Tpa:       tpa,
		Pts:       pts,
		Reb:       reb,
		Ast:       ast,
		Stl:       stl,
		Blk:       blk,
		To:        to,
		Zfg:       zfg,
		Zft:       zft,
		Ztp:       ztp,
		Zpts:      zpts,
		Zreb:      zreb,
		Zast:      zast,
		Zstl:      zstl,
		Zblk:      zblk,
		Zto:       zto,
		Zsum:      zsum,
	}
}

func PrintGameLine(l GameLine) {
	fmt.Printf("%s %s,|,%dm,%d-%d,%d-%d,%d-%d,%d-%d-%d,%d-%d-%d,|,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,%.1f,|,%.1f\n",
		l.FirstName,
		l.LastName,
		l.Min,
		l.Fgm,
		l.Fga,
		l.Ftm,
		l.Fta,
		l.Tpm,
		l.Tpa,
		l.Pts,
		l.Reb,
		l.Ast,
		l.Stl,
		l.Blk,
		l.To,
		l.Zfg,
		l.Zft,
		l.Ztp,
		l.Zpts,
		l.Zreb,
		l.Zast,
		l.Zstl,
		l.Zblk,
		l.Zto,
		l.Zsum,
	)
}
