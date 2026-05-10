# betriebsrat-pp-cli

**German works council advisor — BetrVG rights, deadlines, and decision support — fully offline, in your terminal.**

> **[🇩🇪 Deutsche Version weiter unten](#deutsch)**

---

No API key. No login. No subscription. Everything runs locally from a pre-built knowledge base of BetrVG paragraphs, legal deadlines, and situation-aware decision logic.

Built for works council (*Betriebsrat*) members who need answers in minutes, not hours of reading law textbooks or searching legal websites.

## What it does

| Question | Command |
|---|---|
| Does the BR have a say in this? | `rights-check` |
| What kind of right is it? | `codetermination-type` |
| What should we do, step by step? | `decide` |
| When must we respond? | `deadline` |
| What does § 87 actually mean? | `law 87` |
| What if the employer acts without us? | `consequences` |
| Is our Anhörungsschreiben valid? | `check-anhoerung` |
| How much Sozialplan does this employee get? | `sozialplan-calc` |
| Draft a Betriebsvereinbarung | `bv-template` |
| Can we even write a BV — or does the Tarifvertrag block it? | `tarifvertrag-check` |
| Do these redundancies trigger § 17 KSchG? | `massenentlassung` |
| What is the strongest Widerspruch ground? | `widerspruch-check` |
| Draft a formal information request to the employer | `auskunft` |
| Does this AI/IT system trigger § 87 Nr. 6? | `ki-check` |
| What can employees claim if employer skipped Interessenausgleich? | `nachteilsausgleich` |
| Request training time off under § 37 Abs. 6 | `schulungsantrag` |
| Generate BR meeting minutes | `protokoll` |

---

## Install

### Recommended (CLI + Claude Code skill)

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

### CLI only

```bash
npx -y @mvanhorn/printing-press install betriebsrat --cli-only
```

### From source (Go 1.21+)

```bash
git clone https://github.com/googlarz/betriebsrat-pp-cli
cd betriebsrat-pp-cli
go install ./cmd/betriebsrat-pp-cli/
```

### Pre-built binary

Download from [Releases](../../releases). On macOS, remove quarantine: `xattr -d com.apple.quarantine betriebsrat-pp-cli`. On Linux: `chmod +x betriebsrat-pp-cli`.

### Verify install

```bash
betriebsrat-pp-cli doctor
```

---

## Quick Start

```bash
# 1. Populate local knowledge base (run once)
betriebsrat-pp-cli sync

# 2. Instant co-determination check
betriebsrat-pp-cli rights-check "employer wants to introduce monitoring software"

# 3. Full decision support for a situation
betriebsrat-pp-cli decide "Arbeitgeber kündigt 15 Mitarbeiter"

# 4. Calculate the exact deadline
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from $(date +%Y-%m-%d)

# 5. Look up a paragraph in plain language
betriebsrat-pp-cli law 87

# 6. English output on any command
betriebsrat-pp-cli rights-check "employer wants to introduce monitoring software" --lang en
```

---

## Commands

### Decision support

#### `rights-check` — Does the BR have co-determination rights?

The most common works council question. Maps your situation to BetrVG paragraphs and classifies the right type.

```bash
betriebsrat-pp-cli rights-check "Arbeitgeber will Überwachungssoftware einführen"
betriebsrat-pp-cli rights-check "employer wants to introduce home office policy" --lang en --json
```

Output: applicable §§, co-determination type (erzwingbar / Mitwirkung / Unterrichtung / keine), recommendation.

#### `decide` — Full decision support

Classifies the situation, lists all BR rights, produces a prioritised action plan, and surfaces deadlines.

```bash
betriebsrat-pp-cli decide "Betrieb soll nach München verlagert werden"
betriebsrat-pp-cli decide "mass layoff of 20 employees" --lang en --agent
```

#### `codetermination-type` — What kind of right is it?

Classifies the BR's right as one of: **Mitbestimmung (erzwingbar)** · Zustimmungsvorbehalt · Mitwirkung · Beratung · Unterrichtung · keine.

```bash
betriebsrat-pp-cli codetermination-type "Versetzung"
betriebsrat-pp-cli codetermination-type "Einführung Schichtplan" --json
```

#### `consequences` — What happens if a deadline is missed or the employer acts unilaterally?

```bash
betriebsrat-pp-cli consequences kündigung
betriebsrat-pp-cli consequences betriebsänderung --lang en
betriebsrat-pp-cli consequences software   # software / AI without BV
betriebsrat-pp-cli consequences br-deadline  # BR doesn't respond in time
```

Situations: `kündigung` · `einstellung` · `versetzung` · `betriebsänderung` · `software` · `br-deadline`

#### `checklist` — Step-by-step action checklist

```bash
betriebsrat-pp-cli checklist "Betriebsänderung"
betriebsrat-pp-cli checklist "Massenentlassung" --lang en
```

---

### Deadlines

#### `deadline` — Calculate legal response deadlines

```bash
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from 2026-05-10
betriebsrat-pp-cli deadline "außerordentliche Kündigung" --from 2026-05-10
# Export to calendar (.ics)
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from 2026-05-10 --ical > frist.ics
```

Key deadlines built in: § 102 Anhörung (7 days ordinary, 3 days extraordinary), § 99 Einstellung/Versetzung (1 week), § 17 KSchG Massenentlassung (30 days), and more.

---

### Document generation

#### `bv-template` — Generate a Betriebsvereinbarung skeleton

```bash
betriebsrat-pp-cli bv-template homeoffice --employer "Musterfirma GmbH"
betriebsrat-pp-cli bv-template software --agent
betriebsrat-pp-cli bv-template videoüberwachung --employer "Firma AG"
betriebsrat-pp-cli bv-template leistungsbeurteilung --employer "TechCo GmbH"
```

Topics: `homeoffice` · `software` · `arbeitszeit` · `datenschutz` · `videoüberwachung` · `leistungsbeurteilung`

Each template includes legally required clauses, BR-protective clauses, and negotiation tips inline.

#### `letter` — Draft formal BR letters

```bash
betriebsrat-pp-cli letter kündigung --type widerspruch
betriebsrat-pp-cli letter betriebsänderung --type einigungsstelle
```

#### `auskunft` — Draft a § 80 BetrVG information request

```bash
betriebsrat-pp-cli auskunft --topic sozialdaten --reason "Prüfung Sozialauswahl § 102" --employer "Firma GmbH"
betriebsrat-pp-cli auskunft --topic ki --reason "Einführung KI-Bewertungssystem" --deadline-days 10
betriebsrat-pp-cli auskunft --topic custom --custom "Überstundenaufstellungen der letzten 12 Monate"
```

Topics: `sozialdaten` · `stellenplan` · `gehaelter` · `planung` · `auswahlrichtlinien` · `ki` · `wirtschaft` · `custom`

#### `schulungsantrag` — Draft a § 37 Abs. 6 training request letter

```bash
betriebsrat-pp-cli schulungsantrag --topic betrvg --employer "Musterfirma GmbH"
betriebsrat-pp-cli schulungsantrag --topic kuendigung --provider "ver.di Bildung" --employer "AG GmbH"
betriebsrat-pp-cli schulungsantrag --topic custom --training-name "KI im Betrieb" --employer "TechCo"
```

Topics: `betrvg` · `arbeitsrecht` · `betriebsrat-praxis` · `kuendigung` · `sozialplan` · `datenschutz` · `gesundheit` · `custom`

Includes statutory justification per topic (BAG case law), cost and release-from-work claims.

#### `protokoll` — Generate a BR Sitzungsprotokoll template

```bash
betriebsrat-pp-cli protokoll --topic "Anhörung Kündigung Müller" --br-size 9 --date 2026-05-10
```

---

### Legal checks

#### `check-anhoerung` — Is the Anhörungsschreiben complete? Is the clock running?

```bash
betriebsrat-pp-cli check-anhoerung "Sehr geehrter Betriebsrat, wir beabsichtigen..."
betriebsrat-pp-cli check-anhoerung --file anhörung.txt --type ordentlich
```

Reports missing Sozialdaten, flags whether the 7-day clock has started, severity per gap.

#### `massenentlassung` — Does this trigger § 17 KSchG? Full procedure checklist.

```bash
betriebsrat-pp-cli massenentlassung --employees 120 --planned-dismissals 15
betriebsrat-pp-cli massenentlassung --employees 50 --planned-dismissals 8 --lang en
```

Checks § 17 KSchG thresholds and generates the step-by-step Massenentlassung procedure including BA notification requirements.

#### `widerspruch-check` — What are the strongest § 102 Abs. 3 Widerspruch grounds?

```bash
betriebsrat-pp-cli widerspruch-check --reason "Sozialauswahl fehlerhaft" --age 52 --years 18
betriebsrat-pp-cli widerspruch-check --reason "freier Arbeitsplatz vorhanden" --lang en
```

Ranks grounds by legal strength and drafts the ground text.

#### `ki-check` — Does this AI/IT system trigger § 87 Nr. 6 co-determination?

```bash
betriebsrat-pp-cli ki-check --system "Leistungsmonitoring-Dashboard" --monitors-performance --influences-hr
betriebsrat-pp-cli ki-check --system "AI recruitment screener" --auto-decision --lang en
```

Risk-scores the system, lists required BV clauses, cites relevant BAG rulings.

#### `tarifvertrag-check` — Does the Tarifvertrag block a planned BV?

```bash
betriebsrat-pp-cli tarifvertrag-check --topic lohn --tv-type "Branchentarifvertrag" --tv-covers
betriebsrat-pp-cli tarifvertrag-check --topic homeoffice --no-tv-covers
betriebsrat-pp-cli tarifvertrag-check --topic software --lang en
betriebsrat-pp-cli tarifvertrag-check --topic arbeitszeit --tv-covers --opening-clause
```

Topics: `lohn` · `arbeitszeit` · `urlaub` · `zulagen` · `homeoffice` · `software` · `gesundheit` · `custom`

Always run this before drafting a BV in a TV-regulated area.

---

### Calculations

#### `sozialplan-calc` — Calculate Sozialplan entitlement (Munich formula)

```bash
betriebsrat-pp-cli sozialplan-calc --salary 4500 --years 12 --age 48
betriebsrat-pp-cli sozialplan-calc --salary 6000 --years 20 --age 55 --disabled --factor 0.8
# Batch mode (CSV)
betriebsrat-pp-cli sozialplan-calc --csv employees.csv --agent
```

CSV format: `name,salary,years,age,disabled,children[,factor[,max_cap]]`

#### `nachteilsausgleich` — § 113 BetrVG individual compensation claim

For when the employer implemented a Betriebsänderung without attempting an Interessenausgleich.

```bash
betriebsrat-pp-cli nachteilsausgleich --salary 4500 --years 8 --age 42 --measure "Standortschließung" --no-ia-attempted
betriebsrat-pp-cli nachteilsausgleich --salary 6000 --years 15 --age 55 --measure "relocation" --ia-deviated --lang en
```

---

### Legal reference

#### `law` — Plain-language BetrVG paragraph lookup

```bash
betriebsrat-pp-cli law 87
betriebsrat-pp-cli law 102 --json
betriebsrat-pp-cli law kündigung   # keyword search
```

#### `law` — Meeting preparation

```bash
betriebsrat-pp-cli prepare-meeting "Einführung KI-System"
betriebsrat-pp-cli prepare-meeting "Betriebsänderung" --lang en
```

#### `context` — Store company profile for calibrated advice

```bash
betriebsrat-pp-cli context set --employees 85 --tariff true --br-size 7 --sector "Handel"
betriebsrat-pp-cli context show
```

Filters advice by employee count thresholds: § 111 Betriebsänderung (≥20 AN), § 106 Wirtschaftsausschuss (≥100 AN), full release (≥200 AN).

---

## Output formats

```bash
# Human-readable (default)
betriebsrat-pp-cli rights-check "Überwachungssoftware"

# JSON — for scripting, piping, agents
betriebsrat-pp-cli rights-check "Überwachungssoftware" --json

# Compact JSON — key fields only, minimal tokens for LLMs
betriebsrat-pp-cli rights-check "Überwachungssoftware" --compact

# Filter specific fields
betriebsrat-pp-cli rights-check "Überwachungssoftware" --json --select summary,recommendation

# Agent mode — JSON + compact + no prompts + no color in one flag
betriebsrat-pp-cli rights-check "Überwachungssoftware" --agent

# English output
betriebsrat-pp-cli rights-check "Überwachungssoftware" --lang en

# Dry run — show request without executing
betriebsrat-pp-cli rights-check "Überwachungssoftware" --dry-run
```

Exit codes: `0` success · `2` usage error · `3` not found · `5` API error · `7` rate limited · `10` config error

---

## Use with Claude Code (skill)

Install the agent skill — it drives the CLI automatically:

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

Then in Claude Code, describe your situation in natural language. The skill auto-classifies the situation, runs the right commands, and chains follow-ups automatically (e.g. `ki-check` → draft `auskunft` letter; `widerspruch-check` → draft the Widerspruch + `protokoll`).

**Bilingual:** Write your query in English and add `--lang en` — or just write in English and the skill detects it.

## Use as an MCP server

```bash
# Claude Code
claude mcp add betriebsrat betriebsrat-pp-mcp

# Claude Desktop — add to claude_desktop_config.json
{
  "mcpServers": {
    "betriebsrat": { "command": "betriebsrat-pp-mcp" }
  }
}
```

---

## Offline capability

All decision-support commands run entirely offline from the embedded knowledge base — no network call required after the initial `sync`:

`rights-check` · `decide` · `deadline` · `checklist` · `law` · `codetermination-type` · `consequences` · `letter` · `sozialplan-calc` · `nachteilsausgleich` · `massenentlassung` · `widerspruch-check` · `protokoll` · `auskunft` · `ki-check` · `schulungsantrag` · `tarifvertrag-check` · `bv-template` · `context` · `check-anhoerung`

---

## License

Apache 2.0 — see [LICENSE](LICENSE).

---

---

<a name="deutsch"></a>

# betriebsrat-pp-cli — Deutsche Dokumentation

**Betriebsrat-Berater für die Kommandozeile — BetrVG-Rechte, Fristen und Entscheidungshilfe — vollständig offline.**

Kein API-Key. Kein Login. Kein Abo. Alles läuft lokal aus einer vorberechneten Wissensdatenbank mit BetrVG-Paragrafen, gesetzlichen Fristen und situationsbasierter Entscheidungslogik.

Gebaut für Betriebsratsmitglieder, die Antworten in Minuten brauchen — nicht nach stundenlangem Lesen von Gesetzeskommentaren oder Suchen auf Rechtswebsites.

---

## Was es kann

| Frage | Befehl |
|---|---|
| Hat der BR hier ein Mitbestimmungsrecht? | `rights-check` |
| Welche Art von Recht haben wir? | `codetermination-type` |
| Was sollen wir jetzt tun, Schritt für Schritt? | `decide` |
| Bis wann müssen wir antworten? | `deadline` |
| Was bedeutet § 87 BetrVG eigentlich? | `law 87` |
| Was passiert, wenn der AG ohne uns handelt? | `consequences` |
| Ist unser Anhörungsschreiben vollständig? | `check-anhoerung` |
| Wie viel Sozialplan steht diesem Mitarbeiter zu? | `sozialplan-calc` |
| Betriebsvereinbarung erstellen | `bv-template` |
| Dürfen wir überhaupt eine BV abschließen oder sperrt der TV das? | `tarifvertrag-check` |
| Lösen diese Entlassungen § 17 KSchG aus? | `massenentlassung` |
| Was ist das stärkste Widerspruchsargument? | `widerspruch-check` |
| Formellen Auskunftsantrag an den AG formulieren | `auskunft` |
| Löst dieses KI-/IT-System § 87 Nr. 6 aus? | `ki-check` |
| Was können Arbeitnehmer fordern, wenn AG den IA übergangen hat? | `nachteilsausgleich` |
| Schulungsantrag nach § 37 Abs. 6 stellen | `schulungsantrag` |
| BR-Sitzungsprotokoll erstellen | `protokoll` |

---

## Installation

### Empfohlen (CLI + Claude Code Skill)

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

### Nur CLI

```bash
npx -y @mvanhorn/printing-press install betriebsrat --cli-only
```

### Aus dem Quellcode (Go 1.21+)

```bash
git clone https://github.com/googlarz/betriebsrat-pp-cli
cd betriebsrat-pp-cli
go install ./cmd/betriebsrat-pp-cli/
```

### Installation prüfen

```bash
betriebsrat-pp-cli doctor
```

---

## Schnellstart

```bash
# 1. Wissensdatenbank befüllen (einmalig)
betriebsrat-pp-cli sync

# 2. Mitbestimmungsrecht prüfen
betriebsrat-pp-cli rights-check "Arbeitgeber will Überwachungssoftware einführen"

# 3. Vollständige Entscheidungsunterstützung
betriebsrat-pp-cli decide "Arbeitgeber kündigt 15 Mitarbeiter"

# 4. Genaue Frist berechnen
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from $(date +%Y-%m-%d)

# 5. BetrVG-Paragraf in Alltagssprache
betriebsrat-pp-cli law 87

# 6. Englische Ausgabe bei jedem Befehl
betriebsrat-pp-cli rights-check "Arbeitgeber will Überwachungssoftware einführen" --lang en
```

---

## Befehle im Detail

### Entscheidungsunterstützung

#### `rights-check` — Hat der BR ein Mitbestimmungsrecht?

Die häufigste Frage im Betriebsrat. Ordnet die Situation BetrVG-Paragrafen zu und klassifiziert den Rechtstyp.

```bash
betriebsrat-pp-cli rights-check "Arbeitgeber will Überwachungssoftware einführen"
betriebsrat-pp-cli rights-check "Homeoffice-Regelung" --json
```

Ausgabe: einschlägige §§, Mitbestimmungstyp (erzwingbar / Mitwirkung / Unterrichtung / keine), Handlungsempfehlung.

#### `decide` — Vollständige Entscheidungshilfe

Klassifiziert die Situation, listet alle BR-Rechte, erstellt einen priorisierten Maßnahmenplan und zeigt Fristen.

```bash
betriebsrat-pp-cli decide "Betrieb soll nach München verlagert werden"
betriebsrat-pp-cli decide "Arbeitgeber kündigt 15 Mitarbeiter" --agent
```

#### `consequences` — Was passiert bei Fristversäumnis oder einseitigem Handeln des AG?

```bash
betriebsrat-pp-cli consequences kündigung
betriebsrat-pp-cli consequences betriebsänderung
betriebsrat-pp-cli consequences software     # Software/KI ohne BV eingeführt
betriebsrat-pp-cli consequences br-deadline  # BR antwortet nicht rechtzeitig
```

Situationen: `kündigung` · `einstellung` · `versetzung` · `betriebsänderung` · `software` · `br-deadline`

#### `checklist` — Schritt-für-Schritt-Checkliste

```bash
betriebsrat-pp-cli checklist "Betriebsänderung"
betriebsrat-pp-cli checklist "Massenentlassung"
```

---

### Fristen

#### `deadline` — Gesetzliche Antwortfristen berechnen

```bash
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from 2026-05-10
betriebsrat-pp-cli deadline "außerordentliche Kündigung" --from 2026-05-10
# Als Kalenderdatei exportieren
betriebsrat-pp-cli deadline "ordentliche Kündigung" --from 2026-05-10 --ical > frist.ics
```

Eingebaut: § 102 Anhörung (7 Tage ordentlich, 3 Tage außerordentlich), § 99 Einstellung/Versetzung (1 Woche), § 17 KSchG (30 Tage) u.v.m.

---

### Dokumentenerstellung

#### `bv-template` — Betriebsvereinbarungs-Entwurf generieren

```bash
betriebsrat-pp-cli bv-template homeoffice --employer "Musterfirma GmbH"
betriebsrat-pp-cli bv-template software --agent
betriebsrat-pp-cli bv-template videoüberwachung --employer "Firma AG"
betriebsrat-pp-cli bv-template leistungsbeurteilung --employer "TechCo GmbH"
```

Themen: `homeoffice` · `software` · `arbeitszeit` · `datenschutz` · `videoüberwachung` · `leistungsbeurteilung`

Jede Vorlage enthält Pflichtinhalt, BR-Schutzklauseln und Verhandlungstipps direkt im Text.

#### `auskunft` — Formelles Auskunftsverlangen nach § 80 Abs. 2 BetrVG

```bash
betriebsrat-pp-cli auskunft --topic sozialdaten --reason "Prüfung Sozialauswahl § 102" --employer "Firma GmbH"
betriebsrat-pp-cli auskunft --topic ki --reason "Einführung KI-Bewertungssystem" --deadline-days 10
betriebsrat-pp-cli auskunft --topic custom --custom "Überstundenaufstellungen der letzten 12 Monate"
```

Themen: `sozialdaten` · `stellenplan` · `gehaelter` · `planung` · `auswahlrichtlinien` · `ki` · `wirtschaft` · `custom`

#### `schulungsantrag` — Schulungsantrag nach § 37 Abs. 6 BetrVG

```bash
betriebsrat-pp-cli schulungsantrag --topic betrvg --employer "Musterfirma GmbH"
betriebsrat-pp-cli schulungsantrag --topic kuendigung --provider "ver.di Bildung" --employer "AG GmbH"
betriebsrat-pp-cli schulungsantrag --topic custom --training-name "KI im Betrieb" --employer "TechCo"
```

Themen: `betrvg` · `arbeitsrecht` · `betriebsrat-praxis` · `kuendigung` · `sozialplan` · `datenschutz` · `gesundheit` · `custom`

Enthält gesetzliche Begründung je Thema (BAG-Rechtsprechung), Freistellungs- und Kostenerstattungsanspruch.

#### `protokoll` — BR-Sitzungsprotokoll-Vorlage

```bash
betriebsrat-pp-cli protokoll --topic "Anhörung Kündigung Müller" --br-size 9 --date 2026-05-10
```

---

### Rechtliche Prüfungen

#### `check-anhoerung` — Anhörungsschreiben vollständig? Läuft die Frist?

```bash
betriebsrat-pp-cli check-anhoerung "Sehr geehrter Betriebsrat, wir beabsichtigen..."
betriebsrat-pp-cli check-anhoerung --file anhörung.txt --type ordentlich
```

Meldet fehlende Sozialdaten, zeigt ob die 7-Tage-Frist bereits läuft, Schweregrad je Lücke.

#### `massenentlassung` — § 17 KSchG-Schwellenwert und komplettes Verfahren

```bash
betriebsrat-pp-cli massenentlassung --employees 120 --planned-dismissals 15
betriebsrat-pp-cli massenentlassung --employees 50 --planned-dismissals 8
```

Prüft § 17 KSchG-Schwellenwerte und erstellt das vollständige Massenentlassungsverfahren inkl. Anzeige bei der BA.

#### `widerspruch-check` — Stärkste Widerspruchsgründe nach § 102 Abs. 3 BetrVG

```bash
betriebsrat-pp-cli widerspruch-check --reason "Sozialauswahl fehlerhaft" --age 52 --years 18
betriebsrat-pp-cli widerspruch-check --reason "freier Arbeitsplatz vorhanden"
```

Rankt Gründe nach rechtlicher Stärke und formuliert den Widerspruchstext.

#### `ki-check` — Löst dieses KI-/IT-System § 87 Nr. 6 aus?

```bash
betriebsrat-pp-cli ki-check --system "Leistungsmonitoring-Dashboard" --monitors-performance --influences-hr
betriebsrat-pp-cli ki-check --system "KI-Recruiting-Tool" --auto-decision
```

Bewertet das System, listet erforderliche BV-Klauseln, zitiert einschlägige BAG-Entscheidungen.

#### `tarifvertrag-check` — Sperrt der Tarifvertrag die geplante BV?

```bash
betriebsrat-pp-cli tarifvertrag-check --topic lohn --tv-type "Branchentarifvertrag" --tv-covers
betriebsrat-pp-cli tarifvertrag-check --topic homeoffice --no-tv-covers
betriebsrat-pp-cli tarifvertrag-check --topic software
betriebsrat-pp-cli tarifvertrag-check --topic arbeitszeit --tv-covers --opening-clause
```

Themen: `lohn` · `arbeitszeit` · `urlaub` · `zulagen` · `homeoffice` · `software` · `gesundheit` · `custom`

**Immer zuerst ausführen**, bevor eine BV in einem tarifvertraglich geregelten Bereich abgeschlossen wird.

---

### Berechnungen

#### `sozialplan-calc` — Sozialplan-Abfindung berechnen (Münchener Formel)

```bash
betriebsrat-pp-cli sozialplan-calc --salary 4500 --years 12 --age 48
betriebsrat-pp-cli sozialplan-calc --salary 6000 --years 20 --age 55 --disabled --factor 0.8
# Stapelmodus (CSV)
betriebsrat-pp-cli sozialplan-calc --csv mitarbeiter.csv --agent
```

CSV-Format: `name,gehalt,jahre,alter,schwerbehindert,kinder[,faktor[,max_cap]]`

#### `nachteilsausgleich` — Individueller Nachteilsausgleichsanspruch nach § 113 BetrVG

Für den Fall, dass der AG eine Betriebsänderung ohne Versuch eines Interessenausgleichs durchgeführt hat.

```bash
betriebsrat-pp-cli nachteilsausgleich --salary 4500 --years 8 --age 42 --measure "Standortschließung" --no-ia-attempted
betriebsrat-pp-cli nachteilsausgleich --salary 6000 --years 15 --age 55 --measure "Verlagerung" --ia-deviated
```

---

### Nachschlagewerk

#### `law` — BetrVG-Paragraf in Alltagssprache

```bash
betriebsrat-pp-cli law 87
betriebsrat-pp-cli law 102 --json
betriebsrat-pp-cli law kündigung   # Stichwortsuche
```

#### `context` — Betriebsprofil speichern für kalibrierte Beratung

```bash
betriebsrat-pp-cli context set --employees 85 --tariff true --br-size 7 --sector "Handel"
betriebsrat-pp-cli context show
```

Filtert Empfehlungen nach Mitarbeiterzahl-Schwellenwerten: § 111 (ab 20 AN), § 106 Wirtschaftsausschuss (ab 100 AN), Vollfreistellung (ab 200 AN).

---

## Ausgabeformate

```bash
# Lesbare Ausgabe (Standard)
betriebsrat-pp-cli rights-check "Überwachungssoftware"

# JSON — für Skripte, Pipes, Agenten
betriebsrat-pp-cli rights-check "Überwachungssoftware" --json

# Nur bestimmte Felder ausgeben
betriebsrat-pp-cli rights-check "Überwachungssoftware" --json --select summary,recommendation

# Agent-Modus — JSON + kompakt + keine Prompts in einem Flag
betriebsrat-pp-cli rights-check "Überwachungssoftware" --agent

# Englische Ausgabe
betriebsrat-pp-cli rights-check "Überwachungssoftware" --lang en
```

Exit-Codes: `0` Erfolg · `2` Nutzungsfehler · `3` nicht gefunden · `5` API-Fehler · `7` Rate Limit · `10` Konfigurationsfehler

---

## Offline-Fähigkeit

Alle Entscheidungshilfe-Befehle laufen vollständig offline aus der eingebetteten Wissensdatenbank — nach dem einmaligen `sync` ist kein Netzwerkzugriff mehr nötig:

`rights-check` · `decide` · `deadline` · `checklist` · `law` · `codetermination-type` · `consequences` · `letter` · `sozialplan-calc` · `nachteilsausgleich` · `massenentlassung` · `widerspruch-check` · `protokoll` · `auskunft` · `ki-check` · `schulungsantrag` · `tarifvertrag-check` · `bv-template` · `context` · `check-anhoerung`

---

## Verwendung mit Claude Code

Nach der Installation treibt der Skill die CLI automatisch an. Schreiben Sie Ihre Situation auf Deutsch — der Skill erkennt die Sprache, führt die richtigen Befehle aus und kettet Folgeschritte automatisch (z.B. `ki-check` → Auskunftsverlangen; `widerspruch-check` → Widerspruchstext + Protokoll).

```bash
npx -y @mvanhorn/printing-press install betriebsrat
```

---

## Lizenz

Apache 2.0 — siehe [LICENSE](LICENSE).

---

*Rechtlicher Hinweis: Dieses Tool dient der rechtlichen Orientierung und ersetzt keine anwaltliche Beratung. Bei komplexen Einzelfällen empfehlen wir die Konsultation eines Fachanwalts für Arbeitsrecht oder Ihrer Gewerkschaft.*

*Legal notice: This tool is for legal orientation purposes and does not replace professional legal advice. For complex individual cases we recommend consulting a labour law specialist or your trade union.*
