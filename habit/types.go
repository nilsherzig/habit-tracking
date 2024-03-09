
type QuanityUnit string

// Habit repräsentiert eine Gewohnheit, mit einem Namen und einem Ziel.
type Habit struct {
	Name    string
	Goal    HabitGoal
	History HabitHistory
}

// HabitGoal ist ein Interface für Ziele, die eine Gewohnheit haben kann.
type HabitGoal interface {
	Describe() string // Eine Methode, die eine Beschreibung des Ziels zurückgibt
}

// HabitHistory repräsentiert die Historie einer Gewohnheit.
type HabitHistory struct {
	Actions []HabitAction
}

func (h HabitHistory) Total() int {
	var total int
	for _, action := range h.Actions {
		total += action.Amount
	}
	return total
}

func (h HabitHistory) Describe() string {
	return fmt.Sprintf("Total: %d", h.Total())
}

func (h HabitHistory) Details() string {
	var details string
	for _, action := range h.Actions {
		timeStampAsString := time.Unix(action.When, 0).Format("2006-01-02 15:04:05")
		details += fmt.Sprintf("%d times at %s\n", action.Amount, timeStampAsString)
	}
	return details
}

// newAction fügt eine neue Aktion zur Historie einer Gewohnheit hinzu.
func (h *Habit) newAction(amount int) {
	h.History.Actions = append(h.History.Actions, HabitAction{
		When:   time.Now().Unix(),
		Amount: amount,
	})
}

// HabitAction represents a single action that is part of a habit's history.
type HabitAction struct {
	When   int64
	Amount int // e.g. number of pages read, or distance run
}

// EveryNDaysGoal repräsentiert ein Ziel, das alle N Tage erreicht werden soll.
type EveryNDaysGoal struct {
	IntervalDays int
}

func (edg EveryNDaysGoal) Describe() string {
	// Implementiere die Beschreibung für ein Ziel, das alle N Tage erreicht werden soll
	return fmt.Sprintf("every %d days", edg.IntervalDays)
}

// NTimesInTimeframeGoal repräsentiert ein Ziel, das eine bestimmte Anzahl in einem bestimmten Zeitraum erreicht werden soll.

type NTimesInTimeframeGoal struct {
	Quantity  int
	Timeframe Timeframe
}

func (ntg NTimesInTimeframeGoal) Describe() string {
	return fmt.Sprintf("%d times in %s", ntg.Quantity, ntg.Timeframe.InSentence())
}

// Timeframe repräsentiert einen Zeitraum, wie Tag, Woche, Monat oder Jahr.

type Timeframe struct {
	Name     string
	Duration int // duration in seconds
}

func (t Timeframe) InSentence() string {
	return fmt.Sprintf("a %s", strings.ToLower(t.Name))
}

var (
	Day   = Timeframe{Name: "Day", Duration: 24 * 60 * 60}
	Week  = Timeframe{Name: "Week", Duration: 7 * 24 * 60 * 60}
	Month = Timeframe{Name: "Month", Duration: 30 * 24 * 60 * 60}
	Year  = Timeframe{Name: "Year", Duration: 365 * 24 * 60 * 60}
)

// QuanityGoal repräsentiert ein Ziel, das eine bestimmte Anzahl erreicht werden soll.
type QuanityGoal struct {
	Unit      QuanityUnit
	Quantity  int
	Timeframe Timeframe
}

func (qg QuanityGoal) Describe() string {
	return fmt.Sprintf("%d %s in %s", qg.Quantity, qg.Unit, qg.Timeframe.InSentence())
}
