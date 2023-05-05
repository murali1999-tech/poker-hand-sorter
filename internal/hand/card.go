
package hand

// Suit represents the suit of a playing card.
type Suit int

// Rank represents the rank of a playing card.
type Rank int

// Card represents a playing card with a suit and rank.
type Card struct {
    Suit Suit
    Rank Rank
}

// NewCard creates a new card with the specified suit and rank.
func NewCard(suit Suit, rank Rank) Card {
    return Card{Suit: suit, Rank: rank}
}

// String returns a string representation of the card in the format "rank of suit".
func (c Card) String() string {
    return c.Rank.String() + " of " + c.Suit.String()
}

// Suit constants represent the possible suits of a playing card.
const (
    Spades Suit = iota
    Hearts
    Diamonds
    Clubs
)

// Rank constants represent the possible ranks of a playing card.
const (
    Two Rank = iota + 2
    Three
    Four
    Five
    Six
    Seven
    Eight
    Nine
    Ten
    Jack
    Queen
    King
    Ace
)

// String returns a string representation of the suit.
func (s Suit) String() string {
    switch s {
    case Spades:
        return "Spades"
    case Hearts:
        return "Hearts"
    case Diamonds:
        return "Diamonds"
    case Clubs:
        return "Clubs"
    default:
        return ""
    }
}

// String returns a string representation of the rank.
func (r Rank) String() string {
    switch r {
    case Two:
        return "Two"
    case Three:
        return "Three"
    case Four:
        return "Four"
    case Five:
        return "Five"
    case Six:
        return "Six"
    case Seven:
        return "Seven"
    case Eight:
        return "Eight"
    case Nine:
        return "Nine"
    case Ten:
        return "Ten"
    case Jack:
        return "Jack"
    case Queen:
        return "Queen"
    case King:
        return "King"
    case Ace:
        return "Ace"
    default:
        return ""
    }
}
