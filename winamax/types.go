package winamax

// Cleanning data
type OutcomesId struct {
	Label               string `json:"label"`
	Available           bool   `json:"available"`
	Code                string `json:"code"`
	PercentDistrubution int    `json:"percentDistribution"`
}
type BetId struct {
	BetId                int    `json:"betId"`
	MarketId             int    `json:"marketId"`
	SpecialBetValue      string `json:"specialBetValue"`
	Outcomes             []int  `json:"outcomes"`
	Available            bool   `json:"available"`
	Template             string `json:"template"`
	IsAlternativeMainBet bool   `json:"isAlternativeMainBet"`
	BetTypeIsLive        bool   `json:"betTypeIsLive"`
	BetTitle             string `json:"betTitle"`
	BetTypeName          string `json:"betTypeName"`
	BetType              int    `json:"betType"`
	BetTypeCategoryId    int    `json:"betTypeCategoryId"`
	BetCategories        []int  `json:"betCategories"`
	BetTypeCategory      string `json:"betTypeCategory"`
	BetTypeHelp          string `json:"betTypeHelp"`
	BetGroup             int    `json:"betGroup"`
}
type MatchId struct {
	MatchId                   int    `json:"matchId"`
	Available                 bool   `json:"available"`
	Status                    string `json:"status"`
	MainBetId                 int    `json:"mainBetId"`
	MoreBets                  int    `json:"moreBets"`
	Score                     string `json:"score"`
	PeriodId                  int    `json:"periodId"`
	Period                    string `json:"period"`
	PeriodName                string `json:"periodName"`
	HlType                    string `json:"hlType"`
	SetScores                 string `json:"setScores"`
	GameScore                 string `json:"gameScore"`
	HomeGamescore             int    `json:"homeGamescore"`
	AwayGamescore             int    `json:"awayGamescore"`
	Server                    int    `json:"server"`
	Title                     string `json:"title"`
	RoundName                 string `json:"roundName"`
	TvChannels                string `json:"tvChannels"`
	SportId                   int    `json:"sportId"`
	CategoryId                int    `json:"categoryId"`
	TournamentId              int    `json:"tournamentId"`
	SrTournamentId            string `json:"srTournamentId"`
	SrSeasonId                string `json:"srSeasonId"`
	Competitor1Id             int    `json:"competitor1Id"`
	Competitor1Name           string `json:"competitor1Name"`
	Competitor1Flag           string `json:"competitor1Flag"`
	Competitor2Id             int    `json:"competitor2Id"`
	Competitor2Name           string `json:"competitor2Name"`
	Competitor2Flag           string `json:"competitor2Flag"`
	MatchStart                int    `json:"matchStart"`
	IsBooked                  bool   `json:"isBooked"`
	PrematchDisplayOrder      int    `json:"prematchDisplayOrder"`
	LiveDisplayOrder          string `json:"liveDisplayOrder"`
	Filters                   []int  `json:"filters"`
	StreamHighlightsAvailable bool   `json:"streamHighlightsAvailable"`
	Highlights                []any  `json:"highlights"`
	LineupsAvailable          bool   `json:"lineupsAvailable"`
}
type Data struct {
	SportIds             []int                 `json:"sportIds"`
	Sports               map[string]any        `json:"sports"`
	Categories           map[string]any        `json:"categories"`
	Tournaments          map[string]any        `json:"tournaments"`
	Calendar             map[string]any        `json:"calendar"`
	MainMatchCount       int                   `json:"mainMatchCount"`
	LiveMatchCount       int                   `json:"liveMatchCount"`
	ReallyLiveMatchCount int                   `json:"reallyLiveMatchCount"`
	TvMatchCount         int                   `json:"tvMatchCount"`
	Filters              map[string]any        `json:"filters"`
	Missions             map[string]any        `json:"missions"`
	Achievements         map[string]any        `json:"achievements"` // moyen ça fait bugger
	BetCategories        map[string]any        `json:"betCategories"`
	Teasers              []any                 `json:"teasers"`
	Matches              map[string]MatchId    `json:"matches"` // peut être à changer
	Bets                 map[string]BetId      `json:"bets"`    // peut êterr à changer
	Outcomes             map[string]OutcomesId `json:"outcomes"`
	Odds                 map[string]float64    `json:"odds"`
	Settings             map[string]any        `json:"settings"`
	Cc                   map[string]any        `json:"cc"`
	Browsers             map[string]any        `json:"browser"`
	Notifications        string                `json:"notifications"`
	Terms                bool                  `json:"terms"`
	Events               bool                  `json:"events"`
}

type CleanBet struct {
	Title string
	Odd   float64 // peut être mettre number ???
}
type CleanMatch struct {
	Title string
	Bet1  CleanBet
	Bet2  CleanBet
	Draw  CleanBet
}
type WinamaxData map[string]*CleanMatch
