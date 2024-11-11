package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

var (
	uniformType             string
	hasGloves               bool
	hasBelt                 bool
	serviceStripesAmountStr string
	hasNametape             bool
	badgesAmountStr         string
	extrasAmountStr         string
	hasRope                 bool
	awardsAmountStr         string
	serviceStripesAmount 	int
	badgesAmount 			int
	extrasAmount  			int
	awardsAmount 			int
)

const (
	mccuuPrice          = 35
	dressBluesPrice     = 25
	serviceUniformPrice = 25
	messUniformPrice    = 20
	glovesPrice         = 5
	beltPrice           = 10
	serviceStripePrice  = 3
	maxServiceStripes   = 8
	nametapePrice       = 5
	extraPrice          = 5
	badgePrice          = 3
	ropePrice           = 5
	awardPrice          = 2
)

var (
	glovesAllowed        = []string{"dba", "dbb", "mu"}
	ropeAllowed          = []string{"dba", "dbb", "sa"}
	nametapeAllowed      = []string{"mccuu", "dbc", "dbd", "sb", "sc"}
	beltAllowed          = []string{"mccuu", "dbc", "dbd", "sa", "sb", "sc"}
	serviceStripeAllowed = []string{"dba", "dbb", "sa"}
	extrasAllowed        = []string{"mccuu", "dbd", "sc"}
	awardsAllowed        = []string{"dba", "dbb", "dbc", "dbd", "sa", "sb", "sc", "mu"}
	badgesAllowed        = []string{"mccuu", "dba", "dbb", "dbc", "dbd", "sa", "sb", "sc", "mu"}
)

func uniformSelect() *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[string]().
			Title("Choose the uniform type:").
			Options(
				huh.NewOption("Marine Corps Combat Utility Uniform", "mccuu"),
				huh.NewOption("Dress Blue Alphas", "dba"),
				huh.NewOption("Dress Blue Bravos", "dbb"),
				huh.NewOption("Dress Blue Charlies", "dbc"),
				huh.NewOption("Dress Blue Deltas", "dbd"),
				huh.NewOption("Service Alphas", "sa"),
				huh.NewOption("Service Bravos", "sb"),
				huh.NewOption("Service Charlies", "sc"),
				huh.NewOption("Mess Uniform", "mu"),
			).
			Value(&uniformType),
	)
}

func yesNoSelectHide(question string, value *bool, allowedUniforms []string) *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[bool]().
			Title(question).
			Options(
				huh.NewOption("Yes", true),
				huh.NewOption("No", false),
			).
			Value(value),
	).WithHideFunc(func() bool {
		return !contains(allowedUniforms, uniformType)
	})
}

func numberSelectHide(question string, value *string, allowedUniforms []string, maxNum int) *huh.Group {
	return huh.NewGroup(
		huh.NewInput().
			Title(question).
			Prompt("? ").
			Validate(func(str string) error {
				var i int
				_, err := fmt.Sscanf(str, "%d", &i)

				if err != nil {
					return errors.New("not a number")
				}

				if maxNum <= 0 {
					return nil
				}
				if i > maxNum {
					return errors.New("user input greater than maximum allowed amount")
				}
				return nil
			}).
			Value(value),
	).WithHideFunc(func() bool {
		return !contains(allowedUniforms, uniformType)
	})
}

func contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {

	form := huh.NewForm(
		uniformSelect(),
		yesNoSelectHide("Do you wish to have gloves?", &hasGloves, glovesAllowed),
		yesNoSelectHide("Do you wish to have a rope?", &hasRope, ropeAllowed),
		yesNoSelectHide("Do you wish to have a nametape?", &hasNametape, nametapeAllowed),
		yesNoSelectHide("Do you wish to have a belt?", &hasBelt, beltAllowed),
		numberSelectHide("How many badges do you wish to have?", &badgesAmountStr, badgesAllowed, 0),
		numberSelectHide(
			"How many service stripes do you wish to have?",
			&serviceStripesAmountStr,
			serviceStripeAllowed,
			maxServiceStripes,
		),
		numberSelectHide("How many extra accessories do you wish to have?", &extrasAmountStr, extrasAllowed, 0),
		numberSelectHide("How many awards (medals/ribbons) do you wish to have?", &awardsAmountStr, awardsAllowed, 0),
	)

	err := form.Run()
	if err != nil {
		log.Fatalf("Error while running form: %s", err)
	}

	// these are already error checked, no issues with error checking
	fmt.Sscanf(badgesAmountStr, "%d", &badgesAmount)
	fmt.Sscanf(serviceStripesAmountStr, "%d", &serviceStripesAmount)	
	fmt.Sscanf(extrasAmountStr, "%d", &extrasAmount)
	fmt.Sscanf(awardsAmountStr, "%d", &awardsAmount)

	var price int
	switch uniformType {
	case "mccuu": 
		price = mccuuPrice
	case "dba", "dbb", "dbc", "dbd":
		price = dressBluesPrice
	case "sa", "sb", "sc":
		price = serviceUniformPrice
	case "mu":
		price = messUniformPrice
	default:
		log.Fatalf("Invalid uniform: %s", uniformType)
	}

	if hasBelt {
		price += beltPrice
	}
	if hasGloves {
		price += glovesPrice
	}
	if hasNametape {
		price += nametapePrice
	}
	if hasRope {
		price += ropePrice
	}

	price += badgesAmount * badgePrice
	price += serviceStripesAmount * serviceStripePrice
	price += extrasAmount * extraPrice
	price += awardsAmount * awardPrice

	fmt.Printf("✔️ Uniform Price: %d\r\n", price)
}
