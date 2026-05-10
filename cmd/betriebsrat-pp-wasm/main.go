//go:build js && wasm

package main

import (
	"betriebsrat-pp-cli/internal/betrvg"
	"encoding/json"
	"strings"
	"syscall/js"
)

func main() {
	js.Global().Set("betriebsratAsk", js.FuncOf(askFn))
	select {} // keep the WASM module alive
}

type wasmResult struct {
	Question       string     `json:"question"`
	Lang           string     `json:"lang"`
	Classification string     `json:"classification"`
	RightType      string     `json:"right_type,omitempty"`
	Paragraphs     []wasmPara `json:"paragraphs,omitempty"`
	Answer         string     `json:"answer"`
	Actions        []string   `json:"actions,omitempty"`
	Deadline       string     `json:"deadline,omitempty"`
	Disclaimer     string     `json:"disclaimer"`
}

type wasmPara struct {
	Number    int    `json:"paragraph"`
	Title     string `json:"title"`
	RightType string `json:"right_type"`
	TopicURL  string `json:"topic_url,omitempty"`
}

func askFn(this js.Value, args []js.Value) any {
	question, lang := "", "de"
	if len(args) > 0 {
		question = args[0].String()
	}
	if len(args) > 1 {
		lang = args[1].String()
	}

	words := tokenize(question)
	paras := betrvg.ByKeywords(words)

	result := wasmResult{
		Question:   question,
		Lang:       lang,
		Disclaimer: tr(lang, "Orientierungshilfe, kein Rechtsgutachten. Konsultieren Sie einen Fachanwalt für Arbeitsrecht.", "Legal orientation only, not legal advice. Consult a labour law specialist."),
	}

	result.Classification = classifySituation(lang, question)

	for _, p := range paras {
		result.Paragraphs = append(result.Paragraphs, wasmPara{p.Number, p.Title, string(p.CoDetermType), p.TopicURL})
	}

	if len(paras) > 0 {
		strongest := paras[0].CoDetermType
		result.RightType = string(strongest)
		result.Answer = buildAnswer(lang, strongest, result.Classification)
		result.Actions = buildActions(lang, strongest)
		// Deadline lookup
		for _, rule := range betrvg.Deadlines() {
			for _, w := range words {
				if betrvg.ContainsFold(rule.Situation, w) || betrvg.ContainsFold(w, rule.Situation) {
					if rule.Days > 0 {
						result.Deadline = tr(lang,
							formatDeadlineDE(rule),
							formatDeadlineEN(rule))
					}
					break
				}
			}
			if result.Deadline != "" {
				break
			}
		}
	} else {
		result.Answer = classificationFallback(lang, result.Classification)
	}

	data, _ := json.Marshal(result)
	return string(data)
}

func formatDeadlineDE(r betrvg.DeadlineRule) string {
	return "§ " + itoa(r.Paragraph) + " — " + itoa(r.Days) + " Tage: " + r.Description
}

func formatDeadlineEN(r betrvg.DeadlineRule) string {
	return "§ " + itoa(r.Paragraph) + " — " + itoa(r.Days) + " days: " + r.Description
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}

func tokenize(s string) []string {
	s = strings.ToLower(s)
	var words []string
	for _, w := range strings.Fields(s) {
		w = strings.Trim(w, ".,!?;:\"'()[]{}–—")
		if len(w) >= 3 {
			words = append(words, w)
		}
	}
	return words
}

func tr(lang, de, en string) string {
	if lang == "en" {
		return en
	}
	return de
}

func classifySituation(lang, situation string) string {
	low := strings.ToLower(situation)
	switch {
	case containsAny(low, "software", "analytics", "überwachung", "monitoring", "ki-system",
		"künstliche intelligenz", "artificial intelligence", "surveillance", "tracking", "telemetry"):
		return tr(lang, "Technische Einrichtung – § 87 Abs. 1 Nr. 6", "Technical facility – § 87 Abs. 1 Nr. 6")
	case containsAny(low, "massenentlassung", "sozialplan", "mass dismissal", "mass layoff", "layoff"):
		return tr(lang, "Massenentlassung / Sozialplan (§ 112 BetrVG, § 17 KSchG)", "Mass dismissal / Sozialplan (§ 112 BetrVG, § 17 KSchG)")
	case containsAny(low, "kündigung", "entlassung", "fristlos", "dismissal", "fired", "firing"):
		return tr(lang, "Personelle Angelegenheit – Kündigung (§ 102 BetrVG)", "Personnel matter – Dismissal (§ 102 BetrVG)")
	case containsAny(low, "betriebsänderung", "verlagerung", "stilllegung", "umstrukturierung", "outsourcing", "restructuring"):
		return tr(lang, "Betriebsänderung (§ 111 ff. BetrVG)", "Operational change (§ 111 ff. BetrVG)")
	case containsAny(low, "homeoffice", "remote", "mobiles arbeiten", "telearbeit"):
		return tr(lang, "Mobiles Arbeiten – § 87 Abs. 1 Nr. 1, 2", "Mobile work – § 87 Abs. 1 Nr. 1, 2")
	case containsAny(low, "einstellung", "versetzung", "umgruppierung", "hiring", "transfer"):
		return tr(lang, "Personelle Einzelmaßnahme (§ 99 BetrVG)", "Individual personnel measure (§ 99 BetrVG)")
	default:
		return tr(lang, "Allgemeine Betriebsratsangelegenheit", "General works council matter")
	}
}

func containsAny(s string, terms ...string) bool {
	for _, t := range terms {
		if strings.Contains(s, t) {
			return true
		}
	}
	return false
}

func buildAnswer(lang string, strongest betrvg.CoDeterminationType, classification string) string {
	switch strongest {
	case betrvg.MitbestimmungErzwingbar:
		return tr(lang,
			"Erzwingbare Mitbestimmung: Der Arbeitgeber darf ohne Zustimmung des Betriebsrats oder Spruch der Einigungsstelle nicht handeln. Eine Betriebsvereinbarung ist erforderlich.",
			"Enforceable co-determination: The employer may not act without the works council's consent or a ruling by the conciliation board. A Betriebsvereinbarung is required.")
	case betrvg.Zustimmung:
		return tr(lang,
			"Zustimmungsvorbehalt: Der Betriebsrat muss innerhalb einer Woche zustimmen oder schriftlich verweigern. Ohne Reaktion gilt Zustimmung als erteilt.",
			"Consent required: The works council must consent or refuse in writing within one week. No response means consent is deemed given.")
	case betrvg.Mitwirkung:
		return tr(lang,
			"Mitwirkungsrecht: Der Betriebsrat hat das Recht, Einwände zu erheben und Gegenvorschläge zu machen, kann die Maßnahme aber nicht blockieren.",
			"Participation right: The works council may raise objections and counter-proposals, but cannot block the measure.")
	case betrvg.Beratung:
		return tr(lang,
			"Beratungsrecht: Der Arbeitgeber muss den Betriebsrat ernsthaft konsultieren, behält aber die Entscheidungsbefugnis.",
			"Consultation right: The employer must genuinely consult the works council but retains the final decision.")
	default:
		return tr(lang,
			"Für eine vollständige Analyse nutzen Sie bitte die betriebsrat-pp-cli CLI oder wenden Sie sich an einen Fachanwalt für Arbeitsrecht.",
			"For a complete analysis please use the betriebsrat-pp-cli CLI or consult a labour law specialist.")
	}
}

func buildActions(lang string, strongest betrvg.CoDeterminationType) []string {
	actions := []string{
		tr(lang, "Vollständige Unterrichtung durch Arbeitgeber verlangen (§ 80 Abs. 2 BetrVG)", "Demand full disclosure from the employer (§ 80 Abs. 2 BetrVG)"),
	}
	switch strongest {
	case betrvg.MitbestimmungErzwingbar:
		actions = append(actions,
			tr(lang, "Maßnahme stoppen — ohne BR-Zustimmung oder Einigungsstelle darf der AG nicht handeln", "Stop the measure — employer may not act without BR consent or conciliation board ruling"),
			tr(lang, "Betriebsvereinbarung verhandeln oder Einigungsstelle anrufen (§ 76 BetrVG)", "Negotiate a Betriebsvereinbarung or invoke the conciliation board (§ 76 BetrVG)"),
		)
	case betrvg.Zustimmung:
		actions = append(actions,
			tr(lang, "Frist prüfen: Zustimmung oder Verweigerung binnen 1 Woche (§ 99 Abs. 3)", "Check deadline: consent or refusal within 1 week (§ 99 Abs. 3)"),
			tr(lang, "Verweigerungsgründe nach § 99 Abs. 2 prüfen und schriftlich begründen", "Review refusal grounds under § 99 Abs. 2 and document in writing"),
		)
	case betrvg.Mitwirkung:
		actions = append(actions,
			tr(lang, "Schriftliche Stellungnahme mit konkreten Einwänden einreichen", "Submit written statement with specific objections"),
		)
	case betrvg.Beratung:
		actions = append(actions,
			tr(lang, "Ernsthafte Beratung einfordern — nicht nur formale Unterrichtung akzeptieren (§ 74 BetrVG)", "Demand genuine consultation — do not accept mere formal notification (§ 74 BetrVG)"),
		)
	}
	return actions
}

func classificationFallback(lang, classification string) string {
	low := strings.ToLower(classification)
	switch {
	case strings.Contains(low, "87") || strings.Contains(low, "technical") || strings.Contains(low, "technische"):
		return tr(lang,
			"Technische Einrichtungen mit Überwachungspotenzial unterliegen der erzwingbaren Mitbestimmung (§ 87 Abs. 1 Nr. 6 BetrVG).",
			"Technical facilities capable of monitoring employees trigger enforceable co-determination (§ 87 Abs. 1 Nr. 6 BetrVG).")
	case strings.Contains(low, "kündigung") || strings.Contains(low, "dismissal"):
		return tr(lang,
			"Vor jeder Kündigung muss der Betriebsrat angehört werden (§ 102 BetrVG). Ohne vollständige Anhörung ist die Kündigung unwirksam.",
			"Before any dismissal the works council must be heard (§ 102 BetrVG). Without proper hearing the dismissal is void.")
	default:
		return tr(lang,
			"Für eine genaue Analyse installieren Sie betriebsrat-pp-cli: go install github.com/googlarz/betriebsrat-pp-cli@latest",
			"For a precise analysis install betriebsrat-pp-cli: go install github.com/googlarz/betriebsrat-pp-cli@latest")
	}
}
