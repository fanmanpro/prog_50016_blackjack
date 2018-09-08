package blackjack

type CardDefinition struct {
	Score int
	IsAce bool
}

var cardDefinitions map[string]*CardDefinition

func init() {
	cardDefinitions = map[string]*CardDefinition{
		"SA":  &CardDefinition{Score: 11, IsAce: true},
		"S2":  &CardDefinition{Score: 2},
		"S3":  &CardDefinition{Score: 3},
		"S4":  &CardDefinition{Score: 4},
		"S5":  &CardDefinition{Score: 5},
		"S6":  &CardDefinition{Score: 6},
		"S7":  &CardDefinition{Score: 7},
		"S8":  &CardDefinition{Score: 8},
		"S9":  &CardDefinition{Score: 9},
		"S10": &CardDefinition{Score: 10},
		"SJ":  &CardDefinition{Score: 10},
		"SQ":  &CardDefinition{Score: 10},
		"SK":  &CardDefinition{Score: 10},
		"CA":  &CardDefinition{Score: 11, IsAce: true},
		"C2":  &CardDefinition{Score: 2},
		"C3":  &CardDefinition{Score: 3},
		"C4":  &CardDefinition{Score: 4},
		"C5":  &CardDefinition{Score: 5},
		"C6":  &CardDefinition{Score: 6},
		"C7":  &CardDefinition{Score: 7},
		"C8":  &CardDefinition{Score: 8},
		"C9":  &CardDefinition{Score: 9},
		"C10": &CardDefinition{Score: 10},
		"CJ":  &CardDefinition{Score: 10},
		"CQ":  &CardDefinition{Score: 10},
		"CK":  &CardDefinition{Score: 10},
		"DA":  &CardDefinition{Score: 11, IsAce: true},
		"D2":  &CardDefinition{Score: 2},
		"D3":  &CardDefinition{Score: 3},
		"D4":  &CardDefinition{Score: 4},
		"D5":  &CardDefinition{Score: 5},
		"D6":  &CardDefinition{Score: 6},
		"D7":  &CardDefinition{Score: 7},
		"D8":  &CardDefinition{Score: 8},
		"D9":  &CardDefinition{Score: 9},
		"D10": &CardDefinition{Score: 10},
		"DJ":  &CardDefinition{Score: 10},
		"DQ":  &CardDefinition{Score: 10},
		"DK":  &CardDefinition{Score: 10},
		"HA":  &CardDefinition{Score: 11, IsAce: true},
		"H2":  &CardDefinition{Score: 2},
		"H3":  &CardDefinition{Score: 3},
		"H4":  &CardDefinition{Score: 4},
		"H5":  &CardDefinition{Score: 5},
		"H6":  &CardDefinition{Score: 6},
		"H7":  &CardDefinition{Score: 7},
		"H8":  &CardDefinition{Score: 8},
		"H9":  &CardDefinition{Score: 9},
		"H10": &CardDefinition{Score: 10},
		"HJ":  &CardDefinition{Score: 10},
		"HQ":  &CardDefinition{Score: 10},
		"HK":  &CardDefinition{Score: 10},
	}
}